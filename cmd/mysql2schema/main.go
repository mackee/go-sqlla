package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/serenize/snaker"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ScanStateBeforeColmuns = iota
	ScanStateInColumns
	ScanStateDefineIndex
	ScanStateAfterColmuns
)

var (
	dsn         string
	tables      []string
	outdir      string
	packageName string
)

func main() {
	var err error
	flag.StringVar(&dsn, "dsn", "user:password@tcp(localhost:3306)/dbname", "target data source name (see github.com/go-sql-driver/mysql)")
	flag.StringVar(&outdir, "outdir", "", "output directory")
	flag.StringVar(&packageName, "package", "", "package name when output to file")

	flag.Parse()
	tables = flag.Args()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed connect to data source: %s", err)
	}
	defer db.Close()

	for _, table := range tables {
		out := new(bytes.Buffer)
		if packageName != "" {
			io.WriteString(out, "package "+packageName+"\n\n")
			io.WriteString(out, "//go:generate sqlla\n\n")
		}
		err = outputSchema(db, table, out)
		if err != nil {
			log.Fatalf("failed output source in %s: %s", table, err)
		}
		if outdir != "" {
			filename := filepath.Join(outdir, fmt.Sprintf("%s.schema.go", table))
			f, err := os.Create(filename)
			if err != nil {
				log.Fatalf("fail create output file: %s", err)
			}
			bs, err := imports.Process(filename, out.Bytes(), nil)
			if err != nil {
				log.Fatalf("fail run goimports to output file: %s", err)
			}
			_, err = f.Write(bs)
			if err != nil {
				log.Fatalf("fail write to output file: %s", err)
			}
			f.Close()
		}
	}
}

type Column struct {
	Name       string
	Type       string
	IsUnsigned bool
	IsNull     bool
}

func outputSchema(db *sql.DB, table string, out io.Writer) error {
	query := fmt.Sprintf("SHOW CREATE TABLE %s", table)
	row := db.QueryRow(query)
	var tableName string
	var ddl string
	err := row.Scan(&tableName, &ddl)
	if err != nil {
		return err
	}
	buf := bytes.NewBufferString(ddl)
	s := bufio.NewScanner(buf)

	columns := make([]Column, 0, 16)
	state := ScanStateBeforeColmuns
	for s.Scan() {
		l := s.Text()
		l = strings.TrimSpace(l)
		switch state {
		case ScanStateBeforeColmuns:
			if strings.HasSuffix(l, "(") {
				state = ScanStateInColumns
			}
			continue
		case ScanStateInColumns, ScanStateDefineIndex:
			if strings.HasPrefix(l, "`") {
				state = ScanStateInColumns
			} else {
				state = ScanStateDefineIndex
			}
			if strings.HasPrefix(l, ")") {
				state = ScanStateAfterColmuns
			}
		default:
			break
		}
		if state != ScanStateInColumns {
			continue
		}
		var c Column
		var beforeToken string
		defs := strings.Split(l, " ")
		for _, d := range defs {
			switch {
			case c.Name == "" && strings.HasPrefix(d, "`"):
				c.Name = strings.Trim(d, "`")
			case c.Type == "" && c.Name != "":
				c.Type = d
			case d == "unsigned":
				c.IsUnsigned = true
			case d == "NULL":
				switch beforeToken {
				case "NOT":
					c.IsNull = false
				case "DEFAULT":
					c.IsNull = true
				}
			}
			beforeToken = d
		}
		columns = append(columns, c)
	}

	outBuf := new(bytes.Buffer)
	outBuf.WriteString("//+table: " + table + "\n")
	outBuf.WriteString("type ")
	outBuf.WriteString(snaker.SnakeToCamel(table))
	outBuf.WriteString(" struct {\n")
	for _, c := range columns {
		outBuf.WriteString(snaker.SnakeToCamel(c.Name))
		outBuf.WriteString(" ")
		schemaType, err := sqlTypeToSchemaType(c)
		if err != nil {
			return err
		}
		outBuf.WriteString(schemaType)
		outBuf.WriteString(fmt.Sprintf(" `db:\"%s\"`\n", c.Name))
	}
	outBuf.WriteString("}")

	formated, err := format.Source(outBuf.Bytes())
	if err != nil {
		return fmt.Errorf("go format error: %s", err)
	}
	out.Write(formated)

	return nil
}

func sqlTypeToSchemaType(c Column) (string, error) {
	switch {
	case strings.HasPrefix(c.Type, "bigint"):
		if c.IsNull {
			return "sql.NullInt64", nil
		}
		if c.IsUnsigned {
			return "uint64", nil
		}
		return "int64", nil
	case strings.HasPrefix(c.Name, "is_") || strings.HasPrefix(c.Name, "has_"):
		if c.IsNull {
			return "sql.NullBool", nil
		}
		return "bool", nil
	case strings.HasPrefix(c.Type, "int"):
		if c.IsNull {
			return "sql.NullInt64", nil
		}
		if c.IsUnsigned {
			return "uint32", nil
		}
		return "int32", nil
	case strings.HasPrefix(c.Type, "tinyint"):
		if c.IsNull {
			return "sql.NullInt64", nil
		}
		if c.IsUnsigned {
			return "uint8", nil
		}
		return "int8", nil
	case strings.HasPrefix(c.Type, "varchar"), strings.HasPrefix(c.Type, "text"), strings.HasPrefix(c.Type, "json"):
		if c.IsNull {
			return "sql.NullString", nil
		}
		return "string", nil
	case strings.HasPrefix(c.Type, "datetime"):
		if c.IsNull {
			return "mysql.NullTime", nil
		}
		return "time.Time", nil
	default:
		return "", fmt.Errorf("unexpected column type: %+v\n", c)
	}
}

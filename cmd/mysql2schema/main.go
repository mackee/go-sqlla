package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/serenize/snaker"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ScanStateBeforeColmuns = iota
	ScanStateInColumns
	ScanStateDefineIndex
	ScanStateAfterColmuns
)

var dsn string
var tables []string

func main() {
	var err error
	flag.StringVar(&dsn, "dsn", "user:password@tcp(localhost:3306)/dbname", "target data source name (see github.com/go-sql-driver/mysql)")

	flag.Parse()
	tables = flag.Args()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed connect to data source: %s", err)
	}
	defer db.Close()

	for _, table := range tables {
		err = outputSchema(db, table)
		if err != nil {
			log.Fatalf("failed output source in %s: %s", table, err)
		}
	}
}

type Column struct {
	Name       string
	Type       string
	IsUnsigned bool
	IsNull     bool
}

func outputSchema(db *sql.DB, table string) error {
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

	out := new(bytes.Buffer)
	out.WriteString("type ")
	out.WriteString(snaker.SnakeToCamel(table))
	out.WriteString(" struct {\n")
	for _, c := range columns {
		out.WriteString(snaker.SnakeToCamel(c.Name))
		out.WriteString(" ")
		schemaType, err := sqlTypeToSchemaType(c)
		if err != nil {
			return err
		}
		out.WriteString(schemaType)
		out.WriteString(fmt.Sprintf(" `db:\"%s\"`\n", c.Name))
	}
	out.WriteString("}")

	formated, err := format.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("go format error: %s", err)
	}
	os.Stdout.Write(formated)

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
	case strings.HasPrefix(c.Type, "varchar"):
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

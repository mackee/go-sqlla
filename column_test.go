package sqlla_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/mackee/go-sqlla/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColumnNullT(t *testing.T) {
	cwd, err := os.Getwd()
	t.Cleanup(func() {
		require.NoError(t, os.Chdir(cwd))
	})

	require.NoError(t, err)
	fullpath, err := filepath.Abs(filepath.Join("./testdata", "nullt", "repoa", "schema.go"))
	require.NoError(t, err)
	dir := filepath.Dir(fullpath)
	require.NoError(t, os.Chdir(dir))

	pkgs, err := sqlla.ToPackages(dir)
	require.NoError(t, err)
	var table *sqlla.Table
	for _, pkg := range pkgs {
		files := pkg.Syntax
		for _, f := range files {
			for _, decl := range f.Decls {
				_table, err := sqlla.DeclToTable(pkg, decl, fullpath, true)
				if errors.Is(err, sqlla.ErrNotTargetDecl) {
					continue
				}
				require.NoError(t, err)
				table = _table
			}
		}
	}
	require.NotNil(t, table)
	assert.False(t, table.Columns[0].IsNullT()) // ID
	assert.Equal(t, table.Columns[0].BaseTypeName(), "uint64")
	assert.Equal(t, table.Columns[0].ExprValue(), "sqlla.ExprValue[uint64]")
	assert.Equal(t, table.Columns[0].ExprMultiValue(), "sqlla.ExprMultiValue[uint64]")
	assert.True(t, table.Columns[1].IsNullT()) // ModifiedAt
	assert.Equal(t, table.Columns[1].BaseTypeName(), "time.Time")
	assert.Equal(t, table.Columns[1].ExprValue(), "sqlla.ExprNull[time.Time]")
	assert.Equal(t, table.Columns[1].ExprMultiValue(), "sqlla.ExprMultiValue[time.Time]")
	assert.False(t, table.Columns[2].IsNullT())
	assert.Equal(t, table.Columns[2].BaseTypeName(), "NullableTime")
	assert.Equal(t, table.Columns[2].ExprValue(), "sqlla.ExprValue[NullableTime]")
	assert.Equal(t, table.Columns[2].ExprMultiValue(), "sqlla.ExprMultiValue[NullableTime]")
}

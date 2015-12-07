//go:generate gen -f
package schema

import (
	"database/sql"
	"html/template"
	"regexp"
)

// +gen * slice:"All,Any,Count,DistinctBy,First,GroupBy[string],Shuffle,SortBy,Where"
type Column struct {
	TableCatalog           string         `db:"TABLE_CATALOG"`
	TableSchema            string         `db:"TABLE_SCHEMA"`
	TableName              string         `db:"TABLE_NAME"`
	ColumnName             string         `db:"COLUMN_NAME"`
	OrdinalPosition        int64          `db:"ORDINAL_POSITION"`
	ColumnDefault          sql.NullString `db:"COLUMN_DEFAULT"`
	IsNullable             string         `db:"IS_NULLABLE"`
	DataType               string         `db:"DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       sql.NullInt64  `db:"NUMERIC_PRECISION"`
	NumericScale           sql.NullInt64  `db:"NUMERIC_SCALE"`
	DatetimePrecision      sql.NullString `db:"DATETIME_PRECISION"`
	CharacterSetName       sql.NullString `db:"CHARACTER_SET_NAME"`
	CollationName          sql.NullString `db:"COLLATION_NAME"`
	ColumnType             string         `db:"COLUMN_TYPE"`
	ColumnKey              string         `db:"COLUMN_KEY"`
	Extra                  string         `db:"EXTRA"`
	Privileges             string         `db:"PRIVILEGES"`
	ColumnComment          string         `db:"COLUMN_COMMENT"`
}

func (c *Column) GetColumnComment() template.HTML {
	if len(c.ColumnComment) <= 0 {
		return "　"
	}
	return template.HTML(regexp.MustCompile(`\n|\r\n|\\n|\\r\\n`).ReplaceAllString(c.ColumnComment, "<br>"))
}

func (c *Column) GetColumnKey() string {
	if len(c.ColumnKey) <= 0 {
		return "　"
	}
	return c.ColumnKey
}

func (c *Column) GetExtra() string {
	if len(c.Extra) <= 0 {
		return "　"
	}
	return c.Extra
}

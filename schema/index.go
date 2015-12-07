package schema

import "database/sql"

// +gen * slice:"All,Any,Count,DistinctBy,First,GroupBy[string],Shuffle,SortBy,Where"
type Index struct {
	TableCatalog string         `db:"TABLE_CATALOG"`
	TableSchema  string         `db:"TABLE_SCHEMA"`
	TableName    string         `db:"TABLE_NAME"`
	NonUnique    int32          `db:"NON_UNIQUE"`
	IndexSchema  string         `db:"INDEX_SCHEMA"`
	IndexName    string         `db:"INDEX_NAME"`
	SeqInIndex   int32          `db:"SEQ_IN_INDEX"`
	ColumnName   string         `db:"COLUMN_NAME"`
	Collation    sql.NullString `db:"COLLATION"`
	Cardinality  sql.NullString `db:"CARDINALITY"`
	SubPart      sql.NullInt64  `db:"SUB_PART"`
	Packed       sql.NullString `db:"PACKED"`
	Nullable     string         `db:"NULLABLE"`
	IndexType    string         `db:"INDEX_TYPE"`
	Comment      sql.NullString `db:"COMMENT"`
	IndexComment string         `db:"INDEX_COMMENT"`
}

func (i *Index) GetCollation() string {
	if i.Collation.Valid {
		return i.Collation.String
	}
	return "　"
}

func (i *Index) GetCardinality() string {
	if i.Cardinality.Valid {
		return i.Cardinality.String
	}
	return "　"
}

func (i *Index) GetSubPart() int64 {
	if i.SubPart.Valid {
		return i.SubPart.Int64
	}
	return int64(0)
}

func (i *Index) GetPacked() string {
	if i.Packed.Valid {
		return i.Packed.String
	}
	return "　"
}

func (i *Index) GetComment() string {
	if i.Comment.Valid && len(i.Comment.String) > 0 {
		return i.Comment.String
	}
	return "　"
}

//go:generate gen -f
package schema

// +gen * slice:"All,Any,Count,DistinctBy,First,GroupBy[string],Shuffle,SortBy,Where"
type Schema struct {
	Name        string
	Table       *Table
	ColumnSlice *ColumnSlice
	ValueSlice  *ValueSlice
	IndexSlice  *IndexSlice
}

func NewSchema(Name string, Table *Table, ColumnSlice *ColumnSlice, ValueSlice *ValueSlice, IndexSlice *IndexSlice) *Schema {
	return &Schema{
		Name:        Name,
		Table:       Table,
		ColumnSlice: ColumnSlice,
		ValueSlice:  ValueSlice,
		IndexSlice:  IndexSlice,
	}
}

func (s *Schema) GetExsample(OrdinalPosition int64) string {
	if len(*s.ValueSlice) <= 0 {
		return "　"
	}
	if (*s.ColumnSlice)[OrdinalPosition-1].ColumnType == "datetime" {
		return "2016-01-01 00:00:00"
	}
	return (*s.ValueSlice)[OrdinalPosition-1].(string)
}

//go:generate gen -f
package schema

// +gen * slice:"All,Any,Count,DistinctBy,First,GroupBy[string],Shuffle,SortBy,Where"
type Schema struct {
	Name        string
	ColumnSlice *ColumnSlice
	ValueSlice  *ValueSlice
	IndexSlice  *IndexSlice
}

func NewSchema(Name string, ColumnSlice *ColumnSlice, ValueSlice *ValueSlice, IndexSlice *IndexSlice) *Schema {
	return &Schema{
		Name:        Name,
		ColumnSlice: ColumnSlice,
		ValueSlice:  ValueSlice,
		IndexSlice:  IndexSlice,
	}
}

func (s *Schema) GetExsample(OrdinalPosition int64) string {
	if len(*s.ValueSlice) <= 0 {
		return "　"
	}
	return (*s.ValueSlice)[OrdinalPosition-1].(string)
}

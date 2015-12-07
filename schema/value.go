//go:generate gen -f
package schema

// +gen slice:"All,Any,Count,DistinctBy,First,GroupBy[string],Shuffle,SortBy,Where"
type Value interface{}

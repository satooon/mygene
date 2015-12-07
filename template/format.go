package template

type OutPutFormat string

const (
	MarkDown OutPutFormat = "md"
	Csv      OutPutFormat = "csv"
)

var OutPutFormats = []string{
	MarkDown.String(),
	Csv.String(),
}

func (o OutPutFormat) String() string {
	return string(o)
}

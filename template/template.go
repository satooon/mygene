package template

import (
	"html/template"

	"os"

	"fmt"

	"github.com/satooon/mygene/schema"
)

type Template struct {
	Format      OutPutFormat
	OutPut      string
	SchemaSlice *schema.SchemaSlice
}

func NewTemplate(Format OutPutFormat, OutPut string, SchemaSlice *schema.SchemaSlice) *Template {
	return &Template{
		Format:      Format,
		OutPut:      OutPut,
		SchemaSlice: SchemaSlice,
	}
}

func (t *Template) Print() (err error) {
	if err = os.MkdirAll(t.OutPut, 0777); err != nil {
		return
	}
	switch t.Format {
	case MarkDown:
		if err = outputMarkDown(t); err != nil {
			return
		}
	case Csv:
		if err = outputCsv(t); err != nil {
			return
		}
	}
	return
}

func outputMarkDown(temp *Template) (err error) {
	var file *os.File
	for _, s := range *temp.SchemaSlice {
		file, err = os.Create(temp.OutPut + "/" + fmt.Sprintf("%s.%s", s.Name, temp.Format))
		if err != nil {
			return
		}

		tmp := template.Must(template.New("schema").Funcs(template.FuncMap{
			"GetExsample": s.GetExsample,
		}).Parse(mdTpl))
		if tmp.Execute(file, s); err != nil {
			return
		}
	}
	return
}

func outputCsv(temp *Template) (err error) {
	var file *os.File
	for _, s := range *temp.SchemaSlice {
		file, err = os.Create(temp.OutPut + "/" + fmt.Sprintf("%s.%s", s.Name, temp.Format))
		if err != nil {
			return
		}

		tmp := template.Must(template.New("schema").Funcs(template.FuncMap{
			"GetExsample": s.GetExsample,
		}).Parse(csvTpl))
		if tmp.Execute(file, s); err != nil {
			return
		}
	}
	return
}

var mdTpl string = `# {{.Name}}
|No|名前|型|Null|Key|Extra|Exsample|備考|
|:--|:--|:--:|:--:|:--:|:--:|:--|:--|
{{range $column := .ColumnSlice}}|{{$column.OrdinalPosition}}|{{$column.ColumnName}}|{{$column.ColumnType}}|{{$column.IsNullable}}|{{$column.GetColumnKey}}|{{$column.GetExtra}}|{{$column.OrdinalPosition|GetExsample}}|{{$column.GetColumnComment}}|
{{end}}

` + mdTplIdx

var mdTplIdx string = `## Index
|非ユニーク|キーの名前|インデックス内の順番|カラム名|照合順序|カーディナリティ|部分長|圧縮|コメント|
|:--:|:--|:--:|:--|:--:|:--|:--:|:--:|:--|
{{range $idx := .IndexSlice}}|{{$idx.NonUnique}}|{{$idx.IndexName}}|{{$idx.SeqInIndex}}|{{$idx.ColumnName}}|{{$idx.GetCollation}}|{{$idx.GetCardinality}}|{{$idx.GetSubPart}}|{{$idx.GetPacked}}|{{$idx.GetComment}}|
{{end}}
`

var csvTpl string = `{{range $column := .ColumnSlice}}{{$column.OrdinalPosition}},{{$column.ColumnName}},{{$column.ColumnType}},{{$column.IsNullable}},{{$column.GetColumnKey}},{{$column.GetExtra}},{{$column.OrdinalPosition|GetExsample}},{{$column.GetColumnComment}}
{{end}}
`

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/codegangsta/cli"

	"os"

	"errors"

	"github.com/satooon/mygene/db"
	"github.com/satooon/mygene/schema"
	"github.com/satooon/mygene/template"
)

func main() {
	app := cli.NewApp()
	app.Name = "mygene"
	app.Author = "satooon"
	app.Usage = "MySQL document generator"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "datasource, d",
			Value: "{user}:{password}@{tcp}({host}:{port})/{database_name}",
			Usage: "database connection setting",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "./",
			Usage: "output document path",
		},
		cli.StringFlag{
			Name:  "format, f",
			Value: template.MarkDown.String(),
			Usage: fmt.Sprintf("output file format. %v", template.OutPutFormats),
		},
		cli.BoolFlag{
			Name:  "verbose, vv",
			Usage: "verbose mode",
		},
	}
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) {
	const (
		TableQuery  = "SELECT * FROM information_schema.TABLES WHERE information_schema.TABLES.TABLE_SCHEMA = ?"
		ColumnQuery = "SELECT * FROM information_schema.COLUMNS WHERE information_schema.COLUMNS.TABLE_SCHEMA = ?"
		ValueQuery  = "SELECT * FROM %s LIMIT 1"
		IndexQuery  = "SELECT * FROM information_schema.STATISTICS WHERE information_schema.STATISTICS.TABLE_SCHEMA = ?"
	)

	verbose := c.Bool("verbose")
	datasource := c.String("datasource")
	output := c.String("output")
	fileFormat := template.OutPutFormat(c.String("format"))
	datasourceIdx := strings.Index(datasource, "/")
	if datasourceIdx < 0 {
		fmt.Println(errors.New("not found database"))
		os.Exit(1)
	}
	database := datasource[datasourceIdx+1:]

	if verbose {
		log.Printf("datasource:%v\n", datasource)
		log.Printf("database:%v\n", database)
		log.Printf("output:%v\n", output)
	}

	dbMap, err := db.InitDB(datasource)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if verbose {
		dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	}
	defer dbMap.Db.Close()

	var tableSlice schema.TableSlice
	if _, err := dbMap.Select(&tableSlice, TableQuery, database); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	tableMapByTableName := tableSlice.GroupByString(func(m *schema.Table) string { return m.TableName })

	var columnSlice schema.ColumnSlice
	if _, err := dbMap.Select(&columnSlice, ColumnQuery, database); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	columnMapByTableName := columnSlice.GroupByString(func(m *schema.Column) string { return m.TableName })

	var indexSlice schema.IndexSlice
	if _, err := dbMap.Select(&indexSlice, IndexQuery, database); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	indexMapByTableName := indexSlice.GroupByString(func(m *schema.Index) string { return m.TableName })

	schemaSlice := &schema.SchemaSlice{}
	for tbl, cs := range columnMapByTableName {
		colSlice := cs
		rows, err := dbMap.Db.Query(fmt.Sprintf(ValueQuery, tbl))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		rowColumns, err := rows.Columns()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		rowValues := make([]sql.RawBytes, len(rowColumns))
		scanArgs := make([]interface{}, len(rowValues))
		for i := range rowValues {
			scanArgs[i] = &rowValues[i]
		}
		valSlice := &schema.ValueSlice{}
		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			for _, col := range rowValues {
				if col == nil {
					*valSlice = append(*valSlice, schema.Value("NULL"))
				} else {
					*valSlice = append(*valSlice, schema.Value(string(col)))
				}
			}
		}

		idxSlice, ok := indexMapByTableName[tbl]
		if !ok {
			idxSlice = schema.IndexSlice{}
		}

		tblSlice := tableMapByTableName[tbl]

		*schemaSlice = append(*schemaSlice, schema.NewSchema(tbl, tblSlice[0], &colSlice, valSlice, &idxSlice))
	}

	if verbose {
		for _, v := range *schemaSlice {
			log.Println("----------------------------------------")
			log.Printf("tbl:%s\ncolumn:%v\nval:%v\nindex:%v\n", v.Name, v.ColumnSlice, v.ValueSlice, v.IndexSlice)
			log.Println("----------------------------------------")
		}
	}

	tpl := template.NewTemplate(fileFormat, output, schemaSlice)
	if err := tpl.Print(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	log.Println("Finish")
}

package duncan

import (
	"bytes"
	"html/template"
)

type tableData struct {
	Columns [][]string
	MaxLen  int
}

func Table(columns [][]string) Component {
	return func() (template.HTML, error) {
		tmpl, err := template.New("table").ParseFiles("templates/table.gohtml")
		if err != nil {
			return "", err
		}
		tmpl.Funcs(funcs)

		var maxLen int
		for i := range columns {
			if len(columns[i]) > maxLen {
				maxLen = len(columns[i])
			}
		}
		buf := bytes.Buffer{}
		err = tmpl.Execute(&buf, tableData{
			Columns: columns,
			MaxLen:  maxLen,
		})
		if err != nil {
			return "", err
		}
		return template.HTML(buf.String()), nil
	}

}

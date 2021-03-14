package duncan

import (
	"bytes"
	"html/template"
)

type pageData struct {
	Title    string
	Children []Component
}

type Component func() (template.HTML, error)

func Page(title string, children ...Component) Component {
	return func() (template.HTML, error) {
		tmpl, err := template.ParseFiles("templates/page.gohtml")
		if err != nil {
			return "", err
		}
		tmpl = tmpl.Funcs(funcs)
		buf := bytes.Buffer{}
		err = tmpl.Execute(&buf, pageData{
			Title:    title,
			Children: children,
		})
		if err != nil {
			return "", err
		}
		return template.HTML(buf.String()), nil
	}

}

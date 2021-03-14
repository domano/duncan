package duncan

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func TestPage(t *testing.T) {
	cases := []struct {
		Name       string
		Title      string
		Children   []Component
		ShouldFail bool
	}{
		{
			Name:     "Happy case, just a title",
			Title:    "Happy Title",
			Children: nil,
		},
		{
			Name:  "Happy case, title and a happy child",
			Title: "Happy Title",
			Children: []Component{func() (template.HTML, error) {
				return "Happy Child", nil
			}},
		},
		{
			Name:     "Happy case, title and a happy table",
			Title:    "Happy Title",
			Children: []Component{Table([][]string{{"Header 1", "entry"}, {"Header 2"}, {"Header 3", "e", "e"}})},
		},
		{
			Name:  "Title and an erroneous child",
			Title: "Happy Title",
			Children: []Component{func() (template.HTML, error) {
				return "", errors.New("Bad")
			}},
			ShouldFail: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name,
			func(t *testing.T) {
				page := Page(c.Title, c.Children...)

				result, err := page()
				if c.ShouldFail {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}

				println(result)
			})
	}
}

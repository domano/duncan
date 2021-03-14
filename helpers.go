package duncan

import "html/template"

var funcs = template.FuncMap{
	"numRange": func(end int) []int {
		n := end + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = i
		}
		return result
	},
}

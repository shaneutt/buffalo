package render

import (
	"html/template"
	"net/http"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
)

func init() {
	plush.Helpers.Add("paginator", func(pagination interface{}, opts map[string]interface{}, help plush.HelperContext) (template.HTML, error) {
		if opts["path"] == nil {
			if req, ok := help.Value("request").(*http.Request); ok {
				opts["path"] = req.URL.String()
			}
		}
		t, err := tags.Pagination(pagination, opts)
		if err != nil {
			return "", err
		}
		return t.HTML(), nil
	})
	plush.Helpers.Add("link", link)
}

func link(href string, opts tags.Options, help plush.HelperContext) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	opts["href"] = href
	if help.HasBlock() {
		delete(opts, "body")
		h, err := help.Block()
		if err != nil {
			return "", err
		}
		opts["body"] = h
	}
	a := tags.New("a", opts)
	return a.HTML(), nil
}

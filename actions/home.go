package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func TestQueryParamHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML(c.Param("name")))
}

func TestJsonHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"status": "polki"}))
}

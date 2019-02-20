package actions

import "github.com/gobuffalo/buffalo"

// LaylaShow default implementation.
func LaylaShow(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "You've been running, hiding much too long."}))
}

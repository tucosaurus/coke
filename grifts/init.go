package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tucosaurus/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}

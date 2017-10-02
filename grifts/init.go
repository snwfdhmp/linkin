package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/snwfdhmp/linkin/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}

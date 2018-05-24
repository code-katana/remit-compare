package grifts

import (
	"github.com/code-katana/remitcompare/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}

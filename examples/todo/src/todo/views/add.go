package views

import (
	"io"

	"github.com/lennon-guan/dada"
)

func Add(w io.Writer) {
	base := Base{
		Title: "New TODO",
		Body: func(c *dada.Ctx) {
			c.FORM(dada.Attr{"method": "POST", "action": ""}.Class("form form-horizontal"), func() {
				inputGroup(c, "text", "Title", "title", "")
				inputGroup(c, "text", "Content", "content", "")
				inputGroup(c, "date", "FinishAt", "finat", "")
				c.DIV(dada.Class("form-group"), func() {
					c.DIV(dada.Class("col-sm-10 col-sm-offset-2"), func() {
						c.BUTTON(dada.Attr{
							"class": "btn btn-primary",
						}, "SAVE")
					})
				})
			})
		},
	}
	base.Render(&dada.Ctx{Writer: w})
}

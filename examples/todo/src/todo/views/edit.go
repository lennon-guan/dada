package views

import (
	"io"
	"todo/store"

	"github.com/lennon-guan/dada"
)

func Edit(w io.Writer, rec *store.TodoRec) {
	base := Base{
		Title: "New TODO",
		Body: func(c *dada.Ctx) {
			c.FORM(dada.Attr{"method": "POST", "action": ""}.Class("form form-horizontal"), func() {
				staticGroup(c, "ID", rec.Id)
				inputGroup(c, "text", "Title", "title", rec.Title)
				inputGroup(c, "text", "Content", "content", rec.Content)
				inputGroup(c, "date", "FinishAt", "finat", rec.FinishAt.Format("2006-01-02"))
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
	base.Render(dada.NewCtx(w))
}

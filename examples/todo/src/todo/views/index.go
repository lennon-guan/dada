package views

import (
	"io"
	"strconv"
	"todo/store"

	"github.com/lennon-guan/dada"
)

func Index(w io.Writer, records []*store.TodoRec) {
	base := Base{
		Title: "Todo List",
		Body: func(c *dada.Ctx) {
			c.DIV(func() {
				c.A(dada.Attr{"href": "/add", "class": "btn btn-default"}, "New TODO")
			})
			c.TABLE(dada.Class("table"), func() {
				c.THEAD(func() {
					c.TR(func() {
						c.TH("ID")
						c.TH("TITLE")
						c.TH("")
					})
				})
				c.TBODY(func() {
					for _, rec := range records {
						c.TR(func() {
							c.TD(rec.Id)
							c.TD(rec.Title)
							c.TD(func() {
								c.A(dada.Attr{"href": "/edit?id=" + strconv.Itoa(rec.Id)}, "Edit")
							})
						})
					}
				})
			})
		},
	}
	base.Render(dada.NewCtx(w))
}

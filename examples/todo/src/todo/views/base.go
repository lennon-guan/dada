package views

import "github.com/lennon-guan/dada"

type Base struct {
	Body  func(*dada.Ctx)
	Title string
}

func (b *Base) Render(c *dada.Ctx) {
	c.HTML(func() {
		c.HEAD(func() {
			c.TITLE("TODO - " + b.Title)
			c.ImportCSS("https://cdn.bootcss.com/bootstrap/3.3.6/css/bootstrap.min.css")
		})
		c.BODY(func() {
			c.DIV(dada.Class("container"), func() {
				c.DIV(dada.Class("page-header"), func() {
					c.H1(b.Title)
				})
				if b.Body != nil {
					b.Body(c)
				}
			})
			c.ImportJS("https://cdn.bootcss.com/jquery/2.2.4/jquery.min.js")
			c.ImportJS("https://cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js")
		})
	})
}

package dada

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func renderSimple() string {
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	c := Ctx{Writer: buf}
	c.HTML(func() {
		c.HEAD(func() {
			c.TITLE("test html")
			c.ImportCSS("bootstrap.min.css")
		})
		c.BODY(func() {
			c.DIV(Class("page").Class("page-default"), func() {
				c.H1("test html")
			})
			c.UL(func() {
				for i := 0; i < 10; i++ {
					c.LI(fmt.Sprintf("Item %d", i+1))
				}
			})
			c.ImportJS("jquery.min.js")
			c.ImportJS("bootstrap.min.js")
		})
	})
	return buf.String()
}

func TestRenderSimple(t *testing.T) {
	fmt.Println(renderSimple())
}

func renderPageHead(c *Ctx) {
	c.DIV(Class("panel panel-default"), func() {
		c.DIV(Class("panel-heading"), nil)
	})
}

func TestInclude(t *testing.T) {
	c := &Ctx{Writer: os.Stdout}
	c.BODY(func() {
		renderPageHead(c)
		c.H1("hehehe")
	})
}

type basePage struct {
	Body  func(c *Ctx)
	Sect1 func(c *Ctx)
	Title string
}

func (mp *basePage) Render(c *Ctx) {
	c.HTML(func() {
		c.HEAD(func() {
			c.TITLE(mp.Title)
			c.ImportCSS("bootstrap.min.css")
		})
		c.BODY(func() {
			c.DIV(Class("page").Class("page-default"), func() {
				c.H1(mp.Title)
			})
			if mp.Body != nil {
				mp.Body(c)
			}
			c.ImportJS("jquery.min.js")
			c.ImportJS("bootstrap.min.js")
		})
	})

}

func TestMasterPage(t *testing.T) {
	bp := basePage{
		Body: func(c *Ctx) {
			c.UL(func() {
				for i := 0; i < 10; i++ {
					c.LI(fmt.Sprintf("Item %d", i+1))
				}
			})
		},
		Title: "test html",
	}
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	c := Ctx{Writer: buf}
	bp.Render(&c)
	fmt.Println(buf.String())
}

func BenchmarkRenderSimple(b *testing.B) {
	renderSimple()
}

func BenchmarkRenderGohtml(b *testing.B) {
}

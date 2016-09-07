package views

import "github.com/lennon-guan/dada"

func inputGroup(c *dada.Ctx, type_, label, name, initValue string) {
	c.DIV(dada.Class("form-group"), func() {
		c.LABEL(dada.Attr{"for": name}.Class("col-sm-2 control-label"), label)
		c.DIV(dada.Class("col-sm-5"), func() {
			c.INPUT(dada.Attr{
				"type":  type_,
				"name":  name,
				"id":    name,
				"value": initValue,
				"class": "form-control",
			})
		})
	})
}

func staticGroup(c *dada.Ctx, label, value interface{}) {
	c.DIV(dada.Class("form-group"), func() {
		c.LABEL(dada.Class("col-sm-2 control-label"), label)
		c.DIV(dada.Class("col-sm-5"), func() {
			c.P(dada.Attr{
				"class": "form-control-static",
			}, value)
		})
	})
}

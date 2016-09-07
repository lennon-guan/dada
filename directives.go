package dada

func (c *Ctx) ImportCSS(path string) {
	c.LINK(Attr{
		"rel":  "stylesheet",
		"type": "text/css",
		"href": path,
	})
}

func (c *Ctx) ImportJS(path string) {
	c.SCRIPT(Attr{
		"src": path,
	}, nil)
}

package dada

import (
	"fmt"
	"io"
)

type Ctx struct {
	indent string
	Writer io.Writer
}

func (c *Ctx) Container(name string, attr Attr, children interface{}) {
	if _, err := fmt.Fprint(c.Writer, c.indent+"<"+name); err != nil {
		panic(err)
	}
	if err := attr.Write(c.Writer); err != nil {
		panic(err)
	}
	if children != nil {
		if fn, ok := children.(func()); ok {
			if _, err := fmt.Fprintln(c.Writer, ">"); err != nil {
				panic(err)
			}
			old := c.indent
			c.indent += "  "
			fn()
			c.indent = old
			if _, err := fmt.Fprintln(c.Writer, c.indent+"</"+name+">"); err != nil {
				panic(err)
			}
		} else {
			if _, err := fmt.Fprintf(c.Writer, ">%v</%s>\n", children, name); err != nil {
				panic(err)
			}
		}
	} else {
		if _, err := fmt.Fprintln(c.Writer, "></"+name+">"); err != nil {
			panic(err)
		}
	}
}

func (c *Ctx) Element(name string, attr Attr) {
	if _, err := fmt.Fprintf(c.Writer, "%s<%s", c.indent, name); err != nil {
		panic(err)
	}
	if err := attr.Write(c.Writer); err != nil {
		panic(err)
	}
	if _, err := fmt.Fprintln(c.Writer, "/>"); err != nil {
		panic(err)
	}
}

func (c *Ctx) Text(content string) func() {
	return func() {
		if _, err := fmt.Fprintln(c.Writer, c.indent+content); err != nil {
			panic(err)
		}
	}
}

func (c *Ctx) getArgs(args []interface{}) (Attr, interface{}) {
	if args == nil || len(args) == 0 {
		return nil, nil
	} else if len(args) == 1 {
		if attr, ok := args[0].(Attr); ok {
			return attr, nil
		} else {
			return nil, args[0]
		}
	} else if len(args) == 2 {
		return args[0].(Attr), args[1]
	} else {
		panic("args length error")
	}
}

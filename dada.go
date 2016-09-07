package dada

import (
	"fmt"
	"html"
	"io"
)

type Safe struct {
	Val interface{}
}

type CtxCfg struct {
	IndentStr string
	Safe      bool
}

type Ctx struct {
	indent string
	Writer io.Writer
	Config CtxCfg
}

func NewCtxWithConfig(writer io.Writer, conf *CtxCfg) *Ctx {
	return &Ctx{
		Writer: writer,
		Config: *conf,
	}
}

func NewCtx(writer io.Writer) *Ctx {
	return NewCtxWithConfig(writer, &defaultCfg)
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
			c.indent += c.Config.IndentStr
			fn()
			c.indent = old
			if _, err := fmt.Fprintln(c.Writer, c.indent+"</"+name+">"); err != nil {
				panic(err)
			}
		} else if s, ok := children.(Safe); ok {
			if _, err := fmt.Fprintf(c.Writer, ">%v</%s>\n", s.Val, name); err != nil {
				panic(err)
			}
		} else if !c.Config.Safe {
			raw := fmt.Sprintf("%v", children)
			if _, err := fmt.Fprintf(c.Writer, ">%s</%s>\n", html.EscapeString(raw), name); err != nil {
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

func (c *Ctx) Text(content interface{}) {
	if s, ok := content.(Safe); ok {
		if _, err := fmt.Fprint(c.Writer, s.Val); err != nil {
			panic(err)
		}
	} else if !c.Config.Safe {
		raw := fmt.Sprintf("%v", content)
		if _, err := fmt.Fprintf(c.Writer, html.EscapeString(raw)); err != nil {
			panic(err)
		}
	} else {
		if _, err := fmt.Fprint(c.Writer, content); err != nil {
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

var defaultCfg CtxCfg

func init() {
	defaultCfg.IndentStr = "  "
	defaultCfg.Safe = false
}

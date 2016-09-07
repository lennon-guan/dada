package dada

import (
	"fmt"
	"io"
)

type Attr map[string]string

func Class(classname string) Attr {
	return Attr{"class": classname}
}

func (attr Attr) Class(classname string) Attr {
	if old, exists := attr["class"]; exists {
		attr["class"] = old + " " + classname
	} else {
		attr["class"] = classname
	}
	return attr
}

func (attr Attr) Href(value string) Attr {
	attr["href"] = value
	return attr
}

func (attr *Attr) Write(writer io.Writer) error {
	if attr == nil {
		return nil
	}
	for key, value := range *attr {
		if _, err := fmt.Fprintf(writer, " %s=\"%s\"", key, value); err != nil {
			return err
		}
	}
	return nil
}

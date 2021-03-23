package gen

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Generator struct {
	resolver *Resolver
}

type WriteLine func(s string, p ...interface{})

func (g *Generator) generate() error {
	if err := g.generateModel(); err != nil {
		return err
	}
	if err := g.generateActions(); err != nil {
		return err
	}
	return nil
}

func goTypeName(t *GoType) string {
	if t == nil {
		return "int"
	}
	if t.IsObject() && strings.Index(t.GoType, "map[") == -1 && strings.Index(t.GoType, "*") == -1 {
		return "*" + t.GoType
	}
	if t.GoType == "" {
		t.GoType = "Any"
	}
	return t.GoType
}

func goTypeToString(t *GoType, varName string) string {
	switch t.GoName {
	case "bool":
		return fmt.Sprintf("strconv.FormatBool(%s)", varName)
	case "float64":
		return fmt.Sprintf("strconv.FormatFloat(%s)", varName)
	case "int", "int64":
		return fmt.Sprintf("strconv.FormatInt(int64(%s), 10)", varName)
	case "string":
		return varName
	}
	return ""
}

func (g *Generator) generateActions() error {
	b := &strings.Builder{}
	r := g.resolver

	W := func(s string, p ...interface{}) {
		if len(p) > 0 {
			_, _ = b.WriteString(fmt.Sprintf(s, p...))
		} else {
			_, _ = b.WriteString(s)
		}
		_, _ = b.WriteString("\n")
	}

	W("package rest\n")
	W("import (")
	W("    . \"github.com/kamaiu/ib-cp-go/client/v1/rest/model\"")
	W("    \"github.com/valyala/bytebufferpool\"")
	W("    \"net/url\"")
	W("    \"strconv\"")
	W(")\n")

	/*
		func (c *Client) Logout() (_200 *model.Logout_POST_200, err error) {
			url := bytebufferpool.Get()
			_, _ = url.WriteString(c.host)
			_, _ = url.WriteString("/logout")

			err = c.Do(fasthttp.MethodPost, url, func(status int, body []byte, err error) error {
				switch status {
				case 200:
					_200 = &model.Logout_POST_200{}
					if body != nil {
						return _200.UnmarshalJSON(body)
					} else {
						return err
					}
				}
				return err
			})
			return
		}
	*/

	for _, a := range r.Actions {
		writeComments(a.Action.Description, func(line string) {
			W("// %s", line)
		})
		path, query := a.Params()
		responses := a.ResponseList()
		responsePart := (func() string {
			sb := &strings.Builder{}
			_, _ = sb.WriteString("(\n")
			if len(responses) > 0 {
				count := 0
				for _, resp := range responses {
					if resp.Type == nil || resp.Type.Primitive {
						continue
					}
					//if count > 0 {
					//	_, _ = sb.WriteString(", ")
					//}
					count++
					goName := resp.Type.GoName
					if resp.Type.IsAny() {
						goName = "Any"
					} else if resp.Type.IsObject() {
						goName = "*" + goName
					}
					_, _ = sb.WriteString(fmt.Sprintf("    _%d %s,\n", resp.Code, goName))
				}
				if count == 0 {
					_, _ = sb.WriteString("    err error,\n")
				} else {
					_, _ = sb.WriteString("    err error,\n")
				}
			} else {
				_, _ = sb.WriteString("err error,\n")
			}
			_, _ = sb.WriteString(")")
			return sb.String()
		})()

		if len(path) > 0 || len(query) > 0 {
			W("func (c *Client) %s(", a.Name)
			if len(path) > 0 {
				for _, p := range path {
					writeComments(p.Param.Parameter.Description, func(line string) {
						W("    // %s", line)
					})
					W("    %s %s,", p.Param.Name, goTypeName(p.Param.Type))
				}
			}
			if len(query) > 0 {
				for _, q := range query {
					writeComments(q.Parameter.Description, func(line string) {
						W("    // %s", line)
					})
					W("    %s %s,", q.Name, goTypeName(q.Type))
				}
			}
			W(") %s {", responsePart)

			W("    _u := bytebufferpool.Get()")
			W("    _, _ = _u.WriteString(c.host)")
			for _, p := range a.Url {
				if len(p.Literal) > 0 {
					W("    _, _ = _u.WriteString(\"%s\")", p.Literal)
				} else if p.Param != nil {
					W("    _, _ = _u.WriteString(url.PathEscape(%s))", goTypeToString(p.Param.Type, p.Param.Name))
				}
			}

			if len(query) > 0 {
				for i, q := range query {
					if i > 0 {
						W("    _, _ = _u.WriteString(\"&%s=\")", q.Name)
					} else {
						W("    _, _ = _u.WriteString(\"?%s=\")", q.Name)
					}
					W("    _, _ = _u.WriteString(url.PathEscape(%s))", goTypeToString(q.Type, q.Name))
				}
			}
		} else {
			W("func (c *Client) %s() %s {", a.Name, responsePart)

			W("    _u := bytebufferpool.Get()")
			W("    _, _ = _u.WriteString(c.host)")
			for _, p := range a.Url {
				if len(p.Literal) > 0 {
					W("    _, _ = _u.WriteString(\"%s\")", p.Literal)
				}
			}
		}

		W("    err = c.Do(\"%s\", _u, func(status int, body []byte, err error) error {", strings.ToUpper(a.Verb))
		W("        if err != nil {")
		W("            return err")
		W("        }")
		W("        switch status {")
		for _, resp := range responses {
			if resp.Type == nil {
				continue
			}
			W("        case %d:", resp.Code)
			if resp.Type.Primitive {
				if resp.Code == 200 {
					W("            return nil")
				}
			} else if resp.Type.IsAny() {
				W("            _%d = Any{}", resp.Code)
				W("            if body != nil {")
				W("                return _%d.UnmarshalJSON(body)", resp.Code)
				W("            } else {")
				W("                return err")
				W("            }")
			} else if resp.Type.IsObject() {
				W("            _%d = &%s{}", resp.Code, resp.Type.GoName)
				W("            if body != nil {")
				W("                return _%d.UnmarshalJSON(body)", resp.Code)
				W("            } else {")
				W("                return err")
				W("            }")
			} else if resp.Type.Array {
				W("            _%d = %s{}", resp.Code, resp.Type.GoName)
				W("            if body != nil {")
				W("                return _%d.UnmarshalJSON(body)", resp.Code)
				W("            } else {")
				W("                return err")
				W("            }")
			} else {
				W("     HELP!!!!!")
			}
		}
		W("        default:")
		W("            return StatusCodeError{Code: status}")
		W("        }")
		W("    })")
		W("    return")

		W("}\n")
	}
	ioutil.WriteFile("../actions.go", []byte(b.String()), 0755)
	return nil
}

func (g *Generator) generateModel() error {
	b := &strings.Builder{}
	r := g.resolver

	W := func(s string, p ...interface{}) {
		if len(p) > 0 {
			_, _ = b.WriteString(fmt.Sprintf(s, p...))
		} else {
			_, _ = b.WriteString(s)
		}
		_, _ = b.WriteString("\n")
	}

	W("//go:generate easyjson -all $GOFILE")
	W("package model\n")

	W("//easyjson:json")
	W("type Any map[string]interface{}")
	W("")

	// Write Arrays
	for _, a := range r.Arrays {
		itemType := a.Items
		itemSlice := ""
		if a.Items.IsObject() && a.Items.GoType[0] != 'm' {
			itemSlice = "[]*" + itemType.GoType
		} else {
			if len(itemType.GoType) == 0 {
				itemType.GoType = "map[string]interface{}"
			}
			itemSlice = "[]" + itemType.GoType
		}
		writeComments(a.Schema.Description, func(line string) {
			W("// %s", line)
		})
		W("//easyjson:json")
		W("type %s %s", a.GoName, itemSlice)
		W("")
	}

	// Write Objects
	for _, o := range r.Objects {
		if len(o.GoName) == 0 {
			fmt.Println()
		}
		if len(o.Schema.Description) > 0 {
			writeComments(o.Schema.Description, func(line string) {
				W("// %s", line)
			})
		}
		W("type %s struct {", o.GoName)
		for _, p := range o.Props {
			if p.Type.GoType == "" {
				p.Type.GoType = "map[string]interface{}"
			}
			typeName := p.Type.GoType
			if p.Type.IsObject() && strings.Index(typeName, "map[") == -1 {
				typeName = "*" + typeName
			}
			writeComments(p.Type.Schema.Description, func(line string) {
				W("    // %s", line)
			})
			W("    %s     %s `json:\"%s\"`", p.FieldName, typeName, p.Name)
		}
		W("}")
		W("")
	}

	ioutil.WriteFile("../model/model.go", []byte(b.String()), 0755)
	return nil
}

func writeComments(comments string, fn func(line string)) {
	if len(comments) == 0 {
		return
	}
	lines := wordWrap(comments, 80)
	for _, line := range lines {
		if fn != nil {
			fn(line)
		}
	}
}

func wordWrap(line string, maxLine int) []string {
	var lines []string
	c := 0
	mark := 0
	for i := 0; i < len(line); i++ {
		if c >= maxLine || line[i] == '\n' {
			switch line[i] {
			case ' ', '\t', '\n':
				lines = append(lines, strings.TrimSpace(line[mark:i]))
				mark = i + 1
				c = 0
			default:
				c++
			}
		} else {
			c++
		}
	}

	end := strings.TrimSpace(line[mark:])
	if len(end) > 0 {
		lines = append(lines, end)
	}
	return lines
}

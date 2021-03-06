package gen

import (
	"sort"
	"strings"
)

type GoAction struct {
	Path      string
	Verb      string
	Name      string
	Action    *Action
	Body      *GoType
	Url       []*PathPart
	Query     []*GoParam
	Responses map[int]*GoActionResponse
}

func (a *GoAction) ResponseList() []*GoActionResponse {
	if len(a.Responses) == 0 {
		return nil
	}
	responses := make([]*GoActionResponse, 0, len(a.Responses))
	responsesCodes := make([]int, 0, len(a.Responses))
	for _, resp := range a.Responses {
		responsesCodes = append(responsesCodes, resp.Code)
	}
	sort.Ints(responsesCodes)
	for _, code := range responsesCodes {
		responses = append(responses, a.Responses[code])
	}
	return responses
}

func (a *GoAction) Params() (path []*PathPart, query []*GoParam) {
	for _, p := range a.Url {
		if p.Param != nil {
			path = append(path, p)
		}
	}
	query = a.Query
	return
}

type GoActionResponse struct {
	Code     int
	Type     *GoType
	Response *Response
}

type GoType struct {
	Name      string
	GoName    string
	GoType    string
	Primitive bool
	Array     bool
	Schema    *Schema
	Items     *GoType
	Props     []*Prop
}

func (g *GoType) IsObject() bool {
	return len(g.Props) > 0 && !g.Primitive && !g.Array
}

func (g *GoType) IsAny() bool {
	return strings.Index(g.GoType, "map[string]interface{}") == 0 || len(g.GoType) == 0
	//return strings.Index(g.GoType, "map[") == 0 || strings.Index(g.GoName, "map[") == 0 || len(g.GoType) == 0
}

type Prop struct {
	Name      string
	FieldName string
	Type      *GoType
	Schema    *Schema
}

type GoParam struct {
	Name      string
	GoName    string
	Type      *GoType
	Parameter *Parameter
}

type PathPart struct {
	Literal string
	Param   *GoParam
}

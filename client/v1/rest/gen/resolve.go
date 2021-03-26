package gen

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"sort"
	"strconv"
	"strings"
)

type Resolver struct {
	Spec    *Spec
	Arrays  map[string]*GoType
	Objects map[string]*GoType
	Refs    map[string]*GoType
	Defs    map[string]*Schema
	Actions map[string]*GoAction
}

func NewResolver(spec *Spec) (*Resolver, error) {
	resolver := &Resolver{
		Spec:    spec,
		Arrays:  make(map[string]*GoType),
		Objects: make(map[string]*GoType),
		Refs:    make(map[string]*GoType),
		Defs:    make(map[string]*Schema),
		Actions: make(map[string]*GoAction),
	}
	var err error
	for k, v := range spec.Definitions {
		v.Name = k
	}
	for k, v := range spec.Definitions {
		resolver.Defs["#/definitions/"+k] = v
	}

	for k, v := range spec.Definitions {
		t, err := resolver.typeOf(k, v)
		if err != nil {
			return nil, err
		}
		resolver.Refs["#/definitions/"+k] = t
	}

	addAction := func(path, verb string, v *Action) error {
		name := toGoName(path, verb)

		action := &GoAction{
			Path:      path,
			Verb:      verb,
			Name:      name,
			Action:    v,
			Query:     make([]*GoParam, 0, len(v.Parameters)),
			Responses: make(map[int]*GoActionResponse),
		}
		if resolver.Actions[name] != nil {
			return errors.New("action already exists: " + name)
		}

		pathParams := make(map[string]*Parameter)
		for _, p := range v.Parameters {
			in := strings.ToLower(strings.TrimSpace(p.In))
			switch in {
			case "path":
				pathParams[p.Name] = p

			case "query":
				param, err := pathOrQueryParamOf(p)
				if err != nil {
					return err
				}
				action.Query = append(action.Query, param)

			case "body":
				action.Body, err = resolver.typeOf(action.Name+"_Request", p.Schema)
				if err != nil {
					return err
				}
				if action.Body == nil {
					return errors.New("could not resolve body for: " + path)
				}
			}
		}

		action.Url, err = parsePath(path, pathParams)
		if err != nil {
			return err
		}

		for code, resp := range v.Responses {
			c, err := strconv.Atoi(code)
			if err != nil {
				return err
			}
			if action.Path == "/iserver/scanner/params" {
				fmt.Println()
			}
			t, err := resolver.typeOf(action.Name+"_"+code, resp.Schema)
			if err != nil {
				return err
			}
			action.Responses[c] = &GoActionResponse{
				Code:     c,
				Type:     t,
				Response: resp,
			}
		}
		resolver.Actions[name] = action
		return nil
	}
	for k, v := range spec.Paths {
		if v.Get != nil {
			if err := addAction(k, fasthttp.MethodGet, v.Get); err != nil {
				return nil, err
			}
		}
		if v.Post != nil {
			if err := addAction(k, fasthttp.MethodPost, v.Post); err != nil {
				return nil, err
			}
		}
		if v.Put != nil {
			if err := addAction(k, fasthttp.MethodPut, v.Put); err != nil {
				return nil, err
			}
		}
		if v.Delete != nil {
			if err := addAction(k, fasthttp.MethodDelete, v.Delete); err != nil {
				return nil, err
			}
		}
		if v.Patch != nil {
			if err := addAction(k, fasthttp.MethodPatch, v.Patch); err != nil {
				return nil, err
			}
		}
		if v.Trace != nil {
			if err := addAction(k, fasthttp.MethodTrace, v.Trace); err != nil {
				return nil, err
			}
		}
		if v.Head != nil {
			if err := addAction(k, fasthttp.MethodHead, v.Head); err != nil {
				return nil, err
			}
		}
		if v.Options != nil {
			if err := addAction(k, fasthttp.MethodOptions, v.Options); err != nil {
				return nil, err
			}
		}
	}
	return resolver, nil
}

func (resolver *Resolver) typeOf(name string, schema *Schema) (*GoType, error) {
	if schema == nil {
		return nil, nil
	}
	if len(schema.Ref) > 0 {
		return resolver.typeOfRef(schema.Ref)
	}

	if len(schema.Type) == 0 {
		if schema.Items != nil {
			schema.Type = TypeArray
		} else if len(schema.Properties) > 0 {
			schema.Type = TypeObject
		}
	}

	switch schema.Type {
	case TypeArray:
		n := toGoName(name+"_List", "")
		itemType, err := resolver.typeOf(n+"_Item", schema.Items)
		if err != nil {
			return nil, err
		}
		if itemType == nil {
			return nil, errors.New("could not resolve item type: " + (name + "Item"))
		}
		goType := ""
		if itemType.Primitive {
			goType = "[]" + itemType.GoType
		} else if itemType.Array {
			goType = "[]" + itemType.GoType
		} else {
			if len(itemType.GoType) == 0 {
				goType = "[]map[string]interface{}"
			} else {
				goType = "[]*" + itemType.GoType
			}
		}
		goType = n
		t := &GoType{
			Name:   name,
			GoName: n,
			Schema: schema,
			GoType: goType,
			Array:  true,
			Items:  itemType,
		}
		resolver.Arrays[n] = t
		return t, nil

	case TypeObject:
		n := toGoName(name, "")
		if len(schema.Properties) == 0 {
			return &GoType{
				Name:   name,
				GoName: n,
				Schema: schema,
			}, nil
		}

		propList := make([]*Prop, 0, len(schema.Properties))
		props := make(map[string]*Prop, len(schema.Properties))
		propSortToName := make(map[string]string)
		propNames := make([]string, 0, len(schema.Properties))
		for key, p := range schema.Properties {
			p.Name = key
			var propType *GoType
			var err error

			fieldName := numToFieldName[key]
			if len(fieldName) == 0 {
				fieldName = toGoFieldName(key)
			}
			if len(p.Ref) > 0 {
				propType, err = resolver.typeOfRef(p.Ref)
			} else {
				propType, err = resolver.typeOf(n+"_"+fieldName, p)
			}
			if err != nil {
				return nil, err
			}
			if propType == nil {
				return nil, errors.New("object prop: " + n + "." + key + " could not be resolved")
			}
			sortKey := key
			if len(sortKey) > 0 {
				switch sortKey[0] {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					prepend := ""
					for i := 0; i < 10-len(sortKey); i++ {
						prepend += "0"
					}
					sortKey = prepend + sortKey
				}
			}
			propNames = append(propNames, sortKey)
			propSortToName[sortKey] = key
			props[key] = &Prop{
				Name:      key,
				FieldName: fieldName,
				Type:      propType,
				Schema:    p,
			}
		}

		sort.Strings(propNames)
		for _, name := range propNames {
			propList = append(propList, props[propSortToName[name]])
		}

		if len(name) == 0 {
			name = "map[string]interface{}"
		}

		goType := n
		if len(n) == 0 {
			n = name
		}
		if len(n) == 0 {
			n = "map[string]interface{}"
		}
		if len(props) == 0 {
			n = "map[string]interface{}"
		}
		if len(propList) == 1 && len(propList[0].Schema.Ref) > 0 {
			if propList[0].Type.Array || propList[0].Type.Primitive || propList[0].Type.IsAny() {
				goType = "map[string]" + propList[0].Type.GoType
			} else if propList[0].Type.IsObject() {
				goType = "map[string]*" + propList[0].Type.GoType
			} else {
				goType = "map[string]interface{}"
			}
		}

		t := &GoType{
			Name:   name,
			GoName: n,
			GoType: goType,
			Schema: schema,
			Props:  propList,
		}
		resolver.Objects[n] = t
		return t, nil

	case TypeBool:
		return &GoType{
			Name:      "boolean",
			GoName:    "bool",
			GoType:    "bool",
			Primitive: true,
			Schema:    schema,
		}, nil

	case TypeString:
		return &GoType{
			Name:      "string",
			GoName:    "string",
			GoType:    "string",
			Primitive: true,
			Schema:    schema,
		}, nil

	case TypeNumber:
		return &GoType{
			Name:      "number",
			GoName:    "float64",
			GoType:    "float64",
			Primitive: true,
			Schema:    schema,
		}, nil

	case TypeInt:
		return &GoType{
			Name:      "integer",
			GoName:    "int64",
			GoType:    "int64",
			Primitive: true,
			Schema:    schema,
		}, nil
	}
	return nil, errors.New("not resolved")
}

func (resolver *Resolver) typeOfRef(ref string) (*GoType, error) {
	if len(ref) == 0 {
		return nil, nil
	}
	refType := resolver.Refs[ref]
	if refType == nil {
		def := resolver.Defs[ref]
		if def == nil {
			return nil, nil
		}
		var err error
		refType, err = resolver.typeOf(def.Name, def)
		if err != nil {
			return nil, err
		}
	}
	return refType, nil
}

func toGoName(path, verb string) string {
	parts := strings.Split(path, "/")
	b := &strings.Builder{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) == 0 {
			continue
		}
		if b.Len() > 0 {
			_, _ = b.WriteString("_")
		}
		if part[0] == '{' || part[0] == ':' {
			part = part[1:]

			if len(part) > 0 && part[len(part)-1] == '}' {
				part = part[0 : len(part)-1]
			}
		}
		hypens := strings.Split(part, "-")
		for _, hyphen := range hypens {
			_, _ = b.WriteString(capitalize(hyphen))
		}
	}
	if len(verb) > 0 {
		_, _ = b.WriteString("_" + verb)
	}
	return b.String()
}

func primitiveTypeOf(t Type) *GoType {
	switch t {
	case TypeBool:
		return &GoType{
			Name:      "boolean",
			GoName:    "bool",
			GoType:    "bool",
			Primitive: true,
		}

	case TypeString:
		return &GoType{
			Name:      "string",
			GoName:    "string",
			GoType:    "string",
			Primitive: true,
		}

	case TypeNumber:
		return &GoType{
			Name:      "number",
			GoName:    "float64",
			GoType:    "float64",
			Primitive: true,
		}

	case TypeInt:
		return &GoType{
			Name:      "integer",
			GoName:    "int64",
			GoType:    "int64",
			Primitive: true,
		}
	default:
		return nil
	}
}

var numToFieldName = map[string]string{
	"31":     "LastPrice",
	"55":     "Symbol",
	"58":     "Text",
	"70":     "High",
	"71":     "Low",
	"72":     "Position",
	"73":     "MarketValue",
	"74":     "AveragePrice",
	"75":     "UnrealizedPnL",
	"76":     "FormattedPosition",
	"77":     "FormattedUnrealizedPnL",
	"78":     "DailyPnL",
	"82":     "ChangePrice",
	"83":     "ChangePercent",
	"84":     "BidPrice",
	"85":     "AskSize",
	"86":     "AskPrice",
	"87":     "Volume",
	"87_raw": "VolumeRaw",
	"88":     "BidSize",
	"6004":   "Exchange",
	"6008":   "Conid6008",
	"6070":   "SecurityType",
	"6072":   "Months",
	"6073":   "RegularExpiry",
	"6119":   "Marker",
	"6457":   "UnderlyingConid",
	"6509":   "MarketDataAvailability",
	"7051":   "CompanyName",
	"7059":   "LastSize",
	"7094":   "ConidAndExchange",
	"7219":   "ContractDescription",
	"7220":   "ContractDescription2",
	"7221":   "ListingExchange",
	"7280":   "Industry",
	"7281":   "Category",
	"7282":   "AverageDailyVolume",
	"7284":   "HistoricVolume30d",
	"7285":   "PutCallRatio",
	"7286":   "DividendAmount",
	"7287":   "DividendYieldPercent",
	"7288":   "ExDateOfTheDividend",
	"7289":   "MarketCap",
	"7290":   "PE",
	"7291":   "EPS",
	"7292":   "CostBasis",
	"7293":   "High52Week",
	"7294":   "Low52Week",
	"7295":   "OpenPrice",
	"7296":   "ClosePrice",
	"7308":   "Delta",
	"7309":   "Gamma",
	"7310":   "Theta",
	"7311":   "Vega",
	"7633":   "ImpliedVolatility",
}

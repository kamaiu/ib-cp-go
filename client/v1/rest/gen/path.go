package gen

import "errors"

func pathOrQueryParamOf(param *Parameter) (*GoParam, error) {
	paramType := primitiveTypeOf(param.Type)
	if paramType == nil {
		return nil, errors.New("parameter '" + param.Name + "' is not a primitive type")
	}

	return &GoParam{
		Name:      param.Name,
		GoName:    param.Name,
		Type:      paramType,
		Parameter: param,
	}, nil
}

func urlParts(path string, params map[string]*Parameter) ([]*UrlPart, error) {
	parts := make([]*UrlPart, 0, 4)
	type stateType int
	const (
		stateLiteral stateType = 0
		stateParam   stateType = 1
	)
	state := stateLiteral
	mark := 0
	i := 0

	resolveParam := func(paramName string) error {
		param := params[paramName]
		if param == nil {
			return errors.New("could not resolve param name '" + paramName + "' in path expression: " + path)
		}
		if param.In != "path" {
			return errors.New("parameter '" + paramName + "' is not a path param in path expression: " + path)
		}
		p, err := pathOrQueryParamOf(param)
		if err != nil {
			return errors.New("path expression error " + err.Error() + " for path: " + path)
		}

		parts = append(parts, &UrlPart{
			Param: p,
		})
		return nil
	}

	addLiteral := func(literal string) {
		if len(literal) == 0 {
			return
		}
		parts = append(parts, &UrlPart{
			Literal: literal,
			Param:   nil,
		})
	}

	for ; i < len(path); i++ {
		c := path[i]
		switch state {
		case stateLiteral:
			switch c {
			case '{', ':':
				addLiteral(path[mark:i])
				state = stateParam
				mark = i + 1
			}

		case stateParam:
			switch c {
			case '}', ':', '/':
				paramName := path[mark:i]
				if len(paramName) > 0 {
					if err := resolveParam(paramName); err != nil {
						return nil, err
					}
				}

				switch c {
				case '}':
					mark = i + 1
					state = stateLiteral

				case '/':
					mark = i
					state = stateLiteral

				case ':':
					mark = i + 1
					state = stateParam
				}
			}
		}
	}

	switch state {
	case stateLiteral:
		addLiteral(path[mark:])
	case stateParam:
		if err := resolveParam(path[mark:]); err != nil {
			return nil, err
		}
	}
	return parts, nil
}

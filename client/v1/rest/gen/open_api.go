package gen

type Type string

const (
	TypeObject Type = "object"
	TypeArray  Type = "array"
	TypeString Type = "string"
	TypeBool   Type = "boolean"
	TypeNumber Type = "number"
	TypeInt    Type = "integer"
)

type Spec struct {
	BasePath string   `json:"basePath"`
	Host     string   `json:"host"`
	Produces []string `json:"produces"`
	Name     string   `json:"name"`
	Schemes  []string `json:"schemes"`
	Info     struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Version     string `json:"version"`
	} `json:"info"`
	Paths       map[string]*Path   `json:"paths"`
	Definitions map[string]*Schema `json:"definitions"`
}

type Path struct {
	Get     *Action `json:"get"`
	Post    *Action `json:"post"`
	Put     *Action `json:"put"`
	Patch   *Action `json:"patch"`
	Delete  *Action `json:"delete"`
	Trace   *Action `json:"trace"`
	Head    *Action `json:"head"`
	Options *Action `json:"options"`
	Connect *Action `json:"connect"`
}

type Action struct {
	Summary     string               `json:"summary"`
	Description string               `json:"description"`
	Tags        []string             `json:"tags"`
	Parameters  []*Parameter         `json:"parameters"`
	Responses   map[string]*Response `json:"responses"`
}

type Parameter struct {
	In          string  `json:"in"`
	Name        string  `json:"name"`
	Required    bool    `json:"required"`
	Description string  `json:"description"`
	Type        Type    `json:"type"`
	Schema      *Schema `json:"schema"`
}

type Response struct {
	Description string  `json:"description"`
	Schema      *Schema `json:"schema"`
}

type Schema struct {
	Name        string
	Ref         string             `json:"$ref"`
	Type        Type               `json:"type"`
	Description string             `json:"description"`
	Example     interface{}        `json:"example"`
	Properties  map[string]*Schema `json:"properties"`
	Items       *Schema            `json:"items"`
}

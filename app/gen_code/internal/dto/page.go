package dto

type Page struct {
	Path string
	Name string
}

type PageInfo struct {
	RootID      string
	Path        string
	Name        string
	Nodes       []*Node
	LinkedNodes []string
}

type ComponentInfo struct {
	RootID         string
	Name           string
	ComponentProps []string
	CompNodes      map[string]*CompNode
}

type Node struct {
	ID                        string
	Name                      string
	ComponentType             string // pre comp
	BelongToUserComponentType string // or user-comp
	CorrespondingProp         string // use for child of user component
	Props                     string
	Children                  []string
}

type CompNode struct {
	ID            string
	Name          string
	ComponentType string // pre comp
	Children      []string
}

type ReactElement struct {
	ID                        string
	Props                     string
	Component                 string
	ElementString             string
	Children                  []string
	BelongToUserComponentType string // or user-comp
}

type ImportantProps struct {
	Text string `json:"text"`
}

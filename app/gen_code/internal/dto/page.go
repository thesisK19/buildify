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

type Node struct {
	ID            string
	ComponentType string
	Props         string
	Children      []string
	Page          string
}

type ReactElement struct {
	ID            string
	Props         string
	Component     string
	ElementString string
	Children      []string
}

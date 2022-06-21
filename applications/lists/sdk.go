package lists

import "github.com/steve-care-software/query/domain/lists"

// Application represents a search application
type Application interface {
	Compile(script string) (lists.List, error)
	Execute(data [][]byte, list lists.List) ([]byte, uint, error)
}

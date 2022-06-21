package executions

import (
	"github.com/steve-care-software/query/domain/lists"
	"github.com/steve-care-software/query/domain/retrieves"
	"github.com/steve-care-software/query/domain/searches"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithData(data [][]byte) Builder
	Now() (Application, error)
}

// Application represents an execution application
type Application interface {
	Retrieve(data [][]byte, retrieve retrieves.Retrieve) ([]byte, error)
	Search(data [][]byte, search searches.Search) ([][]byte, error)
	List(data [][]byte, list lists.List) ([]byte, uint, error)
}

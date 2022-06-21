package applications

import (
	"github.com/steve-care-software/query/applications/lists"
	"github.com/steve-care-software/query/applications/retrieves"
	"github.com/steve-care-software/query/applications/searches"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithData(data [][]byte) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Retrieve() retrieves.Application
	Search() searches.Application
	List() lists.Application
}

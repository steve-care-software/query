package compilers

import (
	"github.com/steve-care-software/query/domain/lists"
	"github.com/steve-care-software/query/domain/retrieves"
	"github.com/steve-care-software/query/domain/searches"
)

// Application represents the compiler application
type Application interface {
	Retrieve(script string) (retrieves.Retrieve, error)
	Search(script string) (searches.Search, error)
	List(script string) (lists.List, error)
}

package retrieves

import "github.com/steve-care-software/query/domain/retrieves"

// Application represents a search application
type Application interface {
	Compile(script string) (retrieves.Retrieve, error)
	Execute(data [][]byte, retrieve retrieves.Retrieve) ([]byte, error)
}

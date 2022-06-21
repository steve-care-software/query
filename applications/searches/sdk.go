package searches

import "github.com/steve-care-software/query/domain/searches"

// Application represents a search application
type Application interface {
	Compile(script string) (searches.Search, error)
	Execute(data [][]byte, search searches.Search) ([][]byte, error)
}

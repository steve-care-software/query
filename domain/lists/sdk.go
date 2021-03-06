package lists

import (
	"github.com/steve-care-software/query/domain/commons"
	"github.com/steve-care-software/selector/domain/selectors"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithSelector(selector selectors.Selector) Builder
	WithRange(rnge commons.Range) Builder
	Now() (List, error)
}

// List represents a list query
type List interface {
	Selector() selectors.Selector
	Range() commons.Range
}

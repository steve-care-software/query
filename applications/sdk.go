package applications

import (
	"github.com/steve-care-software/query/applications/compilers"
	"github.com/steve-care-software/query/applications/executions"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithData(data [][]byte) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Compiler() compilers.Application
	Execution() executions.Application
}

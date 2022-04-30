package results

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the  result builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithCursor(cursor uint) Builder
	WithPath(path []uint) Builder
	IsSuccess() Builder
	Now() (Result, error)
}

// Result represents a mistake
type Result interface {
	Index() uint
	Cursor() uint
	Path() []uint
	IsSuccess() bool
}

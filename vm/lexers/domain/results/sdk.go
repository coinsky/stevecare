package results

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithMistake(mistake Mistake) Builder
	WithSuccess(success Success) Builder
	Now() (Result, error)
}

// Result represents a result
type Result interface {
	IsMistake() bool
	Mistake() Mistake
	IsSuccess() bool
	Success() Success
}

// MistakeBuilder represents the  mistake builder
type MistakeBuilder interface {
	Create() MistakeBuilder
	WithIndex(index uint) MistakeBuilder
	WithPath(path []uint) MistakeBuilder
	Now() (Mistake, error)
}

// Mistake represents a mistake
type Mistake interface {
	Index() uint
	Path() []uint
}

// SuccessBuilder represents the success builder
type SuccessBuilder interface {
	Create() SuccessBuilder
	WithIndex(index uint) SuccessBuilder
	Now() (Success, error)
}

// Success represents the success
type Success interface {
	HasIndex() bool
	Index() *uint
}

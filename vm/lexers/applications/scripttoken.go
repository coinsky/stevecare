package applications

import "github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"

type scriptToken struct {
	index uint
	name  string
	lines []*scriptLine
}

type scriptLine struct {
	values []*scriptValue
}

type scriptValue struct {
	pByte       *byte
	tokenName   string
	cardinality cardinality.Cardinality
}

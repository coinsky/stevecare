package tokens

import "github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"

type elementWithCardinalityAdapter struct {
	builder            ElementWithCardinalityBuilder
	elementAdapter     ElementAdapter
	cardinalityAdapter cardinality.Adapter
}

func createElementWithCardinalityAdapter(
	builder ElementWithCardinalityBuilder,
	elementAdapter ElementAdapter,
	cardinalityAdapter cardinality.Adapter,
) ElementWithCardinalityAdapter {
	out := elementWithCardinalityAdapter{
		builder:            builder,
		elementAdapter:     elementAdapter,
		cardinalityAdapter: cardinalityAdapter,
	}

	return &out
}

// ToElementWithCardinality converts data to an ElementWithCardinality
func (app *elementWithCardinalityAdapter) ToElementWithCardinality(data []byte) (ElementWithCardinality, []byte, error) {
	element, remaining, err := app.elementAdapter.ToElement(data)
	if err != nil {
		return nil, nil, err
	}

	cardinality, remainingAfterCard, err := app.cardinalityAdapter.ToCardinality(remaining)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.builder.Create().WithCardinality(cardinality).WithElement(element).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remainingAfterCard, nil
}

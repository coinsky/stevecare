package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

type application struct {
	resultBuilder results.Builder
}

func createApplication(
	resultBuilder results.Builder,
) Application {
	out := application{
		resultBuilder: resultBuilder,
	}

	return &out
}

// Execute executes the lexer application
func (app *application) Execute(token tokens.Token, data []byte) (results.Result, error) {
	lengthData := len(data)
	remaining, path, err := app.executeToken(token, data, []uint{}, map[uint]*tokenData{})
	index := uint(lengthData - len(remaining))
	builder := app.resultBuilder.Create().WithIndex(index).WithPath(path)
	if err != nil {
		return builder.Now()
	}

	return builder.IsSuccess().Now()
}

func (app *application) executeReference(refIndex uint, data []byte, path []uint, prevTokenData map[uint]*tokenData) ([]byte, []uint, error) {
	if tokenData, ok := prevTokenData[refIndex]; ok {
		prevData := tokenData.Data()
		if len(data) == len(prevData) {
			str := fmt.Sprintf("the referenced token (index: %d) is an infinite recursive token", refIndex)
			return nil, path, errors.New(str)
		}

		token := tokenData.Token()
		return app.executeToken(token, data, path, prevTokenData)
	}

	str := fmt.Sprintf("the referenced token (index: %d) is NOT declared", refIndex)
	return nil, path, errors.New(str)
}

func (app *application) executeToken(token tokens.Token, data []byte, path []uint, prevTokenData map[uint]*tokenData) ([]byte, []uint, error) {
	// add the data to the previous token data map:
	index := token.Index()
	path = append(path, index)
	prevTokenData[index] = createTokenData(token, data)

	lines := token.Lines()
	remaining, retPath := app.executeLines(lines, data, path, prevTokenData)
	if len(remaining) != len(data) {
		return remaining, retPath, nil
	}

	str := fmt.Sprintf("the token (index: %d) could not be matched against the given data", token.Index())
	return remaining, retPath, errors.New(str)
}

func (app *application) executeLines(lines tokens.Lines, data []byte, path []uint, prevTokenData map[uint]*tokenData) ([]byte, []uint) {
	lastPath := path
	list := lines.List()
	remainingData := data
	for _, oneLine := range list {
		retRemainingData, retPath, err := app.executeLine(oneLine, remainingData, path, prevTokenData)
		if err != nil {
			continue
		}

		lastPath = retPath
		remainingData = retRemainingData
	}

	return remainingData, lastPath
}

func (app *application) executeLine(line tokens.Line, data []byte, path []uint, prevTokenData map[uint]*tokenData) ([]byte, []uint, error) {
	lastPath := path
	list := line.List()
	remainingData := data
	for index, oneElementWithCard := range list {
		retRemainingData, retPath, err := app.executeElementWithCardinality(oneElementWithCard, remainingData, lastPath, prevTokenData)
		if err != nil {
			str := fmt.Sprintf("there was an error while executing line (index: %d): error: %s", index, err.Error())
			return remainingData, lastPath, errors.New(str)
		}

		lastPath = retPath
		remainingData = retRemainingData
	}

	return remainingData, lastPath, nil
}

func (app *application) executeElementWithCardinality(elementWithCard tokens.ElementWithCardinality, data []byte, path []uint, prevTokenData map[uint]*tokenData) ([]byte, []uint, error) {
	lastPath := path
	remainingData := data
	element := elementWithCard.Element()
	cardinality := elementWithCard.Cardinality()
	if cardinality.IsSpecific() {
		pSpecific := cardinality.Specific()
		specific := int(*pSpecific)
		for i := 0; i < specific; i++ {
			retRemainingData, retPath, err := app.executeElement(element, remainingData, lastPath, prevTokenData)
			if err != nil {
				str := fmt.Sprintf("there was an error while executing the elementWithCardinality at specific cardinality (%d) at index: %d, error: %s", specific, i, err.Error())
				return remainingData, lastPath, errors.New(str)
			}

			lastPath = retPath
			remainingData = retRemainingData
		}

		return remainingData, lastPath, nil
	}

	cpt := uint(0)
	rnge := cardinality.Range()
	min := rnge.Min()
	for {

		if len(remainingData) <= 0 {
			break
		}

		if rnge.HasMax() {
			pMax := rnge.Max()
			if cpt >= *pMax {
				str := fmt.Sprintf("the maximum cardinality (%d) has been reached while trying to execute the elementWithCardinality at cpt index: %d", *pMax, cpt)
				return remainingData, lastPath, errors.New(str)
			}
		}

		retRemainingData, retPath, err := app.executeElement(element, remainingData, lastPath, prevTokenData)
		if err != nil {
			break
		}

		lastPath = retPath
		remainingData = retRemainingData
		cpt++
	}

	if cpt < min {
		str := fmt.Sprintf("the minimum cardinality (%d) has not been reached while trying to execute the elementWithCardinality at cpt index: %d", min, cpt)
		return remainingData, lastPath, errors.New(str)
	}

	return remainingData, lastPath, nil
}

func (app *application) executeElement(element tokens.Element, data []byte, path []uint, prevTokenData map[uint]*tokenData) ([]byte, []uint, error) {
	if element.IsByte() {
		pByte := element.Byte()
		if len(data) > 0 {
			first := data[0]
			if *pByte != first {
				str := fmt.Sprintf("the element byte (%d) could not match the first data byte (%d)", *pByte, first)
				return data[1:], path, errors.New(str)
			}

			return data[1:], path, nil
		}

		str := fmt.Sprintf("the byte (%d) could not be found in the data because the remaining data was empty", *pByte)
		return data, path, errors.New(str)
	}

	if element.IsToken() {
		token := element.Token()
		return app.executeToken(token, data, path, prevTokenData)
	}

	pReference := element.Reference()
	return app.executeReference(*pReference, data, path, prevTokenData)
}

package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

type application struct {
	resultBuilder  results.Builder
	successBuilder results.SuccessBuilder
	mistakeBuilder results.MistakeBuilder
}

func createApplication(
	resultBuilder results.Builder,
	successBuilder results.SuccessBuilder,
	mistakeBuilder results.MistakeBuilder,
) Application {
	out := application{
		resultBuilder:  resultBuilder,
		successBuilder: successBuilder,
		mistakeBuilder: mistakeBuilder,
	}

	return &out
}

// Execute executes the lexer application
func (app *application) Execute(token tokens.Token, data []byte) (results.Result, error) {
	builder := app.resultBuilder.Create()
	_, _, _, err := app.executeToken(token, data)
	if err != nil {
		tokenIndex := token.Index()
		mistake, err := app.mistakeBuilder.Create().WithIndex(0).WithPath([]uint{
			tokenIndex,
		}).Now()

		if err != nil {
			return nil, err
		}

		return builder.WithMistake(mistake).Now()
	}

	success, err := app.successBuilder.Create().Now()
	if err != nil {
		return nil, err
	}

	return builder.WithSuccess(success).Now()
}

func (app *application) executeToken(token tokens.Token, data []byte) (bool, byte, []byte, error) {
	lines := token.Lines()
	searchedByte, remaining := app.executeLines(lines, data)
	if len(remaining) != len(data) {
		return true, searchedByte, remaining, nil
	}

	str := fmt.Sprintf("the token (index: %d) could not be matched against the given data because it could not find the byte: %d", token.Index(), searchedByte)
	return false, searchedByte, remaining, errors.New(str)
}

func (app *application) executeLines(lines tokens.Lines, data []byte) (byte, []byte) {
	var lastSearchedByte byte
	list := lines.List()
	remainingData := data
	for _, oneLine := range list {
		searchedByte, retRemainingData, err := app.executeLine(oneLine, remainingData)
		if err != nil {
			continue
		}

		lastSearchedByte = searchedByte
		remainingData = retRemainingData
	}

	return lastSearchedByte, remainingData
}

func (app *application) executeLine(line tokens.Line, data []byte) (byte, []byte, error) {
	var lastSearchedByte byte
	list := line.List()
	remainingData := data
	for index, oneElementWithCard := range list {
		searchedByte, retRemainingData, err := app.executeElementWithCardinality(oneElementWithCard, remainingData)
		lastSearchedByte = searchedByte
		if err != nil {
			str := fmt.Sprintf("there was an error while executing line (index: %d): error: %s", index, err.Error())
			return lastSearchedByte, remainingData, errors.New(str)
		}

		remainingData = retRemainingData
	}

	return lastSearchedByte, remainingData, nil
}

func (app *application) executeElementWithCardinality(elementWithCard tokens.ElementWithCardinality, data []byte) (byte, []byte, error) {
	var lastSearchedByte byte
	remainingData := data
	element := elementWithCard.Element()
	cardinality := elementWithCard.Cardinality()
	if cardinality.IsSpecific() {
		pSpecific := cardinality.Specific()
		specific := int(*pSpecific)
		for i := 0; i < specific; i++ {
			works, searchedByte, retRemainingData, err := app.executeElement(element, remainingData)
			lastSearchedByte = searchedByte
			if err != nil {
				str := fmt.Sprintf("there was an error while trying to find the byte (%d) at specific cardinality (%d) at index: %d, error: %s", searchedByte, specific, i, err.Error())
				return lastSearchedByte, remainingData, errors.New(str)
			}

			if !works {
				str := fmt.Sprintf("the byte (%d) could not match the data (%d) at specific cardinality (%d) at index: %d", searchedByte, remainingData[0], specific, i)
				return lastSearchedByte, remainingData, errors.New(str)
			}

			remainingData = retRemainingData
		}

		return lastSearchedByte, remainingData, nil
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
				str := fmt.Sprintf("the maximum cardinality (%d) has been reached while trying to find the byte (%d), cpt index: %d", *pMax, lastSearchedByte, cpt)
				return lastSearchedByte, remainingData, errors.New(str)
			}
		}

		works, searchedByte, retRemainingData, err := app.executeElement(element, remainingData)
		lastSearchedByte = searchedByte
		if err != nil {
			break
		}

		if !works {
			break
		}

		remainingData = retRemainingData
		cpt++
	}

	if cpt < min {
		str := fmt.Sprintf("the minimum cardinality (%d) has not been reached while trying to find the byte (%d), cpt index: %d", min, lastSearchedByte, cpt)
		return lastSearchedByte, remainingData, errors.New(str)
	}

	return lastSearchedByte, remainingData, nil
}

func (app *application) executeElement(element tokens.Element, data []byte) (bool, byte, []byte, error) {
	if element.IsByte() {
		pByte := element.Byte()
		if len(data) > 0 {
			first := data[0]
			return *pByte == first, *pByte, data[1:], nil
		}

		return false, *pByte, data, errors.New("empty data")
	}

	token := element.Token()
	return app.executeToken(token, data)
}

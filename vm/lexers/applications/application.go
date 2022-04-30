package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

type application struct {
	resultBuilder results.Builder
	successBuilder results.SuccessBuilder
}

func createApplication(
	resultBuilder results.Builder,
	successBuilder results.SuccessBuilder,
) Application {
	out := application{
		resultBuilder: resultBuilder,
		successBuilder:successBuilder,
	}

	return &out
}

// Execute executes the lexer application
func (app *application) Execute(token tokens.Token, data []byte) (results.Result, error) {
	lines :=token.Lines()
	remaining := app.executeLines(lines, data)
	builder:=app.resultBuilder.Create()
	if len(remaining) != len(data) {
		success,err:=app.successBuilder.Create().Now()
		if err !=nil {
			return nil, err
		}

		builder.WithSuccess(success)
	}

	return builder.Now()
}

func (app *application) executeLines(lines tokens.Lines, data []byte) ([]byte)  {
	list :=lines.List()
	remainingData := data
	for _, oneLine:=range(list){
		retRemainingData,err :=app.executeLine(oneLine, remainingData)
		if err !=nil{
			continue
		}

		remainingData = retRemainingData
	}

	return remainingData
}

func (app *application) executeLine(line tokens.Line, data []byte) ([]byte, error) {
	list :=line.List()
	remainingData := data
	for index,oneElementWithCard:=range(list) {
		retRemainingData, err:= app.executeElementWithCardinality(oneElementWithCard, remainingData)
		if err !=nil{
			str:=fmt.Sprintf("there was an error while executing line (index: %d): error: %s", index,err.Error())
			return remainingData, errors.New(str)
		}

		remainingData = retRemainingData
	}

	return remainingData, nil
}

func (app *application) executeElementWithCardinality(elementWithCard tokens.ElementWithCardinality, data []byte) ([]byte, error) {
	element := elementWithCard.Element()
	cardinality := elementWithCard.Cardinality()
	if cardinality.IsSpecific() {
		pSpecific := cardinality.Specific()
		specific := int(*pSpecific)
		remainingData := data
		for i := 0; i < specific; i++ {
			works, searchedByte, retRemainingData, err := app.executeElement(element, remainingData)
			if err != nil {
				str := fmt.Sprintf("there was an error while trying to find the byte (%d) at specific cardinality (%d) at index: %d, error: %s", searchedByte, specific, i, err.Error())
				return remainingData, errors.New(str)
			}

			if !works {
				str := fmt.Sprintf("the byte (%d) could not match the data (%d) at specific cardinality (%d) at index: %d", searchedByte, remainingData[0], specific, i)
				return remainingData, errors.New(str)
			}

			remainingData = retRemainingData
		}

		return remainingData, nil
	}

	panic(errors.New("finish executeElementWithCardinality with Range Cardinality"))
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

	panic(errors.New("finish executeElement with Token"))
}

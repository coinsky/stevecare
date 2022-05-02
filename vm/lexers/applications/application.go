package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/channels"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/grammars"
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
func (app *application) Execute(grammar grammars.Grammar, data []byte, canHavePrefix bool) (results.Result, error) {
	token := grammar.Root()
	channels := grammar.Channels()
	if canHavePrefix {
		index := uint(0)
		reaminingData := data
		for {
			if len(reaminingData) <= 0 {
				break
			}

			cursor, path, isSuccess := app.executeOnce(token, channels, reaminingData, index)
			if isSuccess {
				return app.resultBuilder.Create().WithIndex(index).WithCursor(cursor).WithPath(path).IsSuccess().Now()
			}

			reaminingData = reaminingData[1:]
			index++
		}
	}

	cursor, path, _ := app.executeOnce(token, channels, data, 0)
	return app.resultBuilder.Create().WithIndex(0).WithCursor(cursor).WithPath(path).Now()
}

func (app *application) executeOnce(
	currentToken tokens.Token,
	chans channels.Channels,
	data []byte,
	index uint,
) (uint, []uint, bool) {
	channelsList := []channels.Channel{}
	if chans != nil {
		channelsList = chans.List()
	}

	lengthData := len(data)
	lengthDataPlusIndex := lengthData + int(index)
	remaining, path, previousToken, prevTokenData, err := app.executeToken(nil, currentToken, channelsList, nil, data, []uint{}, map[uint]*tokenData{})
	remainingLength := len(remaining)
	cursor := uint(lengthDataPlusIndex - remainingLength)
	if err != nil {
		return cursor, path, false
	}

	if len(channelsList) > 0 {
		cursorIndex := uint(lengthData - remainingLength)
		remainingAfterChans, err := app.executeChannels(previousToken, nil, channelsList, data, data[cursorIndex:], path, prevTokenData)
		if err != nil {
			return cursor, path, false
		}

		cursor = uint(lengthData - len(remainingAfterChans))
	}

	return cursor, path, true
}

func (app *application) executeChannels(
	previousToken tokens.Token,
	currentToken tokens.Token,
	channelsList []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, error) {

	executeTokenFn := func(token tokens.Token, currentData []byte) bool {
		_, _, _, _, err := app.executeToken(nil, token, []channels.Channel{}, []byte{}, currentData, path, prevTokenData)
		if err != nil {
			return false
		}

		return true
	}

	executeBothTokenFn := func(
		bothCondition channels.BothCondition,
		nextCurrentData []byte,
		prevCurrentData []byte,
	) (bool, bool) {
		next := bothCondition.Next()
		previous := bothCondition.Previous()
		return executeTokenFn(next, nextCurrentData), executeTokenFn(previous, prevCurrentData)
	}

	previousRemainingData := previousData
	remainingData := currentData
	for _, oneChannel := range channelsList {
		token := oneChannel.Token()
		retRemaining, _, _, _, err := app.executeToken(nil, token, []channels.Channel{}, []byte{}, remainingData, path, prevTokenData)
		if err != nil {
			continue
		}

		loopRemaining := []byte{}
		hasCondition := oneChannel.HasCondition()
		if hasCondition {
			condition := oneChannel.Condition()
			if condition.IsNext() {
				next := condition.Next()
				if !executeTokenFn(next, retRemaining) {
					continue
				}

				loopRemaining = retRemaining
			}

			if condition.IsPrevious() {
				previous := condition.Previous()
				if !executeTokenFn(previous, previousRemainingData) {
					continue
				}

				loopRemaining = retRemaining
			}

			if condition.IsAnd() {
				bothCondition := condition.And()
				isNextMatch, isPrevMatch := executeBothTokenFn(bothCondition, retRemaining, previousRemainingData)
				if !(isNextMatch && isPrevMatch) {
					continue
				}

				loopRemaining = retRemaining
			}

			if condition.IsOr() {
				bothCondition := condition.And()
				isNextMatch, isPrevMatch := executeBothTokenFn(bothCondition, retRemaining, previousRemainingData)
				if !(isNextMatch || isPrevMatch) {
					continue
				}

				loopRemaining = retRemaining
			}

			if condition.IsXor() {
				bothCondition := condition.And()
				isNextMatch, isPrevMatch := executeBothTokenFn(bothCondition, retRemaining, previousRemainingData)
				if !(!isNextMatch && !isPrevMatch) {
					continue
				}

				loopRemaining = retRemaining
			}
		}

		if !hasCondition {
			loopRemaining = retRemaining
		}

		previousRemainingData = remainingData
		remainingData = loopRemaining
	}

	return remainingData, nil
}

func (app *application) executeReference(
	previousToken tokens.Token,
	refIndex uint,
	channels []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, []uint, tokens.Token, map[uint]*tokenData, error) {
	if tokenData, ok := prevTokenData[refIndex]; ok {
		prevData := tokenData.Data()
		if len(currentData) == len(prevData) {
			str := fmt.Sprintf("the referenced token (index: %d) is an infinite recursive token", refIndex)
			return nil, path, previousToken, prevTokenData, errors.New(str)
		}

		token := tokenData.Token()
		return app.executeToken(previousToken, token, channels, previousData, currentData, path, prevTokenData)
	}

	str := fmt.Sprintf("the referenced token (index: %d) is NOT declared", refIndex)
	return nil, path, previousToken, prevTokenData, errors.New(str)
}

func (app *application) executeToken(
	previousToken tokens.Token,
	currentToken tokens.Token,
	channels []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, []uint, tokens.Token, map[uint]*tokenData, error) {
	if len(channels) > 0 {
		remainingData, err := app.executeChannels(previousToken, currentToken, channels, previousData, currentData, path, prevTokenData)
		if err != nil {
			return nil, nil, nil, prevTokenData, err
		}

		currentData = remainingData
	}

	if previousToken == nil {
		previousToken = currentToken
	}

	// add the data to the previous token data map:
	index := currentToken.Index()
	path = append(path, index)
	prevTokenData[index] = createTokenData(currentToken, currentData)

	lines := currentToken.Lines()
	remaining, retPath := app.executeLines(previousToken, lines, channels, previousData, currentData, path, prevTokenData)
	if len(remaining) != len(currentData) {
		return remaining, retPath, currentToken, prevTokenData, nil
	}

	str := fmt.Sprintf("the token (index: %d) could not be matched against the given data", currentToken.Index())
	return remaining, retPath, currentToken, prevTokenData, errors.New(str)
}

func (app *application) executeLines(
	previousToken tokens.Token,
	lines tokens.Lines,
	channels []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, []uint) {
	lastPath := path
	list := lines.List()
	previousRemainingData := previousData
	remainingData := currentData
	for _, oneLine := range list {
		retRemainingData, retPath, err := app.executeLine(previousToken, oneLine, channels, previousRemainingData, remainingData, path, prevTokenData)
		if err != nil {
			continue
		}

		lastPath = retPath
		previousRemainingData = remainingData
		remainingData = retRemainingData
	}

	return remainingData, lastPath
}

func (app *application) executeLine(
	previousToken tokens.Token,
	line tokens.Line,
	channels []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, []uint, error) {
	lastPath := path
	list := line.List()
	previousRemainingData := previousData
	remainingData := currentData
	for index, oneElementWithCard := range list {
		retRemainingData, retPath, err := app.executeElementWithCardinality(previousToken, oneElementWithCard, channels, previousRemainingData, remainingData, lastPath, prevTokenData)
		if err != nil {
			str := fmt.Sprintf("there was an error while executing line (index: %d): error: %s", index, err.Error())
			return remainingData, lastPath, errors.New(str)
		}

		lastPath = retPath
		previousRemainingData = remainingData
		remainingData = retRemainingData
	}

	return remainingData, lastPath, nil
}

func (app *application) executeElementWithCardinality(
	previousToken tokens.Token,
	elementWithCard tokens.ElementWithCardinality,
	channels []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, []uint, error) {
	lastPath := path
	previousRemainingData := previousData
	remainingData := currentData
	element := elementWithCard.Element()
	cardinality := elementWithCard.Cardinality()
	if cardinality.IsSpecific() {
		pSpecific := cardinality.Specific()
		specific := int(*pSpecific)
		for i := 0; i < specific; i++ {
			retRemainingData, retPath, _, _, err := app.executeElement(previousToken, element, channels, previousRemainingData, remainingData, lastPath, prevTokenData)
			if err != nil {
				str := fmt.Sprintf("there was an error while executing the elementWithCardinality at specific cardinality (%d) at index: %d, error: %s", specific, i, err.Error())
				return remainingData, lastPath, errors.New(str)
			}

			lastPath = retPath
			previousRemainingData = remainingData
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
				break
			}
		}

		retRemainingData, retPath, _, _, err := app.executeElement(previousToken, element, channels, previousRemainingData, remainingData, lastPath, prevTokenData)
		if err != nil {
			break
		}

		lastPath = retPath
		previousRemainingData = remainingData
		remainingData = retRemainingData
		cpt++
	}

	if cpt < min {
		str := fmt.Sprintf("the minimum cardinality (%d) has not been reached while trying to execute the elementWithCardinality at cpt index: %d", min, cpt)
		return remainingData, lastPath, errors.New(str)
	}

	return remainingData, lastPath, nil
}

func (app *application) executeElement(
	previousToken tokens.Token,
	element tokens.Element,
	channels []channels.Channel,
	previousData []byte,
	currentData []byte,
	path []uint,
	prevTokenData map[uint]*tokenData,
) ([]byte, []uint, tokens.Token, map[uint]*tokenData, error) {
	if element.IsByte() {
		pByte := element.Byte()
		if len(currentData) > 0 {
			first := currentData[0]
			if *pByte != first {
				str := fmt.Sprintf("the element byte (%d) could not match the first data byte (%d)", *pByte, first)
				return currentData[1:], path, previousToken, prevTokenData, errors.New(str)
			}

			return currentData[1:], path, previousToken, prevTokenData, nil
		}

		str := fmt.Sprintf("the byte (%d) could not be found in the data because the remaining data was empty", *pByte)
		return currentData, path, previousToken, prevTokenData, errors.New(str)
	}

	if element.IsToken() {
		token := element.Token()
		return app.executeToken(previousToken, token, channels, previousData, currentData, path, prevTokenData)
	}

	pReference := element.Reference()
	return app.executeReference(previousToken, *pReference, channels, previousData, currentData, path, prevTokenData)
}

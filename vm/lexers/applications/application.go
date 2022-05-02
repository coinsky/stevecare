package applications

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/channels"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/grammars"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

type application struct {
	grammarBuilder                grammars.Builder
	tokenBuilder                  tokens.TokenBuilder
	linesBuilder                  tokens.LinesBuilder
	lineBuilder                   tokens.LineBuilder
	elementWithCardinalityBuilder tokens.ElementWithCardinalityBuilder
	elementBuilder                tokens.ElementBuilder
	cardinalityBuilder            cardinality.Builder
	cardinalityRangeBuilder       cardinality.RangeBuilder
	resultBuilder                 results.Builder
	rootPrefix                    byte
	rootSuffix                    byte
	tokenNamePrefix               byte
	bytePrefix                    byte
	indexTokenNameSeparator       byte
	linesPrefix                   byte
	linesSuffix                   byte
	lineDelimiter                 byte
	cardinalityNonZeroMultiple    byte
	cardinalityZeroMultiple       byte
	cardinalityOptional           byte
	cardinalityRangePrefix        byte
	cardinalityRangeSuffix        byte
	cardinalityRangeSeparator     byte
	commentPrefix                 byte
	commentSuffix                 byte
	numbersCharacters             []byte
	tokenNameCharacters           []byte
	channelCharacters             []byte
}

func createApplication(
	grammarBuilder grammars.Builder,
	tokenBuilder tokens.TokenBuilder,
	linesBuilder tokens.LinesBuilder,
	lineBuilder tokens.LineBuilder,
	elementWithCardinalityBuilder tokens.ElementWithCardinalityBuilder,
	elementBuilder tokens.ElementBuilder,
	cardinalityBuilder cardinality.Builder,
	cardinalityRangeBuilder cardinality.RangeBuilder,
	resultBuilder results.Builder,
	rootPrefix byte,
	rootSuffix byte,
	tokenNamePrefix byte,
	bytePrefix byte,
	indexTokenNameSeparator byte,
	linesPrefix byte,
	linesSuffix byte,
	lineDelimiter byte,
	cardinalityNonZeroMultiple byte,
	cardinalityZeroMultiple byte,
	cardinalityOptional byte,
	cardinalityRangePrefix byte,
	cardinalityRangeSuffix byte,
	cardinalityRangeSeparator byte,
	commentPrefix byte,
	commentSuffix byte,
	numbersCharacters []byte,
	tokenNameCharacters []byte,
	channelCharacters []byte,
) Application {
	out := application{
		grammarBuilder:                grammarBuilder,
		tokenBuilder:                  tokenBuilder,
		linesBuilder:                  linesBuilder,
		lineBuilder:                   lineBuilder,
		elementWithCardinalityBuilder: elementWithCardinalityBuilder,
		elementBuilder:                elementBuilder,
		cardinalityBuilder:            cardinalityBuilder,
		cardinalityRangeBuilder:       cardinalityRangeBuilder,
		resultBuilder:                 resultBuilder,
		rootPrefix:                    rootPrefix,
		rootSuffix:                    rootSuffix,
		tokenNamePrefix:               tokenNamePrefix,
		bytePrefix:                    bytePrefix,
		indexTokenNameSeparator:       indexTokenNameSeparator,
		linesPrefix:                   linesPrefix,
		linesSuffix:                   linesSuffix,
		lineDelimiter:                 lineDelimiter,
		cardinalityNonZeroMultiple:    cardinalityNonZeroMultiple,
		cardinalityZeroMultiple:       cardinalityZeroMultiple,
		cardinalityOptional:           cardinalityOptional,
		cardinalityRangePrefix:        cardinalityRangePrefix,
		cardinalityRangeSuffix:        cardinalityRangeSuffix,
		cardinalityRangeSeparator:     cardinalityRangeSeparator,
		commentPrefix:                 commentPrefix,
		commentSuffix:                 commentSuffix,
		numbersCharacters:             numbersCharacters,
		tokenNameCharacters:           tokenNameCharacters,
		channelCharacters:             channelCharacters,
	}

	return &out
}

// Compile compiles a script to a grammar instance
func (app *application) Compile(script string) (grammars.Grammar, error) {
	// convert to bytes:
	bytes := []byte(script)

	// remove channel characters:
	remainingAfterChans := app.removeChannelCharacters(bytes)

	// retrieve the root token:
	rootTokenName, remaining, err := app.fetchRootTokenName(remainingAfterChans)
	if err != nil {
		return nil, err
	}

	// retrieve the script tokens:
	tokensMap, err := app.toScriptTokens(remaining)
	if err != nil {
		return nil, err
	}

	// convert the script tokens and rootTokenName to a root Token instance:
	rootToken, err := app.toToken(rootTokenName, tokensMap, []string{})
	if err != nil {
		return nil, err
	}

	return app.grammarBuilder.Create().WithRoot(rootToken).Now()
}

func (app *application) toToken(rootTokenName string, scriptTokensMap map[string]*scriptToken, path []string) (tokens.Token, error) {
	if rootScriptToken, ok := scriptTokensMap[rootTokenName]; ok {
		linesList := []tokens.Line{}
		for _, oneLine := range rootScriptToken.lines {
			elementsList := []tokens.ElementWithCardinality{}
			for _, oneValue := range oneLine.values {
				if oneValue.pByte != nil {
					element, err := app.elementBuilder.Create().WithByte(*oneValue.pByte).Now()
					if err != nil {
						return nil, err
					}

					elementWithCard, err := app.elementWithCardinalityBuilder.Create().WithElement(element).WithCardinality(oneValue.cardinality).Now()
					if err != nil {
						return nil, err
					}

					elementsList = append(elementsList, elementWithCard)
					continue
				}

				isReference := false
				path = append(path, rootTokenName)
				for _, onePrevName := range path {
					if oneValue.tokenName == onePrevName {
						isReference = true
						break
					}
				}

				if isReference {
					if refScriptToken, ok := scriptTokensMap[oneValue.tokenName]; ok {
						element, err := app.elementBuilder.Create().WithReference(refScriptToken.index).Now()
						if err != nil {
							return nil, err
						}

						elementWithCard, err := app.elementWithCardinalityBuilder.Create().WithElement(element).WithCardinality(oneValue.cardinality).Now()
						if err != nil {
							return nil, err
						}

						elementsList = append(elementsList, elementWithCard)
						continue
					}

					str := fmt.Sprintf("the referenced token (name: %s) is undefined", oneValue.tokenName)
					return nil, errors.New(str)

				}

				subToken, err := app.toToken(oneValue.tokenName, scriptTokensMap, path)
				if err != nil {
					return nil, err
				}

				element, err := app.elementBuilder.Create().WithToken(subToken).Now()
				if err != nil {
					return nil, err
				}

				elementWithCard, err := app.elementWithCardinalityBuilder.Create().WithElement(element).WithCardinality(oneValue.cardinality).Now()
				if err != nil {
					return nil, err
				}

				elementsList = append(elementsList, elementWithCard)
				continue
			}

			line, err := app.lineBuilder.Create().WithList(elementsList).Now()
			if err != nil {
				return nil, err
			}

			linesList = append(linesList, line)
		}

		lines, err := app.linesBuilder.Create().WithList(linesList).Now()
		if err != nil {
			return nil, err
		}

		return app.tokenBuilder.Create().WithIndex(rootScriptToken.index).WithLines(lines).Now()
	}

	str := fmt.Sprintf("the root token (name: %s) is undefined", rootTokenName)
	return nil, errors.New(str)
}

func (app *application) fetchRootTokenName(input []byte) (string, []byte, error) {
	if len(input) <= 0 {
		return "", nil, errors.New("the input was NOT expected to be empty while fetching the root token name")
	}

	if input[0] != app.rootPrefix {
		str := fmt.Sprintf("the root prefix (%d) was expected, %d provided", app.rootPrefix, input[0])
		return "", nil, errors.New(str)
	}

	tokenName, remaining, err := app.fetchTokenName(input[1:])
	if err != nil {
		return "", nil, err
	}

	if remaining[0] != app.rootSuffix {
		str := fmt.Sprintf("the root suffix (%d) was expected, %d provided", app.rootSuffix, remaining[0])
		return "", nil, errors.New(str)
	}

	return tokenName, remaining[1:], nil
}

func (app *application) toScriptTokens(input []byte) (map[string]*scriptToken, error) {
	remainingInput := input
	tokens := map[string]*scriptToken{}
	for {

		if len(remainingInput) <= 0 {
			break
		}

		scriptToken, remaining, err := app.toScriptToken(remainingInput)
		if err != nil {
			return nil, err
		}

		remainingInput = remaining
		tokens[scriptToken.name] = scriptToken
	}

	return tokens, nil
}

func (app *application) toScriptToken(input []byte) (*scriptToken, []byte, error) {
	pIndex, remaining, err := app.fetchNumber(input)
	if err != nil {
		return nil, nil, err
	}

	if len(remaining) < 1 {
		return nil, nil, errors.New("the input was NOT expected to be empty while fetching the tokenNameSeparator")
	}

	if remaining[0] != app.indexTokenNameSeparator {
		str := fmt.Sprintf("the first element of the input was expected to be the tokenNameSeparator (%d), %d provided", app.indexTokenNameSeparator, remaining[0])
		return nil, nil, errors.New(str)
	}

	tokenName, remainingAfterTokenName, err := app.fetchTokenName(remaining[1:])
	if err != nil {
		return nil, nil, err
	}

	scriptLines, remainingAfterLines, err := app.fetchLines(remainingAfterTokenName)
	if err != nil {
		str := fmt.Sprintf("there was an error while fetching lines in token (name: %s), error: %s", tokenName, err.Error())
		return nil, nil, errors.New(str)
	}

	return &scriptToken{
		index: *pIndex,
		name:  tokenName,
		lines: scriptLines,
	}, remainingAfterLines, nil
}

/*
func (app *application) fetchLines(input []byte) ([]*scriptLine, []byte, error) {
	cpt := 0
	remainingInput := input[1:]
	lines := []*scriptLine{}
	for {
		values, remaining, err := app.fetchValues(input)
		if err != nil {
			return nil, nil, err
		}

		remainingInput = remaining
		lines = append(lines, &scriptLine{
			values,
		})

		if len(remainingInput) <= 0 {
			str := fmt.Sprintf("the input was NOT expected to be empty while fetching the line (index: %d)", cpt)
			return nil, nil, errors.New(str)
		}

		if remainingInput[0] == app.lineDelimiter {
			fmt.Printf("\n%v\n", values)
			panic(errors.New("yeah"))
			remainingInput = remainingInput[1:]
			cpt++
			continue
		}

		if remainingInput[0] == app.linesSuffix {
			remainingInput = remainingInput[1:]
			break
		}

		str := fmt.Sprintf("the input was expected to contain a linesSuffix (%d) or a lineDelimiter (%d), %d provided", app.linesSuffix, app.lineDelimiter, remainingInput[0])
		return nil, nil, errors.New(str)
	}

	return lines, remainingInput, nil
}
/*/
func (app *application) fetchLines(input []byte) ([]*scriptLine, []byte, error) {
	if len(input) < 1 {
		return nil, nil, errors.New("the input was NOT expected to be empty while fetching the line values")
	}

	if input[0] != app.linesPrefix {
		str := fmt.Sprintf("the first element of the input was expected to be the linesPrefix (%d), %d provided", app.linesPrefix, input[0])
		return nil, nil, errors.New(str)
	}

	remainingInput := input[1:]
	values := []*scriptValue{}
	lines := []*scriptLine{}
	for {
		value, retRemaining, err := app.fetchValue(remainingInput)
		if err != nil {
			str := fmt.Sprintf("there is an error while fetching the line's element (line: %d, element: %d), error: %s", len(lines), len(values), err.Error())
			return nil, nil, errors.New(str)
		}

		if len(retRemaining) <= 0 {
			str := fmt.Sprintf("the first element of the input was NOT expected to be empty while fetching line (index: %d)", len(lines))
			return nil, nil, errors.New(str)
		}

		values = append(values, value)
		if retRemaining[0] == app.lineDelimiter {
			remainingInput = retRemaining[1:]
			lines = append(lines, &scriptLine{
				values,
			})

			values = []*scriptValue{}
			continue
		}

		if retRemaining[0] == app.linesSuffix {
			remainingInput = retRemaining[1:]
			lines = append(lines, &scriptLine{
				values,
			})

			break
		}

		remainingInput = retRemaining
	}

	return lines, remainingInput, nil
}

func (app *application) fetchValue(input []byte) (*scriptValue, []byte, error) {
	pByte, tokenName, remaining, err := app.fetchElement(input)
	if err != nil {
		return nil, nil, err
	}

	cardinality, remainingAfterCardinality, err := app.fetchCardinality(remaining)
	if err != nil {
		return nil, nil, err
	}

	return &scriptValue{
		pByte:       pByte,
		tokenName:   tokenName,
		cardinality: cardinality,
	}, remainingAfterCardinality, nil
}

func (app *application) fetchCardinality(input []byte) (cardinality.Cardinality, []byte, error) {
	if len(input) <= 0 {
		return nil, nil, errors.New("the input was NOT expected to be empty while fetching the element's cardinality")
	}

	remaining := input
	builder := app.cardinalityBuilder.Create()
	if input[0] == app.cardinalityNonZeroMultiple {
		rnge, err := app.cardinalityRangeBuilder.Create().WithMinimum(1).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithRange(rnge)
		remaining = input[1:]
	}

	if input[0] == app.cardinalityZeroMultiple {
		rnge, err := app.cardinalityRangeBuilder.Create().WithMinimum(0).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithRange(rnge)
		remaining = input[1:]
	}

	if input[0] == app.cardinalityOptional {
		rnge, err := app.cardinalityRangeBuilder.Create().WithMinimum(0).WithMaximum(1).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithRange(rnge)
		remaining = input[1:]
	}

	if input[0] == app.cardinalityRangePrefix {
		pSpecific, rnge, retRemaining, err := app.fetchCardinalityRangeOrSpecific(input[1:])
		if err != nil {
			return nil, nil, err
		}

		if pSpecific != nil {
			builder.WithSpecific(*pSpecific)
		}

		if rnge != nil {
			builder.WithRange(rnge)
		}

		remaining = retRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		ins, err = builder.WithSpecific(1).Now()
		if err != nil {
			return nil, nil, err
		}
	}

	return ins, remaining, nil
}

func (app *application) fetchCardinalityRangeOrSpecific(input []byte) (*uint8, cardinality.Range, []byte, error) {
	pFirstNumber, isSpecific, retRemaining, err := app.fetchFirstNumberInRange(input)
	if err != nil {
		return nil, nil, nil, err
	}

	if isSpecific {
		return pFirstNumber, nil, retRemaining, nil
	}

	rangeBuilder := app.cardinalityRangeBuilder.Create().WithMinimum(*pFirstNumber)
	pSecondNumber, _, retRemainingAfterMax, _ := app.fetchFirstNumberInRange(retRemaining)
	if pSecondNumber != nil {
		rangeBuilder.WithMaximum(*pSecondNumber)
	}

	rnge, err := rangeBuilder.Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return nil, rnge, retRemainingAfterMax, nil
}

func (app *application) fetchFirstNumberInRange(input []byte) (*uint8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, false, nil, errors.New("the input was NOT expected to be empty while fetching the element's cardinality range number (min/max)")
	}

	isSpecific := true
	numberBytes := []byte{}
	for _, oneInputByte := range input {
		if oneInputByte == app.cardinalityRangeSeparator {
			isSpecific = false
			break
		}

		if oneInputByte == app.cardinalityRangeSuffix {
			break
		}

		if !app.isBytePresent(oneInputByte, app.numbersCharacters) {
			return nil, false, nil, errors.New("the input elements within a range must be numbers")
		}

		numberBytes = append(numberBytes, oneInputByte)
	}

	if len(numberBytes) <= 0 {
		return nil, false, input[1:], nil
	}

	intNumber, err := strconv.Atoi(string(numberBytes))
	if err != nil {
		return nil, false, nil, err
	}

	if intNumber >= 256 {
		str := fmt.Sprintf("the elements of a cardinality (range, specific) must contain a maximum value of 256, %d provided", intNumber)
		return nil, false, nil, errors.New(str)
	}

	casted := uint8(intNumber)
	return &casted, isSpecific, input[len(numberBytes)+1:], nil
}

func (app *application) fetchElement(input []byte) (*byte, string, []byte, error) {
	if len(input) <= 0 {
		return nil, "", nil, errors.New("the input was NOT expected to be empty while fetching the element")
	}

	if input[0] == app.bytePrefix {
		pUint, remaining, err := app.fetchNumber(input[1:])
		if err != nil {
			return nil, "", nil, err
		}

		if *pUint >= 256 {
			str := fmt.Sprintf("the byte in the given element cannot exceed 256, %d provided", *pUint)
			return nil, "", nil, errors.New(str)
		}

		byteValue := byte(*pUint)
		return &byteValue, "", remaining, nil
	}

	if input[0] == app.tokenNamePrefix {
		str, retRemaining, err := app.fetchTokenName(input[1:])
		if err != nil {
			return nil, "", nil, err
		}

		return nil, str, retRemaining, nil
	}

	str := fmt.Sprintf("the first element of the input was expecting either a bytePrefix (%d) or a tokenNamePrefix (%d) while fetching an element, %d provided", app.bytePrefix, app.tokenNamePrefix, input[0])
	return nil, "", nil, errors.New(str)
}

func (app *application) fetchTokenName(input []byte) (string, []byte, error) {
	nameBytes := []byte{}
	for _, oneInputByte := range input {
		if !app.isBytePresent(oneInputByte, app.tokenNameCharacters) {
			break
		}

		nameBytes = append(nameBytes, oneInputByte)
	}

	if len(nameBytes) <= 0 {
		return "", nil, errors.New("the tokenName must contain at least 1 character, none provided")
	}

	return string(nameBytes), input[len(nameBytes):], nil
}

func (app *application) fetchNumber(input []byte) (*uint, []byte, error) {
	indexBytes := []byte{}
	for _, oneInputByte := range input {
		if !app.isBytePresent(oneInputByte, app.numbersCharacters) {
			break
		}

		indexBytes = append(indexBytes, oneInputByte)
	}

	if len(indexBytes) <= 0 {
		return nil, nil, errors.New("the input does not contain a number")
	}

	intNumber, err := strconv.Atoi(string(indexBytes))
	if err != nil {
		return nil, nil, err
	}

	casted := uint(intNumber)
	return &casted, input[len(indexBytes):], nil
}

func (app *application) isBytePresent(value byte, data []byte) bool {
	isPresent := false
	for _, oneChanByte := range data {
		if value == oneChanByte {
			isPresent = true
			break
		}
	}

	return isPresent
}

func (app *application) removeChannelCharacters(input []byte) []byte {
	output := []byte{}
	for _, oneInputByte := range input {
		if app.isBytePresent(oneInputByte, app.channelCharacters) {
			continue
		}

		output = append(output, oneInputByte)
	}

	return output
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
			return nil, nil, nil, nil, err
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
	min := uint(rnge.Min())
	for {

		if len(remainingData) <= 0 {
			break
		}

		if rnge.HasMax() {
			pMax := rnge.Max()
			if cpt >= uint(*pMax) {
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

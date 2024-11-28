package jtml

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	IsDark          bool
	HasNumberPrefix bool
}

func ParseAny(input any, config Config) (string, error) {
	result, err := convertToHTML(input, config)

	if err != nil {
		return "", err
	}

	return strings.Join(result, "\n"), nil
}

func convertToHTML(input any, config Config) ([]string, error) {
	jsonString, err := prettifyJSON(input)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(jsonString, "\n")
	result := make([]string, 0)

	var bracketColor, textColor, numberPrefixColor string
	if config.IsDark {
		bracketColor = "#cbd5e1"
		numberPrefixColor = "#cbd5e1"
		textColor = "#86efac"
	} else {
		bracketColor = "#111827"
		textColor = "#047857"
		numberPrefixColor = "#111827"
	}

	for _, line := range lines {

		if isBracket(line) {
			newLine := "<div class='bracket' style='color:" + bracketColor + "'>" + line + "</div>"
			result = append(result, newLine)
			continue
		}

		newLine := "<div class='text' style='color: " + textColor + "'>" + line + "</div>"
		result = append(result, newLine)
	}

	if config.HasNumberPrefix {
		result = appendNumberPrefix(result, numberPrefixColor)
	}

	return result, nil
}

func appendNumberPrefix(list []string, numberPrefixColor string) []string {
	numberResult := make([]string, 0)

	for index, line := range list {
		newLine := "<div style='display: flex; flex-direction: row'>" + "<div class='text' style='color: " + numberPrefixColor + "; margin-right: 2em'>" + strconv.Itoa(index+1) + "</div>" + line + "</div>"
		numberResult = append(numberResult, newLine)
	}

	return numberResult
}

func isBracket(input string) bool {
	re := regexp.MustCompile(`^[{}\[\]\s,\n]*$`)
	return re.MatchString(input)
}

func prettifyJSON(input any) (string, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "<span style=\"display: inline-block; width: 1em;\"></span>")
	if err != nil {
		return "", err
	}

	result := out.String()
	return result, nil
}

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

	for _, line := range lines {

		if isBracket(line) {
			newLine := "<div class='bracket' style='color: #15803d'>" + line + "</div>"
			result = append(result, newLine)
			continue
		}

		newLine := "<div class='text' style='color: #15803d'>" + line + "</div>"
		result = append(result, newLine)
	}

	if config.HasNumberPrefix {
		result = appendNumberPrefix(result)
	}

	return result, nil
}

func appendNumberPrefix(list []string) []string {
	numberResult := make([]string, 0)

	for index, line := range list {
		newLine := "<div style='display: flex; flex-direction: row'>" + "<div class='text' style='color: #b91c1c; margin-right: 2em'>" + strconv.Itoa(index+1) + "</div>" + line + "</div>"
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

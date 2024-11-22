package jtml

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
)

func ParseAny(input any) (string, error) {
	result, err := convertToHTML(input)

	if err != nil {
		return "", err
	}

	return strings.Join(result, "\n"), nil
}

func convertToHTML(input any) ([]string, error) {
	jsonString, err := prettifyJSON(input)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(jsonString, "\n")
	result := make([]string, len(lines))
	for _, line := range lines {

		if isBracket(line) {
			newLine := "<div class='bracket'>" + line + "</div>"
			result = append(result, newLine)
			continue
		}

		newLine := "<div class='text'>" + line + "</div>"
		result = append(result, newLine)

	}

	return result, nil
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
	err = json.Indent(&out, b, "", "<span style=\"display: inline-block; width: 2em;\"></span>")
	if err != nil {
		return "", err
	}

	result := out.String()
	return result, nil
}

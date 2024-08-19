package gorapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func GeneratwHtml(body string) string {
	var sb strings.Builder

	sb.WriteString(`<!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Welcome to Gin</title>
        </head>
        <body>\n`)
	sb.WriteString(body)
	sb.WriteString(`</body>
        </html>`)
	return sb.String()
}

func GetDataFromJson(src string) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(src), &data)
	if err != nil {
		return "", err
	}
	return itemToString(&data), nil
}

func itemToString(data *map[string]interface{}) string {
	var sb strings.Builder
	sb.WriteString("<ul>\n")
	for key, value := range *data {

		sb.WriteString(fieldToString(key, value))
	}
	sb.WriteString("</ul>\n")
	return sb.String()
}

func fieldToString(key, value interface{}) string {
	var textVal string
	switch v := value.(type) {
	case string:
		textVal = fmt.Sprintf("<li>%s: %s</li>\n", key, v)
	case int:
		textVal = fmt.Sprintf("<li>%s: %d\n", key, v)
	case float64:
		textVal = fmt.Sprintf("<li>%s: %.2f\n", key, v)
	case bool:
		textVal = fmt.Sprintf("<li>%s: %t\n", key, v)
	case map[string]interface{}:
		if m, ok := value.(map[string]interface{}); ok {
			contents := itemToString(&m)
			textVal = fmt.Sprintf("<li>%s: %s</li>\n", key, contents)
		} else {
			textVal = fmt.Sprintf("<li>Wrong variable %s\n", key)
		}
	case []interface{}:
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("<li>%s:</li>\n", key))
		for _, item := range v {
			sb.WriteString(elementToString(item))
		}
		textVal = sb.String()
	default:
		textVal = fmt.Sprintf("<li>%s: Value (unknown type %s): %v</li>\n", key, reflect.TypeOf(value), v)
	}
	return textVal
}

func elementToString(value interface{}) string {
	var textVal string
	switch v := value.(type) {
	case string:
		textVal = fmt.Sprintf("<li>%s</li>\n", v)
	case int:
		textVal = fmt.Sprintf("<li>%d\n", v)
	case float64:
		textVal = fmt.Sprintf("<li>%.2f\n", v)
	case bool:
		textVal = fmt.Sprintf("<li>%t\n", v)
	case map[string]interface{}:
		if m, ok := value.(map[string]interface{}); ok {
			textVal = itemToString(&m)
		}
	case []interface{}:
		var sb strings.Builder
		for _, item := range v {
			sb.WriteString(elementToString(item))
		}
		textVal = sb.String()
	default:
		textVal = fmt.Sprintf("<li>Value (unknown type %s): %v</li>\n", reflect.TypeOf(value), v)
	}
	return textVal
}

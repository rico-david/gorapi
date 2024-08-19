package gorapi

import (
	"fmt"
	"os"
	"testing"
)

func TestGetDataFromJson(t *testing.T) {
	sampleJson := `
	{
		"name": "Anthony Doe",
		"age": 19,
		"mail": [
		{
			"address": "anthony.doe@mail.com"
		},
		{
			"address": "tony.doe@mail.net"
		}
		]
	}
	`
	result, err := GetDataFromJson(sampleJson)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.OpenFile("out.html", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if result != "" {
		t.Fatal(result)
	}
}

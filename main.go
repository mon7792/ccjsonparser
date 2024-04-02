package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var openBrakets = []string{"{", "["}
var closeBrakets = []string{"}", "]"}

func main() {
	// check if the user has provided the json as an argument
	if len(os.Args) != 2 {
		log.Println("Usage: go run main.go <jsonFilePath>")
		os.Exit(1)
	}

	// Read the json file content -> current iteration reading all file at once.
	fsContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Println("Error reading file: ", err)
		os.Exit(1)
	}

	log.Println("------------------------------------------")
	log.Println("File: ", os.Args[1])
	log.Println("Result: ", validateJson(string(fsContent)))
	log.Println("------------------------------------------")
}

func validateJson(json string) bool {

	preparedJson := strings.TrimSpace(json)
	preparedJsonLen := len(preparedJson)

	// Check if the json is empty
	if preparedJsonLen == 0 {
		return false
	}

	// check if the json starts with array
	if preparedJson[0] == '[' && preparedJson[preparedJsonLen-1] == ']' {
		if preparedJsonLen == 2 {
			return true
		}

		preparedJson = strings.TrimSpace(preparedJson[1 : preparedJsonLen-1])
		preparedJsonLen = len(preparedJson)
		if preparedJson[0] != '{' {
			return validateValue(preparedJson)
		}
	}

	// Check if 1st character is { and last character is }
	if !slices.Contains(openBrakets, string(preparedJson[0])) || !slices.Contains(closeBrakets, string(preparedJson[preparedJsonLen-1])) {
		return false
	}
	// Remove the first and last character
	preparedJson = strings.TrimSpace(preparedJson[1 : preparedJsonLen-1])
	preparedJsonLen = len(preparedJson)

	// Check if the json is empty
	if preparedJsonLen == 0 {
		return true
	}

	// Check if the json has key value pairs
	keyValuePairs := strings.Split(preparedJson, ",")
	if len(keyValuePairs) == 0 {
		return false
	}

	for i := range keyValuePairs {
		prepKVPair := strings.TrimSpace(keyValuePairs[i])
		if len(prepKVPair) == 0 {
			return false
		}
		kvPair := strings.SplitN(prepKVPair, ":", 2)
		if len(kvPair) != 2 {
			return false
		}
		if len(strings.TrimSpace(kvPair[0])) == 0 || len(strings.TrimSpace(kvPair[1])) == 0 {
			return false
		}

		// Check if the key has " and value is either a string or a number
		if kvPair[0][0] != '"' || kvPair[0][len(kvPair[0])-1] != '"' {
			return false
		}

		// check the value is either a string, number, boolean or object
		if !validateValue(strings.TrimSpace(kvPair[1])) {
			return false
		}
	}

	return true
}

func validateValue(value string) bool {
	// check if the value is either a string, number, boolean or object

	// string
	if value[0] == '"' && value[len(value)-1] == '"' {
		return true
	}

	// boolean
	if value == "true" || value == "false" {
		return true
	}

	// number
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}

	// object
	if value == "null" {
		return true
	}

	if (value[0] == '{' && value[len(value)-1] == '}') || (value[0] == '[' && value[len(value)-1] == ']') {
		return validateJson(value)
	}

	return false
}

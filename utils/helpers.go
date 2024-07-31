package utils

import (
	"encoding/json"
	"strings"
	"unicode"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/rand"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func Includes[SliceType constraints.Ordered](array []SliceType, element SliceType) bool {
	for _, item := range array {
		if item == element {
			return true
		}
	}
	return false
}

func CapitalizeFirstLetter(str string) string {
	return cases.Title(language.English, cases.Compact).String(str)
}

func CamelCaseToSentenceCase(str string) string {
	splitWords := []string{}

	currentWord := ""
	for i, letter := range str {
		if unicode.IsUpper(letter) && i > 0 {
			splitWords = append(splitWords, CapitalizeFirstLetter(currentWord))
			currentWord = ""
		}
		currentWord += string(letter)
	}
	splitWords = append(splitWords, CapitalizeFirstLetter(currentWord))
	return strings.Join(splitWords, " ")
}

func SafeStringify(value interface{}) string {
	jsonByte, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(jsonByte)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString32() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

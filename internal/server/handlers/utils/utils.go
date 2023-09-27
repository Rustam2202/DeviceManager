package utils

import "golang.org/x/text/language"


func IsLanguageValid(str string)bool{
	_, err := language.Parse(str)
	if err != nil {
		return false
	}
	return true
}
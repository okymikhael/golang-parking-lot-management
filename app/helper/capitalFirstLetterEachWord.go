package helper

import "strings"

func CapitalFirstLetterEachWord(text string) string {
	temp := ""
	explodeColor := strings.Split(text, " ")
	for _, c := range explodeColor {
		temp += strings.ToUpper(string(c[0])) + strings.ToLower(c[1:]) + " "
	}
	return temp
}
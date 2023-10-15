package rand

import (
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "^|{}<>[]-._~()'!*:@,;"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

// RandomStr will generate a random string with alphanumeric & special chars
func RandomStr(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	rand.Seed(time.Now().UnixNano())
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

// RandomPin will generate a numeric pin to use
func RandomPin(max int) string {
	rand.Seed(time.Now().UnixNano())
	var pin strings.Builder
	for i := 0; i < max; i++ {
		random := rand.Intn(len(numberSet))
		pin.WriteString(string(numberSet[random]))
	}

	return pin.String()
}

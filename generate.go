package main

import (
	"math/rand"
	"strings"
)

func buildDefaultMessage() string {
	maxRepeat := 5

	msg := "ک"

	c := rand.Intn(maxRepeat) + 1
	msg += strings.Repeat("ی", c)

	msg += "ر"
	msg += " "
	msg += "ت"

	c = rand.Intn(maxRepeat) + 1
	msg += strings.Repeat("و", c)

	msg += "ش"

	endings := []string{"", "", "", ".", "!", "؟"}
	msg += endings[rand.Intn(len(endings))]

	return msg
}

func generateMessage() string {
	return buildDefaultMessage()
}

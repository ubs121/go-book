package common

import (
	"testing"
)

func Greeting(prefix string, who ...string) {
	print(prefix, " ")
	for i := 0; i < len(who); i++ {
		print(who[i], " ")
	}
	println()
}

func TestGreeting(t *testing.T) {
	Greeting("Сайн уу", "Сараа", "Нараа", "Навчаа", "Цэцгээ")
}

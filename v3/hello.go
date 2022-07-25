package Hello

import (
	"fmt"

	"rsc.io/quote/v3"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hello %s this is v3", name)
	return message
}

func Proverb() string {
	return quote.Concurrency()
}

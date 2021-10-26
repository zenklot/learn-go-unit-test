package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Goz")

	if result != "Hello Goz" {
		panic("Result Is Not Hello Goz")
	}
}

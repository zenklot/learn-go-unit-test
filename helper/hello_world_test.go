package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Goz")

	if result != "Hello Goz" {
		t.Fail()
	}
	fmt.Println("Test Hello World 1 Done")
}

func TestHelloWorldDua(t *testing.T) {
	result := HelloWorld("Dua")

	if result != "Hello Dua" {
		t.FailNow()
	}
	fmt.Println("Test Hello World 2 Done")
}

func TestHelloWorldTiga(t *testing.T) {
	result := HelloWorld("Goz")

	if result != "Hello Goz" {
		t.Error("harusnya Hello Goz")
	}
	fmt.Println("Test Hello World 3 Done")
}

func TestHelloWorldEmpat(t *testing.T) {
	result := HelloWorld("Goz")

	if result != "Hello Goz" {
		t.Fatal("Harusnya hello Goz")
	}
	fmt.Println("Test Hello World 4 Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Goz")

	assert.Equal(t, "Hello Goz", result, "Result Must Be Hello Goz")
	fmt.Println("Testing Hello World With Assert Done")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Goz")

	require.Equal(t, "Hello Goz", result, "Result Must Be Hello Goz")
	fmt.Println("Testing Hello World With Require")

}

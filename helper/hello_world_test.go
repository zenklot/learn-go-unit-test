package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Gozenx")
	}
}

func BenchmarkHelloWorldZez(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("SUPRIATNA")
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("satu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Gozenx")
		}
	})

	b.Run("Dua", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("SUPRIATNA")
		}
	})
}

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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't Run on Windows")
	}
	result := HelloWorld("Goz")

	require.Equal(t, "Hello Goz", result, "Result Must Be Hello Goz")
	fmt.Println("Testing Hello World With Require")

}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Tes")

	m.Run()

	fmt.Println("Setelah Tes")
}

func TestSubTest(t *testing.T) {
	t.Run("zez", func(t *testing.T) {
		result := HelloWorld("zez")
		require.Equal(t, "Hallo zez", result)
	})

	t.Run("Raisa", func(t *testing.T) {
		result := HelloWorld("raisa")
		require.Equal(t, "Hallo raisa", result)
	})

}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Hallo(Goz)",
			request:  "Goz",
			expected: "Hallo Goz",
		},
		{
			name:     "Hallo(Zez)",
			request:  "Zez",
			expected: "Hallo Zez",
		},
		{
			name:     "Hallo(Raisa)",
			request:  "Raisa",
			expected: "Hallo Raisa",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

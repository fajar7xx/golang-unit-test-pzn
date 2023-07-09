package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// create sub test
func TestSubTest(t *testing.T) {
	t.Run("Fajar", func(t *testing.T) {
		name := "fajar"
		result := HelloWorld(name)
		require.Equal(t, "Hello "+name, result, "Result must be 'Hello fajar'")
	})

	t.Run("Siagian", func(t *testing.T) {
		name := "siagian"
		result := HelloWorld(name)
		require.Equal(t, "Hello "+name, result, "Result must be 'Hello fajar'")
	})
}

// create table test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fajar",
			request:  "Fajar",
			expected: "Hello Fajar",
		},
		{
			name:     "Nadiah",
			request:  "Nadiah",
			expected: "Hello Nadiah",
		},
		{
			name:     "Azhari",
			request:  "Azhari",
			expected: "Hello Azhari",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// testing yang pertama kali di eksekusi
// akan run sekali setiap package di run/exe
// before & after test
func TestMain(m *testing.M) {
	// before testing
	fmt.Println("before testing bisa inisialisasi db atau lainnya")

	m.Run()

	// after testing
	fmt.Println("after unit test")
}

func TestHelloWorld(t *testing.T) {
	name := "fajar"
	result := HelloWorld(name)

	if result != "Hello "+name {
		// unit test failed
		// panic("result is not Hello fajar")
		// t.Fail()
		t.Error("Harusnya Hello " + name)
	}

	fmt.Println("ini test hello world")
}

func TestHelloFajar(t *testing.T) {
	result := HelloWorld("fajar")

	if result != "Hello fajar" {
		// panic("result is not hello fajar")
		// t.FailNow()
		t.Fatal("result must be 'Hello fajar'")
	}

	fmt.Println("ini test hello world fajar")
}

func TestHelloWorldUsingAssert(t *testing.T) {
	name := "fajar"
	result := HelloWorld(name)
	assert.Equal(t, "Hello "+name, result, "Result must be 'Hello fajar'")
	fmt.Println("test hello world using assert")
}

func TestHelloWorldUsingRequire(t *testing.T) {
	name := "fajar"
	result := HelloWorld(name)
	require.Equal(t, "Hello "+name, result, "Result must be 'Hello fajar'")
	fmt.Println("test hello world using require")
}

func TestHelloWordlSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)

	if runtime.GOOS == "linux" {
		t.Skip("can't run on linux OS")
	}

	name := "fajar"
	result := HelloWorld(name)
	require.Equal(t, "Hello "+name, result, "Result must be 'Hello fajar'")
	fmt.Println("test hello world using require")
}

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fajar")
	}
}

func BenchmarkHelloFajar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nadiah")
	}
}

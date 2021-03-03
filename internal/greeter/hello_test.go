package greeter

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Unknown"
	expected := "Hello, " + name + "!"
	if result := Hello(name); result != expected {
		t.Error(fmt.Sprintf("Invalid hello phrase: expect \"%s\", got \"%s\"", expected, result))
	}
}

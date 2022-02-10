package main

import "testing"

func Hello(name string) string {
	if name == "" {
        name = "World"
    }
    return "Hello, " + name
}

func TestHello(t *testing.T) {
	t.Run("Should say Hello for person", func(t *testing.T) {
		result := Hello("Braun")
		expected := "Hello, Braun"
	
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
	t.Run("Should say Hello, world if no name was passed", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World"
	
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}

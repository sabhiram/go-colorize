// Implement tests for the "colorize" library
package colorize

import (
	"fmt"
	"testing"
)

// TestColorString validates the ColorString function
func TestColorString(test *testing.T) {
	// Define our test cases for the colorize with string color cases
	cases := []struct {
		input, color, expected string
	}{
		// Non-standard cases
		{"Test", "BLACK", "\x1b[30mTest\x1b[0m"},
		{"Test", "Black", "\x1b[30mTest\x1b[0m"},
		{"Test", "INVALID", "Test"},
		{"Test", "", "Test"},

		// Typical usage
		{"Test", "black", "\x1b[30mTest\x1b[0m"},
		{"Test", "red", "\x1b[31mTest\x1b[0m"},
		{"Test", "green", "\x1b[32mTest\x1b[0m"},
		{"Test", "yellow", "\x1b[33mTest\x1b[0m"},
		{"Test", "blue", "\x1b[34mTest\x1b[0m"},
		{"Test", "magenta", "\x1b[35mTest\x1b[0m"},
		{"Test", "cyan", "\x1b[36mTest\x1b[0m"},
		{"Test", "white", "\x1b[37mTest\x1b[0m"},
	}

	// Run tests
	for _, tc := range cases {
		actual := ColorString(tc.input, tc.color)
		if actual != tc.expected {
			test.Errorf("ColorString(%q, %q) == %q, expected %q",
				tc.input, tc.color, actual, tc.expected)
		}
	}
}

// TestColorize validates that the Colorize function works
func TestColorize(test *testing.T) {
	// Define our test cases for the colorize with string color cases
	cases := []struct {
		input, expected string
	}{
		{"Test <red>A</red> <blue>B</blue>", "Test \x1b[31mA\x1b[0m \x1b[34mB\x1b[0m"},
		{"Test <red>A</red> <red>B</red>", "Test \x1b[31mA\x1b[0m \x1b[31mB\x1b[0m"},
	}

	// Run tests
	for _, tc := range cases {
		actual := Colorize(tc.input)
		if actual != tc.expected {
			test.Errorf("Colorize(%q) == %q, expected %q",
				tc.input, actual, tc.expected)
		}
	}
}

// ExampleColorString provides sample usage
func ExampleColorString() {
	fmt.Println("Hello - " + ColorString("black", "black"))
	fmt.Println("Hello - " + ColorString("red", "red"))
	fmt.Println("Hello - " + ColorString("green", "green"))
	fmt.Println("Hello - " + ColorString("yellow", "yellow"))
	fmt.Println("Hello - " + ColorString("blue", "blue"))
	fmt.Println("Hello - " + ColorString("magenta", "magenta"))
	fmt.Println("Hello - " + ColorString("cyan", "cyan"))
	fmt.Println("Hello - " + ColorString("white", "white"))

}

// ExampleColorize provides sample usage
func ExampleColorize() {
	fmt.Println(Colorize(`


<red>This text will be red</red> and this is default...

            <blue>This is blue!</blue>

<red>0</red><green>1</green><yellow>3</yellow><blue>4</blue><magenta>5</magenta><cyan>6</cyan><white>7</white>

`))
}

func init() {
	// Force override on this, otherwise depending on the terminal we will fail
	// things like TravisCI etc.
	DisableColor = false
}

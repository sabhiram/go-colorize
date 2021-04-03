package main

import (
	"fmt"

	"github.com/sabhiram/go-colorize"
)

func main() {
	fmt.Println("Hello - " + colorize.ColorString("black", "black"))
	fmt.Println("Hello - " + colorize.ColorString("red", "red"))
	fmt.Println("Hello - " + colorize.ColorString("green", "green"))
	fmt.Println("Hello - " + colorize.ColorString("yellow", "yellow"))
	fmt.Println("Hello - " + colorize.ColorString("blue", "blue"))
	fmt.Println("Hello - " + colorize.ColorString("magenta", "magenta"))
	fmt.Println("Hello - " + colorize.ColorString("cyan", "cyan"))
	fmt.Println("Hello - " + colorize.ColorString("white", "white"))

	fmt.Println(colorize.Colorize(`

<red>This text will be red</red> and this is default...

            <blue>This is blue!</blue>

<red>0</red><green>1</green><yellow>3</yellow><blue>4</blue><magenta>5</magenta><cyan>6</cyan><white>7</white>

`))
}

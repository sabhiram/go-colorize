/****************************************************************************\

`colorize` is a simple package which returns an ascii colorized
string version of an input string

Here is a table of ASCII to color values:

    Intensity   0       1      2       3       4       5       6       7
    Normal      Black   Red    Green   Yellow  Blue    Magenta Cyan    White
    Bright      Black   Red    Green   Yellow  Blue    Magenta Cyan    White

(sourced from: http://en.wikipedia.org/wiki/ANSI_escape_code)

\****************************************************************************/
package colorize

import (
    "fmt"
    "regexp"

    "strings"
)

/****************************************************************************\

We also define a constant <--> color name mapping based on the above table

\****************************************************************************/
var colorToValueMap = map [string] int {
    "black":   0,
    "red":     1,
    "green":   2,
    "yellow":  3,
    "blue":    4,
    "magenta": 5,
    "cyan":    6,
    "white":   7,
}

/*****************************************************************************\

`colorize` implements the function ColorString(...) which takes
in either a color "name" or a color ascii id.

\*****************************************************************************/
func ColorString(input, color string) string {
    color = strings.ToLower(color)
    if colorIndex, valid := colorToValueMap[color]; valid {
        return fmt.Sprintf("\033[3%dm%s\033[0m", colorIndex, input)
    }
    return input
}

/*****************************************************************************\

`colorize` implements the function Colorize(...) which takes in
a string embedded with <color></color>. Any un-matched tags will
be ignored!!

\*****************************************************************************/
func Colorize(input string) string {
    for color, index := range colorToValueMap {
        var (
            searchString  = fmt.Sprintf("(<%s>)(.*?)(</%s>)", color, color)
            replaceString = fmt.Sprintf("\033[3%dm$2\033[0m", index)
            match         = regexp.MustCompile(searchString)
        )
        input = match.ReplaceAllString(input, replaceString)
    }
    return input
}














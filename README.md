# Colorize

[![Build Status](https://travis-ci.org/sabhiram/colorize.svg)](https://travis-ci.org/sabhiram/colorize) [![Coverage Status](https://coveralls.io/repos/sabhiram/colorize/badge.png?branch=master)](https://coveralls.io/r/sabhiram/colorize?branch=master)

A `Go` library to fetch colorized ascii text

## Install

```shell
go get github.com/sabhiram/colorize
```

## (Sample) Usage

```shell
mkdir color_me
cd color_me
touch color_me.go
```

*color_me/color_me.go*:
```go
package main

import (
    "fmt"

    "github.com/sabhiram/colorize"
)

func main() {
    fmt.Println("Hello - " + colorize.Colorize("black",   "black"))
    fmt.Println("Hello - " + colorize.Colorize("red",     "red"))
    fmt.Println("Hello - " + colorize.Colorize("green",   "green"))
    fmt.Println("Hello - " + colorize.Colorize("yellow",  "yellow"))
    fmt.Println("Hello - " + colorize.Colorize("blue",    "blue"))
    fmt.Println("Hello - " + colorize.Colorize("magenta", "magenta"))
    fmt.Println("Hello - " + colorize.Colorize("cyan",    "cyan"))
    fmt.Println("Hello - " + colorize.Colorize("white",   "white"))
}
```

Install and run:
```shell
cd color_me
go install
color_me
```

Outputs:
![](https://raw.githubusercontent.com/sabhiram/public-images/master/colorize/colorize_sample.png)

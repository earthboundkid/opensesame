# opensesame [![GoDoc](https://godoc.org/github.com/carlmjohnson/opensesame?status.svg)](https://godoc.org/github.com/carlmjohnson/opensesame) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/opensesame)](https://goreportcard.com/report/github.com/carlmjohnson/opensesame)
Opensesame is a simple password generator. Passwords guaranteed random or double your money back.

## Installation
First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```shell
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/carlmjohnson/opensesame
```

## Screenshots
```shell
$ opensesame
OCo6X1Py

$ opensesame -h
Usage of opensesame [opts] [alphabet]:

        Creates a password by randomly selecting characters from its alphabet.

        Alphabet is a space separated list of character classes to use.
        At least one character in each class will be output.
        Character classes are either literal sets (like "abc" and "123") or the
        special names "upper", "lower", "digit", and "default".

        Default alphabet is "upper lower digit".

  -length int
        length of password to generate (default 8)

$ opensesame --length 4 '123 ABC xyz &%$'
&Cx3
```

## Web server
A web server for random passwords is also included as `open-sesame-web`. [See it online](https://open-sesame-web.herokuapp.com).

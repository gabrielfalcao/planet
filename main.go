package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

//	"github.com/gabrielfalcao/planet"
)

func GzipString(src string) string {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(src))
	w.Close()
	return b.String()
}

func FlattenStringArray(src []string) string {
	var buffer bytes.Buffer

	for _, i := range src {
		buffer.WriteString(i)
	}
	return buffer.String()
}

func main() {
	// let's make a simple
	flag.Parse()

	resp, _ := http.Get("http://yipit.com/")

	// creating a regex that matches empty lines
	lineRegexp, _ := regexp.Compile("^[ ]*$")
	spaceRegexp, _ := regexp.Compile("[ ]+")

	// closing the response right before returning :) Love go XOXO
	defer resp.Body.Close()

	// reading all the response body in an array of bytes
	body, _ := ioutil.ReadAll(resp.Body)

	// converting `body` which is an array of bytes, to string
	lines := strings.Split(string(body), "\n")
	numLines := len(lines)

	var result = make([]string, numLines)

	for _, line := range lines {
		if !lineRegexp.MatchString(line) {
			minified := spaceRegexp.ReplaceAllString(line, " ")
			trimmed := strings.Trim(minified, " ")
			result = append(result, trimmed)
		}
	}
	fmt.Printf(FlattenStringArray(result))
	fmt.Printf(GzipString(FlattenStringArray(result)))
}

// var width, height int

// fmt.Printf("\033[1;32mPlease type the width:\033[0m")
// fmt.Scanf("%d", &width)

// fmt.Printf("\033[1;32mPlease type the height:\033[0m")
// fmt.Scanf("%d", &height)

// r := planet.Rect{Width: width, Height: height}
// fmt.Println("My rect has the \033[1;33marea\033[0m of", r.Area())

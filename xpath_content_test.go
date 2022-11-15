package main

import (
	"bytes"
	"fmt"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/file"
	"github.com/antchfx/htmlquery"
	"regexp"
	"strings"
	"testing"
)

func TestXpath(t *testing.T) {
	xpath_root, _ := htmlquery.Parse(bytes.NewReader(file.Open("novel/18729784.html", "r", "")))
	content := regexp.MustCompile(`"text": "([^"]+)"`).
		FindAllStringSubmatch(htmlquery.FindOne(xpath_root, `/html/head/script[1]/text()`).Data, -1)
	fmt.Println(content)
	fmt.Println(len(content))
	print(strings.Replace(content[0][1], `"text"`+`:`+`"`, "", -1))

}

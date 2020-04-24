package utils

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type parser interface {
	Query(selector string) ([]string, error)
}

type HTMLParser struct {
	content string
}

func NewHTMLParser(content string) *HTMLParser {
	return &HTMLParser{
		content: content,
	}
}

func (this *HTMLParser) Query(selector string) (ret []string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(this.content))
	if err != nil {
		return
	}
	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		ret = append(ret, selection.Text())
	})
	return
}
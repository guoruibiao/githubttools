package resume

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/guoruibiao/gorequests"
	"strings"
)

type ProfileParser struct {
	Username string
	Homepage string
	Location string
}

func (this *ProfileParser) Parse(selector string) (err error) {
	resp, err := gorequests.NewRequest("GET", this.Homepage).DoRequest()
	if err != nil {
		return
	}
	content, _ := resp.Content()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return
	}
	location := doc.Find(selector).First().Text()
	this.Location = location
	return
}

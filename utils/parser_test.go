package utils

import (
	"github.com/guoruibiao/gorequests"
	"testing"
)
var htmlParser  *HTMLParser
var url string = "http://www.24timemap.com/search" // POST city=xxx
var requester *gorequests.Request
func init() {
	requester = gorequests.NewRequest("POST", url)
	payload := make(map[string]string)
	payload["city"] = "Berlin"
	resp, _ := requester.Form(payload).DoRequest()
	content, _ :=  resp.Content()
	// fmt.Println(content)
	htmlParser = NewHTMLParser(content)
}

func TestHTMLParser_Query(t *testing.T) {
	selector := ".timetable_cont>ul>li"
	ret, _ := htmlParser.Query(selector)
	t.Log(ret)
}

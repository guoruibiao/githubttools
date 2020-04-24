package timemap

import (
	"github.com/guoruibiao/githubtools/utils"
	"github.com/guoruibiao/gorequests"
)

var TIMEMAP_URL string = "http://www.24timemap.com/search"
var SELECTOR string = ".timetable_cont>ul>li"

type TimeMap struct {
	cityname string
	ret []string
}

func NewTimeMap(cityname string) *TimeMap {
	return &TimeMap{cityname: cityname}
}

func (this *TimeMap)DoQuery() (err error) {
	payload := make(map[string]string)
	payload["city"] = this.cityname
	response, err := gorequests.NewRequest("POST", TIMEMAP_URL).Form(payload).DoRequest()
	if err != nil {
		return
	}
	content, err := response.Content()
	if err != nil {
		return
	}
	htmlParser := utils.NewHTMLParser(content)
	ret, err := htmlParser.Query(SELECTOR)
	if err != nil {
		return
	}
	this.ret = ret
	return
}

func (this *TimeMap)GetCityTimes() []string {
	return this.ret
}

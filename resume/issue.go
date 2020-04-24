package resume

import (
	"errors"
	"strings"
)

// parsing user account from `https://github.com/simonhaenisch/md-to-pdf/issues/71`
type IssueParser struct {
	Full string
	Username string
	Homepage string
}

func (this *IssueParser)Parse(issueurl string) (err error) {
	if strings.HasPrefix(issueurl, "http://") || strings.HasPrefix(issueurl, "https://") {
		this.Full = issueurl
		this.Username = strings.Split(issueurl, "/")[3]
		this.Homepage = strings.Join(strings.Split(issueurl, "/")[0:4], "/")
	}else if strings.HasPrefix(issueurl, "github.com") {
		this.Full = "https://" + issueurl
		this.Username = strings.Split(issueurl, "/")[1]
		this.Homepage = strings.Join(strings.Split(issueurl, "/")[:2], "/")
	}else{
		err = errors.New("invalid issue url format.")
	}
	return
}

package resume

import "testing"
var issueParser *IssueParser = &IssueParser{}

func TestIssueParser_Parse(t *testing.T) {
	issueurl := "https://github.com/simonhaenisch/md-to-pdf/issues/71"
	issueParser.Parse(issueurl)
	t.Log("full: " + issueParser.Full)
	t.Log("username: " + issueParser.Username)
	t.Log("homepage: " + issueParser.Homepage)
}

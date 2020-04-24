package resume

import "testing"

func TestProfileParser_Parse(t *testing.T) {
	parser := &ProfileParser{
		Username: "simonhaenisch",
		Homepage: "https://github.com/simonhaenisch",
	}
	parser.Parse("li>span.p-label")
	t.Log(parser.Location)
}

package easy

import (
	"fmt"
	"testing"
)

func TestHtmlentities(t *testing.T) {
	html := "<h1>123</h1><b>23</b>"
	htmlentities := Htmlentities(html)
	fmt.Println(htmlentities)
	decode := Html_entity_decode(htmlentities)
	fmt.Println(decode)
}

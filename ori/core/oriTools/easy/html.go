package easy

import "html"

func Html_entity_decode(str string) string {
	return html.UnescapeString(str)
}

func Htmlentities(str string) string {
	return html.EscapeString(str)
}

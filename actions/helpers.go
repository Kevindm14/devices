package actions

import (
	"net/http"
	"strings"

	"github.com/dustin/go-humanize/english"
	"github.com/gobuffalo/plush"
)

func isPlural(count int, name string) string {
	return english.PluralWord(count, name, "")
}

func isActive(route string, help plush.HelperContext) string {
	request := help.Value("request").(*http.Request)
	requestURL := request.URL.String()

	if strings.Contains(requestURL, route) {
		return "active"
	}

	return ""
}

func order(column string, help plush.HelperContext) {

}

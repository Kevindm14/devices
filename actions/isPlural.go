package actions

import "github.com/dustin/go-humanize/english"

func isPlural(devices int, name string) string {
	if name == "Device" {
		return english.PluralWord(devices, "Device", "")
	}

	return english.PluralWord(devices, "User", "")
}

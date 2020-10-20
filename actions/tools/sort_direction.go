package tools

func SortDirection(sort string) string {
	directions := []string{"asc", "desc"}

	for _, direction := range directions {
		if direction == sort {
			return direction
		}
	}

	return "asc"
}

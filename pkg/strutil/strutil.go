package strutil

// find look for a string in a slice. Returns the position and true if found, -1 and false otherwise.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// removeDups remove duplicates.
func RemoveDups(slice []string) []string {
	keys := make(map[string]bool)
	newSlice := []string{}
	for _, key := range slice {
		if _, value := keys[key]; !value {
			keys[key] = true
			newSlice = append(newSlice, key)
		}
	}
	return newSlice
}

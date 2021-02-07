package main

import "regexp"

func getUUID(path string) string {
	re := regexp.MustCompile(`(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
	if m := re.FindStringSubmatch(path); len(m) > 0 {
		return m[1]
	}
	return ""
}

package ett

import (
	"fmt"
	"strings"
	"testing"
)

func TestInsert(t *testing.T) {
	urls := [...]string{"/hello/b/c", "/hello/*filepath"}

	parts := parsePatternTest(urls[1])
	fmt.Println(parts)

	pattern := urls[1]

	tireNode := &node{}

	tireNode.insert(pattern, parts, 0)

	searchUrls := "/hello/abc"
	search := parsePattern(searchUrls)
	result := tireNode.search(search, 0)
	if result != nil {
		t.Log("Yes, result:", result.pattern)
	} else {
		t.Log("No")
	}
}

func parsePatternTest(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

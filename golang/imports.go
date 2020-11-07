package golang

import (
	"fmt"
	"strings"
)

type imports struct {
	standard []string
	app      []string
	vendor   []string
}

func (i imports) hasImports() bool {
	return len(i.standard) > 0 || len(i.app) > 0 || len(i.vendor) > 0
}

func (i imports) Render() []byte {
	if !i.hasImports() {
		return []byte{}
	}

	appendSection := func(heap []string, section []string) []string {
		if len(section) > 0 {
			var sRes []string
			for _, i := range section {
				sRes = append(sRes, fmt.Sprintf(`	"%s"`, i))
			}
			heap = append(heap, strings.Join(sRes, "\n"))
		}
		return heap
	}

	var sectionImports []string
	sectionImports = appendSection(sectionImports, i.standard)
	sectionImports = appendSection(sectionImports, i.app)
	sectionImports = appendSection(sectionImports, i.vendor)

	result := []string{"import ("}

	// Separate each section with an additional line break.
	result = append(result, strings.Join(sectionImports, "\n\n"))

	result = append(result, ")")

	return []byte(strings.Join(result, "\n"))
}

func mergeImports(target, additional imports) imports {
	target.standard = append(target.standard, additional.standard...)
	target.app = append(target.app, additional.app...)
	target.vendor = append(target.vendor, additional.vendor...)
	return imports{
		standard: removeDuplicateStrings(target.standard),
		app:      removeDuplicateStrings(target.app),
		vendor:   removeDuplicateStrings(target.vendor),
	}
}

func removeDuplicateStrings(items []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, i := range items {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			result = append(result, i)
		}
	}
	return result
}

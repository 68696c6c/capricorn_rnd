package golang

import (
	"fmt"
	"strings"
)

type Imports struct {
	Standard []string
	App      []string
	Vendor   []string
}

func (i Imports) hasImports() bool {
	return len(i.Standard) > 0 || len(i.App) > 0 || len(i.Vendor) > 0
}

func (i Imports) Render() []byte {
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
	sectionImports = appendSection(sectionImports, i.Standard)
	sectionImports = appendSection(sectionImports, i.App)
	sectionImports = appendSection(sectionImports, i.Vendor)

	result := []string{"import ("}

	// Separate each section with an additional line break.
	result = append(result, strings.Join(sectionImports, "\n\n"))

	result = append(result, ")")

	return []byte(strings.Join(result, "\n"))
}

func mergeImports(target, additional Imports) Imports {
	target.Standard = append(target.Standard, additional.Standard...)
	target.App = append(target.App, additional.App...)
	target.Vendor = append(target.Vendor, additional.Vendor...)
	return Imports{
		Standard: removeDuplicateStrings(target.Standard),
		App:      removeDuplicateStrings(target.App),
		Vendor:   removeDuplicateStrings(target.Vendor),
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

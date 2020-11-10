package golang

import (
	"fmt"
	"strings"
)

type imports struct {
	standardImports []string
	appImports      []string
	vendorImports   []string
}

func newImports() *imports {
	return &imports{
		standardImports: []string{},
		appImports:      []string{},
		vendorImports:   []string{},
	}
}

func (i imports) hasImports() bool {
	return len(i.standardImports) > 0 || len(i.appImports) > 0 || len(i.vendorImports) > 0
}

func (i imports) Render() string {
	if !i.hasImports() {
		return ""
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
	sectionImports = appendSection(sectionImports, i.standardImports)
	sectionImports = appendSection(sectionImports, i.appImports)
	sectionImports = appendSection(sectionImports, i.vendorImports)

	result := []string{"import ("}

	// Separate each section with an additional line break.
	result = append(result, strings.Join(sectionImports, "\n\n"))

	result = append(result, ")")

	return strings.Join(result, "\n")
}

func (i *imports) AddImportsStandard(pkgImport ...string) {
	i.standardImports = append(i.standardImports, pkgImport...)
}

func (i *imports) AddImportsApp(pkgImport ...string) {
	i.appImports = append(i.appImports, pkgImport...)
}

func (i *imports) AddImportsVendor(pkgImport ...string) {
	i.vendorImports = append(i.vendorImports, pkgImport...)
}

func mergeImports(target, additional imports) imports {
	target.standardImports = append(target.standardImports, additional.standardImports...)
	target.appImports = append(target.appImports, additional.appImports...)
	target.vendorImports = append(target.vendorImports, additional.vendorImports...)
	return imports{
		standardImports: removeDuplicateStrings(target.standardImports),
		appImports:      removeDuplicateStrings(target.appImports),
		vendorImports:   removeDuplicateStrings(target.vendorImports),
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

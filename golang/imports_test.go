package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_imports_Render(t *testing.T) {
	input := Imports{
		Standard: []string{"standard-one", "standard-two"},
		App:      []string{"app-one", "app-two"},
		Vendor:   []string{"vendor-one", "vendor-two"},
	}

	result := input.Render()
	expected := `import (
	"standard-one"
	"standard-two"

	"app-one"
	"app-two"

	"vendor-one"
	"vendor-two"
)`

	assert.Equal(t, expected, string(result))
}

func Test_imports_Render_none(t *testing.T) {
	input := Imports{}

	result := input.Render()
	expected := ""

	assert.Equal(t, expected, string(result))
}

func Test_imports_Render_noStandard(t *testing.T) {
	input := Imports{
		App:    []string{"app-one", "app-two"},
		Vendor: []string{"vendor-one", "vendor-two"},
	}

	result := input.Render()
	expected := `import (
	"app-one"
	"app-two"

	"vendor-one"
	"vendor-two"
)`

	assert.Equal(t, expected, string(result))
}

func Test_imports_Render_noApp(t *testing.T) {
	input := Imports{
		Standard: []string{"standard-one", "standard-two"},
		Vendor:   []string{"vendor-one", "vendor-two"},
	}

	result := input.Render()
	expected := `import (
	"standard-one"
	"standard-two"

	"vendor-one"
	"vendor-two"
)`

	assert.Equal(t, expected, string(result))
}

func Test_imports_Render_noVendor(t *testing.T) {
	input := Imports{
		Standard: []string{"standard-one", "standard-two"},
		App:      []string{"app-one", "app-two"},
	}

	result := input.Render()
	expected := `import (
	"standard-one"
	"standard-two"

	"app-one"
	"app-two"
)`

	assert.Equal(t, expected, string(result))
}

func Test_mergeImports(t *testing.T) {
	stack := Imports{
		Standard: []string{"one"},
		App:      []string{"one", "two"},
		Vendor:   []string{"one", "two", "three"},
	}

	additional := Imports{
		Standard: []string{},
		App:      []string{},
		Vendor:   []string{"one", "two", "three", "four"},
	}

	stack = mergeImports(stack, additional)

	assert.Len(t, stack.Standard, 1)
	assert.Len(t, stack.App, 2)
	assert.Len(t, stack.Vendor, 4)
}

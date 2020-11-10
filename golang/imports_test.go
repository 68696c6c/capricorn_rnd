package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_imports_Render(t *testing.T) {
	input := imports{
		standardImports: []string{"standard-one", "standard-two"},
		appImports:      []string{"app-one", "app-two"},
		vendorImports:   []string{"vendor-one", "vendor-two"},
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
	input := imports{}

	result := input.Render()
	expected := ""

	assert.Equal(t, expected, string(result))
}

func Test_imports_Render_noStandard(t *testing.T) {
	input := imports{
		appImports:    []string{"app-one", "app-two"},
		vendorImports: []string{"vendor-one", "vendor-two"},
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
	input := imports{
		standardImports: []string{"standard-one", "standard-two"},
		vendorImports:   []string{"vendor-one", "vendor-two"},
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
	input := imports{
		standardImports: []string{"standard-one", "standard-two"},
		appImports:      []string{"app-one", "app-two"},
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
	stack := imports{
		standardImports: []string{"one"},
		appImports:      []string{"one", "two"},
		vendorImports:   []string{"one", "two", "three"},
	}

	additional := imports{
		standardImports: []string{},
		appImports:      []string{},
		vendorImports:   []string{"one", "two", "three", "four"},
	}

	stack = mergeImports(stack, additional)

	assert.Len(t, stack.standardImports, 1)
	assert.Len(t, stack.appImports, 2)
	assert.Len(t, stack.vendorImports, 4)
}

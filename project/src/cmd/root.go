package cmd

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func buildRoot(pkg golang.IPackage, o config.CmdOptions, p *config.Project) {
	file := pkg.AddGoFile(o.RootFileName)

	rootVar := makeRootVar(o.RootVarName, o.RootCommandUse, p.Name)
	file.AddVar(rootVar)

	rootFunc := makeRootFunc(p.Author, p.License)
	file.AddFunction(rootFunc)
}

func makeRootFunc(author config.AuthorMeta, license string) *golang.Function {
	result := golang.NewFunction("init")
	t := `
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("author", "{{ .AuthorName }} <{{ .AuthorEmail }}>")
	viper.SetDefault("license", "{{ .License }}")
`
	result.SetBodyTemplate(t, struct {
		AuthorName  string
		AuthorEmail string
		License     string
	}{
		AuthorName:  author.Name,
		AuthorEmail: author.Email,
		License:     license,
	})

	result.AddImportsStandard("strings")
	result.AddImportsVendor(goat.ImportViper)

	result.AddImportsVendor(goat.ImportCobra)

	return result
}

func makeRootVar(rootVarName, commandUse, projectName string) *golang.Var {
	result := golang.NewVar(rootVarName, "", goat.MakeTypeCobraCommand(), false)
	t := `&cobra.Command{
	Use:   "{{ .CommandUsage }}",
	Short: "Root command for {{ .ProjectName }}",
}`
	result.SetValueTemplate(t, struct {
		CommandUsage string
		ProjectName  string
	}{
		CommandUsage: commandUse,
		ProjectName:  projectName,
	})
	return result
}

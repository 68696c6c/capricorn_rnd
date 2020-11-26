package cmd

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
)

func buildRoot(pkg golang.IPackage, meta *config.CmdMeta) {
	file := pkg.AddGoFile(meta.RootFileName)

	rootVar := makeRootVar(meta)
	file.AddVar(rootVar)

	rootFunc := makeRootFunc(meta)
	file.AddFunction(rootFunc)
}

func makeRootFunc(meta *config.CmdMeta) *golang.Function {
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
		AuthorName:  meta.Author.Name,
		AuthorEmail: meta.Author.Email,
		License:     meta.License,
	})

	result.AddImportsStandard("strings")
	result.AddImportsVendor(goat.ImportViper)

	result.AddImportsVendor(goat.ImportCobra)

	return result
}

func makeRootVar(meta *config.CmdMeta) *golang.Var {
	result := golang.NewVar(meta.RootVarName, "", goat.MakeTypeCobraCommand(), false)
	t := `&cobra.Command{
	Use:   "{{ .CommandUsage }}",
	Short: "Root command for {{ .ProjectName }}",
}`
	result.SetValueTemplate(t, struct {
		CommandUsage string
		ProjectName  string
	}{
		CommandUsage: meta.RootCommandUse,
		ProjectName:  meta.ProjectName,
	})
	return result
}

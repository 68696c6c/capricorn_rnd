package cmd

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/http"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func Build(pkg golang.IPackage, p *config.Project, o config.CmdOptions, a *app.App, h *http.Http) {
	pkgCmd := pkg.AddPackage(o.PkgName)

	buildRoot(pkgCmd, o, p)
	buildServer(pkgCmd, o, p.Name, a, h)
	buildMigrate(pkgCmd, o)

	for _, c := range p.Commands {
		cmdFile := pkgCmd.AddGoFile(c.Name)
		cmdFunc := golang.NewFunction("")
		cmdFunc.AddArg(o.CmdArgName, goat.MakeTypeCobraCommand())
		cmdFunc.AddArg(o.ArgsArgName, golang.MakeTypeStringSlice(false))
		cmdFile.AddFunction(makeCommandFunc(commandFuncMeta{
			rootVarName: o.RootVarName,
			use:         c.Name,
			runFunc:     cmdFunc,
		}))
	}
}

type commandFuncMeta struct {
	rootVarName string
	use         string
	short       string
	long        string
	example     string
	runFunc     *golang.Function
}

func makeCommandFunc(meta commandFuncMeta) *golang.Function {
	result := golang.NewFunction("init")
	t := `
	{{ .RootCommandName }}.AddCommand(&cobra.Command{
		Use:   "{{ .Use }}",
		Short: "{{ .Short }}",{{ .Long }}{{ .Example }}
		Run: {{ .RunFunc.Render }},
	})
`
	long := ""
	if meta.long != "" {
		long = "\n		Long: `" + meta.long + "`,"
	}
	example := ""
	if meta.example != "" {
		example = "\n		Example: `" + meta.example + "`,"
	}
	result.SetBodyTemplate(t, struct {
		RootCommandName string
		Use             string
		Short           string
		Long            string
		RunFunc         utils.Renderable
		Example         string
	}{
		RootCommandName: meta.rootVarName,
		Use:             meta.use,
		Short:           meta.short,
		Long:            long,
		RunFunc:         meta.runFunc,
		Example:         example,
	})
	result.AddImportsVendor(goat.ImportCobra)
	golang.MergeImports(result, meta.runFunc)
	return result
}

package cmd

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Command struct {
	*golang.File
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

func Build(pkg golang.IPackage, commands []Command, meta *config.CmdMeta) {
	pkgCmd := pkg.AddPackage("cmd")

	buildRoot(pkgCmd, meta)
	buildServer(pkgCmd, meta)
	buildMigrate(pkgCmd, meta)

	for _, c := range commands {
		cmdFile := pkgCmd.AddGoFile(c.Name)
		cmdFunc := golang.NewFunction("")
		cmdFunc.AddArg("cmd", goat.MakeTypeCobraCommand())
		cmdFunc.AddArg("args", golang.MakeTypeStringSlice(false))
		cmdFile.AddFunction(makeCommandFunc(commandFuncMeta{
			rootVarName: meta.RootVarName,
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

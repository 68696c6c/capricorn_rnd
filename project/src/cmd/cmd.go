package cmd

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/http"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Commands struct {
	*golang.Package
	commands []*golang.Var
}

func (c *Commands) Render() string {
	result := []string{""}
	for _, command := range c.commands {
		result = append(result, fmt.Sprintf("		%s,", command.GetReference()))
	}
	return strings.Join(result, "\n")
}

func Build(pkg golang.IPackage, p *config.Project, o config.CmdOptions, a *app.App, h *http.Http) *Commands {
	pkgCmd := pkg.AddPackage(o.PkgName)

	serverVar := buildServer(pkgCmd, o, p.Name, a, h)
	migrateVar := buildMigrate(pkgCmd, o)

	commands := []*golang.Var{serverVar, migrateVar}
	for _, c := range p.Commands {
		cmdFile := pkgCmd.AddGoFile(c.Name)
		cmdFunc := golang.NewFunction("")
		cmdFunc.AddArg(o.CmdArgName, goat.MakeTypeCobraCommand())
		cmdFunc.AddArg(o.ArgsArgName, golang.MakeTypeStringSlice(false))
		cmdVar := makeCommandVar(commandFuncMeta{
			name:    utils.Pascal(c.Name),
			use:     c.Name,
			runFunc: cmdFunc,
		})
		cmdFile.AddVar(cmdVar)
		commands = append(commands, cmdVar)
	}

	return &Commands{
		Package:  pkgCmd,
		commands: commands,
	}
}

type commandFuncMeta struct {
	name    string
	use     string
	short   string
	long    string
	example string
	runFunc *golang.Function
}

func makeCommandVar(meta commandFuncMeta) *golang.Var {
	result := golang.NewVar(meta.name, "", goat.MakeTypeCobraCommand(), false)
	t := `&cobra.Command{
	Use:   "{{ .Use }}",
	Short: "{{ .Short }}",{{ .Long }}{{ .Example }}
	Run: {{ .RunFunc.Render }},
}
`
	long := ""
	if meta.long != "" {
		long = "\n	Long: `" + meta.long + "`,"
	}
	example := ""
	if meta.example != "" {
		example = "\n	Example: `" + meta.example + "`,"
	}
	result.SetValueTemplate(t, struct {
		Use     string
		Short   string
		Long    string
		RunFunc utils.Renderable
		Example string
	}{
		Use:     meta.use,
		Short:   meta.short,
		Long:    long,
		RunFunc: meta.runFunc,
		Example: example,
	})
	result.AddImportsVendor(goat.ImportCobra)
	golang.MergeImports(result, meta.runFunc)
	return result
}

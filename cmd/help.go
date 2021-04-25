package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type Command = cobra.Command

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "Help about commands",
	ValidArgsFunction: func(c *Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var completions []string
		cmd, _, e := c.Root().Find(args)
		if e != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		if cmd == nil {
			// Root help command.
			cmd = c.Root()
		}
		for _, subCmd := range cmd.Commands() {
			if subCmd.IsAvailableCommand() {
				if strings.HasPrefix(subCmd.Name(), toComplete) {
					completions = append(completions, fmt.Sprintf("%s\t%s", subCmd.Name(), subCmd.Short))
				}
			}
		}
		return completions, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(c *Command, args []string) {
		cmd, _, e := c.Root().Find(args)
		if cmd == nil || e != nil {
			c.Printf("Unknown help topic %#q\n", args)
			cobra.CheckErr(c.Root().Usage())
		} else {
			cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
			cobra.CheckErr(cmd.Help())
		}
	},
}

const helpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}
{{if or .Runnable .HasSubCommands}}
{{.UsageString}}
{{end}}`

const usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`

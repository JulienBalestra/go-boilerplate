package cli

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/JulienBalestra/go-boilerplate/pkg/example"
)

const programName = "go-boilerplate"

var viperConfig = viper.New()

func NewCommand() (*cobra.Command, *int) {
	var verbose int
	var exitCode int

	rootCommand := &cobra.Command{
		Use:   fmt.Sprintf("%s command line", programName),
		Short: "Use this command to do something",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			flag.Lookup("alsologtostderr").Value.Set("true")
			flag.Lookup("v").Value.Set(strconv.Itoa(verbose))
		},
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(`
# generate the private key and the csr
%s my-argument --super-flag my-flag 
`, programName),
		Run: func(cmd *cobra.Command, args []string) {

			e := newExamplePackage(args[0])
			err := e.Hello()
			if err != nil {
				glog.Errorf("Command returns error: %v", err)
				exitCode = 2
				return
			}
		},
	}

	rootCommand.PersistentFlags().IntVarP(&verbose, "verbose", "v", 2, "verbose level")

	viperConfig.SetDefault("super-flag", "")
	rootCommand.PersistentFlags().String("super-flag", viperConfig.GetString("super-flag"), "This is a super flag")
	viperConfig.BindPFlag("super-flag", rootCommand.PersistentFlags().Lookup("super-flag"))

	return rootCommand, &exitCode
}

func newExamplePackage(arg string) *example.Example {
	return &example.Example{
		Argument:  arg,
		SuperFlag: viperConfig.GetString("super-flag"),
	}
}

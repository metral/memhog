package cmd

import (
	"fmt"
	"os"
	"strings"

	goflag "flag"

	"github.com/golang/glog"
	"github.com/metral/memhog/pkg/memhog"
	"github.com/spf13/cobra"
)

var (
	cmdName = "memhog"
	usage   = fmt.Sprintf("%s -c config.toml", cmdName)
)

// Define a type for the options of MemHog
type MemHogOptions struct {
	ConfigFile string
}

func AddConfigFileFlag(cmd *cobra.Command, value *string) {
	cmd.Flags().StringVarP(value, "config", "c", *value, "<insert description>")
}

// Fatal prints the message (if provided) and then exits. If V(2) or greater,
// glog.Fatal is invoked for extended information.
func fatal(msg string) {
	if glog.V(2) {
		glog.FatalDepth(2, msg)
	}
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(1)
}

// NewCmdOptions creates an options Cobra command to return usage
func NewCmdOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use: "options",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	return cmd
}

// Create a new command for the memhog. This cmd includes logging,
// and cmd option parsing from flags.
func NewCmdMemHog() (*cobra.Command, error) {
	// Define the options for MemHog command
	options := MemHogOptions{}

	// Create a new command
	cmd := &cobra.Command{
		Use:   usage,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(RunMemHog(cmd, &options), fatal)
		},
	}

	// Bind & parse flags defined by external projects.
	// e.g. This imports the golang/glog pkg flags into the cmd flagset
	cmd.Flags().AddGoFlagSet(goflag.CommandLine)
	goflag.CommandLine.Parse([]string{})

	// Define the flags allowed in this command & store each option provided
	// as a flag, into the MemHogOptions
	AddConfigFileFlag(cmd, &options.ConfigFile)

	return cmd, nil
}

func checkErr(err error, handleErr func(string)) {
	if err == nil {
		return
	}
	handleErr(err.Error())
}

// Run the memhog
func RunMemHog(cmd *cobra.Command, options *MemHogOptions) error {
	return run(options)
}

func run(options *MemHogOptions) error {
	// Create a new RAM memhog
	bufferLen := 20
	mh, err := memhog.NewMemHog(bufferLen)
	if err != nil {
		return err
	}

	// Create a slice of buffers and randomly alloc its elements with data to
	// increase & throttle RAM usage
	mh.HogMemory()

	return nil
}

// Print & log a message
func log(msg string) {
	glog.V(2).Infof("%s", msg)
}

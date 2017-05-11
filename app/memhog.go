package app

import (
	"github.com/metral/memhog/pkg/cmd"
	k8slogsutil "k8s.io/kubernetes/pkg/util/logs"
)

// Run the memhog command
func Run() error {
	// Init logging
	k8slogsutil.InitLogs()
	defer k8slogsutil.FlushLogs()

	// Create & execute new command
	cmd, err := cmd.NewCmdMemHog()
	if err != nil {
		return err
	}

	return cmd.Execute()
}

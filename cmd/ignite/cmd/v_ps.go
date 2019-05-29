package cmd

import (
	"github.com/luxas/ignite/cmd/ignite/run"
	"github.com/luxas/ignite/pkg/errutils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
)

// NewCmdVMPs lists running VMs
func NewCmdVMPs(out io.Writer) *cobra.Command {
	po := &run.PsOptions{}

	cmd := &cobra.Command{
		Use:     "ps",
		Short:   "List running VMs",
		Aliases: []string{"ls", "list"},
		Run: func(cmd *cobra.Command, args []string) {
			errutils.Check(func() error {
				var err error
				if po.VMs, err = matchAllVMs(po.All); err != nil {
					return err
				}
				return run.Ps(po)
			}())
		},
	}

	addVMPsFlags(cmd.Flags(), po)
	return cmd
}

func addVMPsFlags(fs *pflag.FlagSet, po *run.PsOptions) {
	fs.BoolVarP(&po.All, "all", "a", false, "Show all VMs, not just running ones")
}
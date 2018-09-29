package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var vmStartCmd = &cobra.Command{
	Use:   "start <vm name> [vm name] ...",
	Short: "Start virtual machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		errs := []error{}
		for _, v := range args {
			if err := startVirtualMachine(v); err != nil {
				errs = append(errs, fmt.Errorf("could not start %q: %s", v, err))
			}
		}

		if len(errs) == 1 {
			return errs[0]
		}
		if len(errs) > 1 {
			var b strings.Builder
			for _, err := range errs {
				if _, e := fmt.Fprintln(&b, err); e != nil {
					return e
				}
			}
			return errors.New(b.String())
		}

		return nil
	},
}

// startVirtualMachine start a virtual machine instance Async
func startVirtualMachine(vmName string) error {
	vm, err := getVMWithNameOrID(vmName)
	if err != nil {
		return err
	}

	state := (string)(egoscale.VirtualMachineStopped)
	if vm.State != state {
		return fmt.Errorf("%q is not in a %s state, got: %s", vmName, state, vm.State)
	}

	_, err = asyncRequest(&egoscale.StartVirtualMachine{ID: vm.ID}, fmt.Sprintf("Starting %q ", vm.Name))
	return err
}

func init() {
	vmCmd.AddCommand(vmStartCmd)
}

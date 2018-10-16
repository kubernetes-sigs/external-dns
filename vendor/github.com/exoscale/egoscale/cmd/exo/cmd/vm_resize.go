package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// vmResetCmd represents the stop command
var vmResizeCmd = &cobra.Command{
	Use:   "resize <vm name> [vm name] ...",
	Short: "resize disk virtual machine instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		diskValue, err := cmd.Flags().GetInt64("disk")
		if err != nil {
			return err
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return err
		}

		errs := []error{}
		for _, v := range args {
			if err := resizeVirtualMachine(v, diskValue, force); err != nil {
				errs = append(errs, fmt.Errorf("could not resize %q: %s", v, err))
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

// resizeVirtualMachine stop a virtual machine instance
func resizeVirtualMachine(vmName string, diskValue int64, force bool) error {
	vm, err := getVMWithNameOrID(vmName)
	if err != nil {
		return err
	}

	state := (string)(egoscale.VirtualMachineStopped)
	if vm.State != state {
		return fmt.Errorf("this operation is not permitted if your VM is not stopped")
	}

	if !force {
		if !askQuestion(fmt.Sprintf("sure you want to resize %q virtual machine", vm.Name)) {
			return nil
		}
	}

	volume := &egoscale.Volume{
		VirtualMachineID: vm.ID,
		Type:             "ROOT",
	}

	if err = cs.GetWithContext(gContext, volume); err != nil {
		return err
	}

	_, err = asyncRequest(&egoscale.ResizeVolume{ID: volume.ID, Size: diskValue}, fmt.Sprintf("Resizing %q ", vm.Name))
	return err
}

func init() {
	vmCmd.AddCommand(vmResizeCmd)
	vmResizeCmd.Flags().Int64P("disk", "d", 0, "Disk size in GB")
	vmResizeCmd.Flags().BoolP("force", "f", false, "Attempt to resize vitual machine without prompting for confirmation")
	if err := vmResizeCmd.MarkFlagRequired("disk"); err != nil {
		log.Fatal(err)
	}
}

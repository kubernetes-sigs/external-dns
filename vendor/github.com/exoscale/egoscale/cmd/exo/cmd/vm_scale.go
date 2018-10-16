package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/exoscale/egoscale"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var vmScaleCmd = &cobra.Command{
	Use:   "scale <vm name> [vm name] ...",
	Short: "Scale virtual machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		so, err := cmd.Flags().GetString("service-offering")
		if err != nil {
			return err
		}

		serviceoffering, err := getServiceOfferingByName(so)
		if err != nil {
			return err
		}

		errs := []error{}
		for _, v := range args {
			if err := scaleVirtualMachine(v, *serviceoffering.ID); err != nil {
				errs = append(errs, fmt.Errorf("could not scale %q: %s", v, err))
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

// scaleVirtualMachine scale a virtual machine instance Async with context
func scaleVirtualMachine(vmName string, serviceofferingID egoscale.UUID) error {
	vm, err := getVMWithNameOrID(vmName)
	if err != nil {
		return err
	}

	if vm.State != (string)(egoscale.VirtualMachineStopped) {
		return fmt.Errorf("this operation is not permitted if your VM is not stopped")
	}

	_, err = asyncRequest(&egoscale.ScaleVirtualMachine{ID: vm.ID, ServiceOfferingID: &serviceofferingID}, fmt.Sprintf("Scaling %q ", vm.Name))
	return err
}

func init() {
	vmCmd.AddCommand(vmScaleCmd)
	vmScaleCmd.Flags().StringP("service-offering", "o", "", "<name | id> (micro|tiny|small|medium|large|extra-large|huge|mega|titan")
	if err := vmScaleCmd.MarkFlagRequired("service-offering"); err != nil {
		log.Fatal(err)
	}
}

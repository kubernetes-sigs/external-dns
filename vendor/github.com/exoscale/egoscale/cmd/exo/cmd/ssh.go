package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh <vm name | id>",
	Short: "SSH into a virtual machine instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}

		ipv6, err := cmd.Flags().GetBool("ipv6")
		if err != nil {
			return err
		}

		isInfo, err := cmd.Flags().GetBool("info")
		if err != nil {
			return err
		}

		isConnectionSTR, err := cmd.Flags().GetBool("print")
		if err != nil {
			return err
		}

		info, err := getSSHInfo(args[0], ipv6)
		if err != nil {
			return err
		}

		if isConnectionSTR {
			printSSHConnectSTR(info)
			return nil
		}

		if isInfo {
			printSSHInfo(info)
			return nil
		}
		return connectSSH(info)
	},
}

type sshInfo struct {
	sshKeys  string
	userName string
	ip       net.IP
	vmName   string
	vmID     *egoscale.UUID
}

func getSSHInfo(name string, isIpv6 bool) (*sshInfo, error) {
	vm, err := getVMWithNameOrID(name)
	if err != nil {
		return nil, err
	}

	sshKeyPath := path.Join(gConfigFolder, "instances", vm.ID.String(), "id_rsa")

	nic := vm.DefaultNic()
	if nic == nil {
		return nil, fmt.Errorf("this instance %q has no default NIC", vm.ID)
	}

	vmIP := vm.IP()

	if isIpv6 {
		if nic.IP6Address == nil {
			return nil, fmt.Errorf("missing IPv6 address on the instance %q", vm.ID)
		}
		vmIP = &nic.IP6Address
	}

	if vmIP == nil {
		return nil, fmt.Errorf("no valid IP address found")
	}

	template := &egoscale.Template{
		ID:         vm.TemplateID,
		IsFeatured: true,
		ZoneID:     vm.ZoneID,
	}

	if err := cs.GetWithContext(gContext, template); err != nil {
		return nil, err
	}

	tempUser, ok := template.Details["username"]
	if !ok {
		return nil, fmt.Errorf("missing username information in Template %q", template.ID)
	}

	return &sshInfo{
		sshKeys:  sshKeyPath,
		userName: tempUser,
		ip:       *vmIP,
		vmName:   vm.Name,
		vmID:     vm.ID,
	}, nil

}

func printSSHConnectSTR(info *sshInfo) {
	sshArgs := ""

	if _, err := os.Stat(info.sshKeys); err == nil {
		sshArgs = fmt.Sprintf("-i %q ", info.sshKeys)
	}

	fmt.Printf("ssh %s%s@%s\n", sshArgs, info.userName, info.ip)
}

func printSSHInfo(info *sshInfo) {
	fmt.Println("Host", info.vmName)
	fmt.Println("\tHostName", info.ip.String())
	fmt.Println("\tUser", info.userName)
	if _, err := os.Stat(info.sshKeys); err == nil {
		fmt.Println("\tIdentityFile", info.sshKeys)
	}
}

func connectSSH(info *sshInfo) error {

	args := make([]string, 0, 3)

	if _, err := os.Stat(info.sshKeys); os.IsNotExist(err) {
		log.Printf("Warning: Identity file %s not found or not accessible.", info.sshKeys)
	} else {
		args = append(args, "-i")
		args = append(args, info.sshKeys)
	}

	args = append(args, info.userName+"@"+info.ip.String())

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func init() {
	sshCmd.Flags().BoolP("info", "i", false, "Print SSH config information")
	sshCmd.Flags().BoolP("print", "p", false, "Print SSH command")
	sshCmd.Flags().BoolP("ipv6", "6", false, "Use IPv6 for SSH")
	RootCmd.AddCommand(sshCmd)
}

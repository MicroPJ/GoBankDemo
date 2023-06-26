package gobankdemo

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Vsam(input []string) (result string) {

	var cmd *exec.Cmd

	fmt.Printf("*---[One] Running\n")
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("*---[VSAM] Windows identified\n")

	default: //Mac & Linux
		fmt.Printf("*---[VSAM] Linux identified\n")
	}

	if len(input) == 0 {
		fmt.Printf("*---[VSAM] No Parameter provided\n")
	} else {
		fmt.Printf("*---[VSAM] Parameter provided: %v\n", input)
	}

	//Clone BankDemo Repo
	fmt.Printf("*---[VSAM] Start Clone BankDemo GitHub.com Repo\n")
	var repo = "https://github.com/MicroFocus/BankDemo.git"
	cmd = exec.Command("git", "clone", repo, "--progress")
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	fmt.Printf("*---[VSAM] End Clone BankDemo GitHub.com Repo\n")

	//Run python MF_Provision_Region.py vsam
	fmt.Printf("*---[VSAM] Start python MF_Provision_Region.py vsam\n")

	cmd = exec.Command("python", "MF_Provision_Region.py", "vsam")
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	fmt.Printf("*---[VSAM] End python MF_Provision_Region.py vsam\n")
	fmt.Printf("*---[VSAM] Completed\n")
	return "Done"
}

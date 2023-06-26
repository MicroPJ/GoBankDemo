package gobankdemo

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"syscall"
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

	//Delete BankDemo Folder
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-nologo", "-noprofile")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			fmt.Fprintln(stdin, "Remove-Item 'BankDemo' -Force -Recurse")
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			//log.Fatal(err)
		}
		_ = out
		//fmt.Printf("%s\n", out)
	default: //Mac & Linux
		fmt.Printf("*---[VSAM] Delete BankDemo Folder in Linux not yet implemented\n")
	}

	//Clone BankDemo Repo
	fmt.Printf("*---[VSAM] Start Clone BankDemo GitHub.com Repo\n")
	var repo = "https://github.com/MicroFocus/BankDemo.git"
	cmd = exec.Command("git", "clone", repo, "--progress")
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	fmt.Printf("*---[VSAM] End Clone BankDemo GitHub.com Repo\n")

	//Fix Github typos
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-nologo", "-noprofile")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			fmt.Fprintln(stdin, "(Get-Content -path BankDemo\\scripts\\config\\ports.json -Raw) -replace '9023', '3270' | Set-Content -Path BankDemo\\scripts\\config\\ports.json")
			fmt.Fprintln(stdin, "(Get-Content -path BankDemo\\scripts\\config\\ports.json -Raw) -replace '8001', '5001 | Set-Content -Path BankDemo\\scripts\\config\\ports.json")
			fmt.Fprintln(stdin, "(Get-Content -path BankDemo\\scripts\\options\\vsam.json -Raw) -replace '\"is64bit\": false,','\"is64bit\": true,' | Set-Content -Path BankDemo\\scripts\\options\\vsam.json")
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		_ = out
		//fmt.Printf("%s\n", out)
	default: //Mac & Linux
		fmt.Printf("*---[VSAM] Fix Github typos in Linux not yet implemented\n")
	}

	//Run python MF_Provision_Region.py vsam
	fmt.Printf("*---[VSAM] Start python MF_Provision_Region.py vsam\n")
	syscall.Chdir("BankDemo\\scripts\\")
	cmd = exec.Command("python", "MF_Provision_Region.py", "vsam")
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	fmt.Printf("*---[VSAM] End python MF_Provision_Region.py vsam\n")
	fmt.Printf("*---[VSAM] Completed\n")
	return "ESCWA: http://localhost:10086\n3270: localhost:3270\n"
}

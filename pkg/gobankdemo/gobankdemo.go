package gobankdemo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func Vsam(input []string, verbose bool) (result string) {

	var cmd *exec.Cmd

	fmt.Printf("*---[VSAM] Running\n")
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

	if verbose {
		fmt.Printf("*---[VSAM] Verbose true\n")
	} else {
		fmt.Printf("*---[VSAM] Verbose false\n")
	}

	//Delete BankDemo Folder
	fmt.Printf("*---[VSAM] Checking for left over BankDemo clone folder\n")
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
		out, _ := cmd.CombinedOutput()

		//_ = out
		if verbose {
			fmt.Printf("*---[VSAM] %s\n", out)
		}
		if err != nil {
			fmt.Printf("*---[VSAM] BankDemo clone folder not found\n")
		} else {
			fmt.Printf("*---[VSAM] BankDemo clone folder found & deleted\n")
		}
		//fmt.Printf("%s\n", out)
	default: //Mac & Linux
		fmt.Printf("*---[VSAM] Delete BankDemo Folder in Linux not yet implemented\n")
	}
	fmt.Printf("*---[VSAM] End Checking for left over BankDemo clone folder\n")

	//Clone BankDemo Repo
	fmt.Printf("*---[VSAM] Start Clone BankDemo GitHub.com Repo\n")

	var repo = "https://github.com/MicroFocus/BankDemo.git"
	cmd = exec.Command("git", "clone", repo, "--progress")
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Error executing command: %s......\n", err.Error())
		return err.Error()
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......\n", err.Error())
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
			fmt.Fprintln(stdin, "(Get-Content -path BankDemo\\scripts\\config\\ports.json -Raw) -replace '9023', '9023' | Set-Content -Path BankDemo\\scripts\\config\\ports.json")
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
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	fmt.Printf("*---[VSAM] End python MF_Provision_Region.py vsam\n")

	//Run HASESSION Server
	//fmt.Printf("*---[VSAM] Start setup environment for HASession Server\n")
	//syscall.Chdir("C:\\Program Files (x86)\\Micro Focus\\Enterprise Developer")
	//cmd = exec.Command("setupenv.bat", "", "")
	//if verbose {
	//	cmd.Stdout = os.Stdout
	//	cmd.Stderr = os.Stderr
	//}
	//if err := cmd.Run(); err != nil {
	//	//return err.Error()
	//}
	//fmt.Printf("*---[VSAM] End setup environment for HASession Server\n")

	fmt.Printf("*---[VSAM] Start HASession Server\n")
	//syscall.Chdir("C:\\Program Files (x86)\\Micro Focus\\Enterprise Developer")
	//cmd = exec.Command("hacloudserviceinstall 64", "", "")
	//if verbose {
	//	cmd.Stdout = os.Stdout
	//	cmd.Stderr = os.Stderr
	//}
	//if err := cmd.Run(); err != nil {
	//return err.Error()
	//}
	//cmd = exec.Command("net start mfhacloud", "", "")
	//if verbose {
	//	cmd.Stdout = os.Stdout
	//	cmd.Stderr = os.Stderr
	//}
	//if err := cmd.Run(); err != nil {
	//	return err.Error()
	//}
	//fmt.Printf("*---[VSAM] End HASession Server\n")

	//fmt.Printf("*---[VSAM] Completed\n")
	//return "ESCWA: http://localhost:10086\n3270: localhost:9023\nHA: localhost:7443\n"
	return "ESCWA: http://localhost:10086\n3270: localhost:9023\n"
}

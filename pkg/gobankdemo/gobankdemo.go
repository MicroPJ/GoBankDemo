package gobankdemo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/common-nighthawk/go-figure"
	"golang.org/x/exp/slices"
)

func Deploy(input []string, verbose bool) (result string) {

	var cmd *exec.Cmd
	var option string
	if len(input) < 1 {
		option = "vsam"
		fmt.Printf("*---[%v] No Option provided, using default VSAM\n", input)
	} else {
		option = input[0]
	}

	valid_options := []string{"vsam", "vsam_postgres", "sql_postgres"}
	if slices.Contains(valid_options, option) {
		fmt.Printf("*---[%v] Valid Option\n", input)
	} else {
		err1 := errors.New("[ERROR]: INVALID OPTION. Valid options are ['vsam', 'vsam_postgres', 'sql_postgres']")
		log.Fatal(err1)
	}

	fmt.Printf("*---[%v] Running\n", input)
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("*---[%v] Windows identified\n", input)

	default: //Mac & Linux
		fmt.Printf("*---[%v] Linux identified\n", input)
	}

	if len(input) == 0 {
		fmt.Printf("*---[%v] No Parameter provided\n", input)
	} else {
		fmt.Printf("*---[%v] Parameter provided: %v\n", input)
	}

	if verbose {
		fmt.Printf("*---[%v] Verbose true\n", input)
	} else {
		fmt.Printf("*---[%v] Verbose false\n", input)
	}

	//Delete BankDemo Folder
	fmt.Printf("*---[%v] Checking for left over BankDemo clone folder\n", input)
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
			fmt.Printf("*---[%v] %s\n", input, out)
		}
		if err != nil {
			fmt.Printf("*---[%v] BankDemo clone folder not found\n", input)
		} else {
			fmt.Printf("*---[%v] BankDemo clone folder found & deleted\n", input)
		}
		//fmt.Printf("%s\n", out)
	default: //Mac & Linux
		fmt.Printf("*---[%v] Delete BankDemo Folder in Linux not yet implemented\n", input)
	}
	fmt.Printf("*---[%v] End Checking for left over BankDemo clone folder\n", input)

	//Clone BankDemo Repo
	fmt.Printf("*---[%v] Start Clone BankDemo GitHub.com Repo\n", input)

	var repo = "https://github.com/MicroFocus/BankDemo.git"
	cmd = exec.Command("git", "clone", repo, "--progress", "--branch", "main")
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

	fmt.Printf("*---[%v] End Clone BankDemo GitHub.com Repo\n", input)

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
		fmt.Printf("*---[%v] Fix Github typos in Linux not yet implemented\n", input)
	}

	//Run python MF_Provision_Region.py vsam

	fmt.Printf("*---[%v] Start python MF_Provision_Region.py %v\n", input)
	syscall.Chdir("BankDemo\\scripts\\")
	//option := input[0]
	cmd = exec.Command("python", "MF_Provision_Region.py", option)
	syscall.Chdir("BankDemo\\scripts\\")
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	fmt.Printf("*---[%v] End python MF_Provision_Region.py vsam\n", input)

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

	//fmt.Printf("*---[VSAM] Start HASession Server\n")
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
	myFigure := figure.NewFigure("Finshed", "", true)
	myFigure.Print()

	return "\nESCWA: http://localhost:10086\n3270: localhost:9023\n"

}

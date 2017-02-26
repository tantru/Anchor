package main

import (
	//"bytes"
	"fmt"
	//"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	//	"sys"
	//"net/http"
	//"github.com/gorilla/mux"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
func getSSOO() {
	switch runtime.GOOS {
	case "windows": // ...
	case "linux":
		fmt.Println("mierda")
	case "freebsd": // ...
	}

}

func main() {

	//HACER FUNCION QUE TE DE SSOO Y DISTRO
	//getSSOO

	// Create an *exec.Cmd
	cmd := exec.Command("apk", "-v", "info")

	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
}

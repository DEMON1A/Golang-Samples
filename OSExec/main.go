package main 

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	OS := runtime.GOOS

	if OS == "linux" {
		CMD := "ls"
		fmt.Println("You are running this script on" , OS)
		Output , error := exec.Command(CMD).Output()
		if error != nil {
			fmt.Println(error)
		}

		fmt.Println("Your Command Exected Successfully. The Comnand Was:" , CMD)
		CMDOutput := string(Output[:])
		fmt.Println(CMDOutput)
	}
}
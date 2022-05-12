package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	user := os.Args[1]
	err := os.Setenv("cfuser", user)
	if err != nil {
		fmt.Println("ERROR:SECRET_KEY SET----", err.Error())
	}

	fmt.Println(user, "欢迎使用 终端堡垒机 获取帮助请安 help")
	r1, err := readline.New(">> ")
	if err != nil {
		panic(err)
	}
	defer r1.Close()

	for {

		input, err := r1.Readline()
		if err != nil { // io.EOF
			break
		}
		line := strings.Fields(input)
		//	fmt.Println(len(line))
		if len(line) == 0 {
			t1("./cf")

		}
		if len(line) == 1 {
			t1("./cf")

		}
		if len(line) > 1 && line[0] == "cf" {

			//		t3(line[0], line[1:]...)
			t3("./cf", line[1:]...)
		}
		//		println(line)
		if input == "top" {
			t2()
		}
		if input == "ls" {
			t3("/usr/bin/ls", "-al")
		}

		if strings.HasPrefix(input, "ls ") {
		}

	}
}

func t1(arg string) {
	cmd := exec.Command(arg)
	//cmd.Path = "/"
	cmd.Stdout = os.Stdout //
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func t2() {
	cmd := exec.Command("/usr/bin/top")
	cmd.Stdout = os.Stdout //
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

/*
line := strings.Fields(input)
fmt.Println(len(line))
*/
func t3(name string, args ...string) {
	fmt.Println("name", name)
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
	cmd := exec.Command(name, args...)
	//cmd.Path = "/tmp"
	cmd.Stdout = os.Stdout //
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	/*err := cmd.Run()
	  //recode log
	if err != nil {
		fmt.Println(err)
	}
	*/

}

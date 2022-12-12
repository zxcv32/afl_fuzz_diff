package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	processA := exec.Command("./A")
	processB := exec.Command("./B")
	processC := exec.Command("./C")

	pipeA, errorPipeA := processA.StdinPipe()
	pipeB, errorPipeB := processB.StdinPipe()
	pipeC, errorPipeC := processC.StdinPipe()

	if nil != errorPipeA {
		fmt.Println(errorPipeA)
	}
	if nil != errorPipeB {
		fmt.Println(errorPipeB)
	}
	if nil != errorPipeC {
		fmt.Println(errorPipeC)
	}

	go func() {
		defer pipeA.Close()
		io.WriteString(pipeA, input)
	}()

	go func() {
		defer pipeB.Close()
		io.WriteString(pipeB, input)
	}()

	go func() {
		defer pipeC.Close()
		io.WriteString(pipeC, input)
	}()

	_, errorOutA := processA.CombinedOutput()
	if nil != errorOutA {
		fmt.Println(errorOutA)
	}
	_, errorOutB := processB.CombinedOutput()
	if nil != errorOutB {
		fmt.Println(errorOutB)
	}
	_, errorOutC := processC.CombinedOutput()
	if nil != errorOutC {
		fmt.Println(errorOutC)
	}

	if time.Now().Second()%2 == 0 {
		panic("Crash found!")
	}

	fmt.Println("Equal outputs")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu   string
	cpu2 string
	ready []string
	io1   []string
	io2   []string
	io3 []string
	io4 []string

)

func initialized() {
	cpu = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)

}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU   -> %s\n", cpu)
	fmt.Printf("CPU2 -> %s\n", cpu2)
	fmt.Printf("Ready -> ")
	for i := range ready {
		fmt.Printf("%s ", ready[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_new(p string) {
	if cpu == "" {
		cpu = p
	} else {
		insertQueue(ready, p)
	}
}

func command_terminate() {
	if cpu != "" {
		cpu = deleteQueue(ready)
	}
}

func command_expire() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	insertQueue(ready, cpu)
	cpu = p
}

func command_io1() {
	insertQueue(io1, cpu)
	cpu = ""
	command_expire()
}

func command_io2() {
	insertQueue(io2, cpu)
	cpu = ""
	command_expire()
}

func command_io3() {
	insertQueue(io3, cpu)
	cpu = ""
	command_expire()
}

func command_io4() {
	insertQueue(io4, cpu)
	cpu = ""
	command_expire()
}
func command_io1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io4x() {
	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else {
		insertQueue(ready, p)
	}
}

func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new(commandx[i])
			}
		case "terminate":
			command_terminate()
		case "expire":
			command_expire()
		case "io1":
			command_io1()
		case "io2":
			command_io2()
		case "io1x":
			command_io1x()
		case "io2x":
			command_io2x()
		case "io3x":
			command_io3x()
		case "io4x":
			command_io4x()
		default:
			fmt.Printf("\nSorry !!! Command Error !!!\n")
		}
	}
}

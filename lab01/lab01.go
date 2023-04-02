package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu   string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
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

	fmt.Printf("\nCPU1  -> %s\n", cpu)
	fmt.Printf("CPU2  -> %s\n", cpu2)
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
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}

}

func command_terminate_cpu() {
	if cpu != "" {
		cpu = deleteQueue(ready)
	} else if cpu2 != "" {
		cpu2 = deleteQueue(ready)
	}
}

func command_terminate_cpu2() {
	if cpu2 != "" {
		cpu2 = deleteQueue(ready)
	}
}

func command_expire1() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	insertQueue(ready, cpu)
	cpu = p
}
func command_expire2() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	insertQueue(ready, cpu2)
	cpu2 = p
}

func command_io1_cpu() {
	insertQueue(io1, cpu)
	cpu = ""
	command_expire1()
}

func command_io2_cpu() {
	insertQueue(io2, cpu)
	cpu = ""
	command_expire1()
}

func command_io3_cpu() {
	insertQueue(io3, cpu)
	cpu = ""
	command_expire1()
}

func command_io4_cpu() {
	insertQueue(io4, cpu)
	cpu = ""
	command_expire1()
}

// --------Cpu2----------
func command_io1_cpu2() {
	insertQueue(io1, cpu2)
	cpu2 = ""
	command_expire2()
}

func command_io2_cpu2() {
	insertQueue(io2, cpu2)
	cpu2 = ""
	command_expire2()
}

func command_io3_cpu2() {
	insertQueue(io3, cpu2)
	cpu2 = ""
	command_expire2()
}

func command_io4_cpu2() {
	insertQueue(io4, cpu2)
	cpu2 = ""
	command_expire2()
}

func command_io1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else if cpu2 == "" {
		cpu2 = p
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
	} else if cpu2 == "" {
		cpu2 = p
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
	} else if cpu2 == "" {
		cpu2 = p
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
	} else if cpu2 == "" {
		cpu2 = p
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
			//terminate cpu1
		case "ter1":
			command_terminate_cpu()
			//terminate cpu2
		case "ter2":
			command_terminate_cpu2()
		case "expire1":
			command_expire1()
		case "expire2":
			command_expire2()
		//command io(1,2,3,4)ตามด้วย(cpu1 or cpu2)
		case "io1cpu1":
			command_io1_cpu()
		case "io2cpu1":
			command_io2_cpu()
		case "io3cpu1":
			command_io3_cpu()
		case "io4cpu1":
			command_io4_cpu()
		case "io1cpu2":
			command_io1_cpu2()
		case "io2cpu2":
			command_io2_cpu2()
		case "io3cpu2":
			command_io3_cpu2()
		case "io4cpu2":
			command_io4_cpu2()
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

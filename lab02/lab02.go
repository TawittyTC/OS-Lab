package main

// Import the necessary packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Declare the global variables
var (
	cpu1  string
	cpu2  string
	ready1 []string
	ready2 []string
	ready3 []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

// Initialize the global variables
func initialized() {
	cpu1 = ""
	cpu2 = ""
	ready1 = make([]string, 10)
	ready2 = make([]string, 10)
	ready3 = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)

}

// Show the process
func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU1   -> %s\n", cpu1)
	fmt.Printf("CPU2 -> %s\n", cpu2)
	fmt.Printf("ready1 -> ")
	for i := range ready1 {
		fmt.Printf("%s ", ready1[i])
	}
	fmt.Printf("ready2 -> ")
	for i := range ready2 {
		fmt.Printf("%s ", ready2[i])
	}
	fmt.Printf("ready3 -> ")
	for i := range ready3 {
		fmt.Printf("%s ", ready3[i])
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

// Get the command
func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

// Command functions
func command_new(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready1, p)
	}
}

// Command functions Terminate Process in CPU1
func command_terminate1() {
	if cpu1 != "" {
		cpu1 = deleteQueue(ready1)
	}
}

// Command functions Terminate Process in CPU2
func command_terminate2() {
	if cpu2 != "" {
		cpu2 = deleteQueue(ready1)
	}
}

// Command functions Expire in ready1 Queue
func command_expire1() {
	p := deleteQueue(ready1)
	if p == "" {
		return
	}
	insertQueue(ready1, cpu1)
	cpu1 = p
}

func command_expire2() {
	p := deleteQueue(ready1)
	if p == "" {
		return
	}
	insertQueue(ready1, cpu2)
	cpu2 = p
}

// Command function Insert in I/O Queue 1 from CPU 1
func command_io1_c1() {
	insertQueue(io1, cpu1)
	cpu1 = ""
	command_expire1()
}

// Command function Insert in I/O Queue 1 from CPU 2
func command_io1_c2() {
	insertQueue(io1, cpu2)
	cpu2 = ""
	command_expire2()
}

// Command function Insert in I/O Queue 2 from CPU 1
func command_io2_c1() {
	insertQueue(io2, cpu1)
	cpu1 = ""
	command_expire1()
}

// Command function Insert in I/O Queue 2 from CPU 2
func command_io2_c2() {
	insertQueue(io2, cpu2)
	cpu2 = ""
	command_expire2()
}

// Command function Insert in I/O Queue 3 from CPU 1
func command_io3_c1() {
	insertQueue(io3, cpu1)
	cpu1 = ""
	command_expire1()
}

// Command function Insert in I/O Queue 3 from CPU 2
func command_io3_c2() {
	insertQueue(io3, cpu2)
	cpu2 = ""
	command_expire2()
}

// Command function Insert in I/O Queue 4 from CPU 1
func command_io4_c1() {
	insertQueue(io4, cpu1)
	cpu1 = ""
	command_expire1()
}

// Command function Insert in I/O Queue 4 from CPU 2
func command_io4_c2() {
	insertQueue(io4, cpu2)
	cpu2 = ""
	command_expire2()
}

// Command function Delete in I/O Queue 1
func command_io1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue(ready1, p)
	}
}

// Command function Delete in I/O Queue 2
func command_io2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue(ready1, p)
	}
}

// Command function Delete in I/O Queue 3
func command_io3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue(ready1, p)
	}
}

// Command function Delete in I/O Queue 4
func command_io4x() {
	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue(ready1, p)
	}
}

// Function to insert in queue
func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

// Function to delete in queue
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
		case "terminate1":
			command_terminate1()
		case "terminate2":
			command_terminate2()
		case "expire1":
			command_expire1()
		case "expire2":
			command_expire2()
		case "io11":
			command_io1_c1()
		case "io12":
			command_io1_c2()
		case "io21":
			command_io2_c1()
		case "io22":
			command_io2_c2()
		case "io31":
			command_io3_c1()
		case "io32":
			command_io3_c2()
		case "io41":
			command_io4_c1()
		case "io42":
			command_io4_c2()
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

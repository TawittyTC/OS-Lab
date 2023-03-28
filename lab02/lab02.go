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
	cpu1   string
	cpu2   string
	ready1 []string
	ready2 []string
	ready3 []string
	io1    []string
	io2    []string
	io3    []string
	io4    []string

	plt1 = 0
	plt2 = 0
	plt3 = 0
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
	fmt.Printf("CPU1 -> %s\n", cpu1)
	fmt.Printf("CPU2 -> %s\n", cpu2)
	fmt.Printf("ready1 -> ")
	for i := range ready1 {
		fmt.Printf("%s", ready1[i])
	}
	fmt.Printf("\nready2 -> ")
	for i := range ready2 {
		fmt.Printf("%s ", ready2[i])
	}
	fmt.Printf("\nready3 -> ")
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

func priority(p string){
	if p == "1"{
		plt1++
	} else if p == "2"{
		plt2++
	} else if p == "3"{
		plt3++
	}
}

func resetPriority(){
	if plt1 == 3 {
		plt1 = 0
	}else if plt2 == 3 {
		plt2 = 0
	}else if plt3 == 3 {
		plt3 = 0
	}
}

// Command functions
func command_new1(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue1(ready1, p)
	}
}

func command_new2(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue2(ready2, p)
	}
}

func command_new3(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue3(ready3, p)
	}
}

// Command functions Terminate Process in CPU1
func command_terminate1() {
	if cpu1 != "" {
		cpu1 = deleteQueue1(ready1)
	}
}

// Command functions Terminate Process in CPU2
func command_terminate2() {
	if cpu2 != "" {
		cpu2 = deleteQueue1(ready1)
	}
}

// Command functions Expire in ready1 Queue
func command_expire11() {
	p := deleteQueue1(ready1)
	if p == "" {
		return
	}
	insertQueue1(ready1, cpu1)
	cpu1 = p
}

func command_expire12() {
	p := deleteQueue2(ready1)
	//delete queue from ready1 because we move process from ready to cpu1 and cpu1 is empty
	if p == "" {
		return
	}
	insertQueue2(ready2, cpu1)
	cpu1 = p
}

func command_expire13() {
	p := deleteQueue3(ready1)
	if p == "" {
		return
	}
	insertQueue3(ready3, cpu1)
	cpu1 = p
}

func command_expire21() {
	p := deleteQueue1(ready1)
	if p == "" {
		return
	}
	insertQueue1(ready1, cpu2)
	cpu2 = p
}


func command_expire22() {
	p := deleteQueue2(ready1)
	//delete queue from ready1 because we move process from ready to cpu1 and cpu1 is empty
	if p == "" {
		return
	}
	insertQueue2(ready2, cpu2)
	cpu1 = p
}

func command_expire23() {
	p := deleteQueue3(ready1)
	if p == "" {
		return
	}
	insertQueue3(ready3, cpu2)
	cpu1 = p
}

// Command function Insert in I/O Queue 1 from CPU 1
func command_io1_c1() {
	insertQueue1(io1, cpu1)
	cpu1 = ""
	command_expire11()
}

// Command function Insert in I/O Queue 1 from CPU 2
func command_io1_c2() {
	insertQueue1(io1, cpu2)
	cpu2 = ""
	command_expire21()
}

// Command function Insert in I/O Queue 2 from CPU 1
func command_io2_c1() {
	insertQueue1(io2, cpu1)
	cpu1 = ""
	command_expire11()
}

// Command function Insert in I/O Queue 2 from CPU 2
func command_io2_c2() {
	insertQueue1(io2, cpu2)
	cpu2 = ""
	command_expire21()
}

// Command function Insert in I/O Queue 3 from CPU 1
func command_io3_c1() {
	insertQueue1(io3, cpu1)
	cpu1 = ""
	command_expire11()
}

// Command function Insert in I/O Queue 3 from CPU 2
func command_io3_c2() {
	insertQueue1(io3, cpu2)
	cpu2 = ""
	command_expire21()
}

// Command function Insert in I/O Queue 4 from CPU 1
func command_io4_c1() {
	insertQueue1(io4, cpu1)
	cpu1 = ""
	command_expire11()
}

// Command function Insert in I/O Queue 4 from CPU 2
func command_io4_c2() {
	insertQueue1(io4, cpu2)
	cpu2 = ""
	command_expire21()
}

// Command function Delete in I/O Queue 1
func command_io1x1() {
	p := deleteQueue1(io1)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue1(ready1, p)
	}
}

func command_io1x2() {
	p := deleteQueue2(io1)
	if p == "" {
		return
	}
	if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue2(ready2, p)
	}
}

func command_io1x3() {
	p := deleteQueue3(io1)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue3(ready3, p)
	}
}




// Command function Delete in I/O Queue 2
func command_io2x1() {
	p := deleteQueue1(io2)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue1(ready1, p)
	}
}

func command_io2x2(){
	p := deleteQueue2(io2)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue2(ready2, p)
	}
}

func command_io2x3(){
	p := deleteQueue3(io2)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue3(ready3, p)
	}
}


// Command function Delete in I/O Queue 3
func command_io3x1() {
	p := deleteQueue1(io3)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue1(ready1, p)
	}
}

func command_io3x2() {
	p := deleteQueue2(io3)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue2(ready2, p)
	}
}

func command_io3x3() {
	p := deleteQueue3(io3)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue3(ready3, p)
	}
}


// Command function Delete in I/O Queue 4
func command_io4x1() {
	p := deleteQueue1(io4)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue1(ready1, p)
	}
}

func command_io4x2() {
	p := deleteQueue2(io4)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue2(ready2, p)
	}
}

func command_io4x3() {
	p := deleteQueue3(io4)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else {
		insertQueue3(ready3, p)
	}
}


// Function to insert in queue
func insertQueue1(q1 []string, data string) {
	for i := range q1 {
		if q1[i] == "" {
			q1[i] = data
			break
		}
	}
}

func insertQueue2(q2 []string, data string) {
	for i := range q2 {
		if q2[i] == "" {
			q2[i] = data
			break
		}
	}
}

func insertQueue3(q3 []string, data string) {
	for i := range q3 {
		if q3[i] == "" {
			q3[i] = data
			break
		}
	}
}

// Function to delete in queue
func deleteQueue1(q1 []string) string {
	result := q1[0]
	for i := range q1 {
		if i == 0 {
			continue
		}
		q1[i-1] = q1[i]
	}
	q1[9] = ""
	return result
}

func deleteQueue2(q2 []string) string {
	result := q2[0]
	for i:= range q2 {
		if i == 0 {
			continue
		}
		q2[i-1] = q2[i]
	}
	q2[9] = ""
	return	result
}

func deleteQueue3(q3 []string) string {
	result := q3[0]
	for i := range q3 {
		if i == 0 {
			continue
		}
		q3[i-1] = q3[i]
	}
	q3[9] = ""
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
		case "new1":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new1(commandx[i])
			}
		case "new2":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new2(commandx[i])
			}
		case "new3":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new3(commandx[i])
			}
		case "terminate1":
			command_terminate1()
		case "terminate2":
			command_terminate2()
		case "expire11":
			command_expire11()
		case "expire12":
			command_expire12()
		case "expire13":
			command_expire13()
		case "expire21":
			command_expire21()
		case "expire22":
			command_expire22()
		case "expire23":
			command_expire23()

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
		case "io1x1":
			command_io1x1()
		case "io1x2":
			command_io1x2()
		case "io1x3":
			command_io1x3()
		case "io2x1":
			command_io2x1()
		case "io2x2":
			command_io2x2()
		case "io2x3":
			command_io2x3()
		case "io3x1":
			command_io3x1()
		case "io3x2":
			command_io3x2()
		case "io3x3":
			command_io3x3()
		case "io4x1":
			command_io4x1()
		case "io4x2":
			command_io4x2()
		case "io4x3":
			command_io4x3()
		default:
			fmt.Printf("\nSorry !!! Command Error !!!\n")
		}
	}
}

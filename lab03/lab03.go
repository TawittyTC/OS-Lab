package main

import (
	"fmt"         // Package สำหรับการแสดงผลทาง Console
	"math/rand"   // Package สำหรับ Random Number Generator
	"time"        // Package สำหรับเวลา
)

var (
	mess = make(chan int)  // Channel ชื่อ mess สำหรับส่งข้อมูลระหว่าง Go routines
)

// Producer function รับพารามิเตอร์ id และสร้างตัวเลขสุ่มแล้วส่งไปที่ Channel mess
func Producer(id int) {
	s1 := rand.NewSource(time.Now().UnixNano())  // กำหนด Seed ของ Random Number Generator
	r1 := rand.New(s1)                           // สร้าง Random Number Generator
	for {
		a := r1.Intn(20) + 1                // สุ่มตัวเลขระหว่าง 1-20
		sp := r1.Intn(4) + 1               // สุ่มเวลาหน่วงก่อนส่งข้อมูล
		fmt.Printf("Producer %d produce %d sleep %d\n\n", id, a, sp) // แสดงผลข้อมูล Producer
		mess <- a                          // ส่งข้อมูลไปที่ Channel mess
		time.Sleep(time.Duration(sp) * time.Second)  // หน่วงเวลาก่อนสร้างตัวเลขใหม่
	}
}

// Fibonacci function สำหรับคำนวณตัวเลข Fibonacci
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Consumer function รับพารามิเตอร์ id และรับข้อมูลจาก Channel mess แล้วคำนวณเลข Fibonacci และแสดงผล
func Consumer(id int) {
	s1 := rand.NewSource(time.Now().UnixNano())  // กำหนด Seed ของ Random Number Generator
	r1 := rand.New(s1)                           // สร้าง Random Number Generator
	sp := r1.Intn(4) + 1                        // สุ่มเวลาหน่วงก่อนคำนวณ
	for {
		data, ok := <-mess    // รับข้อมูลจาก Channel mess
		if ok {
			f := fib(data)                    // คำนวณเลข

			fmt.Printf("Consumer %d Fib %d = %d sleep %d\n\n", id, data, f, sp)  // แสดงผลลัพธ์ Consumer
			time.Sleep(time.Duration(sp) * time.Second)  // หน่วงเวลาก่อนรับข้อมูลต่อ
		} else {
			fmt.Printf("Consumer %d no data\n", id)  // แสดงผลลัพธ์เมื่อไม่มีข้อมูลใน Channel
		}

	}
}

// main function เรียก Goroutine ของ Producer 2 ตัวและ Consumer 3 ตัว
func main() {
	go Producer(1)   // สร้าง Goroutine ของ Producer 1
	go Producer(2)   // สร้าง Goroutine ของ Producer 2
	go Consumer(1)   // สร้าง Goroutine ของ Consumer 1
	go Consumer(2)   // สร้าง Goroutine ของ Consumer 2
	go Consumer(3)   // สร้าง Goroutine ของ Consumer 3
	select {}  // รอจนกว่า Channel จะถูกปิด
}

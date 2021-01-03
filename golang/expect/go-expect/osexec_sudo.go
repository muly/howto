package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	expect "github.com/Netflix/go-expect"
)

func main() {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	log.Println("1")

	cmd := exec.Command("sudo", "ls")
	// cmd := exec.Command("echo")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	// go func() {
	// 	c.ExpectEOF()
	// }()

	log.Println("2")

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("3")

	// s, err := c.ExpectString("Password:")
	// if err != nil {
	// 	log.Fatalf("c.ExpectString error::::::::::::::::::: %v",err)
	// }
	// log.Println("######",s)

	// s, err := c.Expect(expect.String("Password:"), expect.WithTimeout(2*time.Second))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("$$$$$$$$$$$$", s)

	// // time.Sleep(2*time.Second)
	// c.SendLine("beg9tAll")
	// // time.Sleep(time.Second)
	// // c.Send("dd")
	// // time.Sleep(time.Second)
	// // c.SendLine(":wq")

	time.Sleep(time.Second)
	c.ExpectString("Password")
	// c.Expect(expect.String("Password"), expect.WithTimeout(time.Second))
	log.Println("sucessfully waited")

	// c.Tty().Close()
	// time.Sleep(time.Second)
	c.SendLine("beg9tAll\n")

	// c.Tty().Close()
	// c.ExpectEOF()

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
	"os/exec"

	expect "github.com/Netflix/go-expect"
)

func main() {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("ssh user@hostip")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	s, err := c.ExpectString("Password: ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("####", s)

	// time.Sleep(time.Second)
	// c.Send("iHello world\x1b")
	// time.Sleep(time.Second)
	// c.Send("dd")
	// time.Sleep(time.Second)
	// c.SendLine(":wq")

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

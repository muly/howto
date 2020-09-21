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

	cmd := exec.Command("python")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	c.ExpectString(">>>")

	c.SendLine(`print("hello")`)
	c.ExpectString("hello")

	c.SendLine("exit()")

	// c.ExpectString("mulys-new-mbp:go-expect muly$ ")
	// log.Println("execution complete")

	c.Tty().Close()
	c.ExpectEOF()

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

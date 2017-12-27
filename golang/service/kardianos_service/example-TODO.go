//TODO: need to review code, test, confirm, and add more documentation on how to use it.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kardianos/service"
)

var args []string = os.Args

type program struct{}

func main() {

	ServiceArg := make([]string, 1)
	ServiceArg[0] = "service"

	var sc = &service.Config{
		Name:        "YourServiceName",
		DisplayName: "YourServiceName",
		Description: "Monitors the YourServiceName applications.",
		Arguments:   ServiceArg,
	}

	prg := &program{}
	s, err := service.New(prg, sc)
	if err != nil {
		fmt.Println(err)
	}

	strCmdArg := getCommandLineargs()

	switch strCmdArg {

	case "install":
		err = s.Install()
		if err != nil {
			fmt.Println("Application service not Installed successfully, ErrorInfo : ", err)
		}

		err = s.Start()
		if err != nil {
			fmt.Println("Application service not Started , ErrorInfo : ", err)
		}
		os.Exit(0)

	case "start":
		err = s.Start()
		if err != nil {
			fmt.Println("Application service not Started , ErrorInfo : ", err)
		}
		os.Exit(0)

	case "stop":
		err = s.Stop()
		if err != nil {
			fmt.Println("Application service not stopped , ErrorInfo : ", err)
		}

		os.Exit(0)

	case "uninstall":
		err = s.Stop()
		if err != nil {
			fmt.Println("Application service not stopped , ErrorInfo : ", err)
		}

		err = s.Uninstall()
		if err != nil {
			fmt.Println("Application service not uninstalled successfully, ErrorInfo : ", err)
		}
		os.Exit(0)
	case "service":
		err = s.Run()
		if err != nil {
			fmt.Println("Application not running properly, ErrorInfo : ", err)
		}

	default:
		fmt.Println("Invalid Argument for executable")
		os.Exit(0)
	}

}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	fmt.Println("########################################## Stopping Application Service ##########################################")
	return nil
}

func getCommandLineargs() (result string) {
	result = args[1]
	return result
}

func (p *program) run() {

	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred at run() and Recovered in run(), Error Info: ", errD)
		}
	}()
	fmt.Println("########################################## Application started as Service ##########################################")
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Process is running as service")
	}
}

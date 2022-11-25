package main

import (
	"fmt"
	"os"
	"plugin"
)

type Msngr interface {
	Welcome()
}

func main() {
	// determine plugin to load
	userType := "user"
	if len(os.Args) == 2 {
		userType = os.Args[1]
	}
	var mod string
	switch userType {
	case "user":
		mod = "./plugin/user.so"
	case "admin":
		mod = "./plugin/admin.so"
	default:
		fmt.Println("Not supported")
		return
	}

	// load module

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		return
	}

	// look up a symbol
	symMsngr, err := plug.Lookup("User")
	if err != nil {
		fmt.Println(err)
		return
	}

	var msngr Msngr
	msngr, ok := symMsngr.(Msngr)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		return
	}

	msngr.Welcome()

}

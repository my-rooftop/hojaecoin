package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/nomadcoders/nomadcoin/explorer"
	"github.com/nomadcoders/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to HojaeCoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:	  	Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Sets the port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

}

package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

// MAJOR RELEASE . MAJOR FEATURE/MAJOR FIX . MINOR FEATURE/MINOR FIX
var CURRENT_VERSION = "1.0.0"
var DEBUG	= flag.BoolP("debug","d",false, "Turns debug on (turns off logging)")
var VERSION	= flag.BoolP("version","V",false, "Displays the current version")
var HELP    = flag.BoolP("help","h",false, "Displays this help message")
var VERBOSE = flag.BoolP("verbose","v",false, "Gives more information about breaches")

const APIURI = "https://haveibeenpwned.com/api/v2/breachedaccount"

func main() {

	// set the usage
	flag.Usage = func() {
		fmt.Printf("%v Version: %v\n\nUsage: %s [options]\n\n",get_prog_name(),CURRENT_VERSION,get_prog_name())
		fmt.Printf("Optional Args:\n")
		flag.PrintDefaults()
		fmt.Printf("\nRequired Args:\n\t<email>\t\tEmail address you want to lookup\n")
		fmt.Println("")
	}
	//parse the command-line flags
	flag.Parse()

	// the pflag library has a stupid output if you don't declare
	// your own "help", so that's why we do it here
	if *HELP {
		flag.Usage()
		os.Exit(0)
	}

	// if they are checking the version, just print the version and exit
	if *VERSION == true {
		fmt.Printf("%v Version: %v\n",get_prog_name(), CURRENT_VERSION)
		os.Exit(0)
	}
	// parse if we are doing login or logout
	var email string
	if len(flag.Args()) > 0 {
		email = flag.Args()[0]
	}
	if email == "" || len(email) < 1 {
		print_error(" Must supply an email address")
	}
	d_log(fmt.Sprintf("email: %v",email))

	get_response(email)

}

package main

import (
	"strings"
	"os"
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"net/url"
)

type Breach struct {
	Title        string   `json:"Title"`
	Name         string   `json:"Name"`
	Domain       string   `json:"Domain"`
	BreachDate   string   `json:"BreachDate"`
	AddedDate    string   `json:"AddedDate"`
	ModifiedDate string   `json:"ModifiedDate"`
	PwnCount     int64    `json:"PwnCount"`
	Description  string   `json:"Description"`
	DataClasses  []string `json:"DataClasses"`
	IsVerified   bool     `json:"IsVerified"`
	IsFabricated bool     `json:"IsFabricated"`
	IsSensitive  bool     `json:"IsSensitive"`
	IsActive     bool     `json:"IsActive"`
	IsRetired    bool     `json:"IsRetired"`
	IsSpamList   bool     `json:"IsSpamList"`
	LogoType     string   `json:"LogoType"`
}

func get_prog_name () string {
	name_slice := strings.Split(os.Args[0],"/")
	// kinda like a pop
	name 	   := name_slice[len(name_slice)-1]
	return name
}

func print_error (error string) {
	fmt.Printf("[ERROR]: %v\n",error)
	os.Exit(1)
}

func d_log (message string) {
	if *DEBUG {
		fmt.Printf("[DEBUG]: %v\n",message)
	}
}

func verbose_print (message string) {
	if *VERBOSE {
		fmt.Println(message)
	}
}

func get_response (email string) {
	var breach *Breach

	endpoint := fmt.Sprintf(APIURI,url.QueryEscape(email))
	d_log(fmt.Sprintf("endpoint: %v", endpoint))
	client := &http.Client {
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(endpoint)

	if err != nil {
		print_error(fmt.Sprintf("Failed to get information: %v",err))
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	d_log(fmt.Sprintf("len(body): %v", len(body)))

	if len(body) == 0 {
		fmt.Printf("Unable to find breaches for %v!\n",email)
		os.Exit(0)
	}

	if err = json.Unmarshal(body, &breach); err != nil {
		//attempt to decode status and data field to check for error message
		var objmap []Breach

		err := json.Unmarshal(body, &objmap)
		if err != nil {
			print_error(fmt.Sprintf("Unable to unmarshal: %v",err))
		}
		d_log(fmt.Sprintf("objmap: %v", objmap))

		print_response(objmap, email)

	} else {
		print_error(fmt.Sprintf("Unable to unmarshal: %v\n", err))
	}
}


func print_response (breaches []Breach, email string) {

	num_of_breaches := len(breaches)

	fmt.Printf("Number of breaches for %v: %v\n\n", email, num_of_breaches)
	breach_num := 1
	for _, breach := range breaches {

		fmt.Printf("Breach: %v\n", breach_num)
		verbose_print(fmt.Sprintf("\tTitle: %v", breach.Name))
		fmt.Printf("\tName: %v\n", breach.Name)
		fmt.Printf("\tDomain: %v\n", breach.Domain)
		fmt.Printf("\tBreach Date: %v\n", breach.BreachDate)
		verbose_print(fmt.Sprintf("\tAdded Date: %v", breach.AddedDate))
		verbose_print(fmt.Sprintf("\tModified Date: %v", breach.ModifiedDate))
		verbose_print(fmt.Sprintf("\tPwn Count: %v", breach.PwnCount))
		verbose_print(fmt.Sprintf("\tDescription: %v", breach.Description))
		verbose_print(fmt.Sprintf("\tData Classes: %v", get_data_classes(breach.DataClasses)))
		fmt.Printf("\tIs Verified: %v\n", breach.IsVerified)
		verbose_print(fmt.Sprintf("\tIs Fabricated: %v", breach.IsFabricated))
		verbose_print(fmt.Sprintf("\tIs Sensitive: %v", breach.IsSensitive))
		verbose_print(fmt.Sprintf("\tIs Active: %v", breach.IsActive))
		verbose_print(fmt.Sprintf("\tIs Retired: %v", breach.IsRetired))
		verbose_print(fmt.Sprintf("\tIs Spam List: %v", breach.IsSpamList))
		fmt.Println("")

		breach_num ++
	}

}

func get_data_classes (DataClasses []string) (data_classes string) {
	var tmp string
	for _, event := range DataClasses {
		tmp = tmp + event  + ", "
	}

	data_classes = strings.TrimRight(tmp,", ")

	return data_classes
}
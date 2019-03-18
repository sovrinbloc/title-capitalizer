package main

import (
"fmt"
"strings"

xj "github.com/basgys/goxml2json"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

const songCaseApi = "http://songcase.com/xml.asp?q=%s"



func main() {
	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.

	if len(os.Args) < 2 {
		fmt.Println("Please enter a title in quotations")
		return
	}
	arg := os.Args[1]

	title := getTitle(arg)
	fmt.Println(XmltoStruct(title))
}

func XmltoStruct(title string) string {
	xml := strings.NewReader(title)
	result, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}
	type Capitalizer struct {
		Capitalize struct {
			Input  string `bson:"input" json:"input"`
			Output string `bson:"output" json:"output"`
		} `json:"capitalize"`
	}
	resultMap := Capitalizer{}
	err = json.Unmarshal([]byte(result.String()), &resultMap)
	return resultMap.Capitalize.Output
}


func getTitle(title string) string {
	title = strings.Replace(title, " ", "%20", -1)
	url := fmt.Sprintf(songCaseApi, title)
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
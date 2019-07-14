package json_hndl

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var osFileMode = os.FileMode.Perm(0644)

var ErrInvalidArgs = errors.New("invalid arguments")

type Job struct {
	Arg1 int `json:"arg1"`
	Arg2 int `json:"arg2"`
}

type JobReport struct {
	Title string `json:"title"`
	Div   int    `json:"div"`
}

// This function handles user console input consisting of input and output file names.
// Client of this function should provide callback that
// will be called each time user input successfully handled.
// Jobs that don't match JSON schema will be filled with zeros.
func HandleJSON(callback func(jobs []Job, output string)) {
	fmt.Println("Please specify input and output files in respective order")
	fmt.Println("Mind that input file should be in JSON format")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inpParams := strings.Split(scanner.Text(), " ")
		if len(inpParams) < 2 {
			fmt.Println(ErrInvalidArgs.Error())
			continue
		}

		input, output := inpParams[0], inpParams[1]
		file, err := ioutil.ReadFile(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		var jobs []Job
		err = json.Unmarshal([]byte(file), &jobs)
		if err != nil {
			fmt.Println("JSON unmarshal: ", err.Error())
			continue
		}
		callback(jobs, output)
	}
}

// This function writes report in JSON format to specified output file.
func WriteJSON(report *[]JobReport, output string) {
	json, err := json.MarshalIndent(report, "", "\t")
	if err != nil {
		fmt.Println("JSON marshal: ", err)
		return
	}
	err = ioutil.WriteFile(output, json, osFileMode)
	if err != nil {
		fmt.Println("write JSON file: ", err)
		return
	}
}
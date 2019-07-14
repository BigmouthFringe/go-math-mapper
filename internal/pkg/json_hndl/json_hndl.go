// This package handles user console input
// consisting of input and output file names.
// Specifically, it decodes JSON and converts it to Jobs.
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

var (
	ErrInvalidArgs = errors.New("invalid arguments")
)

type Job struct {
	Arg1 int
	Arg2 int
}

// Client of this function should provide callback that
// will be called each time user input successfully handled.
// It will ignore Jobs that don't match Job struct spec
// TODO: Write tests
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
			fmt.Println(err.Error())
			continue
		}
		callback(jobs, output)
	}
}
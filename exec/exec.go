package exec

import (
	"os/exec"
	"strings"
)

// Output is the struct that hold a command's output
type Output struct {
	// Stdout will have byte output of exec.Command
	Stdout []byte
	// Stderr will have error object of exec.Command
	Stderr error
}

// run will process the specified cmd, execute it and return the Output struct
func Run(cmd string) *Output {
	cmdName, cmdArgs := processCmdStr(cmd)
	outStruct := new(Output)
	outStruct.Stdout, outStruct.Stderr = execute(cmdName, cmdArgs...)
	return outStruct
}

// processCmdStr will split full command string to command and arguments slice
func processCmdStr(cmd string) (string, []string) {
	cmdParts := strings.Split(cmd, " ")
	return cmdParts[0], cmdParts[1:]
}

// execute will run the specified command with arguments
// and return output, error
func execute(cmd string, args ...string) ([]byte, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

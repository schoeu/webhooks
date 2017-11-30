package exec

import (
	"os/exec"
	"strings"
	"path/filepath"
)

// Output is the struct that hold a command's output
type Output struct {
	// Stdout will have byte output of exec.Command
	Stdout []byte
	// Stderr will have error object of exec.Command
	Stderr error
}

var curDir = ""

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
func execute(c string, args ...string) ([]byte, error) {
	cmd := exec.Command(c, args...)
	preDir := curDir
	isChangeDir := c == "cd"
	if isChangeDir {
		p := args[0]
		if filepath.IsAbs(p) {
			curDir = p
		} else {
			curDir = filepath.Join(curDir, p)
		}
		// After change direction should show the file list.
		cmd = exec.Command("ls", curDir)
	}
	cmd.Dir = curDir

	out, err := cmd.Output()
	if err != nil {
		// If catch an error when change direction
		// should rollback the path.
		if isChangeDir {
			curDir = preDir
		}

		return nil, err
	}

	return out, nil
}

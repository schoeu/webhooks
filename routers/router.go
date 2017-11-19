package routers

import (
	"github.com/schoeu/webhooks/config"
	"github.com/schoeu/webhooks/utils"
	"io"
	"net/http"
)

// Action for task.
func TaskRouter(w http.ResponseWriter, infos [][]string, c config.ConfigMap) bool {
	hit := false
	action := infos[0][1]
	allActions := c.GetAll()
	cmdStr := allActions[action]
	if cmdStr != "" {
		rs := utils.ExecCmds(cmdStr)
		hit = true
		io.WriteString(w, rs)
	}
	return hit
}

// Action for run script.
func RunRouter(w http.ResponseWriter, infos [][]string) {
	runCmd := infos[0][1]
	//o := execCmd(runCmd, "run")
	rs := utils.ExecCmds(runCmd)
	io.WriteString(w, rs)
}

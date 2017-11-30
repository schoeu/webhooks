package routers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/schoeu/webhooks/config"
	"github.com/schoeu/webhooks/utils"
)

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	query   url.Values
}

// Action for task.
func TaskRouter(ctx utils.Context, infos [][]string, c config.ConfigMap) bool {
	hit := false
	action := infos[0][1]
	allActions := c.GetAll()
	cmdStr := allActions[action]
	isJson := ctx.Query.Get("json") != ""
	if cmdStr != "" {
		rs := utils.ExecCmds(cmdStr)
		hit = true
		if isJson {
			utils.ReturnJson(ctx.Writer, rs)
		} else {
			io.WriteString(ctx.Writer, rs)
		}

	}
	return hit
}

// Action for run script.
func RunRouter(ctx utils.Context, infos [][]string) {
	runCmd := infos[0][1]
	//o := execCmd(runCmd, "run")
	rs := utils.ExecCmds(runCmd)
	isJson := ctx.Query.Get("json") != ""
	if isJson {
		utils.ReturnJson(ctx.Writer, rs)
	} else {
		io.WriteString(ctx.Writer, rs)
	}

}

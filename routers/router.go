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
	isAsync := ctx.Query.Get("async") != ""
	if cmdStr != "" {
		var rs string
		if isAsync {
			go utils.ExecCmds(cmdStr)
			rs = "done"
		} else {
			rs = utils.ExecCmds(cmdStr)
		}
		if isJson {
			utils.ReturnJson(ctx.Writer, rs)
		} else {
			io.WriteString(ctx.Writer, rs)
		}
		hit = true
	}
	return hit
}

// Action for run script.
func RunRouter(ctx utils.Context, infos [][]string) {
	runCmd := infos[0][1]
	isJson := ctx.Query.Get("json") != ""
	isAsync := ctx.Query.Get("async") != ""
	if runCmd != "" {
		var rs string
		if isAsync {
			go utils.ExecCmds(runCmd)
			rs = "done"
		} else {
			rs = utils.ExecCmds(runCmd)
		}
		if isJson {
			utils.ReturnJson(ctx.Writer, rs)
		} else {
			io.WriteString(ctx.Writer, rs)
		}
	}
}

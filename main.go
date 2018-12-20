package beego_pt

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log/syslog"
	"fmt"
)

var customAddr string
var customService string
var customWriter *syslog.Writer

//func InitReport(ctx *context.Context) {
//}

func ProcessReport(ctx *context.Context) {
	msg := fmt.Sprintf("method = %s path = %s IP = %s refer = %s agent = %s",
			ctx.Input.Method(), ctx.Input.URI(), ctx.Input.IP(), ctx.Input.Refer(), ctx.Input.UserAgent(),
		)
	customWriter.Write([]byte(msg))
}

// InitPaperTrailFilter("logs4.papertrailapp.com:49763", "znzn")
func InitPaperTrailFilter(addr string, service string) (error) {

	// Initialize PaperTrail parameters
	customAddr = addr
	customService = service

	// Initialize PaperTrail writer
	w, err := syslog.Dial("udp", addr, syslog.LOG_EMERG | syslog.LOG_KERN, service)
	if err != nil {
		return err
	}
	customWriter = w

	// Initialize BeeGo request handler
//	beego.InsertFilter("*", beego.BeforeRouter, InitReport)
	beego.InsertFilter("*", beego.FinishRouter, ProcessReport, false)

	// Complete report
	beego.Info("PaperTrail BeeGo reporter system initialized")

	return nil
}

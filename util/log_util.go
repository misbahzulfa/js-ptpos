package util

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogRequest(ctx *fiber.Ctx, requestString string, uuid string) string {
	currentTime := time.Now()
	return "[Start][RequestId]= " + uuid + ", [Path]= " + ctx.Route().Path + ", [IP]= " + ctx.IP() + ", [IPCLIENT]= " + strings.Join(ctx.IPs(), ",") + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString
}

func LogResponse(ctx *fiber.Ctx, responseString string, uuid string) string {
	currentTime := time.Now()
	return "[Stop][RequestId]= " + uuid + ", [Path]= " + ctx.Route().Path + ", [IP]= " + ctx.IP() + ", [IPCLIENT]= " + strings.Join(ctx.IPs(), ",") + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString
}

package emaaslogs

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

type LogInfo struct {
	Action      string
	UserAccount string
	Result      string
}

func TestMyhooks(t *testing.T) {

	LogInit("", "dev", nil)
	info := LogInfo{}
	info.UserAccount = "test1"
	info.Result = "ok"
	info.Action = "test1"

	log.Info(ToMap(info))
	log.Error(ToMap(info))
	log.Warn(ToMap(info))
	log.Fatal(ToMap(info))
}

package logrus

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func TestLogger2(t *testing.T) {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("replica_"+time.Now().Format("2006-01-02_15:04:05")+".log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Fatalf("create file log.txt failed: %v", err)
	}
	logrus.SetFormatter(&nested.Formatter{NoColors: true})

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Infof("info msg %d test %d test %s", 1, 2, "k")
}

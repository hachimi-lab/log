package aliyunlog

import (
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

type Config struct {
	ProjectName     string
	LogStoreName    string
	LogTopic        string
	LogSource       string
	LogContentKey   string
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

type Logger struct {
	Config
	client sls.ClientInterface
}

func New(config Config) *Logger {
	provider := sls.NewStaticCredentialsProvider(config.AccessKeyId, config.AccessKeySecret, config.SecurityToken)
	return &Logger{
		Config: config,
		client: sls.CreateNormalInterfaceV2(config.Endpoint, provider),
	}
}

func (slf *Logger) Write(data []byte) (int, error) {
	contents := []*sls.LogContent{{
		Key:   proto.String("__log__"),
		Value: proto.String(string(data)),
	}}
	logs := []*sls.Log{{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: contents,
	}}
	logGroup := &sls.LogGroup{
		Logs:   logs,
		Topic:  proto.String(slf.LogTopic),
		Source: proto.String(slf.LogSource),
	}
	if err := slf.client.PutLogs(slf.ProjectName, slf.LogStoreName, logGroup); err != nil {
		return 0, errors.Wrap(err, "failed to put logs to aliyun sls")
	}
	return len(data), nil
}

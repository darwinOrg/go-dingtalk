package dgding

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	dingtalkClient := DingtalkClient{
		Webhook:   "https://oapi.dingtalk.com/robot/send?access_token=4cf7baa88b551146726697ad5901aee429c5a7660c18c0e45c083abeaec42ef9",
		AtAll:     false,
		AtMobiles: []string{"15901431753"},
	}

	ctx := &dgctx.DgContext{TraceId: "trace-005"}
	message, _ := dingtalkClient.SendTextMessage(ctx, "（反馈问题）test")
	dglogger.Infoln(ctx, message)
}

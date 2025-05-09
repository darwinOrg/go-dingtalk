package dgding

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	"strings"
)

type DingtalkClient struct {
	Webhook   string
	AtMobiles []string
	AtAll     bool
	Keyword   string
}

func (dc *DingtalkClient) SendTextMessage(ctx *dgctx.DgContext, content string) (string, error) {
	if dc.Keyword != "" && !strings.Contains(content, dc.Keyword) {
		content = "[" + dc.Keyword + "]" + content
	}

	params := map[string]any{
		"msgtype": "text",
		"text":    map[string]string{"content": content},
	}

	if dc.AtAll {
		params["at"] = map[string]any{"isAtAll": true}
	} else if dc.AtMobiles != nil && len(dc.AtMobiles) > 0 {
		params["at"] = map[string]any{"atMobiles": dc.AtMobiles, "isAtAll": false}
	}

	data, err := dghttp.GlobalHttpClient.DoPostJson(ctx, dc.Webhook, params, map[string]string{})
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func SimpleSendTextMessage(ctx *dgctx.DgContext, webhook, content string) (string, error) {
	dglogger.Infof(ctx, "simple send text message, content: %s", content)
	data, err := dghttp.GlobalHttpClient.DoPostJson(ctx, webhook, map[string]any{
		"msgtype": "text",
		"text":    map[string]string{"content": content},
	}, nil)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

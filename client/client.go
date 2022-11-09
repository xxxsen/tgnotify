package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tgnotify/model"

	"github.com/xxxsen/common/errs"
)

type NotifyClient struct {
	c *config
}

func New(opts ...Option) (*NotifyClient, error) {
	c := &config{}
	for _, opt := range opts {
		opt(c)
	}
	if len(c.host) == 0 {
		return nil, errs.New(errs.ErrParam, "nil host")
	}
	if len(c.ak) == 0 || len(c.sk) == 0 {
		return nil, errs.New(errs.ErrParam, "nil user")
	}
	return &NotifyClient{
		c: c,
	}, nil
}

func (c *NotifyClient) SendMessage(ctx context.Context, msg string) error {
	return c.SendMessageWithType(ctx, MessageTypeText, msg)
}

func (c *NotifyClient) buildURL(api string) string {
	return c.c.host + api
}

type jsonFrame struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *NotifyClient) jsonPost(ctx context.Context, api string, req, rsp interface{}) error {
	raw, err := json.Marshal(req)
	if err != nil {
		return errs.Wrap(errs.ErrMarshal, "encode fail", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.buildURL(api), bytes.NewReader(raw))
	if err != nil {
		return errs.Wrap(errs.ErrParam, "build http request fail", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("user", c.c.ak)
	httpReq.Header.Set("code", c.c.sk)
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return errs.Wrap(errs.ErrIO, "call http fail", err)
	}
	defer httpRsp.Body.Close()
	if httpRsp.StatusCode != http.StatusOK {
		return errs.New(errs.ErrUnknown, "code not ok, code:%d", httpRsp.StatusCode)
	}

	data, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return errs.Wrap(errs.ErrIO, "read rsp fail", err)
	}
	rspFrame := &jsonFrame{
		Data: rsp,
	}
	if err := json.Unmarshal(data, rspFrame); err != nil {
		return errs.Wrap(errs.ErrUnknown, "decode rsp fail", err)
	}
	if rspFrame.Code != 0 {
		return errs.New(rspFrame.Code, rspFrame.Message)
	}
	return nil
}

func (c *NotifyClient) SendMessageWithType(ctx context.Context, typ string, msg string) error {
	req := &model.SendMessageRequest{
		Message:     msg,
		MessageType: typ,
	}
	if err := c.jsonPost(ctx, apiSendMesg, req, nil); err != nil {
		return errs.Wrap(errs.ErrServiceInternal, "json post fail", err)
	}
	return nil
}

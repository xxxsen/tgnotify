package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"tgnotify/message"
	"tgnotify/model"
	"tgnotify/sender"

	"github.com/gin-gonic/gin"
	"github.com/xxxsen/common/logutil"
	"github.com/xxxsen/common/webapi/proxyutil"
	"go.uber.org/zap"
)

const (
	defaultMessageTypeHeader = "x-message-type"
)

func createMessageByKind(typ string, msg string) (message.IMessage, error) {
	switch typ {
	case "", message.MKindText:
		return message.NewTextMessage(msg), nil
	case message.MKindMarkdown:
		return message.NewMarkdownMessage(msg), nil
	case message.MKindHTML:
		return message.NewHTMLMessage(msg), nil
	case message.MKindTGMarkdown:
		return message.NewTGMarkdownMessage(msg), nil
	}
	return nil, fmt.Errorf("unsupported messsage kind:%s", typ)
}

func readJsonMessage(c *gin.Context) (string, string, error) {
	res := &model.SendMessageRequest{}
	if err := c.ShouldBindBodyWithJSON(res); err != nil {
		return "", "", err
	}
	return res.MessageType, res.Message, nil
}

func readDefaultMessage(c *gin.Context) (string, string, error) {
	typ := c.GetHeader(defaultMessageTypeHeader)
	raw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "", "", err
	}
	return typ, string(raw), nil
}

func readMessageByContentType(c *gin.Context) (string, string, error) {
	switch strings.ToLower(c.ContentType()) {
	case "application/json":
		return readJsonMessage(c)
	default:
		return readDefaultMessage(c)
	}
}

func SendMessage(c *gin.Context) {
	ctx := c.Request.Context()
	kind, msg, err := readMessageByContentType(c)
	if err != nil {
		proxyutil.FailJson(c, http.StatusBadRequest, fmt.Errorf("read message failed, err:%w", err))
		return
	}
	logutil.GetLogger(ctx).Debug("read message succ", zap.String("kind", kind), zap.String("msg", msg))
	msgblk, err := createMessageByKind(kind, msg)
	if err != nil {
		proxyutil.FailJson(c, http.StatusInternalServerError, fmt.Errorf("create message failed, err:%w", err))
		return
	}
	if err := sender.SendMessage(ctx, msgblk); err != nil {
		proxyutil.FailJson(c, http.StatusInternalServerError, fmt.Errorf("send msg failed, err:%w", err))
		return
	}
	logutil.GetLogger(ctx).Info("send message succ", zap.String("kind", kind))
}

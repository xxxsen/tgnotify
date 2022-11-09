package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	c, err := New(WithHost("http://127.0.0.1:10010"), WithUser("abc", "123"))
	assert.NoError(t, err)
	err = c.SendMessage(context.Background(), "test")
	assert.NoError(t, err)
}

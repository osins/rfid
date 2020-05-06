package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewENV(t *testing.T) {
	env := NewENV()
	env.Load()

	assert.Equal(t, os.Getenv("TEST"), "true", "测试是否能够加载到.env中的TEST信息")
}

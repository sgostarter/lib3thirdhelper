package feishu

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendTextNotify(t *testing.T) {
	result, err := SendTextNotify(os.Getenv("fstoken"), "hello2\nhello")
	assert.Nil(t, err)
	t.Log(result)
}

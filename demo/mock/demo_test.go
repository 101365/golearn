package mock

import (
	"bou.ke/monkey"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

// 模拟外部http请求
func TestMockServer(t *testing.T) {

	aq := assert.New(t)

	monkey.Patch(http.Get, func(url string) (resp *http.Response, err error) {
		resp = &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader([]byte("hello"))),
		}
		return resp, nil
	})
	defer monkey.UnpatchAll()

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	aq.True(string(body) == "hello")
}

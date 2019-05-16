package mock

import (
	"bou.ke/monkey"
	"bytes"
	"fmt"
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

func TestSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]    // [2 3 4]
	s2 := s1[2:6:7] // [4 5 6 7]

	fmt.Println(s1, s2, len(s2), cap(s2))
	//[2 3 4] [4 5 6 7] 4 5
}

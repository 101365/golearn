package group

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroup(t *testing.T) {
	aq := assert.New(t)
	arr := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(s), &arr)
	aq.Nil(err)

	var sum float64
	for _, m := range arr {
		v := m["点击量"]
		vv := v.(float64)
		sum += vv
	}

	for _, m := range arr {
		v := m["点击量"]
		vv := v.(float64)
		ratio := vv / sum
		m["比例"] = ratio
	}

	fmt.Println("时段", "	", "点击量", "	", "比例")
	for _, m := range arr {
		v := m["比例"]
		vv := v.(float64)
		fmt.Println(m["小时"], "	", m["点击量"], "	", fmt.Sprintf("%.2f%%", vv*100))
	}

}

var s = `[
    {
        "小时": "052800",
        "点击量": 610
    },
    {
        "小时": "052801",
        "点击量": 307
    },
    {
        "小时": "052802",
        "点击量": 175
    },
    {
        "小时": "052803",
        "点击量": 150
    },
    {
        "小时": "052804",
        "点击量": 227
    },
    {
        "小时": "052805",
        "点击量": 466
    },
    {
        "小时": "052806",
        "点击量": 761
    },
    {
        "小时": "052807",
        "点击量": 636
    },
    {
        "小时": "052808",
        "点击量": 438
    },
    {
        "小时": "052809",
        "点击量": 375
    },
    {
        "小时": "052810",
        "点击量": 289
    },
    {
        "小时": "052811",
        "点击量": 137
    }
]`

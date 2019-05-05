package compare

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

// 对比两个文件中重复的UID
func TestCompare(t *testing.T) {

	zhunxingMap := getZhunxingUIDList("/Users/xxx/Downloads/xxx.txt")
	// 从其他地方读取
	rawMap := getZhunxingUIDList("/Users/xxx/Downloads/xxx.log")

	//fmt.Println(len(rawMap))

	sameM := make(map[string]bool)
	for k := range zhunxingMap {
		uid, _ := strconv.ParseInt(k, 10, 64)
		pid := int(uid) & 0x0000FFFFFFFFFFFF
		if v, exist := rawMap[fmt.Sprintf("%v", pid)]; exist && v {
			sameM[k] = true
		}
	}

	fmt.Println(len(sameM))

	WriteMaptoFile(sameM, fmt.Sprintf("/Users/xxx/Downloads/native_same_pid_list_%v.txt", len(sameM)))

}

func WriteMaptoFile(m map[string]bool, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k := range m {
		fmt.Fprintln(w, k)
	}
	return w.Flush()
}

func getZhunxingUIDList(f string) map[string]bool {
	m := make(map[string]bool)
	lines := readLines(f)
	for _, line := range lines {
		if line == "" {
			continue
		}
		m[line] = true
	}
	return m
}

func readLines(f string) []string {
	ret := make([]string, 0)
	fi, err := os.Open(f)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ret
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		ret = append(ret, string(a))
	}
	return ret
}

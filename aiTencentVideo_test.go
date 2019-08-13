package aiTencentVideo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestGetVideoUrlForVid(t *testing.T) {
	vid := "k0912t1gqas"
	originUrl := "http://vv.video.qq.com/getinfo?vids=%v&platform=101001&charge=0&otype=json&defn=shd"
	resp, err := http.Get(fmt.Sprintf(originUrl, vid))
	if err != nil {
		t.Fatal(err)

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)[len("QZOutputJson=") : len(string(body))-1]
	url := jsoniter.Get([]byte(bodyStr), "vl").Get("vi").Get(0).
		Get("ul").Get("ui").Get(0).Get("url").ToString()
	fn := jsoniter.Get([]byte(bodyStr), "vl").Get("vi").Get(0).Get("fn").ToString()
	fvKey := jsoniter.Get([]byte(bodyStr), "vl").Get("vi").Get(0).Get("fvkey").ToString()
	if url == "" || fn == "" || fvKey == "" {
		t.Fatal("json解析数据失败")
	}
}

package ins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DlIp struct {
	Ret  string `json:"Ret"`
	Data string `json:"Data"`
	Msg  string `json:"Msg"`
}

func ChangeIP(ipproxy string) {
	var shuju DlIp
	resp, err := http.Get("http://refresh.rola.info/refresh?user=" + ipproxy + "&country=ru")
	if err != nil {
		fmt.Println("获取代理网址失败")
	}
	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("读取网页数据失败")
	}

	fmt.Println("调试查看---------", string(all))

	json.Unmarshal(all, &shuju)

	if shuju.Ret == "SUCCESS" {
		fmt.Println("更换IP成功")

	} else {
		fmt.Println("更换IP失败")

	}
}

func CountTime(num int) {
	if num > 0 {
		fmt.Println(num)
		time.Sleep(time.Duration(1) * time.Second)
		CountTime(num - 1)
	} else {
		fmt.Println("倒计时完成")
	}
}

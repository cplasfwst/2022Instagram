package ins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

func GetRandNum() int {
	/*
	   rand.Seed:
	       还函数是用来创建随机数的种子,如果不执行该步骤创建的随机数是一样的，因为默认Go会使用一个固定常量值来作为随机种子。

	   time.Now().UnixNano():
	       当前操作系统时间的毫秒值
	*/
	rand.Seed(time.Now().UnixNano())
	//fmt.Println("数组的长度是",len(Tiezi))
	a := rand.Intn(len(Tiezi)) //实际随机生成的数字范围[0,99]
	//fmt.Printf("a的类型为[%T],a的随机数值为:[%d]\n", a, a)
	fmt.Println("随机数是", a)
	return a
}
func ReadTiezi(path string) []string {
	var tiezitest []string
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	bytes, err := ioutil.ReadFile(wd + "/data/text/" + path)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println("Bytes read: ", len(bytes))
	//fmt.Println("String read: ", string(bytes))
	str := string(bytes)
	split := strings.Split(str, "|")

	for i := 0; i < len(split); i++ {
		tiezitest = append(tiezitest, split[i])
	}

	return tiezitest
}

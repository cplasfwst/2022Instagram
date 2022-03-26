package ins

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type DlIp struct {
	Ret  string `json:"Ret"`
	Data string `json:"Data"`
	Msg  string `json:"Msg"`
}

//更换IP
func ChangeIP(ipproxy string) {
	var shuju DlIp
	resp, err := http.Get("http://refresh.rola.info/refresh?user=" + ipproxy + "&country=tw")
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
		fmt.Println(ipproxy, "更换IP成功")

	} else {
		fmt.Println(ipproxy, "更换IP失败")

	}
}

//倒计时
func CountTime(num int, data sync.Map, renwushu int) {
	if num > 0 {
		//fmt.Println(num)
		//data["INSdaojishi"] = strconv.Itoa(num)
		data.Store("INSdaojishi", strconv.Itoa(num))
		renwushustr := strconv.Itoa(renwushu)
		//data["INSzhuangtai"] = "倒计时还有" + data["INSdaojishi"] + "秒,完成任务总数为:" + renwushustr
		data.Store("INSzhuangtai", "倒计时还有"+MapRead(data, "INSdaojishi")+"秒,完成任务总数为:"+renwushustr+"违规次数是："+MapRead(data, "INSweigui"))
		time.Sleep(time.Duration(1) * time.Second)
		CountTime(num-1, data, renwushu)
	} else {
		fmt.Println("倒计时完成")
		data.Store("INSzhuangtai", "倒计时完成")
	}
	return
}

//获取随机数
func GetRandNum(num int) int {
	/*
	   rand.Seed:
	       还函数是用来创建随机数的种子,如果不执行该步骤创建的随机数是一样的，因为默认Go会使用一个固定常量值来作为随机种子。

	   time.Now().UnixNano():
	       当前操作系统时间的毫秒值
	*/
	rand.Seed(time.Now().UnixNano())
	//fmt.Println("数组的长度是",len(Tiezi_huashu))  num是随机数的最大值
	a := rand.Intn(num) //实际随机生成的数字范围[0,99]
	//fmt.Printf("a的类型为[%T],a的随机数值为:[%d]\n", a, a)
	//fmt.Println("随机数是", a)
	return a
}

//读取对应的发帖文本
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

//导入用户所有信息
func ImportuserMap(filename string) ([]sync.Map, error) {
	var data []sync.Map

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	file, err := os.Open(wd + "/data/user/" + filename)
	if err != nil {
		return data, errors.New("打开账号密码文件失败.")
	}
	defer file.Close()
	bf := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := bf.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				return data, errors.New("读取账号密码文件失败")
			}
			break
		}
		if isPrefix {
			return data, errors.New("行数有点问题")
		}
		str := string(line)
		s := strings.Split(str, "----")

		var user sync.Map
		user.Store("DLhost", s[0])
		user.Store("DLzhanghao", s[1])
		user.Store("DLmima", s[2])
		user.Store("INSzhanghao", s[3])
		user.Store("INSmima", s[4])
		user.Store("cookies", s[5])
		user.Store("UserAgent", s[6])
		user.Store("INSweigui", 0)
		data = append(data, user)
	}

	//fmt.Println(data)
	return data, nil
}

//获取图片数量
func GetPictureCount() int {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	files, _ := ioutil.ReadDir(wd + "/data/jpg")
	//fmt.Println(len(files))
	return len(files)
}

//获取热门关键词
func Inshot() string {

	var Inshottest []string
	var Inshotmain string
	Inshotmain = INSgudinghot
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	bytes, err := ioutil.ReadFile(wd + "/data/text/" + "inshot.txt")
	if err != nil {
		log.Println(err)
	}

	//fmt.Println("Bytes read: ", len(bytes))
	//fmt.Println("String read: ", string(bytes))
	str := string(bytes)
	split := strings.Split(str, "|")

	for i := 0; i < len(split); i++ {
		Inshottest = append(Inshottest, split[i])
	}
	maxTest := len(Inshottest)
	//最后一步添加随机热门词
	for i := 0; i < INSsuijihot; i++ {
		Inshotmain = Inshotmain + Inshottest[GetRandNum(maxTest)] + " "
		time.Sleep(time.Millisecond)
	}

	return Inshotmain

}

//读取sync.Map的值并且转化为string
func MapRead(data sync.Map, key string) string {
	load, ok := data.Load(key)
	if ok != true {

	}

	s := strval(load)

	return s
}

func strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

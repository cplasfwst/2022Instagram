package ins

import "sync"

//完成任务总数和倒计时
var Renwu int
var Insyanchi int

//是否显示页面
var Isheadless bool

//图片的数量
var PictureCount int

//帖子话术
var Tiezi_huashu []string

//账号密码文本
var UserData []sync.Map

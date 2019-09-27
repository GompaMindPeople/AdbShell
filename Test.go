/*
@Time : 2019/9/26 11:46 
@Author : 一条小咸鱼
@File : 
@Software: GoLand
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	url:="http://api.17173g.cn:28080/platform-register/user-register/register.do"
	str := "{\"sign\"=\"33a037a0832fda21a387029c5b5a95b5\",\"style\" = \"2\",\"imei\" = \"666660724468133\",\"package\" = \"org.cocos2d.sqzz\",\"appid\" = \"1efcc1ddd61afead4c699e66baea8fd8\",\"channellevel1\" = \"MHT\",\"androidID\" = \"780e2c5023c135b5\",\"imeiSDK\" = \"666660724468133\",\"channellevel2\" = \"\",\"serial\" = \"null\" }"
	jsonStr, _ := json.Marshal(str)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(jsonStr))


	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(resp)

}




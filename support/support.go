/*
针对adb 的shell 指令进行的一层封装.
@Time : 2019/9/25 17:38
@Author : 一条小咸鱼
@File :
@Software: GoLand
*/
package support

import (
	"io/ioutil"
	"log"
	"os/exec"
)

type CommonInfo struct {

}

type AndroidInfo struct {

}


func TestA(){
	command := exec.Command("ping","192.168.5.234")
	stdout, e := command.StdoutPipe()
	if e != nil{
		log.Fatal(e)
	}

	defer stdout.Close()
	if err := command.Start(); err != nil {   // 运行命令
		log.Fatal(err)

	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {  // 读取输出结果

		log.Fatal(err)

	} else {

		log.Println(string(opBytes))

	}

}
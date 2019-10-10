/*
@Time : 2019/9/27 10:34
@Author : 一条小咸鱼
@File :
@Software: GoLand
*/
package support

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"os/exec"
	"testing"
)

func TestAB(t *testing.T) {

	command := exec.Command("cmd", "/c", "adb shell input")

	enc := mahonia.NewEncoder("gbk")
	stdout, e := command.StdoutPipe()
	//stdin, e := command.StdinPipe()
	fmt.Println("要运行的指令：", command.Args)
	if e != nil {
		log.Fatal(e)
	}

	defer stdout.Close()
	if err := command.Start(); err != nil { // 运行命令
		log.Fatal(err)

	}
	//_, e = stdin.Write([]byte("input"))

	if e != nil {
		log.Fatal(e)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果

		log.Fatal(err)

	} else {

		log.Println(enc.ConvertString(string(opBytes)))
		log.Println("-------------------------------------------------------")
		log.Println(string(opBytes))
	}
	command.Wait()
}

func TestB(t *testing.T) {
	var err error
	cmd := exec.Command("cmd")
	enc := mahonia.NewEncoder("gbk")
	// cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out, errout bytes.Buffer
	cmd.Stdout = &out //绑定输出
	cmd.Stderr = &errout
	go func() {
		// start stop restart
		in.WriteString("adb -s S9B7N17929004960 shell input tap 500 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 600 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 700 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 800 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 600 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 800 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 500 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 400 500\n") //写入你的命令，可以有多行，"\n"表示回车
		in.WriteString("adb -s S9B7N17929004960 shell input tap 100 500\n") //写入你的命令，可以有多行，"\n"表示回车
	}()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cmd---->", cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	rt := out.String() //
	//  mahonia.NewDecoder("gbk").ConvertString(out.String()) //
	fmt.Println(enc.ConvertString(string(rt)))
	fmt.Println("err------------->", enc.ConvertString(string(errout.String())))
	//if strings.ContainsAny(rt, "成功"){
	//         fmt.Fprintf(w, "{\"msg\":\"%s\"}", "操作成功")
	//     }else{
	//         fmt.Fprintf(w, "{\"msg\":\"%s\"}", rt)
	//    }

}

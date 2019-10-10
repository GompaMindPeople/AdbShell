/*
针对 cmd 进行的一层封装
@Time : 2019/9/25 17:38
@Author : 一条小咸鱼
@File :
@Software: GoLand
*/
package support

import (
	"bytes"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"os/exec"
)

type Common struct {
	Cmd                                  *exec.Cmd
	enc                                  *mahonia.Encoder
	InputStream, OutputStream, ErrStream *bytes.Buffer
}

func (c *Common) SetDefault() {
	enc := mahonia.NewEncoder("gbk")
	c.enc = &enc
}

//设置 输入流  输出流  错误流
func (c *Common) ResetStream() {
	c.Cmd.Stdin = c.InputStream
	c.Cmd.Stdout = c.OutputStream
	c.Cmd.Stderr = c.ErrStream
}

func NewCMD() *Common {
	cmd := exec.Command("cmd")
	common := Common{}
	common.Cmd = cmd
	return &common
}

//type AndroidInfo struct {
//
//}

func TestA() {
	command := exec.Command("ping", "192.168.5.234")
	stdout, e := command.StdoutPipe()
	if e != nil {
		log.Fatal(e)
	}

	defer stdout.Close()
	if err := command.Start(); err != nil { // 运行命令
		log.Fatal(err)

	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果

		log.Fatal(err)

	} else {

		log.Println(string(opBytes))

	}

}

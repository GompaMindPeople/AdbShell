/*
@Time : 2019/9/27 10:34 
@Author : 一条小咸鱼
@File : 
@Software: GoLand
*/
package support
import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"os/exec"
	"testing"
)

func TestAB(t *testing.T) {
	command := exec.Command("adb","shell input")

	enc:=mahonia.NewEncoder("gbk")
	stdout, e := command.StdoutPipe()
	//stdin, e := command.StdinPipe()
	fmt.Println("要运行的指令：",command.Args)
	if e != nil{
		log.Fatal(e)
	}

	defer stdout.Close()
	if err := command.Start(); err != nil {   // 运行命令
		log.Fatal(err)

	}
	//_, e = stdin.Write([]byte("input"))

	if e != nil{
		log.Fatal(e)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {  // 读取输出结果

		log.Fatal(err)

	} else {

		log.Println(enc.ConvertString(string(opBytes)))
		log.Println("-------------------------------------------------------")
		log.Println(string(opBytes))
	}
	command.Wait()
}




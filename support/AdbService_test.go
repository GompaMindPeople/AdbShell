package support

import (
	"bytes"
	"fmt"
	"testing"
)

func TestADBService_GetDevices(t *testing.T) {

	service := ADBService{}
	//for a:=0;a<5 ;a++  {
	cmd := NewCMD()
	service.CMD = cmd
	var out, errout bytes.Buffer
	cmd.InputStream = bytes.NewBuffer(nil)
	cmd.OutputStream = &out
	cmd.ErrStream = &errout
	cmd.SetDefault()
	cmd.ResetStream()
	enc := cmd.enc
	err := cmd.Cmd.Start()
	if err != nil {
		fmt.Print("err-->", err)
		return
	}
	//service.GetDevices()

	//service.SetSimulator("abc5fe6b")
	////service.TestA()
	//i  := 0
	//service.SetSimulator("abc5fe6b")
	//fmt.Println(enc.ConvertString(string(out.String())))
	//err = cmd.Cmd.Start()

	//service.Execute("","adb shell\n",&i)
	////service.Execute("","adb -s S9B7N17929004960 shell input tap 600 500\n",&i)
	////service.Execute("","adb -s S9B7N17929004960 shell input tap 550 500\n",&i)
	//service.Execute("","adb -s S9B7N17929004960 shell input tap 700 500\n",&i)
	//service.Execute("","input 100 100 \n",&i)
	cmd.InputStream.WriteString("adb shell\n")
	cmd.InputStream.WriteString("input tap 500 500 \n")
	//time.Sleep(time.Duration(2) * time.Second)
	cmd.InputStream.WriteString("input tap 600 600 \n")
	//service.TestA()
	//service.GetDevices()
	err = cmd.Cmd.Wait()
	fmt.Println(enc.ConvertString(string(out.String())))
	fmt.Println("errout-->", enc.ConvertString(string(errout.String())))
	fmt.Println("err----->", err)
	//}

}

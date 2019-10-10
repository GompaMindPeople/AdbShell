package support

type ADBService struct {
	CMD *Common
}

func (adb *ADBService) SetSimulator(SimulatorName string) {
	InputBuffer := adb.CMD.InputStream
	inst := SET_SIMULATOR + SimulatorName + " shell\n"
	InputBuffer.WriteString(inst)
}

func (adb *ADBService) GetDevices() {
	//err := adb.CMD.Cmd.Start()
	InputBuffer := adb.CMD.InputStream
	inst := GET_DEVICES + "\n"
	InputBuffer.WriteString(inst)
}

func (adb *ADBService) Execute(simulatorName, instr string, a *int) {
	//err := adb.CMD.Cmd.Start()
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//
	//}

	adb.CMD.InputStream.WriteString(instr)
	//err = adb.CMD.Cmd.Wait()
	//adb.CMD.Cmd.Process = nil
	//
	//value := reflect.ValueOf(adb.CMD.Cmd).Elem()
	////TODO: fuck 私有属性 无法赋值 真是日了!
	//fmt.Println(value.Field(15).CanSet())

	//adb.CMD.Cmd.
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//
	//}

}

//func

func (adb *ADBService) TestA() {

	InputBuffer := adb.CMD.InputStream
	InputBuffer.WriteString("dumpsys window displays\n")

}

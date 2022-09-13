package tools

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"io/ioutil"
	"os/exec"
	"runtime"
	"video_storage/model"
)

// 执行命令
func Execute(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: " + err.Error())
		return "", err
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return "", err
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if nil != err {
		return "", err
	}
	bytesOut, err := ioutil.ReadAll(stdout)
	if nil != err {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return "", err
	}
	if 0 != len(bytesErr) {
		return "", errors.New(string(bytesErr))
	}

	return string(bytesOut), nil
}

// 获得设备基础信息
func GetDeviceInfo() *model.DeviceInfo {
	runTimeDeviceInfo := &model.DeviceInfo{}
	runTimeDeviceInfo.PlatformFamily = runtime.GOOS
	runTimeDeviceInfo.KernelArch = runtime.GOARCH
	runTimeDeviceInfo.CPUs = runtime.GOMAXPROCS(0)
	return runTimeDeviceInfo
}

func GetDeviceFullInfo() *model.RunTimeDeviceInfo {
	runTimeDeviceInfo := &model.RunTimeDeviceInfo{}
	c, _ := cpu.Info()
	v, _ := mem.VirtualMemory()
	n, _ := net.Interfaces()

	// process cpu info
	var processorArr = make([]model.Processor, len(c))
	for i, v := range c {
		processorArr[i].Cores = v.Cores
		processorArr[i].Name = v.ModelName
		processorArr[i].CacheSize = v.CacheSize
		processorArr[i].Microcode = v.Microcode
	}
	runTimeDeviceInfo.CPUs = processorArr

	// process memory info
	runTimeDeviceInfo.Memory.Total = v.Total
	runTimeDeviceInfo.Memory.Available = v.Available
	runTimeDeviceInfo.Memory.Used = v.Used
	runTimeDeviceInfo.Memory.UsedPercent = v.UsedPercent
	runTimeDeviceInfo.Memory.Free = v.Free

	// process net interface
	for _, netInterface := range n {

		for i, v := range n {
			if 0 != len(v.Addrs) && "" != v.HardwareAddr {
				if len(v.Addrs) > len(netInterface.Addrs) {
					netInterface = n[i]
				}
			}
		}

		var addrGroup = make([]string, len(netInterface.Addrs))
		for i, a := range netInterface.Addrs {
			addrGroup[i] = a.Addr
		}
		runTimeDeviceInfo.NetInterface = append(runTimeDeviceInfo.NetInterface, struct {
			MTU       int      `json:"mtu"`
			Name      string   `json:"name"`
			Flags     []string `json:"flags"`
			AddrGroup []string `json:"addrGroup"`
		}{
			netInterface.MTU,
			netInterface.Name,
			netInterface.Flags,
			addrGroup,
		})
	}

	runTimeDeviceInfo.MacID = Sha256Hex(n[0].HardwareAddr)

	return runTimeDeviceInfo
}

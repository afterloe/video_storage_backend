package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type (
	receive struct {
		command    string
		response   string
		errMessage string
	}
)

func (that *receive) Splice(charsets ...string) []string {
	if len(charsets) != 0 {
		return strings.Split(that.response, charsets[0])
	}
	return strings.Split(that.response, "\n")
}

func (that *receive) ToString(charsets ...string) string {
	if len(charsets) != 0 {
		var newChar string
		if len(charsets) > 1 {
			newChar = charsets[1]
		} else {
			newChar = ""
		}
		return strings.ReplaceAll(that.response, charsets[0], newChar)
	}
	return that.response
}

func (that *receive) GetError() error {
	if that.errMessage != "" {
		fmt.Println("[ERROR]: Execute ", that.command)
		return errors.New(that.errMessage)
	}
	return nil
}

func (that *receive) HasError() bool {
	if that.errMessage != "" {
		return true
	}
	return false
}

func ExecuteWithOutError(command string, args ...interface{}) *receive {
	r := new(receive)
	if len(args) != 0 {
		command = fmt.Sprintf(command, args...)
	}
	cmd := exec.Command("/bin/sh", "-c", command)
	r.command = command
	stdout, _ := cmd.StdoutPipe()
	_ = cmd.Start()
	bytesOut, _ := ioutil.ReadAll(stdout)
	r.response = string(bytesOut)
	return r
}

func Execute(command string, args ...interface{}) *receive {
	r := new(receive)
	if len(args) != 0 {
		command = fmt.Sprintf(command, args...)
	}
	r.command = command
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		r.errMessage = err.Error()
		return r
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		r.errMessage = err.Error()
		return r
	}
	if err := cmd.Start(); err != nil {
		r.errMessage = err.Error()
		return r
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if nil != err {
		r.errMessage = err.Error()
		return r
	}
	bytesOut, err := ioutil.ReadAll(stdout)
	if nil != err {
		r.errMessage = err.Error()
		return r
	}
	if err := cmd.Wait(); err != nil {
		r.errMessage = err.Error()
		return r
	}
	if len(bytesErr) != 0 {
		r.errMessage = string(bytesErr)
		return r
	}
	r.response = string(bytesOut)
	return r
}

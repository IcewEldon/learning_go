//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/9/3 11:45 AM
//Copyright 2021 www.ucloud.cn

package error

import (
	"errors"
)

var (
	//pre errors information
	//Uncertain mistakes， just declared
	ErrParmeterMissing    = errors.New("error : missing parameter")
	ErrOsArgMissingOrLess = errors.New("error : missing or lack of parameter")
	ErrConnectTimeOut	  = errors.New("error : connection timeout")
)

//const variable
const (
	//pre errors template
	ErrExpression = "error level -> : %v, info -> : %v"
)

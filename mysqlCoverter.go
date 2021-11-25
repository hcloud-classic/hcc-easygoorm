package easygoorm

import (
	"reflect"
	"time"
)

// convertInterfaceToModelType :
func convertInterfaceToModelType(interfaceValue interface{}) (bool, interface{}) {
	switch reflect.TypeOf(interfaceValue).String() {
	case "int64":
		arg := (interfaceValue).(int64)
		retVal := (int)(arg)
		return true, &retVal
	case "[]uint8":
		arg := (interfaceValue).([]uint8)
		retVal := (string)(arg)
		return true, &retVal
	case "time.Time":
		retVal := (interfaceValue).(time.Time)
		return true, &retVal
	default:
		goto ERROR
	}
ERROR:
	return false, nil
}

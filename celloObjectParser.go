package easygoorm

import (
	"fmt"
	"time"

	"innogrid.com/hcloud-classic/model"
)

// ModelInterfaceMapping :
func CelloSqlModelInterfaceMapping(columName string, interfaceValue interface{}, volume *model.Volume) bool {
	_, retVal := convertInterfaceToModelType(interfaceValue)
	switch columName {
	case "uuid":
		volume.UUID = (*retVal.(*string))
	case "group_id":
		volume.GroupID = (int)(*retVal.(*int))
	case "size":
		volume.Size = (int)(*retVal.(*int))
	case "filesystem":
		volume.Filesystem = *retVal.(*string)
	case "server_uuid":
		volume.ServerUUID = *retVal.(*string)
	case "use_type":
		volume.UseType = *retVal.(*string)
	case "user_uuid":
		volume.UserUUID = *retVal.(*string)
	case "lun_num":
		volume.LunNum = (int)(*retVal.(*int))
	case "state":
		volume.State = *retVal.(*string)
	case "pool":
		volume.Pool = *retVal.(*string)
	case "created_at":
		volume.CreatedAt = *retVal.(*time.Time)
	default:
		goto ERROR
	}
	return true
ERROR:
	return false
}

func celloSQLValueBuilder(count int, in *model.Volume) []interface{} {
	var inetrnalCount int
	retInterface := make([]interface{}, count)
	fmt.Println("retSql in=> ", in)
	if in.UUID != "" {
		retInterface[inetrnalCount] = in.UUID
		inetrnalCount++
	}
	if in.GroupID > 0 {
		retInterface[inetrnalCount] = in.GroupID
		inetrnalCount++
	}
	if in.Size > 0 {
		retInterface[inetrnalCount] = in.Size
		inetrnalCount++
	}
	if in.Filesystem != "" {
		retInterface[inetrnalCount] = in.Filesystem
		inetrnalCount++
	}
	if in.ServerUUID != "" {
		retInterface[inetrnalCount] = in.ServerUUID
		inetrnalCount++
	}
	if in.UseType != "" {
		retInterface[inetrnalCount] = in.UseType
		inetrnalCount++
	}
	if in.UserUUID != "" {
		retInterface[inetrnalCount] = in.UserUUID
		inetrnalCount++
	}
	if in.LunNum >= 0 {
		retInterface[inetrnalCount] = in.LunNum
		inetrnalCount++
	}
	if in.Pool != "" {
		retInterface[inetrnalCount] = in.Pool
		inetrnalCount++
	}
	if in.State != "" {
		retInterface[inetrnalCount] = in.State
		inetrnalCount++
	}

	// fmt.Println("retSql => ", retSql)
	return retInterface
}

func CelloInsertSQLBuilder(in *model.Volume) (string, []interface{}) {
	var (
		insertColums       string
		insertStringValues string
		retSql             string
		count              int
	)
	retInterface := make([]interface{}, 13)
	if in.UUID != "" {
		insertColums += "uuid, "
		insertStringValues += "?, "
		retInterface[count] = in.UUID
		count++
	}
	if in.GroupID > 0 {
		insertColums += "group_id, "
		insertStringValues += "?, "
		retInterface[count] = in.GroupID
		count++
	} else {
		goto ERROR
	}
	if in.Size > 0 {
		insertColums += "size, "
		insertStringValues += "?, "
		retInterface[count] = in.Size
		count++
	}
	if in.Filesystem != "" {
		insertColums += "filesystem, "
		insertStringValues += "?, "
		retInterface[count] = in.Filesystem
		count++
	}
	if in.ServerUUID != "" {
		insertColums += "server_uuid, "
		insertStringValues += "?, "
		retInterface[count] = in.ServerUUID
		count++
	}
	if in.UseType != "" {
		insertColums += "use_type, "
		insertStringValues += "?, "
		retInterface[count] = in.UseType
		count++
	}
	if in.UserUUID != "" {
		insertColums += "user_uuid, "
		insertStringValues += "?, "
		retInterface[count] = in.UserUUID
		count++
	}
	if in.LunNum >= 0 {
		insertColums += "lun_num, "
		insertStringValues += "?, "
		retInterface[count] = in.LunNum
		count++
	} else {
		goto ERROR
	}
	if in.Pool != "" {
		insertColums += "pool, "
		insertStringValues += "?, "
		retInterface[count] = in.Pool
		count++
	}
	if in.State != "" {
		insertColums += "state, "
		insertStringValues += "?, "
		retInterface[count] = in.State
		count++
	}

	retSql = "insert into volume(" + insertColums + "created_at) values (" + insertStringValues + "now())"
	fmt.Println("retSql => ", retSql)
	return retSql, celloSQLValueBuilder(count, in)
ERROR:
	return "", nil
}

func CelloUpdateSQLBuilder(in *model.Volume) (string, []interface{}) {
	var (
		insertColums       string
		insertStringValues string
		retSql             string
		count              int
	)
	retInterface := make([]interface{}, 13)
	if in.UUID != "" {
		insertColums += "uuid, "
		insertStringValues += "?, "
		retInterface[count] = in.UUID
		count++
	}
	if in.GroupID > 0 {
		insertColums += "group_id, "
		insertStringValues += "?, "
		retInterface[count] = in.GroupID
		count++
	} else {
		goto ERROR
	}
	if in.Size > 0 {
		insertColums += "size, "
		insertStringValues += "?, "
		retInterface[count] = in.Size
		count++
	}
	if in.Filesystem != "" {
		insertColums += "filesystem, "
		insertStringValues += "?, "
		retInterface[count] = in.Filesystem
		count++
	}
	if in.ServerUUID != "" {
		insertColums += "server_uuid, "
		insertStringValues += "?, "
		retInterface[count] = in.ServerUUID
		count++
	}
	if in.UseType != "" {
		insertColums += "use_type, "
		insertStringValues += "?, "
		retInterface[count] = in.UseType
		count++
	}
	if in.UserUUID != "" {
		insertColums += "user_uuid, "
		insertStringValues += "?, "
		retInterface[count] = in.UserUUID
		count++
	}
	if in.LunNum >= 0 {
		insertColums += "lun_num, "
		insertStringValues += "?, "
		retInterface[count] = in.LunNum
		count++
	} else {
		goto ERROR
	}
	if in.Pool != "" {
		insertColums += "pool, "
		insertStringValues += "?, "
		retInterface[count] = in.Pool
		count++
	}
	if in.State != "" {
		insertColums += "state, "
		insertStringValues += "?, "
		retInterface[count] = in.State
		count++
	}

	retSql = "insert into volume(" + insertColums + "created_at) values (" + insertStringValues + "now())"
	fmt.Println("retSql => ", retSql)
	return retSql, celloSQLValueBuilder(count, in)
ERROR:
	return "", nil
}

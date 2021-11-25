# Concept

## Version Control

Flyway를 통한 DB의 버전 관리를 지원 한다.

Flyway는 java진영에서 사용하기 때문에 Spring Framework에서는 코드내에서 Entity의 관리가 가능하지만, Golang에서는 이를 지원하지 않는다.

하지만 Flyway의 Core기능은 DB의 버전 관리 이기 때문에 스키마에 대해서는 버전 관리가 가능하다.

사용자가 필수 데이터에 대한 정의만 해둔다면 초기 상태에서의 데이터 입력도 가능하다.

## ORM(easygoorm)

Golang 에서는 오픈소스 혹은 자체적으로 ORM을 지원하지 않는다.

매번 데이터베이스의 컬럼 혹은 구조가 바뀔때 마다 각 모델과 쿼리 호출 인자를 변경하는것은 상당히 불편하기 때문에 Object 처럼 다룰 수 있도록 기능을 지원한다.

# Usage

## ORM

easyGoORM라이브러리는 Golang 의 Interface로 받아지는 Query의 조회 결과를 Model의 자료구조에 맞게 변환해 준다.

### Prerequitsite

>  쿼리의 결과값을 모델로 변환 받기 위해 해당 모델의 정의가 필요



### func ModelInterfaceMapping()

해당 함수는 쿼리를 호출 할때 리턴되는 결과값을 인터페이스로 받아 이를 파싱해준다.

다음 3가지를 인자 값으로 사용한다.

* columName string

  조회된 컬럼의 이름

* interfaceValue interface{}

  조회된 값의 이차원 인터페이스 슬라이스

* someModel *model.someModel

  변환받고자 하는 모델의 주소

`Function Basic Structure`

해당 함수에서 각 마이크로 서비스에 맞게 모델에 적합한 인자값으로 추가 및 수정 해 주면된다.

```go

// ModelInterfaceMapping :
func ModelInterfaceMapping(columName string, interfaceValue interface{}, someModel *model.someModel) bool {
	_, retVal := convertInterfaceToModelType(interfaceValue)
	switch columName {
	case "uuid":
		someModel.SomeValue = (*retVal.(*someType))
	default:
		goto ERROR
	}
	return true
ERROR:
	return false
}
```

### func  convertInterfaceToModelType()

현재 개발된 서비스들의 모델을 고려 할 때, Mysql 에서 리턴되는 타입은 크게 3가지 이다.

int64는 int 이며, []uint8은 string, time.time

`Function Basic Structure`

```go
// convertInterfaceToModelType :
func convertInterfaceToModelType(interfaceValue interface{}) (bool, interface{}) {
	switch reflect.TypeOf(interfaceValue).String() {
	case "int64":
		arg := (interfaceValue).(int64)
		retVal, _ := strconv.Atoi(string(arg))
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
```



위 두 함수를 사용해서 인터페이스로 리턴된 값을 적절한 모델에 맞게 파싱해 준다.

이를 사용하기 위해서는 쿼리를 호출하는 부분에 몇가지 추가해야 될 과정이 필요 하다.

### func  SomeQuery()

```go
func SomeQuery() {
    var (
    someModel model.SomeModel
    someModels []model.SomeModel
    )
	sql := "select * from someTable"
	sql += " order by created_at asc limit ? offset ?"
	stmt, err = mysql.Db.Query(sql, row, 0)
	if err != nil {
		logger.Logger.Println(err)
		ErrStr = "[SQL] SomeQuery() " + err.Error()
		ErrCode = hcc_errors.SomeServiceSQLOperationFail
		goto ERROR
	}
	defer func() {
		_ = stmt.Close()
	}()

	cols, _ = stmt.Columns()
	vals = make([]interface{}, len(cols))
	valsptr = make([]interface{}, len(cols))
	for i := range vals {
		// var ii interface{}
		valsptr[i] = &vals[i]
	}
	for stmt.Next() {
		err := stmt.Scan(valsptr...)
		colType, _ := stmt.ColumnTypes()
		if err != nil {
			logger.Logger.Println(err.Error())
             ErrStr = "[SQL] SomeQuery() " + err.Error()
             ErrCode = hcc_errors.SomeServiceSQLOperationFail
			goto ERROR
		}
		for i, v := range vals {
			modelInterfaceMapping(colType[i].Name(), v, &someModel)
		}
		someModels = append(someModels, someModel)
	}
	return someModels, 0, ""
ERROR:
	return someModels, ErrCode, ErrStr
```

중요하게 봐야 할 부분은 쿼리의 호출 결과를 받은 Ineterface형의 Slice 이다.

* Interface Slice Initialize
  하단의 코드를 살펴보면, `valsptr[]` slice를 확인 할 수 있다.

  쿼리를 호출하게 되면 다양한 타입의 컬럼 값들이 리턴 되기 때문에 2차원의 slice를 생성하여 값을 받아야 한다.

  ```go
  	cols, _ = stmt.Columns()
  	vals = make([]interface{}, len(cols))
  	valsptr = make([]interface{}, len(cols))
  	for i := range vals {
  		// var ii interface{}
  		valsptr[i] = &vals[i]
  	}
  ```

* Interface Slice in Query Scanner

  ModelInterfaceMapping 함수를 사용하기 위해서는 3가지 인자가 필요 하다.

  현재 조회된 컬럼의 이름, 조회된 결과의 이차원 Interface, 변환하고자 하는 모델의 주소 이다.

  ```go
  	for stmt.Next() {
  		err := stmt.Scan(valsptr...)
  		colType, _ := stmt.ColumnTypes()
  		...
  		for i, v := range vals {
  			ModelInterfaceMapping(colType[i].Name(), v, &someModel)
  		}
  		...
  	}
  ```

  


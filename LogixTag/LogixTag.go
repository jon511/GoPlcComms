package LogixTag

import "fmt"

type Controller struct {

}

type ErrorCode struct {
	Code int
	ExtCode int
	Description string
}

type DataType struct {
	Val int
	Size int
	fragReadSize int
	fragWritSize int
}

var (
	BOOL = DataType{0xc1, 1, 0, 0}
	SINT = DataType{0xc2, 1, 490, 200}
	INT = DataType{0xc3, 2,244,100}
	DINT = DataType{0xc4, 4,122, 50}
	REAL = DataType{0xca, 4, 122, 50}
	DWORD = DataType{0xd3, 4, 122, 50}
	LINT = DataType{0xc5,8, 60, 60}
)

type LogixTag struct {
	Name string
	Status int
	Value interface{}
	UserData interface{}
	Length int
	Controller
	DataType
	ErrorCode

}

func (l LogixTag) Read(){
	if l.Controller == (Controller{}){
		fmt.Println("not implemented")
	}
}
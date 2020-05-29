package vm

import (
	"regexp"
	"strconv"
)

type FieldType int

const (
	StringType FieldType = iota
	NumberType
)

type Field struct {
	Type FieldType
	Val  string
}

const digitRegexp = `(\-|\+)?\d+(\.\d*)?((e|E)(\-|\+)?\d+)?|(\-|\+)?\.\d+((e|E)(\-|\+)?\d+)?`

func NewField(val string) Field {
	var ft FieldType
	re := regexp.MustCompile(`^` + digitRegexp + `$`)
	//re := regexp.MustCompile(`^` + digitRegexp)
	numStr := re.FindString(val)
	if len(numStr) == len(val) {
		ft = NumberType
	} else {
		ft = StringType
	}
	return Field{
		Type: ft,
		Val:  val,
	}
}

func NewStringField(val string) Field {
	return Field{
		Type: StringType,
		Val:  val,
	}
}

func (f Field) String() string {
	return f.Val
}

func (f Field) Number() float64 {
	re := regexp.MustCompile(`^` + digitRegexp)
	numStr := re.FindString(f.Val)
	if len(numStr) == 0 {
		return 0
	}
	num, _ := strconv.ParseFloat(numStr, 64)
	return num
}

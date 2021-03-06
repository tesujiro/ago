package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		{script: "", result: nil},
		{script: "{x}", result: nil},
		{script: "{x=x+1;}", result: nil},
		//{script: "${x;}", result: nil},
		{script: "{x;x;}", result: nil},
		{script: "BEGIN{x;}{x=1;y=x+1;}END{x;}", result: nil},
		{script: "a==1{x}", result: nil},
		{script: "$1==1{x}", result: nil},
		{script: "$NR==1{x}", result: nil},
		{script: "$(NF+1)==1{x}", result: nil},
	}
	for _, test := range tests {
		//fmt.Println("*************************\nTEST SCRIPT:", test.script)
		_, parseError := ParseSrc(test.script)
		if parseError != nil {
			if test.errMessage == "" || parseError.Error() != test.errMessage {
				t.Errorf("Run error:%#v want%#v - script:%v\n", parseError, test.errMessage, test.script)
			}
			continue
		}
		//Dump(ast)
	}

}

package vm

import (
	"errors"
	"fmt"
	"regexp"
)

type builtin struct {
	//ARGC,ARGV,FILENAME
	NF, NR  int
	FS, OFS string
	ORS     string
	SUBSEP  string
	//OFMT,FNR,
	RLENGTH, RSTART int
	//RS
	//ENVIRON
	//CONVFMT
	field []string
}

func NewBuiltIn() *builtin {
	return &builtin{
		SUBSEP: string([]byte{0x1c}),
		ORS:    "\n",
		OFS:    " ",
	}
}

func (e *Env) setNR(i int) {
	e.builtin.NR = i
}

func (e *Env) SetNF() {
	e.builtin.NF = len(e.builtin.field) - 1
}

func (e *Env) SetFS(fs string) {
	e.builtin.FS = fs
	//e.Dump()
}

func (e *Env) SetOFS(fs string) {
	e.builtin.OFS = fs
	//e.Dump()
}

func (e *Env) SetORS(fs string) {
	e.builtin.ORS = fs
	//e.Dump()
}

func (e *Env) SetSUBSEP(ss string) {
	e.builtin.SUBSEP = ss
}

func (e *Env) GetField(i int) (string, error) {
	// TODO: out of index
	if i < 0 || i >= len(e.builtin.field) {
		return "", nil
	}
	return e.builtin.field[i], nil
}

func (e *Env) SetFieldZero() error {
	//fmt.Println("SetFieldZero:", e.builtin.field)
	if len(e.builtin.field) <= 1 {
		e.builtin.field[0] = ""
		return nil
	}
	str := e.builtin.field[1]
	//fmt.Println("len:", len(e.builtin.field))
	for i := 2; i < len(e.builtin.field); i++ {
		str += e.builtin.OFS + e.builtin.field[i]
	}
	e.builtin.field[0] = str
	e.SetNF()
	return nil
}

func (e *Env) SetField(index int, str string) error {
	if index <= 0 {
		return fmt.Errorf("Field Index Out of Range:%v", index)
	}
	if index > len(e.builtin.field) {
		newField := make([]string, index+1)
		for i := 1; i < len(e.builtin.field); i++ {
			newField[i] = e.builtin.field[i]
		}
		e.builtin.field = newField
	}
	e.builtin.field[index] = str
	e.SetFieldZero()
	return nil
}

var re_org_awk_truncate = regexp.MustCompile("^[ \t]*([^ \t].*[^ \t])[ \t]*$")

func (e *Env) SetFieldFromLine(line string) error {
	split := func(regex, line string) {
		re := regexp.MustCompile(regex) //TODO: STORE PRE COMPILED VALUE TO ENV FOR PERFORMANCE
		result := re.Split(line, -1)
		e.builtin.field = make([]string, len(result)+1)
		for i, f := range result {
			e.builtin.field[i+1] = f
		}
	}
	switch e.builtin.FS {
	case "":
		return errors.New("Field Seaparotor not set")
	case " ":
		//THIS IS SPECIAL CASE FOR ORIGINAL AWK
		//fmt.Printf("line %v:[%v]\n", e.builtin.NR, line)
		line = re_org_awk_truncate.ReplaceAllString(line, "$1")
		//fmt.Printf("line %v:[%v]\n", e.builtin.NR, line)
		split("[ \t]+", line)
	default:
		//fmt.Printf("line %v:FS[%v]\n", e.builtin.NR, e.builtin.FS)
		split(e.builtin.FS, line)
	}
	e.builtin.field[0] = line
	//e.SetFieldZero()

	return nil
}

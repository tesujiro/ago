package vm

import (
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
	field           []string
	inStartStopLoop bool
}

func NewBuiltIn() *builtin {
	return &builtin{
		SUBSEP: string([]byte{0x1c}),
		ORS:    "\n",
		OFS:    " ",
	}
}

func (e *Env) SetNR(i int) {
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

func (e *Env) SetRLENGTH(i int) {
	e.builtin.RLENGTH = i
}

func (e *Env) SetRSTART(i int) {
	e.builtin.RSTART = i
}

func (e *Env) GetField(i int) (string, error) {
	// TODO: out of index
	if i < 0 || i >= len(e.builtin.field) {
		return "", nil
	}
	return e.builtin.field[i], nil
}

/*
func (e *Env) GetFieldPtr(i int) (*string, error) {
	// TODO: out of index
	if i < 0 || i >= len(e.builtin.field) {
		return nil, nil
	}
	return &e.builtin.field[i], nil
}
*/

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
	if index < 0 {
		return fmt.Errorf("Field Index Out of Range:%v", index)
	}
	if index >= len(e.builtin.field) {
		newField := make([]string, index+1)
		for i := 1; i < len(e.builtin.field); i++ {
			//newField[i] = getField(i)
			//f, _ := e.GetField(i)
			//newField[i] = f
			newField[i] = e.builtin.field[i]
		}
		e.builtin.field = newField
	}
	e.builtin.field[index] = str
	if index > 0 {
		e.SetFieldZero()
	}
	return nil
}

var re_org_awk_truncate = regexp.MustCompile("^[ \t]*([^ \t].*[^ \t])[ \t]*$")

func (e *Env) SetFieldFromLine(line string) error {
	split := func(regex, line string) {
		re := regexp.MustCompile(regex) //TODO: STORE PRE COMPILED VALUE TO ENV FOR PERFORMANCE
		result := re.Split(line, -1)
		//fmt.Printf("SetFieldFromLine split result:%#v\n", result)
		if len(result) == 1 && result[0] == "" {
			result = result[:0]
		}
		if e.builtin.FS == " " && len(result) == 2 && result[0] == "" && result[1] == "" {
			result = result[:0]
		}
		e.builtin.field = make([]string, len(result)+1)
		for i, f := range result {
			e.builtin.field[i+1] = f
		}
	}
	switch e.builtin.FS {
	case "":
		e.builtin.field = make([]string, len(line)+1)
		for i, r := range line {
			//fmt.Printf("c:%s\n", string(c))
			e.builtin.field[i+1] = string(r)
		}
	case " ":
		//THIS IS SPECIAL CASE FOR ORIGINAL AWK
		line = re_org_awk_truncate.ReplaceAllString(line, "$1")
		split("[ \t]+", line)
	default:
		//fmt.Printf("line %v:FS[%v]\n", e.builtin.NR, e.builtin.FS)
		split(e.builtin.FS, line)
	}
	e.builtin.field[0] = line
	//e.SetFieldZero()
	e.SetNF()

	return nil
}

func (e *Env) GetLoop() bool {
	return e.builtin.inStartStopLoop
}

func (e *Env) SetLoop(b bool) {
	e.builtin.inStartStopLoop = b
	return
}

package vm

import (
	"fmt"
	"regexp"
	"strconv"
)

type builtin struct {
	//ARGC,ARGV,FILENAME
	NF, NR  int
	FS, OFS string
	ORS     string
	SUBSEP  string
	OFMT    string
	//FNR,
	RLENGTH, RSTART int
	RS              string
	//ENVIRON
	//CONVFMT
	field           []string
	inStartStopLoop bool
}

// newBuiltIn returns new builtin variables.
func newBuiltIn() *builtin {
	return &builtin{
		SUBSEP: string([]byte{0x1c}),
		ORS:    "\n",
		OFS:    " ",
		OFMT:   "%.6g",
	}
}

// TODO: repeated names
func (e *Env) isBuiltin(k string) bool {
	switch k {
	case "NF", "NR", "FS", "OFS", "ORS", "SUBSEP", "OFMT", "RLENGTH", "RSTART", "RS":
		return true
	default:
		return false
	}
}

// SetNR sets built in variable NR, number of records.
func (e *Env) SetNR(i int) {
	e.builtin.NR = i
}

// SetNF sets built in variable NF, number of fields.
func (e *Env) SetNF() {
	l := len(e.builtin.field)
	if l > 0 {
		e.builtin.NF = len(e.builtin.field) - 1
	} else {
		e.builtin.NF = 0
	}
}

// SetFS sets built in variable FS, field separator.
func (e *Env) SetFS(fs string) {
	e.builtin.FS = fs
	//e.Dump()
}

// SetRLENGTH sets built in variable RLENGTH, the length of the match.
func (e *Env) SetRLENGTH(i int) {
	e.builtin.RLENGTH = i
}

// SetRSTART sets built in variable RSTART, the location in the string of the search pattern.
func (e *Env) SetRSTART(i int) {
	e.builtin.RSTART = i
}

// GetField gets the field 0 value or $0.
func (e *Env) GetFieldZero() (string, error) {
	return e.builtin.field[0], nil
}

// GetField gets the field value with specified index. ex: $1, $NF, $i
func (e *Env) GetField(i int) (interface{}, error) {
	if i < 0 || i >= len(e.builtin.field) {
		return "", nil
	}
	field := e.builtin.field[i]
	digit := `(\-|\+)?\d+(\.\d*)?([e|E]\d+)?|(\-|\+)?\.\d+([e|E]\d+)?`
	re := regexp.MustCompile(`^` + digit + `$`)
	numStr := re.FindString(field)
	if len(numStr) == 0 || numStr != field {
		return field, nil
	}
	fnum, _ := strconv.ParseFloat(numStr, 64)
	inum, err := strconv.ParseInt(numStr, 10, 64)
	if err == nil {
		return inum, nil
	}
	return fnum, nil
}

// SetFieldZero sets the value of the field zero or $0.
func (e *Env) SetFieldZero() error {
	str := e.builtin.field[1]
	for i := 2; i < len(e.builtin.field); i++ {
		str += e.builtin.OFS + e.builtin.field[i]
	}
	e.builtin.field[0] = str
	e.SetNF()
	return nil
}

// SetField sets the value of the field with the specified index.
func (e *Env) SetField(index int, str string) error {
	if index < 0 {
		return fmt.Errorf("Field Index Out of Range:%v", index)
	}
	if index >= len(e.builtin.field) {
		newField := make([]string, index+1)
		for i := 1; i < len(e.builtin.field); i++ {
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

var regexOrgAwkTruncate = regexp.MustCompile("^[ \t]*([^ \t].*[^ \t])[ \t]*$")

// SetFieldFromLine split line string and sets the value of the fields.
func (e *Env) SetFieldFromLine(line string) {
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
	/*
		case "":
			e.builtin.field = make([]string, len([]rune(line))+1)
			for i, r := range []rune(line) {
				e.builtin.field[i+1] = string(r)
			}
	*/
	case " ":
		// SPECIAL CASE FOR ORIGINAL AWK
		newline := regexOrgAwkTruncate.ReplaceAllString(line, "$1")
		split("[ \t]+", newline)
	default:
		//fmt.Printf("line %v:FS[%v]\n", e.builtin.NR, e.builtin.FS)
		split(e.builtin.FS, line)
	}
	e.builtin.field[0] = line
	e.SetNF()

	return
}

// GetLoop returns if current line is inside start and stop scope.
func (e *Env) GetLoop() bool {
	return e.builtin.inStartStopLoop
}

// SetLoop reverses the env flag represents inside start and stop scope.
func (e *Env) SetLoop(b bool) {
	e.builtin.inStartStopLoop = b
	return
}

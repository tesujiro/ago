package vm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func openNextFile(e *Env) error {
	//fmt.Printf("curFileIndex=%v\n", e.fileInfo.curFileIndex)
	var fileReader *os.File
	var file string
	//fmt.Printf("len(files)==%v\n", len(e.fileInfo.files))
	if len(e.fileInfo.files) == 0 {
		fileReader = os.Stdin
		file = "-"
	} else {
		if e.fileInfo.curFileCloser != nil {
			if err := (*e.fileInfo.curFileCloser).Close(); err != nil {
				return err
			}
			e.fileInfo.curFileCloser = nil
			e.fileInfo.curFileScanner = nil
		}
		if e.fileInfo.curFileIndex >= len(e.fileInfo.files) {
			//fmt.Println("return EOF")
			return io.EOF
		}
		file = e.fileInfo.files[e.fileInfo.curFileIndex]
		e.fileInfo.curFileIndex++
		//fmt.Printf("file=%v\n", file)
		var err error
		fileReader, err = os.Open(file)
		if err != nil {
			return fmt.Errorf("open %v error: %v", file, err)
		}
	}
	e.SetFILENAME(file)
	e.ResetFNR()

	readCloser := io.ReadCloser(fileReader)
	e.fileInfo.curFileCloser = &readCloser
	e.fileInfo.curFileScanner = bufio.NewScanner(io.Reader(readCloser))
	err := e.setScannerSplit("")
	if err != nil {
		return err
	}
	//fmt.Println("openNextFile finished normally")
	return nil
}

// SetPseudoStdin set pseudo stdin, read from string
func (e *Env) SetPseudoStdin(data string) error {
	e.fileInfo.curFileScanner = bufio.NewScanner(strings.NewReader(data))
	return nil
}

// GetLine read one line from input files. If EOF and no file to read, return io.EOF.
func (e *Env) NextFile() error {
	return openNextFile(e)
}

// GetLine read one line from input files. If EOF and no file to read, return io.EOF.
func (e *Env) GetLine() (string, error) {
	//fmt.Println("GetLine")
	if e.fileInfo.curFileScanner == nil {
		//fmt.Println("Getline openNextFile")
		err := openNextFile(e)
		if err != nil {
			return "", err
		}
	}
	for !(e.fileInfo.curFileScanner).Scan() {
		//fmt.Printf("Getline scan finished -> openNextFile: Scanner.Err=%v\n", e.fileInfo.curFileScanner.Err())
		if err := e.fileInfo.curFileScanner.Err(); err != nil {
			return "", err
		}
		if len(e.fileInfo.files) == 0 {
			// when read from Stdin
			return "", io.EOF
		}
		err := openNextFile(e)
		if err != nil {
			return "", err
		}
	}
	e.IncFNR()
	e.IncNR()
	return e.fileInfo.curFileScanner.Text(), nil
}

// GetLine read one line from specified file. If EOF, return io.EOF.
func (e *Env) GetLineFrom(file string) (string, error) {
	openFile := func(file string) error {
		var fileReader *os.File
		var err error
		if file == "" {
			return fmt.Errorf("cannot open null string file")
		}
		fileReader, err = os.Open(file)
		if err != nil {
			return fmt.Errorf("open %v error: %v", file, err)
		}
		readCloser := io.ReadCloser(fileReader)
		e.fileInfo.readCloser[file] = &readCloser
		e.fileInfo.scanner[file] = bufio.NewScanner(io.Reader(readCloser))
		err = e.setScannerSplit(file)
		if err != nil {
			return err
		}
		return nil
	}

	//fmt.Printf("filename:%v\n", file)
	if e.fileInfo.scanner[file] == nil {
		err := openFile(file)
		if err != nil {
			return "", err
		}
	}
	if !e.fileInfo.scanner[file].Scan() {
		return "", io.EOF
	}
	return e.fileInfo.scanner[file].Text(), nil
}

func (e *Env) setScannerSplit(key string) error {
	rs, err := e.Get("RS") // Record Separater
	if err == ErrUnknownSymbol {
		return nil
	} else if err != nil {
		return err
	}
	if rs == "" {
		return nil
	}

	var splitHelper func(int, []byte, []byte, []byte) (int, []byte, error)
	splitHelper = func(advance int, token []byte, data []byte, pat []byte) (int, []byte, error) {
		if len(pat) == 0 {
			return advance, token, nil
		}
		if len(data) == 0 {
			return advance, token, bufio.ErrFinalToken
		}
		/*
			if len(pat) == 0 || len(data) == 0 {
				return advance, token, nil
			}
		*/
		if data[0] == pat[0] {
			return splitHelper(advance+1, append(token, data[0]), data[1:], pat[1:])
		}
		return splitHelper(advance+1, append(token, data[0]), data[1:], pat)
	}
	split := func(data []byte, atEOF bool) (int, []byte, error) {
		i, bs, err := splitHelper(0, []byte{}, data, []byte(rs.(string)))
		if err != nil {
			return i, bs, err
		} else if len(data) == len(bs) {
			return i, bs[:len(bs)-len(rs.(string))], bufio.ErrFinalToken
		} else {
			//fmt.Printf("data=%s\tbs=[%s]\n", data, bs[:len(bs)-len(rs.(string))])
			return i, bs[:len(bs)-len(rs.(string))], nil
		}
	}
	if key == "" {
		// set split func to current scanner
		if e.fileInfo.curFileScanner != nil {
			scanner := e.fileInfo.curFileScanner
			if len(scanner.Text()) == 0 {
				scanner.Split(split)
			}
		}
	} else {
		// set split func to speified scanner
		scanner, ok := e.fileInfo.scanner[key]
		if !ok {
			return fmt.Errorf("file key %v not found", key)
		}
		scanner.Split(split)
	}
	return nil
}

// CloseFile close a file.
func (e *Env) CloseFile(k string) error {
	f, ok := e.fileInfo.readCloser[k]
	if !ok {
		return ErrUnknownSymbol
	}
	stdin := io.ReadCloser(os.Stdin)
	if f != &stdin {
		if e := (*f).Close(); e != nil {
			return e
		}
	}
	delete(e.fileInfo.readCloser, k)
	delete(e.fileInfo.scanner, k)
	return nil
}

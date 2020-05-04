package vm

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// SetFile defines a new file.
func (e *Env) SetFile(k string, f *io.ReadCloser) (*bufio.Scanner, error) {
	if k == "" {
		k = "-" // stdin
	}
	_, ok := e.readCloser[k]
	if ok {
		//fmt.Printf("SetFile(%v)\n", k)
		return nil, ErrAlreadyKnownSymbol
	}
	scanner := bufio.NewScanner(io.Reader(*f))
	e.readCloser[k] = f
	e.scanner[k] = scanner
	err := e.setScannerSplit(k)
	if err != nil {
		return nil, err
	}
	return scanner, nil
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
		// set split func to all the scanners
		for _, scanner := range e.scanner {
			// scanner.Split() panics when used after Scan()
			// No interface to check Scan() is called .
			if len(scanner.Text()) == 0 {
				scanner.Split(split)
			}
		}
	} else {
		// set split func to speified  scanner
		scanner, ok := e.scanner[key]
		if !ok {
			return fmt.Errorf("file key %v not found", key)
		}
		scanner.Split(split)
	}
	return nil
}

// GetScanner returns the scanner with a specified name.
func (e *Env) GetScanner(k string) (*bufio.Scanner, error) {
	s, ok := e.scanner[k]
	if !ok {
		return nil, ErrUnknownSymbol
	}
	return s, nil
}

// CloseFile close a file.
func (e *Env) CloseFile(k string) error {
	if k == "" {
		k = "-"
	}

	f, ok := e.readCloser[k]
	if !ok {
		return ErrUnknownSymbol
	}
	stdin := io.ReadCloser(os.Stdin)
	if f != &stdin {
		if e := (*f).Close(); e != nil {
			return e
		}
	}
	delete(e.readCloser, k)
	delete(e.scanner, k)
	return nil
}

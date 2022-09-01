package chscht

import (
	"bufio"
	"errors"
	"os"
	"sync"
)

type Dict struct {
	l           sync.RWMutex
	data        map[rune]rune
	dataReverse map[rune]rune
	dictPath    string
}

func New(dictPath ...string) (*Dict, error) {
	d := &Dict{
		data:        make(map[rune]rune),
		dataReverse: make(map[rune]rune),
	}
	if len(dictPath) >= 1 {
		d.dictPath = "./dict.txt"
	}
	var (
		err  error
		file *os.File
	)
	var simplified []rune
	var traditional []rune
	if d.dictPath != "" {
		file, err = os.Open(d.dictPath)
		if err != nil {
			return nil, err
		}
		buf := bufio.NewScanner(file)
		var i int
		for buf.Scan() {
			text := buf.Text()
			switch i {
			case 0:
				simplified = []rune(text)
			case 1:
				traditional = []rune(text)
			}
			i++
		}
	} else {
		simplified = defaultSimplified
		traditional = defaultTraditional
	}

	if len(simplified) != len(traditional) {
		return nil, errors.New("the length of simplified varies from the length of traditional")
	}
	for i := 0; i < len(simplified) && i < len(traditional); i++ {
		d.set(traditional[i], simplified[i])
	}
	return d, nil
}

func (d *Dict) set(on, under rune) {
	d.l.Lock()
	defer d.l.Unlock()
	d.data[on] = under
	d.dataReverse[under] = on
}

func (d *Dict) getData(char rune) rune {
	d.l.RLock()
	defer d.l.RUnlock()
	if s, ok := d.data[char]; ok {
		return s
	}
	return char
}

func (d *Dict) getDataR(char rune) rune {
	d.l.RLock()
	defer d.l.RUnlock()
	if s, ok := d.dataReverse[char]; ok {
		return s
	}
	return char
}

func read(s string, f func(char rune) rune) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		r[i] = f(r[i]) //d.getData(r[i])
	}
	return string(r)
}

func (d *Dict) Simplified(s string) string {
	return read(s, d.getData)
}

func (d *Dict) Traditional(s string) string {
	return read(s, d.getDataR)
}

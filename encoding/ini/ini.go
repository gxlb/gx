//    CopyRight @Ally Dale 2018
//    Author  : Ally Dale(vipally@gmail.com)

//ini file reader
package ini

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type content struct {
	content string
	comment string
}

type sectionData struct {
	name    string
	comment string
	data    map[string]content
}

type iniData struct {
	comment string
	data    map[string]sectionData
}

type section map[string]string

type IniFile struct {
	sections map[string]section
}

func New(path string) (*IniFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		f.Close()
	}()
	r := Load(f)
	return r, nil
}

func Load(f io.Reader) *IniFile {
	m := make(map[string]section)
	r := bufio.NewReader(f)
	sec := ""
	var line string
	var err error
	for err == nil {
		line, err = r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" || line[0] == ';' {
			continue
		}
		if line[0] == '[' && line[len(line)-1] == ']' {
			sec = line[1 : len(line)-1]
			_, ok := m[sec]
			if !ok {
				m[sec] = make(section)
			}
			continue
		}
		if sec == "" {
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		val := ""
		if len(pair) < 1 || len(pair) > 2 {
			continue
		}
		if len(pair) == 2 {
			val = pair[1]
		}
		key := strings.TrimSpace(pair[0])
		val = strings.TrimSpace(val)
		if key == "" {
			continue
		}
		m[sec][key] = val
	}
	return &IniFile{m}
}

func (p *IniFile) Sections() []string {
	s := make([]string, len(p.sections))
	i := 0
	for k, _ := range p.sections {
		s[i] = k
		i++
	}
	return s
}

func (p *IniFile) Keys(sec string) []string {
	m, ok := p.sections[sec]
	if !ok {
		return nil
	}
	keys := make([]string, len(m))
	i := 0
	for key, _ := range m {
		keys[i] = key
		i++
	}
	return keys
}

func (p *IniFile) GetString(sec, key, def string) string {
	m, ok := p.sections[sec]
	if !ok {
		return def
	}
	v, ok := m[key]
	if !ok {
		return def
	}
	return v
}

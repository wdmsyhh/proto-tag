package tag

import (
	"bufio"
	"bytes"
	"container/list"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init() {
	generator.RegisterPlugin(new(Tag))
}

type Tag struct {
	gen         *generator.Generator
	tags        map[string]string
	fieldMaxLen int
	tagMaxLen   int
}

// Name returns the name of this plugin, "settag"
func (r *Tag) Name() string {
	return "tag"
}

// Init initializes the plugin.
func (r *Tag) Init(gen *generator.Generator) {
	r.gen = gen
}

func (r *Tag) P(args ...interface{}) { r.gen.P(args...) }

// Generate generates code for the services in the given file.
func (r *Tag) Generate(file *generator.FileDescriptor) {
	r.GetStructTags(*file.Name)
	r.Tag("")
}

// GenerateImports generates the import declaration for this file.
func (r *Tag) GenerateImports(file *generator.FileDescriptor) {}

func (r *Tag) GetStructTags(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	r.tags = make(map[string]string)
	var comment bool
	msgNameStack := NewStack()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		//skip empty line in message
		if len(line) <= 0 {
			continue
		}

		strLine := string(line)

		if strings.HasPrefix(strings.TrimSpace(strLine), "/*") {
			comment = true
		}

		if comment && strings.Contains(strLine, "*/") {
			comment = false
			continue
		}

		if comment {
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(strLine), "message") {
			pop := msgNameStack.GetPOP()
			if pop != "" {
				msgNameStack.PUSH(pop + "_" + strings.Fields(string(line))[1])
			} else {
				msgNameStack.PUSH(strings.Fields(strLine)[1])
			}
			continue
		}

		pop := msgNameStack.GetPOP()
		if pop != "" && strings.TrimSpace(strLine)[0] == '}' {
			msgNameStack.POP()
			continue
		}

		pop = msgNameStack.GetPOP()
		if pop != "" {
			if strings.HasPrefix(strings.TrimSpace(strLine), "//") {
				continue
			}

			k, v := getFieldTag(strLine, msgNameStack.GetPOP())

			if k == "" && v == "" {
				continue
			}

			r.tags[strings.ToLower(k)] = v

			if len(strings.Split(k, ".")[1]) > r.fieldMaxLen {
				r.fieldMaxLen = len(strings.Split(k, ".")[1])
			}

			tags := strings.Fields(v)
			for _, tag := range tags {
				if len(strings.Split(tag, ":")[1])-2 > r.tagMaxLen {
					r.tagMaxLen = len(strings.Split(tag, ":")[1]) - 2
				}
			}
		}
	}
}

func getFieldTag(line string, msgName string) (field string, tag string) {
	fts := strings.Split(line, "//")
	if len(fts) <= 1 {
		return "", ""
	}

	tag = fts[1]
	fs := strings.Fields(fts[0])
	fsl := len(fs)
	field = msgName + "."
	for i := 0; i < fsl; i++ {
		if i == fsl-1 {
			field += fs[i]
			break
		} else {
			if fs[i+1] == "=" {
				field += fs[i]
				break
			}
		}
	}

	tag = strings.TrimSpace(tag)
	tag = strings.Trim(tag, "`")
	tag = trimInside(tag)

	return
}

func trimInside(s string) string {
	for {
		if strings.Contains(s, "  ") {
			s = strings.Replace(s, "  ", " ", -1)
		} else {
			break
		}
	}

	return s
}

func (r *Tag) Tag(pbgoFileName string) {
	if len(r.tags) <= 0 {
		return
	}

	reader := &bufio.Reader{}
	// 仅在测试的时候传
	if pbgoFileName != "" {
		file, err := os.Open(pbgoFileName)
		if err != nil {
			return
		}
		defer file.Close()
		reader = bufio.NewReader(file)
		r.gen = &generator.Generator{
			Buffer: &bytes.Buffer{},
		}
	} else {
		readbuf := bytes.NewBuffer([]byte{})
		readbuf.Write(r.gen.Buffer.Bytes())
		reader = bufio.NewReader(readbuf)
	}

	buf := bytes.NewBuffer([]byte{})
	var comment bool
	msgNameStack := NewStack()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			buf.WriteString("\n")
			break
		}

		strLine := string(line)

		if strings.HasPrefix(strings.TrimSpace(strLine), "/*") {
			comment = true
		}

		if comment && strings.Contains(strLine, "*/") {
			comment = false
			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if comment {
			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if r.needtag(strings.TrimSpace(strLine)) {
			if msgNameStack.GetPOP() != "" {
				msgNameStack.PUSH(msgNameStack.GetPOP() + "_" + strings.Fields(strLine)[1])
			} else {
				msgNameStack.PUSH(strings.Fields(strLine)[1])
			}

			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if msgNameStack.GetPOP() != "" && strings.TrimSpace(strLine)[0] == '}' {
			msgNameStack.POP()
			buf.Write(line)
			buf.WriteString("\n")
			continue
		}

		if msgNameStack.GetPOP() != "" {
			if strings.HasPrefix(strings.TrimSpace(strLine), "//") {
				buf.Write(line)
				buf.WriteString("\n")
				continue
			}

			fields := strings.Fields(strings.TrimSpace(strLine))
			key := msgNameStack.GetPOP() + "." + fields[0]
			tag := r.tags[strings.ToLower(key)]
			newline := resetTag(strLine, fields[0], tag, r.fieldMaxLen, r.tagMaxLen)
			buf.WriteString(newline)
			buf.WriteString("\n")
			continue
		}
		buf.Write(line)
		buf.WriteString("\n")
	}

	r.gen.Buffer.Reset()
	data := buf.Bytes()
	r.gen.Buffer.Write(data)

	if pbgoFileName != "" {
		dir, _ := os.Getwd()
		err := ioutil.WriteFile(dir+"/test/example/example.pb_resetTag.go", r.gen.Buffer.Bytes(), 0644)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (r *Tag) needtag(line string) bool {
	for k := range r.tags {
		ks := strings.Split(k, ".")
		sub := "type " + ks[0] + " struct"
		if strings.HasPrefix(strings.ToLower(line), sub) {
			return true
		}
	}

	return false
}
func resetTag(line string, field string, tag string, maxlenField, maxlenTag int) string {
	if tag == "" {
		return line
	}
	//reset default json
	res := strings.Trim(strings.TrimRight(strings.TrimRight(line, "\n"), " "), "`")
	if strings.Contains(line, "json:") && strings.Contains(tag, "json:") {
		substr := " json:\"" + field + ",omitempty\""
		res = strings.Replace(res, substr, "", -1)
	}

	fs := strings.Fields(res)
	for i := 2; i < len(fs); i++ {
		if i == 2 {
			res = strings.Replace(res, fs[i], "`", -1)
			fs[i] = fs[i][1:]
		} else {
			res = strings.Replace(res, fs[i], "", -1)
		}
	}

	fs = append(fs, strings.Fields(tag)...)

	res = strings.TrimRight(res, " ")

	for i := 2; i < len(fs); i++ {
		if i == 2 {
			res += fs[i]
		} else {
			res += " " + fs[i]
		}
	}

	res += "`"

	return res
}

type stack struct {
	s *list.List
}

func NewStack() *stack {
	return &stack{
		s: list.New(),
	}
}
func (s *stack) PUSH(value string) {
	s.s.PushBack(value)
}

func (s *stack) POP() string {
	res := s.s.Back()
	if res == nil {
		return ""
	}
	s.s.Remove(res)
	return res.Value.(string)
}

func (s *stack) GetPOP() string {
	res := s.s.Back()
	if res == nil {
		return ""
	}

	return res.Value.(string)
}

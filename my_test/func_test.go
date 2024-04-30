package my

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"
)

type Pipeline struct {
	line     []interface{}
	index    int8
	maxIndex int8
	val      map[string]interface{}
	isAbort  bool
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		val: make(map[string]interface{}), // 初始化 val 字段为一个空的 map
	}
}
func (p *Pipeline) Set(key string, val interface{}) {
	p.val[key] = val
}

func (p *Pipeline) Get(key string) (bool, interface{}) {
	v, ok := p.val[key]
	return ok, v
}

type Exec interface{}

func (p *Pipeline) Use(fn Exec) *Pipeline {
	p.line = append(p.line, fn)
	p.maxIndex = int8(len(p.line))
	return p
}

func (p *Pipeline) Abort() {
	p.isAbort = true
}

func (p *Pipeline) Exec(arg interface{}) {
	//var arg1,arg2,arg3,arg4,arg5,arg6,arg7,arg8,arg9,arg10 interface{}
	var arg1, arg2, arg3 interface{}
	arg1 = arg
	for _, v := range p.line {
		if p.isAbort {
			break
		}
		switch f := v.(type) {
		case func(interface{}):
			f(arg1)
		case func(interface{}) interface{}:
			arg1 = f(arg1)
		case func(interface{}) (interface{}, interface{}):
			arg1, arg2 = f(arg1)
		case func(interface{}, interface{}):
			f(arg1, arg2)
		case func(interface{}, interface{}) (interface{}, interface{}):
			arg1, arg2 = f(arg1, arg2)
		case func(interface{}, interface{}) (interface{}, interface{}, interface{}):
			arg1, arg2, arg3 = f(arg1, arg2)
		case func(interface{}, interface{}, interface{}) (interface{}, interface{}, interface{}):
			f(arg1, arg2, arg3)
		default:
			log.Println("error")
		}
	}
}

func TestPipeLine(t *testing.T) {
	pip := NewPipeline()
	pip.Set("Assss", "ssss")
	pip.Use(func(arg interface{}) interface{} {
		str, _ := arg.(string)
		log.Println(str)
		pip.Set("fsd", "dddd")
		return "this is first"
	}).Use(func(arg interface{}) (interface{}, interface{}) {
		str, _ := arg.(string)
		log.Println(str)

		pip.Abort()
		return "this is second", "The second"
	}).Use(func(arg1 interface{}, arg2 interface{}) (interface{}, interface{}, interface{}) {
		str1, _ := arg1.(string)
		str2, _ := arg2.(string)
		log.Println(str1)
		log.Println(str2)
		return "this is three", "ssss", "ddd"
	})

	pip.Exec("begin")
}

func TestT1(t *testing.T) {
	formulaTemplate := `{{.A}} * {{.B}} + {{.C}}`

	// 准备模板数据
	data := struct {
		A float64
		B float64
		C float64
	}{
		A: 10,
		B: 5,
		C: 3,
	}

	// 解析模板字符串
	tmpl, err := template.New("formula").Parse(formulaTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// 应用模板并计算数学公式
	var result float64
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println(" = ", result)
}

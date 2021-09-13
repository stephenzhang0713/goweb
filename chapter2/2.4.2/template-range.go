package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	rangeTemplate := `
{{if .Kind}}
{{range $i, $v := .MapContent}}
{{$i}} => {{$v}} , {{$.OutsideContent}}
{{end}}
{{else}}
{{range .MapContent}}
{{.}} , {{$.OutsideContent}}
{{end}}    
{{end}}`

	str1 := []string{"第一次range", "用index和value"}
	str2 := []string{"第二次range", "没有用index和value"}

	type Content struct {
		MapContent     []string
		OutsideContent string
		Kind           bool
	}
	var contents = []Content{
		{str1, "第一次外面的内容", true},
		{str2, "第二次外面的内容", false},
	}

	t := template.Must(template.New("range").Parse(rangeTemplate))

	for _, c := range contents {
		err := t.Execute(os.Stdout, c)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

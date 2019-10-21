package console

import (
	"fmt"
	"io"
	"text/template"
)

type Writer struct {
	Out io.Writer
}

func (w Writer) DisplayText(text string, data ...map[string]interface{}) {
	var keys interface{}
	if len(data) > 0 {
		keys = data[0]
	}

	formattedTemplate := template.Must(template.New("Display Text").Parse(text + "\n"))
	formattedTemplate.Execute(w.Out, keys)
}

func (w Writer) DisplayTable(headers []string, data [][]string) {
	for _, header := range headers {
		fmt.Fprintf(w.Out, "%s\t", header)
	}

	fmt.Fprintln(w.Out)
	fmt.Fprintln(w.Out, "---")

	for _, row := range data {
		for _, item := range row {
			fmt.Fprintf(w.Out, "%s\t", item)
		}
		fmt.Fprintln(w.Out)
	}

}

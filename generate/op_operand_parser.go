package main

import (
	"encoding/binary"
	"encoding/hex"
	"git.agehadev.com/elliebelly/jamboy/internal/code"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const opTemplate = `package code

var (
	Ops = map[int]Op{
	{{- range $op := .Ops }}
		0x{{ $op.CodeHex }}: {
			Type: Op{{ $op.Name }},
			Code: 0x{{ $op.CodeHex }},
		{{- if $op.Operands }}
			Operands: []Operand{
			{{- range $operand := $op.Operands }}
				{
					Location: {{ $operand.Location.String }},
					Type: {{ $operand.Type.String }},
				},
			{{- end }}
			},
		{{- end }}
		},
	{{- end }}
	}
)
`

type OpTemplateList struct {
	Ops []OpTemplateData
}

type OpTemplateData struct {
	Name     string
	CodeHex  string
	Operands []code.Operand
}

func main() {
	opdata, err := ioutil.ReadFile("notes/fullinstructions")

	if err != nil {
		panic(err)
	}

	extractValLocationRegexp := regexp.MustCompile("\\(?(.*?)\\)?")
	getValSizeRegexp := regexp.MustCompile("([a-z])([0-9]+)")

	lines := strings.Split(string(opdata), "\n")

	codeI := uint16(0)

	opTemplateList := OpTemplateList{
		Ops: make([]OpTemplateData, 0),
	}

	for _, line := range lines {
		splitOp := strings.Split(line, " ")

		name := splitOp[0]

		codeByteBuffer := make([]byte, 2)

		binary.BigEndian.PutUint16(codeByteBuffer, codeI)

		opTemplateData := OpTemplateData{
			Name:    name,
			CodeHex: hex.EncodeToString(codeByteBuffer),
		}

		if len(splitOp) > 1 && name != "PREFIX" {
			splitOperands := strings.Split(splitOp[1], ",")

			opTemplateData.Operands = make([]code.Operand, 0)

			for _, operandString := range splitOperands {
				operand := code.Operand{}

				if string(([]rune(operandString))[0]) == "(" {
					operand.Type = code.RetrievePointer
				} else {
					operand.Type = code.RetrieveVal
				}

				valLocationString := extractValLocationRegexp.FindStringSubmatch(operandString)[1]

				if match, _ := regexp.MatchString("[a-z][0-9]+", valLocationString); match {
					submatch := getValSizeRegexp.FindStringSubmatch(valLocationString)

					valType := submatch[1]
					size, err := strconv.Atoi(submatch[2])

					if err != nil {
						panic(err)
					}

					switch valType {
					case "a":
						operand.Location = code.ValAddress
					case "r":
						operand.Location = code.ValRegister
					default:
						if size == 8 {
							operand.Location = code.Val8
						} else {
							operand.Location = code.Val16
						}
					}
				}

				opTemplateData.Operands = append(opTemplateData.Operands, operand)
			}
		}
		opTemplateList.Ops = append(opTemplateList.Ops, opTemplateData)

		codeI++
	}

	tmpl := template.New("ops")

	outputFile, err := os.OpenFile("internal/code/ops_generated.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	_, err = tmpl.Parse(opTemplate)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(outputFile, &opTemplateList)

	if err != nil {
		panic(err)
	}
}

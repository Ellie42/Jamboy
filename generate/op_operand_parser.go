package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
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
					Value: {{ $operand.ValueString }},
					RetrieveType: {{ $operand.RetrieveType.String }},
					ValueType: {{ $operand.ValueType.String }},
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

var registerValMap = map[string]code.ValueLocation{
	"A":  code.ValA,
	"B":  code.ValB,
	"C":  code.ValC,
	"D":  code.ValD,
	"E":  code.ValE,
	"F":  code.ValF,
	"H":  code.ValH,
	"L":  code.ValL,
	"AF": code.ValAF,
	"BC": code.ValBC,
	"DE": code.ValDE,
	"HL": code.ValHL,
	"SP": code.ValSP,
	"PC": code.ValPC,
}

var keywordValMap = map[string]code.ValueKeyword{
	"Z":  code.ValKeywordZ,
	"NZ": code.ValKeywordNZ,
	"CB": code.ValKeywordCB,
}

func main() {
	opdata, err := ioutil.ReadFile("generate/fullinstructions")

	if err != nil {
		panic(err)
	}

	extractValLocationRegexp := regexp.MustCompile("\\(?(.*?)\\)?$")
	getValSizeRegexp := regexp.MustCompile("([a-z])([0-9]+)")
	getHexValStringRegexp := regexp.MustCompile("([0-9]+)H")

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

		if len(splitOp) > 1 {
			splitOperands := strings.Split(splitOp[1], ",")

			opTemplateData.Operands = make([]code.Operand, 0)

			for _, operandString := range splitOperands {
				operand := code.Operand{}

				if string(([]rune(operandString))[0]) == "(" {
					operand.RetrieveType = code.RetrievePointer
				} else {
					operand.RetrieveType = code.RetrieveVal
				}

				valString := extractValLocationRegexp.FindStringSubmatch(operandString)[1]

				if match, _ := regexp.MatchString("[a-z][0-9]+", valString); match {

					// We are expecting to read this value from the next byte(s)
					operand.ValueType = code.ValTypeRead

					submatch := getValSizeRegexp.FindStringSubmatch(valString)

					valType := submatch[1]
					size, err := strconv.Atoi(submatch[2])

					if err != nil {
						panic(err)
					}

					switch valType {
					case "a":
						operand.Value = int(code.ValAddress)
					case "r":
						operand.Value = int(code.ValRegister)
					default:
						if size == 8 {
							operand.Value = int(code.Val8)
						} else {
							operand.Value = int(code.Val16)
						}
					}
				} else if register, ok := registerValMap[valString]; ok {
					operand.Value = int(register)
					operand.ValueType = code.ValTypeRegister
				} else if match, _ := regexp.MatchString("[0-9]+", valString); match {
					var constVal int
					var err error

					if matches := getHexValStringRegexp.FindStringSubmatch(valString); matches != nil {
						hexBytes, err := hex.DecodeString(matches[1])

						if err != nil {
							panic(err)
						}

						if len(hexBytes) == 1 {
							constVal = int(hexBytes[0])
						} else {
							constVal = int(binary.LittleEndian.Uint16(hexBytes))
						}
					} else {
						constVal, err = strconv.Atoi(valString)

						if err != nil {
							panic(err)
						}
					}

					operand.Value = constVal
					operand.ValueType = code.ValTypeConst
				} else if keyword, ok := keywordValMap[valString]; ok {
					operand.Value = int(keyword)
					operand.ValueType = code.ValTypeKeyword
				}

				operand.ValueString = GetValueTypeString(operand.Value, operand.ValueType)

				if operand.ValueType != code.ValTypeConst {
					operand.ValueString = fmt.Sprintf("int(%s)", operand.ValueString)
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

func GetValueTypeString(value int, valueType code.ValueType) string {
	switch valueType {
	case code.ValTypeRead:
		fallthrough
	case code.ValTypeRegister:
		return code.ValueLocation(value).String()
	case code.ValTypeConst:
		return fmt.Sprintf("%d", value)
	case code.ValTypeKeyword:
		return code.ValueKeyword(value).String()
	default:
		return "UNKNOWN"
	}
}

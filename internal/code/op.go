package code

type Op struct {
	Type     OpType
	Code     int
	Operands []Operand
}

type Operand struct {
	Value        int
	RetrieveType RetrieveType
	ValueType    ValueType
	ValueString  string
}

//go:generate stringer -type ValueLocation
type ValueLocation int

const (
	ValA ValueLocation = iota
	ValB
	ValC
	ValD
	ValE
	ValF
	ValH
	ValL
	ValAF
	ValBC
	ValDE
	ValHL
	ValSP
	ValPC
	Val8
	Val16
	ValAddress
	ValRegister
)

//go:generate stringer -type ValueKeyword
type ValueKeyword int

const (
	ValKeywordZ ValueKeyword = iota
	ValKeywordNZ
	ValKeywordCB
)

//go:generate stringer -type ValueType
type ValueType int

const (
	ValTypeRegister ValueType = iota
	ValTypeRead
	ValTypeConst
	ValTypeKeyword
)

//go:generate stringer -type RetrieveType
type RetrieveType int

const (
	RetrieveVal RetrieveType = iota
	RetrievePointer
)

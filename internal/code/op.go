package code

type Op struct {
	Type     OpType
	Code     int
	Operands []Operand
}

type Operand struct {
	Location ValueLocation
	Type     RetrieveType
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

//go:generate stringer -type RetrieveType
type RetrieveType int

const (
	RetrieveVal RetrieveType = iota
	RetrievePointer
)

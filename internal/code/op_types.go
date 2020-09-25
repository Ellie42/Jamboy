package code

//go:generate stringer -type=OpType
type OpType uint16

const (
	OpADC OpType = iota
	OpADD
	OpAND
	OpBIT
	OpCALL
	OpCCF
	OpCP
	OpCPL
	OpDAA
	OpDEC
	OpDI
	OpEI
	OpHALT
	OpINC
	OpJP
	OpJR
	OpLD
	OpLDH
	OpNOP
	OpOR
	OpPOP
	OpPREFIX
	OpPUSH
	OpPrefix
	OpRES
	OpRET
	OpRETI
	OpRL
	OpRLA
	OpRLC
	OpRLCA
	OpRR
	OpRRA
	OpRRC
	OpRRCA
	OpRST
	OpSBC
	OpSCF
	OpSET
	OpSLA
	OpSRA
	OpSRL
	OpSTOP
	OpSUB
	OpSWAP
	OpXOR
)

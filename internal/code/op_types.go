package code

//go:generate stringer -type=OpType -linecomment
type OpType uint16

const (
	OpADC    OpType = iota //ADC
	OpADD                  //ADD
	OpAND                  //AND
	OpBIT                  //BIT
	OpCALL                 //CALL
	OpCCF                  //CCF
	OpCP                   //CP
	OpCPL                  //CPL
	OpDAA                  //DAA
	OpDEC                  //DEC
	OpDI                   //DI
	OpEI                   //EI
	OpHALT                 //HALT
	OpINC                  //INC
	OpJP                   //JP
	OpJR                   //JR
	OpLD                   //LD
	OpLDH                  //LDH
	OpNOP                  //NOP
	OpOR                   //OR
	OpPOP                  //POP
	OpPREFIX               //PREFIX
	OpPUSH                 //PUSH
	OpPrefix               //Prefix
	OpRES                  //RES
	OpRET                  //RET
	OpRETI                 //RETI
	OpRL                   //RL
	OpRLA                  //RLA
	OpRLC                  //RLC
	OpRLCA                 //RLCA
	OpRR                   //RR
	OpRRA                  //RRA
	OpRRC                  //RRC
	OpRRCA                 //RRCA
	OpRST                  //RST
	OpSBC                  //SBC
	OpSCF                  //SCF
	OpSET                  //SET
	OpSLA                  //SLA
	OpSRA                  //SRA
	OpSRL                  //SRL
	OpSTOP                 //STOP
	OpSUB                  //SUB
	OpSWAP                 //SWAP
	OpXOR                  //XOR
)

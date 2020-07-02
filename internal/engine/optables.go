package engine

type OpCode uint8

var BaseOpJumpTable = []func(jb *Jamboy, opcode OpCode) (err error){
	NOP,             // 0x00
	LDd16,           // 0x01
	LDAToRAM,        // 0x02
	INCB,            // 0x03
	INCA,            // 0x04
	DECA,            // 0x05
	LDd8A,           // 0x06
	RLCA,            // 0x07
	LD,              // 0x08
	ADD,             // 0x09
	LDRAMToA,        // 0x0A
	DEC,             // 0x0B
	INCB,            // 0x0C
	DECB,            // 0x0D
	LDd8B,           // 0x0E
	RRCA,            // 0x0F
	STOP,            // 0x10
	LDd16,           // 0x11
	LDAToRAM,        // 0x12
	INCB,            // 0x13
	INCA,            // 0x14
	DECA,            // 0x15
	LDd8A,           // 0x16
	RLA,             // 0x17
	JR,              // 0x18
	ADD,             // 0x19
	LDRAMToA,        // 0x1A
	DEC,             // 0x1B
	INCB,            // 0x1C
	DECB,            // 0x1D
	LDd8B,           // 0x1E
	RRA,             // 0x1F
	JR,              // 0x20
	LDd16,           // 0x21
	LDAToRAM,        // 0x22
	INCB,            // 0x23
	INCA,            // 0x24
	DECA,            // 0x25
	LDd8A,           // 0x26
	DAA,             // 0x27
	JR,              // 0x28
	ADD,             // 0x29
	LDRAMToA,        // 0x2A
	DEC,             // 0x2B
	INCB,            // 0x2C
	DECB,            // 0x2D
	LDd8B,           // 0x2E
	CPL,             // 0x2F
	JR,              // 0x30
	LDd16,           // 0x31
	LDAToRAM,        // 0x32
	INCB,            // 0x33
	INCA,            // 0x34
	DECA,            // 0x35
	LDd8A,           // 0x36
	SCF,             // 0x37
	JR,              // 0x38
	ADD,             // 0x39
	LDRAMToA,        // 0x3A
	DEC,             // 0x3B
	INCB,            // 0x3C
	DECB,            // 0x3D
	LDd8B,           // 0x3E
	CCF,             // 0x3F
	LD,              // 0x40
	LD,              // 0x41
	LD,              // 0x42
	LD,              // 0x43
	LD,              // 0x44
	LD,              // 0x45
	LD,              // 0x46
	LD,              // 0x47
	LD,              // 0x48
	LD,              // 0x49
	LD,              // 0x4A
	LD,              // 0x4B
	LD,              // 0x4C
	LD,              // 0x4D
	LD,              // 0x4E
	LD,              // 0x4F
	LD,              // 0x50
	LD,              // 0x51
	LD,              // 0x52
	LD,              // 0x53
	LD,              // 0x54
	LD,              // 0x55
	LD,              // 0x56
	LD,              // 0x57
	LD,              // 0x58
	LD,              // 0x59
	LD,              // 0x5A
	LD,              // 0x5B
	LD,              // 0x5C
	LD,              // 0x5D
	LD,              // 0x5E
	LD,              // 0x5F
	LD,              // 0x60
	LD,              // 0x61
	LD,              // 0x62
	LD,              // 0x63
	LD,              // 0x64
	LD,              // 0x65
	LD,              // 0x66
	LD,              // 0x67
	LD,              // 0x68
	LD,              // 0x69
	LD,              // 0x6A
	LD,              // 0x6B
	LD,              // 0x6C
	LD,              // 0x6D
	LD,              // 0x6E
	LD,              // 0x6F
	LD,              // 0x70
	LD,              // 0x71
	LD,              // 0x72
	LD,              // 0x73
	LD,              // 0x74
	LD,              // 0x75
	HALT,            // 0x76
	LD,              // 0x77
	LD,              // 0x78
	LD,              // 0x79
	LD,              // 0x7A
	LD,              // 0x7B
	LD,              // 0x7C
	LD,              // 0x7D
	LD,              // 0x7E
	LD,              // 0x7F
	ADD,             // 0x80
	ADD,             // 0x81
	ADD,             // 0x82
	ADD,             // 0x83
	ADD,             // 0x84
	ADD,             // 0x85
	ADD,             // 0x86
	ADD,             // 0x87
	ADC,             // 0x88
	ADC,             // 0x89
	ADC,             // 0x8A
	ADC,             // 0x8B
	ADC,             // 0x8C
	ADC,             // 0x8D
	ADC,             // 0x8E
	ADC,             // 0x8F
	SUB,             // 0x90
	SUB,             // 0x91
	SUB,             // 0x92
	SUB,             // 0x93
	SUB,             // 0x94
	SUB,             // 0x95
	SUB,             // 0x96
	SUB,             // 0x97
	SBC,             // 0x98
	SBC,             // 0x99
	SBC,             // 0x9A
	SBC,             // 0x9B
	SBC,             // 0x9C
	SBC,             // 0x9D
	SBC,             // 0x9E
	SBC,             // 0x9F
	AND,             // 0xA0
	AND,             // 0xA1
	AND,             // 0xA2
	AND,             // 0xA3
	AND,             // 0xA4
	AND,             // 0xA5
	AND,             // 0xA6
	AND,             // 0xA7
	XOR,             // 0xA8
	XOR,             // 0xA9
	XOR,             // 0xAA
	XOR,             // 0xAB
	XOR,             // 0xAC
	XOR,             // 0xAD
	XOR,             // 0xAE
	XOR,             // 0xAF
	OR,              // 0xB0
	OR,              // 0xB1
	OR,              // 0xB2
	OR,              // 0xB3
	OR,              // 0xB4
	OR,              // 0xB5
	OR,              // 0xB6
	OR,              // 0xB7
	CP,              // 0xB8
	CP,              // 0xB9
	CP,              // 0xBA
	CP,              // 0xBB
	CP,              // 0xBC
	CP,              // 0xBD
	CP,              // 0xBE
	CP,              // 0xBF
	RET,             // 0xC0
	POP,             // 0xC1
	JP,              // 0xC2
	JP,              // 0xC3
	CALL,            // 0xC4
	PUSH,            // 0xC5
	ADD,             // 0xC6
	RST,             // 0xC7
	RET,             // 0xC8
	RET,             // 0xC9
	JP,              // 0xCA
	PREFIX,          // 0xCB
	CALL,            // 0xCC
	CALL,            // 0xCD
	ADC,             // 0xCE
	RST,             // 0xCF
	RET,             // 0xD0
	POP,             // 0xD1
	JP,              // 0xD2
	NOP,             // 0xD3
	CALL,            // 0xD4
	PUSH,            // 0xD5
	SUB,             // 0xD6
	RST,             // 0xD7
	RET,             // 0xD8
	RETI,            // 0xD9
	JP,              // 0xDA
	NOP,             // 0xDB
	CALL,            // 0xDC
	NOP,             // 0xDD
	SBC,             // 0xDE
	RST,             // 0xDF
	LDH,             // 0xE0
	POP,             // 0xE1
	LD,              // 0xE2
	NOP,             // 0xE3
	NOP,             // 0xE4
	PUSH,            // 0xE5
	AND,             // 0xE6
	RST,             // 0xE7
	ADD,             // 0xE8
	JP,              // 0xE9
	LDAToRAMPointer, // 0xEA
	NOP,             // 0xEB
	NOP,             // 0xEC
	NOP,             // 0xED
	XOR,             // 0xEE
	RST,             // 0xEF
	LDH,             // 0xF0
	POP,             // 0xF1
	LD,              // 0xF2
	DI,              // 0xF3
	NOP,             // 0xF4
	PUSH,            // 0xF5
	OR,              // 0xF6
	RST,             // 0xF7
	LD,              // 0xF8
	LD,              // 0xF9
	LDRAMPointerToA, // 0xFA
	EI,              // 0xFB
	NOP,             // 0xFC
	NOP,             // 0xFD
	CP,              // 0xFE
	RST,             // 0xFF
}

var ExtraOpJumpTable = []func(jb *Jamboy, opcode OpCode) (err error){
	RLC,  // 0x0
	RLC,  // 0x1
	RLC,  // 0x2
	RLC,  // 0x3
	RLC,  // 0x4
	RLC,  // 0x5
	RLC,  // 0x6
	RLC,  // 0x7
	RRC,  // 0x8
	RRC,  // 0x9
	RRC,  // 0xA
	RRC,  // 0xB
	RRC,  // 0xC
	RRC,  // 0xD
	RRC,  // 0xE
	RRC,  // 0xF
	RL,   // 0x10
	RL,   // 0x11
	RL,   // 0x12
	RL,   // 0x13
	RL,   // 0x14
	RL,   // 0x15
	RL,   // 0x16
	RL,   // 0x17
	RR,   // 0x18
	RR,   // 0x19
	RR,   // 0x1A
	RR,   // 0x1B
	RR,   // 0x1C
	RR,   // 0x1D
	RR,   // 0x1E
	RR,   // 0x1F
	SLA,  // 0x20
	SLA,  // 0x21
	SLA,  // 0x22
	SLA,  // 0x23
	SLA,  // 0x24
	SLA,  // 0x25
	SLA,  // 0x26
	SLA,  // 0x27
	SRA,  // 0x28
	SRA,  // 0x29
	SRA,  // 0x2A
	SRA,  // 0x2B
	SRA,  // 0x2C
	SRA,  // 0x2D
	SRA,  // 0x2E
	SRA,  // 0x2F
	SWAP, // 0x30
	SWAP, // 0x31
	SWAP, // 0x32
	SWAP, // 0x33
	SWAP, // 0x34
	SWAP, // 0x35
	SWAP, // 0x36
	SWAP, // 0x37
	SRL,  // 0x38
	SRL,  // 0x39
	SRL,  // 0x3A
	SRL,  // 0x3B
	SRL,  // 0x3C
	SRL,  // 0x3D
	SRL,  // 0x3E
	SRL,  // 0x3F
	BIT,  // 0x40
	BIT,  // 0x41
	BIT,  // 0x42
	BIT,  // 0x43
	BIT,  // 0x44
	BIT,  // 0x45
	BIT,  // 0x46
	BIT,  // 0x47
	BIT,  // 0x48
	BIT,  // 0x49
	BIT,  // 0x4A
	BIT,  // 0x4B
	BIT,  // 0x4C
	BIT,  // 0x4D
	BIT,  // 0x4E
	BIT,  // 0x4F
	BIT,  // 0x50
	BIT,  // 0x51
	BIT,  // 0x52
	BIT,  // 0x53
	BIT,  // 0x54
	BIT,  // 0x55
	BIT,  // 0x56
	BIT,  // 0x57
	BIT,  // 0x58
	BIT,  // 0x59
	BIT,  // 0x5A
	BIT,  // 0x5B
	BIT,  // 0x5C
	BIT,  // 0x5D
	BIT,  // 0x5E
	BIT,  // 0x5F
	BIT,  // 0x60
	BIT,  // 0x61
	BIT,  // 0x62
	BIT,  // 0x63
	BIT,  // 0x64
	BIT,  // 0x65
	BIT,  // 0x66
	BIT,  // 0x67
	BIT,  // 0x68
	BIT,  // 0x69
	BIT,  // 0x6A
	BIT,  // 0x6B
	BIT,  // 0x6C
	BIT,  // 0x6D
	BIT,  // 0x6E
	BIT,  // 0x6F
	BIT,  // 0x70
	BIT,  // 0x71
	BIT,  // 0x72
	BIT,  // 0x73
	BIT,  // 0x74
	BIT,  // 0x75
	BIT,  // 0x76
	BIT,  // 0x77
	BIT,  // 0x78
	BIT,  // 0x79
	BIT,  // 0x7A
	BIT,  // 0x7B
	BIT,  // 0x7C
	BIT,  // 0x7D
	BIT,  // 0x7E
	BIT,  // 0x7F
	RES,  // 0x80
	RES,  // 0x81
	RES,  // 0x82
	RES,  // 0x83
	RES,  // 0x84
	RES,  // 0x85
	RES,  // 0x86
	RES,  // 0x87
	RES,  // 0x88
	RES,  // 0x89
	RES,  // 0x8A
	RES,  // 0x8B
	RES,  // 0x8C
	RES,  // 0x8D
	RES,  // 0x8E
	RES,  // 0x8F
	RES,  // 0x90
	RES,  // 0x91
	RES,  // 0x92
	RES,  // 0x93
	RES,  // 0x94
	RES,  // 0x95
	RES,  // 0x96
	RES,  // 0x97
	RES,  // 0x98
	RES,  // 0x99
	RES,  // 0x9A
	RES,  // 0x9B
	RES,  // 0x9C
	RES,  // 0x9D
	RES,  // 0x9E
	RES,  // 0x9F
	RES,  // 0xA0
	RES,  // 0xA1
	RES,  // 0xA2
	RES,  // 0xA3
	RES,  // 0xA4
	RES,  // 0xA5
	RES,  // 0xA6
	RES,  // 0xA7
	RES,  // 0xA8
	RES,  // 0xA9
	RES,  // 0xAA
	RES,  // 0xAB
	RES,  // 0xAC
	RES,  // 0xAD
	RES,  // 0xAE
	RES,  // 0xAF
	RES,  // 0xB0
	RES,  // 0xB1
	RES,  // 0xB2
	RES,  // 0xB3
	RES,  // 0xB4
	RES,  // 0xB5
	RES,  // 0xB6
	RES,  // 0xB7
	RES,  // 0xB8
	RES,  // 0xB9
	RES,  // 0xBA
	RES,  // 0xBB
	RES,  // 0xBC
	RES,  // 0xBD
	RES,  // 0xBE
	RES,  // 0xBF
	SET,  // 0xC0
	SET,  // 0xC1
	SET,  // 0xC2
	SET,  // 0xC3
	SET,  // 0xC4
	SET,  // 0xC5
	SET,  // 0xC6
	SET,  // 0xC7
	SET,  // 0xC8
	SET,  // 0xC9
	SET,  // 0xCA
	SET,  // 0xCB
	SET,  // 0xCC
	SET,  // 0xCD
	SET,  // 0xCE
	SET,  // 0xCF
	SET,  // 0xD0
	SET,  // 0xD1
	SET,  // 0xD2
	SET,  // 0xD3
	SET,  // 0xD4
	SET,  // 0xD5
	SET,  // 0xD6
	SET,  // 0xD7
	SET,  // 0xD8
	SET,  // 0xD9
	SET,  // 0xDA
	SET,  // 0xDB
	SET,  // 0xDC
	SET,  // 0xDD
	SET,  // 0xDE
	SET,  // 0xDF
	SET,  // 0xE0
	SET,  // 0xE1
	SET,  // 0xE2
	SET,  // 0xE3
	SET,  // 0xE4
	SET,  // 0xE5
	SET,  // 0xE6
	SET,  // 0xE7
	SET,  // 0xE8
	SET,  // 0xE9
	SET,  // 0xEA
	SET,  // 0xEB
	SET,  // 0xEC
	SET,  // 0xED
	SET,  // 0xEE
	SET,  // 0xEF
	SET,  // 0xF0
	SET,  // 0xF1
	SET,  // 0xF2
	SET,  // 0xF3
	SET,  // 0xF4
	SET,  // 0xF5
	SET,  // 0xF6
	SET,  // 0xF7
	SET,  // 0xF8
	SET,  // 0xF9
	SET,  // 0xFA
	SET,  // 0xFB
	SET,  // 0xFC
	SET,  // 0xFD
	SET,  // 0xFE
	SET,  // 0xFF
}

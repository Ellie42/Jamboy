package code

var (
	Ops = map[int]Op{
		0x0000: {
			Type: OpNOP,
			Code: 0x0000,
		},
		0x0001: {
			Type: OpLD,
			Code: 0x0001,
			Operands: []Operand{
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0002: {
			Type: OpLD,
			Code: 0x0002,
			Operands: []Operand{
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0003: {
			Type: OpINC,
			Code: 0x0003,
			Operands: []Operand{
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0004: {
			Type: OpINC,
			Code: 0x0004,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0005: {
			Type: OpDEC,
			Code: 0x0005,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0006: {
			Type: OpLD,
			Code: 0x0006,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0007: {
			Type: OpRLCA,
			Code: 0x0007,
		},
		0x0008: {
			Type: OpLD,
			Code: 0x0008,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0009: {
			Type: OpADD,
			Code: 0x0009,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x000a: {
			Type: OpLD,
			Code: 0x000a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x000b: {
			Type: OpDEC,
			Code: 0x000b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x000c: {
			Type: OpINC,
			Code: 0x000c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x000d: {
			Type: OpDEC,
			Code: 0x000d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x000e: {
			Type: OpLD,
			Code: 0x000e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x000f: {
			Type: OpRRCA,
			Code: 0x000f,
		},
		0x0010: {
			Type: OpSTOP,
			Code: 0x0010,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0011: {
			Type: OpLD,
			Code: 0x0011,
			Operands: []Operand{
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0012: {
			Type: OpLD,
			Code: 0x0012,
			Operands: []Operand{
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0013: {
			Type: OpINC,
			Code: 0x0013,
			Operands: []Operand{
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0014: {
			Type: OpINC,
			Code: 0x0014,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0015: {
			Type: OpDEC,
			Code: 0x0015,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0016: {
			Type: OpLD,
			Code: 0x0016,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0017: {
			Type: OpRLA,
			Code: 0x0017,
		},
		0x0018: {
			Type: OpJR,
			Code: 0x0018,
			Operands: []Operand{
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0019: {
			Type: OpADD,
			Code: 0x0019,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x001a: {
			Type: OpLD,
			Code: 0x001a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x001b: {
			Type: OpDEC,
			Code: 0x001b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x001c: {
			Type: OpINC,
			Code: 0x001c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x001d: {
			Type: OpDEC,
			Code: 0x001d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x001e: {
			Type: OpLD,
			Code: 0x001e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x001f: {
			Type: OpRRA,
			Code: 0x001f,
		},
		0x0020: {
			Type: OpJR,
			Code: 0x0020,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0021: {
			Type: OpLD,
			Code: 0x0021,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0022: {
			Type: OpLD,
			Code: 0x0022,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 1,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0023: {
			Type: OpINC,
			Code: 0x0023,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0024: {
			Type: OpINC,
			Code: 0x0024,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0025: {
			Type: OpDEC,
			Code: 0x0025,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0026: {
			Type: OpLD,
			Code: 0x0026,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0027: {
			Type: OpDAA,
			Code: 0x0027,
		},
		0x0028: {
			Type: OpJR,
			Code: 0x0028,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0029: {
			Type: OpADD,
			Code: 0x0029,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x002a: {
			Type: OpLD,
			Code: 0x002a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 1,
					ValueSize:      2,
				},
			},
		},
		0x002b: {
			Type: OpDEC,
			Code: 0x002b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x002c: {
			Type: OpINC,
			Code: 0x002c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x002d: {
			Type: OpDEC,
			Code: 0x002d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x002e: {
			Type: OpLD,
			Code: 0x002e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x002f: {
			Type: OpCPL,
			Code: 0x002f,
		},
		0x0030: {
			Type: OpJR,
			Code: 0x0030,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0031: {
			Type: OpLD,
			Code: 0x0031,
			Operands: []Operand{
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0032: {
			Type: OpLD,
			Code: 0x0032,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: -1,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0033: {
			Type: OpINC,
			Code: 0x0033,
			Operands: []Operand{
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0034: {
			Type: OpINC,
			Code: 0x0034,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0035: {
			Type: OpDEC,
			Code: 0x0035,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0036: {
			Type: OpLD,
			Code: 0x0036,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0037: {
			Type: OpSCF,
			Code: 0x0037,
		},
		0x0038: {
			Type: OpJR,
			Code: 0x0038,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0039: {
			Type: OpADD,
			Code: 0x0039,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x003a: {
			Type: OpLD,
			Code: 0x003a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: -1,
					ValueSize:      2,
				},
			},
		},
		0x003b: {
			Type: OpDEC,
			Code: 0x003b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x003c: {
			Type: OpINC,
			Code: 0x003c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x003d: {
			Type: OpDEC,
			Code: 0x003d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x003e: {
			Type: OpLD,
			Code: 0x003e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x003f: {
			Type: OpCCF,
			Code: 0x003f,
		},
		0x0040: {
			Type: OpLD,
			Code: 0x0040,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0041: {
			Type: OpLD,
			Code: 0x0041,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0042: {
			Type: OpLD,
			Code: 0x0042,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0043: {
			Type: OpLD,
			Code: 0x0043,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0044: {
			Type: OpLD,
			Code: 0x0044,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0045: {
			Type: OpLD,
			Code: 0x0045,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0046: {
			Type: OpLD,
			Code: 0x0046,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0047: {
			Type: OpLD,
			Code: 0x0047,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0048: {
			Type: OpLD,
			Code: 0x0048,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0049: {
			Type: OpLD,
			Code: 0x0049,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x004a: {
			Type: OpLD,
			Code: 0x004a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x004b: {
			Type: OpLD,
			Code: 0x004b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x004c: {
			Type: OpLD,
			Code: 0x004c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x004d: {
			Type: OpLD,
			Code: 0x004d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x004e: {
			Type: OpLD,
			Code: 0x004e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x004f: {
			Type: OpLD,
			Code: 0x004f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0050: {
			Type: OpLD,
			Code: 0x0050,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0051: {
			Type: OpLD,
			Code: 0x0051,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0052: {
			Type: OpLD,
			Code: 0x0052,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0053: {
			Type: OpLD,
			Code: 0x0053,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0054: {
			Type: OpLD,
			Code: 0x0054,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0055: {
			Type: OpLD,
			Code: 0x0055,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0056: {
			Type: OpLD,
			Code: 0x0056,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0057: {
			Type: OpLD,
			Code: 0x0057,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0058: {
			Type: OpLD,
			Code: 0x0058,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0059: {
			Type: OpLD,
			Code: 0x0059,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x005a: {
			Type: OpLD,
			Code: 0x005a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x005b: {
			Type: OpLD,
			Code: 0x005b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x005c: {
			Type: OpLD,
			Code: 0x005c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x005d: {
			Type: OpLD,
			Code: 0x005d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x005e: {
			Type: OpLD,
			Code: 0x005e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x005f: {
			Type: OpLD,
			Code: 0x005f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0060: {
			Type: OpLD,
			Code: 0x0060,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0061: {
			Type: OpLD,
			Code: 0x0061,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0062: {
			Type: OpLD,
			Code: 0x0062,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0063: {
			Type: OpLD,
			Code: 0x0063,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0064: {
			Type: OpLD,
			Code: 0x0064,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0065: {
			Type: OpLD,
			Code: 0x0065,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0066: {
			Type: OpLD,
			Code: 0x0066,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0067: {
			Type: OpLD,
			Code: 0x0067,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0068: {
			Type: OpLD,
			Code: 0x0068,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0069: {
			Type: OpLD,
			Code: 0x0069,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x006a: {
			Type: OpLD,
			Code: 0x006a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x006b: {
			Type: OpLD,
			Code: 0x006b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x006c: {
			Type: OpLD,
			Code: 0x006c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x006d: {
			Type: OpLD,
			Code: 0x006d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x006e: {
			Type: OpLD,
			Code: 0x006e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x006f: {
			Type: OpLD,
			Code: 0x006f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0070: {
			Type: OpLD,
			Code: 0x0070,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0071: {
			Type: OpLD,
			Code: 0x0071,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0072: {
			Type: OpLD,
			Code: 0x0072,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0073: {
			Type: OpLD,
			Code: 0x0073,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0074: {
			Type: OpLD,
			Code: 0x0074,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0075: {
			Type: OpLD,
			Code: 0x0075,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0076: {
			Type: OpHALT,
			Code: 0x0076,
		},
		0x0077: {
			Type: OpLD,
			Code: 0x0077,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0078: {
			Type: OpLD,
			Code: 0x0078,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0079: {
			Type: OpLD,
			Code: 0x0079,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x007a: {
			Type: OpLD,
			Code: 0x007a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x007b: {
			Type: OpLD,
			Code: 0x007b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x007c: {
			Type: OpLD,
			Code: 0x007c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x007d: {
			Type: OpLD,
			Code: 0x007d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x007e: {
			Type: OpLD,
			Code: 0x007e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x007f: {
			Type: OpLD,
			Code: 0x007f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0080: {
			Type: OpADD,
			Code: 0x0080,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0081: {
			Type: OpADD,
			Code: 0x0081,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0082: {
			Type: OpADD,
			Code: 0x0082,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0083: {
			Type: OpADD,
			Code: 0x0083,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0084: {
			Type: OpADD,
			Code: 0x0084,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0085: {
			Type: OpADD,
			Code: 0x0085,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0086: {
			Type: OpADD,
			Code: 0x0086,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0087: {
			Type: OpADD,
			Code: 0x0087,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0088: {
			Type: OpADC,
			Code: 0x0088,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0089: {
			Type: OpADC,
			Code: 0x0089,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x008a: {
			Type: OpADC,
			Code: 0x008a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x008b: {
			Type: OpADC,
			Code: 0x008b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x008c: {
			Type: OpADC,
			Code: 0x008c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x008d: {
			Type: OpADC,
			Code: 0x008d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x008e: {
			Type: OpADC,
			Code: 0x008e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x008f: {
			Type: OpADC,
			Code: 0x008f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0090: {
			Type: OpSUB,
			Code: 0x0090,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0091: {
			Type: OpSUB,
			Code: 0x0091,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0092: {
			Type: OpSUB,
			Code: 0x0092,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0093: {
			Type: OpSUB,
			Code: 0x0093,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0094: {
			Type: OpSUB,
			Code: 0x0094,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0095: {
			Type: OpSUB,
			Code: 0x0095,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0096: {
			Type: OpSUB,
			Code: 0x0096,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0097: {
			Type: OpSUB,
			Code: 0x0097,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0098: {
			Type: OpSBC,
			Code: 0x0098,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0099: {
			Type: OpSBC,
			Code: 0x0099,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x009a: {
			Type: OpSBC,
			Code: 0x009a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x009b: {
			Type: OpSBC,
			Code: 0x009b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x009c: {
			Type: OpSBC,
			Code: 0x009c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x009d: {
			Type: OpSBC,
			Code: 0x009d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x009e: {
			Type: OpSBC,
			Code: 0x009e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x009f: {
			Type: OpSBC,
			Code: 0x009f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a0: {
			Type: OpAND,
			Code: 0x00a0,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a1: {
			Type: OpAND,
			Code: 0x00a1,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a2: {
			Type: OpAND,
			Code: 0x00a2,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a3: {
			Type: OpAND,
			Code: 0x00a3,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a4: {
			Type: OpAND,
			Code: 0x00a4,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a5: {
			Type: OpAND,
			Code: 0x00a5,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a6: {
			Type: OpAND,
			Code: 0x00a6,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00a7: {
			Type: OpAND,
			Code: 0x00a7,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a8: {
			Type: OpXOR,
			Code: 0x00a8,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00a9: {
			Type: OpXOR,
			Code: 0x00a9,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00aa: {
			Type: OpXOR,
			Code: 0x00aa,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ab: {
			Type: OpXOR,
			Code: 0x00ab,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ac: {
			Type: OpXOR,
			Code: 0x00ac,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ad: {
			Type: OpXOR,
			Code: 0x00ad,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ae: {
			Type: OpXOR,
			Code: 0x00ae,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00af: {
			Type: OpXOR,
			Code: 0x00af,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b0: {
			Type: OpOR,
			Code: 0x00b0,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b1: {
			Type: OpOR,
			Code: 0x00b1,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b2: {
			Type: OpOR,
			Code: 0x00b2,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b3: {
			Type: OpOR,
			Code: 0x00b3,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b4: {
			Type: OpOR,
			Code: 0x00b4,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b5: {
			Type: OpOR,
			Code: 0x00b5,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b6: {
			Type: OpOR,
			Code: 0x00b6,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00b7: {
			Type: OpOR,
			Code: 0x00b7,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b8: {
			Type: OpCP,
			Code: 0x00b8,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00b9: {
			Type: OpCP,
			Code: 0x00b9,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ba: {
			Type: OpCP,
			Code: 0x00ba,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00bb: {
			Type: OpCP,
			Code: 0x00bb,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00bc: {
			Type: OpCP,
			Code: 0x00bc,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00bd: {
			Type: OpCP,
			Code: 0x00bd,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00be: {
			Type: OpCP,
			Code: 0x00be,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00bf: {
			Type: OpCP,
			Code: 0x00bf,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00c0: {
			Type: OpRET,
			Code: 0x00c0,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
			},
		},
		0x00c1: {
			Type: OpPOP,
			Code: 0x00c1,
			Operands: []Operand{
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00c2: {
			Type: OpJP,
			Code: 0x00c2,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00c3: {
			Type: OpJP,
			Code: 0x00c3,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00c4: {
			Type: OpCALL,
			Code: 0x00c4,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00c5: {
			Type: OpPUSH,
			Code: 0x00c5,
			Operands: []Operand{
				{
					ValueStatic:    int(ValBC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00c6: {
			Type: OpADD,
			Code: 0x00c6,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00c7: {
			Type: OpRST,
			Code: 0x00c7,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00c8: {
			Type: OpRET,
			Code: 0x00c8,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
			},
		},
		0x00c9: {
			Type: OpRET,
			Code: 0x00c9,
		},
		0x00ca: {
			Type: OpJP,
			Code: 0x00ca,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00cb: {
			Type: OpPREFIX,
			Code: 0x00cb,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordCB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
			},
		},
		0x00cc: {
			Type: OpCALL,
			Code: 0x00cc,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordZ),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00cd: {
			Type: OpCALL,
			Code: 0x00cd,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00ce: {
			Type: OpADC,
			Code: 0x00ce,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00cf: {
			Type: OpRST,
			Code: 0x00cf,
			Operands: []Operand{
				{
					ValueStatic:    8,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00d0: {
			Type: OpRET,
			Code: 0x00d0,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
			},
		},
		0x00d1: {
			Type: OpPOP,
			Code: 0x00d1,
			Operands: []Operand{
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00d2: {
			Type: OpJP,
			Code: 0x00d2,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00d3: {
			Type: OpNOP,
			Code: 0x00d3,
		},
		0x00d4: {
			Type: OpCALL,
			Code: 0x00d4,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordNC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00d5: {
			Type: OpPUSH,
			Code: 0x00d5,
			Operands: []Operand{
				{
					ValueStatic:    int(ValDE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00d6: {
			Type: OpSUB,
			Code: 0x00d6,
			Operands: []Operand{
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00d7: {
			Type: OpRST,
			Code: 0x00d7,
			Operands: []Operand{
				{
					ValueStatic:    16,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00d8: {
			Type: OpRET,
			Code: 0x00d8,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00d9: {
			Type: OpRETI,
			Code: 0x00d9,
		},
		0x00da: {
			Type: OpJP,
			Code: 0x00da,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00db: {
			Type: OpNOP,
			Code: 0x00db,
		},
		0x00dc: {
			Type: OpCALL,
			Code: 0x00dc,
			Operands: []Operand{
				{
					ValueStatic:    int(ValKeywordC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeKeyword,
					IncDecModifier: 0,
					ValueSize:      0,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00dd: {
			Type: OpNOP,
			Code: 0x00dd,
		},
		0x00de: {
			Type: OpSBC,
			Code: 0x00de,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00df: {
			Type: OpRST,
			Code: 0x00df,
			Operands: []Operand{
				{
					ValueStatic:    24,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00e0: {
			Type: OpLDH,
			Code: 0x00e0,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00e1: {
			Type: OpPOP,
			Code: 0x00e1,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00e2: {
			Type: OpLD,
			Code: 0x00e2,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00e3: {
			Type: OpNOP,
			Code: 0x00e3,
		},
		0x00e4: {
			Type: OpNOP,
			Code: 0x00e4,
		},
		0x00e5: {
			Type: OpPUSH,
			Code: 0x00e5,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00e6: {
			Type: OpAND,
			Code: 0x00e6,
			Operands: []Operand{
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00e7: {
			Type: OpRST,
			Code: 0x00e7,
			Operands: []Operand{
				{
					ValueStatic:    32,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00e8: {
			Type: OpADD,
			Code: 0x00e8,
			Operands: []Operand{
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00e9: {
			Type: OpJP,
			Code: 0x00e9,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00ea: {
			Type: OpLD,
			Code: 0x00ea,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00eb: {
			Type: OpNOP,
			Code: 0x00eb,
		},
		0x00ec: {
			Type: OpNOP,
			Code: 0x00ec,
		},
		0x00ed: {
			Type: OpNOP,
			Code: 0x00ed,
		},
		0x00ee: {
			Type: OpXOR,
			Code: 0x00ee,
			Operands: []Operand{
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ef: {
			Type: OpRST,
			Code: 0x00ef,
			Operands: []Operand{
				{
					ValueStatic:    40,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00f0: {
			Type: OpLDH,
			Code: 0x00f0,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00f1: {
			Type: OpPOP,
			Code: 0x00f1,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAF),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00f2: {
			Type: OpLD,
			Code: 0x00f2,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00f3: {
			Type: OpDI,
			Code: 0x00f3,
		},
		0x00f4: {
			Type: OpNOP,
			Code: 0x00f4,
		},
		0x00f5: {
			Type: OpPUSH,
			Code: 0x00f5,
			Operands: []Operand{
				{
					ValueStatic:    int(ValAF),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00f6: {
			Type: OpOR,
			Code: 0x00f6,
			Operands: []Operand{
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00f7: {
			Type: OpRST,
			Code: 0x00f7,
			Operands: []Operand{
				{
					ValueStatic:    48,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00f8: {
			Type: OpLD,
			Code: 0x00f8,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValSigned),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00f9: {
			Type: OpLD,
			Code: 0x00f9,
			Operands: []Operand{
				{
					ValueStatic:    int(ValSP),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00fa: {
			Type: OpLD,
			Code: 0x00fa,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValAbsolute),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x00fb: {
			Type: OpEI,
			Code: 0x00fb,
		},
		0x00fc: {
			Type: OpNOP,
			Code: 0x00fc,
		},
		0x00fd: {
			Type: OpNOP,
			Code: 0x00fd,
		},
		0x00fe: {
			Type: OpCP,
			Code: 0x00fe,
			Operands: []Operand{
				{
					ValueStatic:    int(ValNumber),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRead,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x00ff: {
			Type: OpRST,
			Code: 0x00ff,
			Operands: []Operand{
				{
					ValueStatic:    56,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0100: {
			Type: OpRLC,
			Code: 0x0100,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0101: {
			Type: OpRLC,
			Code: 0x0101,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0102: {
			Type: OpRLC,
			Code: 0x0102,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0103: {
			Type: OpRLC,
			Code: 0x0103,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0104: {
			Type: OpRLC,
			Code: 0x0104,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0105: {
			Type: OpRLC,
			Code: 0x0105,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0106: {
			Type: OpRLC,
			Code: 0x0106,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0107: {
			Type: OpRLC,
			Code: 0x0107,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0108: {
			Type: OpRRC,
			Code: 0x0108,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0109: {
			Type: OpRRC,
			Code: 0x0109,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x010a: {
			Type: OpRRC,
			Code: 0x010a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x010b: {
			Type: OpRRC,
			Code: 0x010b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x010c: {
			Type: OpRRC,
			Code: 0x010c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x010d: {
			Type: OpRRC,
			Code: 0x010d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x010e: {
			Type: OpRRC,
			Code: 0x010e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x010f: {
			Type: OpRRC,
			Code: 0x010f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0110: {
			Type: OpRL,
			Code: 0x0110,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0111: {
			Type: OpRL,
			Code: 0x0111,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0112: {
			Type: OpRL,
			Code: 0x0112,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0113: {
			Type: OpRL,
			Code: 0x0113,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0114: {
			Type: OpRL,
			Code: 0x0114,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0115: {
			Type: OpRL,
			Code: 0x0115,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0116: {
			Type: OpRL,
			Code: 0x0116,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0117: {
			Type: OpRL,
			Code: 0x0117,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0118: {
			Type: OpRR,
			Code: 0x0118,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0119: {
			Type: OpRR,
			Code: 0x0119,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x011a: {
			Type: OpRR,
			Code: 0x011a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x011b: {
			Type: OpRR,
			Code: 0x011b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x011c: {
			Type: OpRR,
			Code: 0x011c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x011d: {
			Type: OpRR,
			Code: 0x011d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x011e: {
			Type: OpRR,
			Code: 0x011e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x011f: {
			Type: OpRR,
			Code: 0x011f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0120: {
			Type: OpSLA,
			Code: 0x0120,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0121: {
			Type: OpSLA,
			Code: 0x0121,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0122: {
			Type: OpSLA,
			Code: 0x0122,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0123: {
			Type: OpSLA,
			Code: 0x0123,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0124: {
			Type: OpSLA,
			Code: 0x0124,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0125: {
			Type: OpSLA,
			Code: 0x0125,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0126: {
			Type: OpSLA,
			Code: 0x0126,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0127: {
			Type: OpSLA,
			Code: 0x0127,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0128: {
			Type: OpSRA,
			Code: 0x0128,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0129: {
			Type: OpSRA,
			Code: 0x0129,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x012a: {
			Type: OpSRA,
			Code: 0x012a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x012b: {
			Type: OpSRA,
			Code: 0x012b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x012c: {
			Type: OpSRA,
			Code: 0x012c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x012d: {
			Type: OpSRA,
			Code: 0x012d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x012e: {
			Type: OpSRA,
			Code: 0x012e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x012f: {
			Type: OpSRA,
			Code: 0x012f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0130: {
			Type: OpSWAP,
			Code: 0x0130,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0131: {
			Type: OpSWAP,
			Code: 0x0131,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0132: {
			Type: OpSWAP,
			Code: 0x0132,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0133: {
			Type: OpSWAP,
			Code: 0x0133,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0134: {
			Type: OpSWAP,
			Code: 0x0134,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0135: {
			Type: OpSWAP,
			Code: 0x0135,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0136: {
			Type: OpSWAP,
			Code: 0x0136,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0137: {
			Type: OpSWAP,
			Code: 0x0137,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0138: {
			Type: OpSRL,
			Code: 0x0138,
			Operands: []Operand{
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0139: {
			Type: OpSRL,
			Code: 0x0139,
			Operands: []Operand{
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x013a: {
			Type: OpSRL,
			Code: 0x013a,
			Operands: []Operand{
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x013b: {
			Type: OpSRL,
			Code: 0x013b,
			Operands: []Operand{
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x013c: {
			Type: OpSRL,
			Code: 0x013c,
			Operands: []Operand{
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x013d: {
			Type: OpSRL,
			Code: 0x013d,
			Operands: []Operand{
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x013e: {
			Type: OpSRL,
			Code: 0x013e,
			Operands: []Operand{
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x013f: {
			Type: OpSRL,
			Code: 0x013f,
			Operands: []Operand{
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0140: {
			Type: OpBIT,
			Code: 0x0140,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0141: {
			Type: OpBIT,
			Code: 0x0141,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0142: {
			Type: OpBIT,
			Code: 0x0142,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0143: {
			Type: OpBIT,
			Code: 0x0143,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0144: {
			Type: OpBIT,
			Code: 0x0144,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0145: {
			Type: OpBIT,
			Code: 0x0145,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0146: {
			Type: OpBIT,
			Code: 0x0146,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0147: {
			Type: OpBIT,
			Code: 0x0147,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0148: {
			Type: OpBIT,
			Code: 0x0148,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0149: {
			Type: OpBIT,
			Code: 0x0149,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x014a: {
			Type: OpBIT,
			Code: 0x014a,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x014b: {
			Type: OpBIT,
			Code: 0x014b,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x014c: {
			Type: OpBIT,
			Code: 0x014c,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x014d: {
			Type: OpBIT,
			Code: 0x014d,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x014e: {
			Type: OpBIT,
			Code: 0x014e,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x014f: {
			Type: OpBIT,
			Code: 0x014f,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0150: {
			Type: OpBIT,
			Code: 0x0150,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0151: {
			Type: OpBIT,
			Code: 0x0151,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0152: {
			Type: OpBIT,
			Code: 0x0152,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0153: {
			Type: OpBIT,
			Code: 0x0153,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0154: {
			Type: OpBIT,
			Code: 0x0154,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0155: {
			Type: OpBIT,
			Code: 0x0155,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0156: {
			Type: OpBIT,
			Code: 0x0156,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0157: {
			Type: OpBIT,
			Code: 0x0157,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0158: {
			Type: OpBIT,
			Code: 0x0158,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0159: {
			Type: OpBIT,
			Code: 0x0159,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x015a: {
			Type: OpBIT,
			Code: 0x015a,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x015b: {
			Type: OpBIT,
			Code: 0x015b,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x015c: {
			Type: OpBIT,
			Code: 0x015c,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x015d: {
			Type: OpBIT,
			Code: 0x015d,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x015e: {
			Type: OpBIT,
			Code: 0x015e,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x015f: {
			Type: OpBIT,
			Code: 0x015f,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0160: {
			Type: OpBIT,
			Code: 0x0160,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0161: {
			Type: OpBIT,
			Code: 0x0161,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0162: {
			Type: OpBIT,
			Code: 0x0162,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0163: {
			Type: OpBIT,
			Code: 0x0163,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0164: {
			Type: OpBIT,
			Code: 0x0164,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0165: {
			Type: OpBIT,
			Code: 0x0165,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0166: {
			Type: OpBIT,
			Code: 0x0166,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0167: {
			Type: OpBIT,
			Code: 0x0167,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0168: {
			Type: OpBIT,
			Code: 0x0168,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0169: {
			Type: OpBIT,
			Code: 0x0169,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x016a: {
			Type: OpBIT,
			Code: 0x016a,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x016b: {
			Type: OpBIT,
			Code: 0x016b,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x016c: {
			Type: OpBIT,
			Code: 0x016c,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x016d: {
			Type: OpBIT,
			Code: 0x016d,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x016e: {
			Type: OpBIT,
			Code: 0x016e,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x016f: {
			Type: OpBIT,
			Code: 0x016f,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0170: {
			Type: OpBIT,
			Code: 0x0170,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0171: {
			Type: OpBIT,
			Code: 0x0171,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0172: {
			Type: OpBIT,
			Code: 0x0172,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0173: {
			Type: OpBIT,
			Code: 0x0173,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0174: {
			Type: OpBIT,
			Code: 0x0174,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0175: {
			Type: OpBIT,
			Code: 0x0175,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0176: {
			Type: OpBIT,
			Code: 0x0176,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0177: {
			Type: OpBIT,
			Code: 0x0177,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0178: {
			Type: OpBIT,
			Code: 0x0178,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0179: {
			Type: OpBIT,
			Code: 0x0179,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x017a: {
			Type: OpBIT,
			Code: 0x017a,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x017b: {
			Type: OpBIT,
			Code: 0x017b,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x017c: {
			Type: OpBIT,
			Code: 0x017c,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x017d: {
			Type: OpBIT,
			Code: 0x017d,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x017e: {
			Type: OpBIT,
			Code: 0x017e,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x017f: {
			Type: OpBIT,
			Code: 0x017f,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0180: {
			Type: OpRES,
			Code: 0x0180,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0181: {
			Type: OpRES,
			Code: 0x0181,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0182: {
			Type: OpRES,
			Code: 0x0182,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0183: {
			Type: OpRES,
			Code: 0x0183,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0184: {
			Type: OpRES,
			Code: 0x0184,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0185: {
			Type: OpRES,
			Code: 0x0185,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0186: {
			Type: OpRES,
			Code: 0x0186,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0187: {
			Type: OpRES,
			Code: 0x0187,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0188: {
			Type: OpRES,
			Code: 0x0188,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0189: {
			Type: OpRES,
			Code: 0x0189,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x018a: {
			Type: OpRES,
			Code: 0x018a,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x018b: {
			Type: OpRES,
			Code: 0x018b,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x018c: {
			Type: OpRES,
			Code: 0x018c,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x018d: {
			Type: OpRES,
			Code: 0x018d,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x018e: {
			Type: OpRES,
			Code: 0x018e,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x018f: {
			Type: OpRES,
			Code: 0x018f,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0190: {
			Type: OpRES,
			Code: 0x0190,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0191: {
			Type: OpRES,
			Code: 0x0191,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0192: {
			Type: OpRES,
			Code: 0x0192,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0193: {
			Type: OpRES,
			Code: 0x0193,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0194: {
			Type: OpRES,
			Code: 0x0194,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0195: {
			Type: OpRES,
			Code: 0x0195,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0196: {
			Type: OpRES,
			Code: 0x0196,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x0197: {
			Type: OpRES,
			Code: 0x0197,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0198: {
			Type: OpRES,
			Code: 0x0198,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x0199: {
			Type: OpRES,
			Code: 0x0199,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x019a: {
			Type: OpRES,
			Code: 0x019a,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x019b: {
			Type: OpRES,
			Code: 0x019b,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x019c: {
			Type: OpRES,
			Code: 0x019c,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x019d: {
			Type: OpRES,
			Code: 0x019d,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x019e: {
			Type: OpRES,
			Code: 0x019e,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x019f: {
			Type: OpRES,
			Code: 0x019f,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a0: {
			Type: OpRES,
			Code: 0x01a0,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a1: {
			Type: OpRES,
			Code: 0x01a1,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a2: {
			Type: OpRES,
			Code: 0x01a2,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a3: {
			Type: OpRES,
			Code: 0x01a3,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a4: {
			Type: OpRES,
			Code: 0x01a4,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a5: {
			Type: OpRES,
			Code: 0x01a5,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a6: {
			Type: OpRES,
			Code: 0x01a6,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01a7: {
			Type: OpRES,
			Code: 0x01a7,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a8: {
			Type: OpRES,
			Code: 0x01a8,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01a9: {
			Type: OpRES,
			Code: 0x01a9,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01aa: {
			Type: OpRES,
			Code: 0x01aa,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ab: {
			Type: OpRES,
			Code: 0x01ab,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ac: {
			Type: OpRES,
			Code: 0x01ac,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ad: {
			Type: OpRES,
			Code: 0x01ad,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ae: {
			Type: OpRES,
			Code: 0x01ae,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01af: {
			Type: OpRES,
			Code: 0x01af,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b0: {
			Type: OpRES,
			Code: 0x01b0,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b1: {
			Type: OpRES,
			Code: 0x01b1,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b2: {
			Type: OpRES,
			Code: 0x01b2,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b3: {
			Type: OpRES,
			Code: 0x01b3,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b4: {
			Type: OpRES,
			Code: 0x01b4,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b5: {
			Type: OpRES,
			Code: 0x01b5,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b6: {
			Type: OpRES,
			Code: 0x01b6,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01b7: {
			Type: OpRES,
			Code: 0x01b7,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b8: {
			Type: OpRES,
			Code: 0x01b8,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01b9: {
			Type: OpRES,
			Code: 0x01b9,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ba: {
			Type: OpRES,
			Code: 0x01ba,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01bb: {
			Type: OpRES,
			Code: 0x01bb,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01bc: {
			Type: OpRES,
			Code: 0x01bc,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01bd: {
			Type: OpRES,
			Code: 0x01bd,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01be: {
			Type: OpRES,
			Code: 0x01be,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01bf: {
			Type: OpRES,
			Code: 0x01bf,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c0: {
			Type: OpSET,
			Code: 0x01c0,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c1: {
			Type: OpSET,
			Code: 0x01c1,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c2: {
			Type: OpSET,
			Code: 0x01c2,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c3: {
			Type: OpSET,
			Code: 0x01c3,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c4: {
			Type: OpSET,
			Code: 0x01c4,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c5: {
			Type: OpSET,
			Code: 0x01c5,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c6: {
			Type: OpSET,
			Code: 0x01c6,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01c7: {
			Type: OpSET,
			Code: 0x01c7,
			Operands: []Operand{
				{
					ValueStatic:    0,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c8: {
			Type: OpSET,
			Code: 0x01c8,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01c9: {
			Type: OpSET,
			Code: 0x01c9,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ca: {
			Type: OpSET,
			Code: 0x01ca,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01cb: {
			Type: OpSET,
			Code: 0x01cb,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01cc: {
			Type: OpSET,
			Code: 0x01cc,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01cd: {
			Type: OpSET,
			Code: 0x01cd,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ce: {
			Type: OpSET,
			Code: 0x01ce,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01cf: {
			Type: OpSET,
			Code: 0x01cf,
			Operands: []Operand{
				{
					ValueStatic:    1,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d0: {
			Type: OpSET,
			Code: 0x01d0,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d1: {
			Type: OpSET,
			Code: 0x01d1,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d2: {
			Type: OpSET,
			Code: 0x01d2,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d3: {
			Type: OpSET,
			Code: 0x01d3,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d4: {
			Type: OpSET,
			Code: 0x01d4,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d5: {
			Type: OpSET,
			Code: 0x01d5,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d6: {
			Type: OpSET,
			Code: 0x01d6,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01d7: {
			Type: OpSET,
			Code: 0x01d7,
			Operands: []Operand{
				{
					ValueStatic:    2,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d8: {
			Type: OpSET,
			Code: 0x01d8,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01d9: {
			Type: OpSET,
			Code: 0x01d9,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01da: {
			Type: OpSET,
			Code: 0x01da,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01db: {
			Type: OpSET,
			Code: 0x01db,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01dc: {
			Type: OpSET,
			Code: 0x01dc,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01dd: {
			Type: OpSET,
			Code: 0x01dd,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01de: {
			Type: OpSET,
			Code: 0x01de,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01df: {
			Type: OpSET,
			Code: 0x01df,
			Operands: []Operand{
				{
					ValueStatic:    3,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e0: {
			Type: OpSET,
			Code: 0x01e0,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e1: {
			Type: OpSET,
			Code: 0x01e1,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e2: {
			Type: OpSET,
			Code: 0x01e2,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e3: {
			Type: OpSET,
			Code: 0x01e3,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e4: {
			Type: OpSET,
			Code: 0x01e4,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e5: {
			Type: OpSET,
			Code: 0x01e5,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e6: {
			Type: OpSET,
			Code: 0x01e6,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01e7: {
			Type: OpSET,
			Code: 0x01e7,
			Operands: []Operand{
				{
					ValueStatic:    4,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e8: {
			Type: OpSET,
			Code: 0x01e8,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01e9: {
			Type: OpSET,
			Code: 0x01e9,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ea: {
			Type: OpSET,
			Code: 0x01ea,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01eb: {
			Type: OpSET,
			Code: 0x01eb,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ec: {
			Type: OpSET,
			Code: 0x01ec,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ed: {
			Type: OpSET,
			Code: 0x01ed,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01ee: {
			Type: OpSET,
			Code: 0x01ee,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01ef: {
			Type: OpSET,
			Code: 0x01ef,
			Operands: []Operand{
				{
					ValueStatic:    5,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f0: {
			Type: OpSET,
			Code: 0x01f0,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f1: {
			Type: OpSET,
			Code: 0x01f1,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f2: {
			Type: OpSET,
			Code: 0x01f2,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f3: {
			Type: OpSET,
			Code: 0x01f3,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f4: {
			Type: OpSET,
			Code: 0x01f4,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f5: {
			Type: OpSET,
			Code: 0x01f5,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f6: {
			Type: OpSET,
			Code: 0x01f6,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01f7: {
			Type: OpSET,
			Code: 0x01f7,
			Operands: []Operand{
				{
					ValueStatic:    6,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f8: {
			Type: OpSET,
			Code: 0x01f8,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValB),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01f9: {
			Type: OpSET,
			Code: 0x01f9,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValC),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01fa: {
			Type: OpSET,
			Code: 0x01fa,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValD),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01fb: {
			Type: OpSET,
			Code: 0x01fb,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValE),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01fc: {
			Type: OpSET,
			Code: 0x01fc,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValH),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01fd: {
			Type: OpSET,
			Code: 0x01fd,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValL),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
		0x01fe: {
			Type: OpSET,
			Code: 0x01fe,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValHL),
					RetrieveType:   RetrievePointer,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      2,
				},
			},
		},
		0x01ff: {
			Type: OpSET,
			Code: 0x01ff,
			Operands: []Operand{
				{
					ValueStatic:    7,
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeConst,
					IncDecModifier: 0,
					ValueSize:      1,
				},
				{
					ValueStatic:    int(ValA),
					RetrieveType:   RetrieveVal,
					ValueType:      ValTypeRegister,
					IncDecModifier: 0,
					ValueSize:      1,
				},
			},
		},
	}
)

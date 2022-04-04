package ast

type Exp interface {}

type NilExp			struct { line int }
type TrueExp		struct { line int }
type FalseExp		struct { line int }
type VarargExp		struct { line int }
type IntegerExp		struct { line int; Val int64   }
type FloatExp		struct { line int; Val float64 }
type StringExp		struct { line int; Str string  }
type NameExp		struct { line int; Name string }


type UnopExp struct {
	Line int
	Op	 int
	Exp  Exp 
}

type BinopExp struct {
	Line int
	Op   int
	Exp1 Exp
	Exp2 Exp
}

type ConcatExp struct {
	Line int
	Exps []Exp
}

type TableConstructorExp struct {
	Line	 int
	LastLine int
	KeyExps  []Exp 
	ValExps  []Exp 
}

type FuncDefExp struct {
	Line     int
	LastLine int
	ParList  []string
	IsVararg bool
	Block    *Block
}

type ParensExp struct {
	Exp Exp 
}

type TableAccessExp struct {
	LastLine  int
	PrefixExp Exp 
	KeyExp    Exp
}

type FuncCallExp struct {
	Line      int //'('
	LastLine  int //')'
	PrefixExp Exp 
	NameExp   *StringExp 
	Args      []Exp 
}


package ast

type Stat interface{}

type EmptyStat struct{}

type BreakStat struct{
	Line int
}

type LabelStat struct{
	Name string
}

type GotoStat struct{
	Name string
}

type DoStat struct{
	Block *Block
}

type FuncCallStat = FuncCallExp


type WhileStat struct {
	Exp   Exp
	Block *Block
	//EBNF: while exp do block end
}

type RepeatStat struct {
	Block *Block
	Exp   Exp 
	//EBNF: repeat block until exp
}

type IfStat struct {
	Exps   []Exp
	Blocks []*Block
	//EBNF: if exp then block { else if exp then block } end
}

type ForNumStat struct {
	//EBNF: for Name '=' exp ',' exp [',' exp] do block end
	LineOfFor int
	LineOfDo  int
	//for code generation

	VarName   string
	InitExp   Exp 
	LimitExp  Exp 
	StepExp   Exp 
	Block     *Block
}

type ForInStat struct {
	//EBNF: for namelist in explist do block end
	//      namelist ::= Name {',' Name}
	//		explist  ::= exp  {',' exp}


	LineOfDo int		//for code generation
	NameList []string
	ExpList  []Exp
	Block    *Block
}

type LocalVarDeclStat struct {
	//EBNF: local namelist ['=' explist]
	//      namelist ::= Name {',' Name}
	//		explist  ::= exp  {',' exp }
	LastLine int
	NameList []string
	ExpList  []Exp 
}

type AssignStat struct {
	LastLine int
	VarList  []Exp 
	ExpList  []Exp 
}

type LocalFuncDefStat struct {
	Name string
	Exp  *FuncDefExp
}
package mylang

import (
	"fmt"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func Codegen(prg *program) string {
	m := ir.NewModule()
	main := m.NewFunc("main", types.I8)
	entry := main.NewBlock("")

	printf := m.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr))
	printf.Sig.Variadic = true

	printfn := func(val value.Value) {
		format := constant.NewCharArrayFromString("%d\n")
		if types.IsFloat(val.Type()) {
			format = constant.NewCharArrayFromString("%f\n")
		}

		str := m.NewGlobalDef("s", format)
		str.Immutable = true
		str.UnnamedAddr = enum.UnnamedAddrUnnamedAddr
		s := constant.NewGetElementPtr(
			types.NewArray(format.Typ.Len, types.I8),
			str,
			constant.NewInt(types.I64, 0),
			constant.NewInt(types.I64, 0),
		)
		s.InBounds = true
		entry.NewCall(printf, s, val)
	}

	codegenStatement(entry, prg, printfn)
	entry.NewRet(constant.NewInt(types.I8, 0))

	return m.String()
}

func codegenStatement(entry *ir.Block, prg *program, printfn func(val value.Value)) {
	vars := make(map[string]*ir.InstAlloca)

	for _, statement := range prg.statements {
		switch stmt := statement.(type) {
		case *assignStatement:
			val := codegenExpression(entry, stmt.expr)
			a := entry.NewAlloca(val.Type())
			a.SetName(stmt.identifier.literal)
			vars[stmt.identifier.literal] = a
			entry.NewStore(val, a)
			continue

		case *printStatement:
			v := vars[stmt.identifier.literal]
			l := entry.NewLoad(v.ElemType, v)
			printfn(l)
			continue
		default:
			fmt.Println("???", stmt)
		}

		panic("unexpected statement")
	}
}

func codegenExpression(entry *ir.Block, expression expression) value.Value {
	switch expr := expression.(type) {
	case basicLiteral:
		switch expr.kind {
		case FLOAT:
			f, err := strconv.ParseFloat(expr.literal, 64)
			if err != nil {
				panic(err)
			}
			return constant.NewFloat(types.Double, float64(f))
		case INT:
			i, err := strconv.ParseInt(expr.literal, 10, 64)
			if err != nil {
				panic(err)
			}
			return constant.NewInt(types.I64, i)
		}

	case binaryExpression:
		left := codegenExpression(entry, expr.left)
		right := codegenExpression(entry, expr.right)
		if !left.Type().Equal(right.Type()) {
			panic(fmt.Sprintf("unmatched type: %s, %s", left.String(), right.String()))
		}

		switch expr.operator {
		case ADD:
			if types.IsInt(left.Type()) {
				return entry.NewAdd(left, right)
			}
			return entry.NewFAdd(left, right)
		case SUB:
			if types.IsInt(left.Type()) {
				return entry.NewSub(left, right)
			}
			return entry.NewFSub(left, right)
		case MUL:
			if types.IsInt(left.Type()) {
				return entry.NewMul(left, right)
			}
			return entry.NewFMul(left, right)
		case DIV:
			if types.IsInt(left.Type()) {
				return entry.NewSDiv(left, right)
			}
			return entry.NewFDiv(left, right)
		case REM:
			if types.IsInt(left.Type()) {
				return entry.NewSRem(left, right)
			}
			return entry.NewFRem(left, right)
		default:
			panic("???")
		}
	}
	panic("???")
}

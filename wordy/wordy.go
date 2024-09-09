package wordy

import (
	"strconv"
	"strings"
)

type Term interface {
	Compute() int
}

type Constant struct {
	value int
}

func (c Constant) Compute() int {
	return c.value
}

type Expression struct {
	first  Term
	second Term
	op     Operator
}

func (e Expression) Compute() int {
	return e.op.apply(e.first.Compute(), e.second.Compute())
}

type Operator interface {
	apply(n1 int, n2 int) int
}

type minus struct {
}

func (m minus) apply(n1 int, n2 int) int {
	return n1 - n2
}

type plus struct {
}

func (p plus) apply(n1 int, n2 int) int {
	return n1 + n2
}

type divide struct {
}

func (p divide) apply(n1 int, n2 int) int {
	return n1 / n2
}

type multiply struct {
}

func (p multiply) apply(n1 int, n2 int) int {
	return n1 * n2
}

func Answer(question string) (int, bool) {
	expression := getExpression(question)

	if len(expression) != 1 && len(expression) != 3 && len(expression) != 5 {
		return 0, false
	}

	firstNumber, err := strconv.Atoi(expression[0])
	if err != nil {
		return 0, false
	}
	firstConstant := Constant{firstNumber}

	if len(expression) == 1 {
		return firstConstant.value, true
	}

	secondNumber, err := strconv.Atoi(expression[2])
	if err != nil {
		return 0, false
	}
	secondConstant := Constant{secondNumber}

	exp, expOk := buildExpression(firstConstant, secondConstant, expression[1])
	if !expOk {
		return 0, false
	}

	if len(expression) == 3 {
		return exp.Compute(), true
	}

	thirdNumber, ok := strconv.Atoi(expression[4])
	if ok != nil {
		return 0, false
	}
	fExp, fExpOk := buildExpression(exp, Constant{value: thirdNumber}, expression[3])
	if !fExpOk {
		return 0, false
	}
	return fExp.Compute(), true
}

func getExpression(question string) []string {
	question = strings.Replace(question, "What is ", "", -1)
	question = strings.Replace(question, "?", "", -1)
	question = strings.Replace(question, "plus", "+", -1)
	question = strings.Replace(question, "minus", "-", -1)
	question = strings.Replace(question, "multiplied by", "*", -1)
	question = strings.Replace(question, "divided by", "/", -1)

	words := strings.Split(question, " ")
	return words
}

func buildExpression(term1 Term, term2 Term, opStr string) (Expression, bool) {
	operator, ok := getOperator(opStr)
	if !ok {
		return Expression{}, false
	}

	exp := Expression{
		first:  term1,
		second: term2,
		op:     operator,
	}
	return exp, true
}

func getOperator(opStr string) (Operator, bool) {
	oMap := map[string]Operator{
		"+": plus{},
		"-": minus{},
		"/": divide{},
		"*": multiply{},
	}
	operator, ok := oMap[opStr]
	return operator, ok
}

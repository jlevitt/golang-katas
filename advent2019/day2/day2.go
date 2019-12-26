package day2

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program = []int
type programCounter = int

const InstructionAdd = 1
const InstructionMul = 2
const InstructionHalt = 99

type opcode interface {
	execute(programCounter, program) programCounter
}

type binaryOpCode struct {
	leftAddr int
	rightAddr int
	outputAddr int
}

func (op binaryOpCode) args(p program) (int, int) {
	return p[op.leftAddr], p[op.rightAddr]
}

func (op binaryOpCode) advance(pc programCounter) int {
	return pc + 4
}

func (op binaryOpCode) store(p program, addr int, value int) {
	p[addr] = value
}

type add struct {
	binaryOpCode
}

func (op add) execute(pc programCounter, p program) programCounter {
	left, right := op.args(p)
	result := left + right
	op.store(p, op.outputAddr, result)

	return op.advance(pc)
}


type multiply struct {
	binaryOpCode
}

func (op multiply) execute(pc programCounter, p program) programCounter {
	left, right := op.args(p)
	result := left * right
	op.store(p, op.outputAddr, result)

	return op.advance(pc)
}

type halt struct {
}

func (op halt) execute(pc programCounter, p program) programCounter {
	fmt.Print(toString(p))
	os.Exit(0)

	return -1
}

func parseOpcode(pc programCounter, p program) (opcode, error) {
	opcodeInt := p[pc]

	if opcodeInt == InstructionAdd {
		op := add{
			binaryOpCode{
				leftAddr:   p[pc+1],
				rightAddr:  p[pc+2],
				outputAddr: p[pc+3],
			},
		}

		return op, nil
	} else if opcodeInt == InstructionMul {
			op := multiply{
				binaryOpCode{
					leftAddr: p[pc + 1],
					rightAddr: p[pc + 2],
					outputAddr: p[pc + 3],
				},
			}

		return op, nil
	} else if opcodeInt == InstructionHalt {
		return halt{}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Unknown opcode: %v", opcodeInt))
	}
}

func SetProgramAlarm(p program) {
	p[1] = 12
	p[2] = 2
}

func RunProgram(p program) (string, error) {
	pc := 0

	for pc < len(p) {
		op, err := parseOpcode(pc, p)

		if err != nil {
			return "", err
		}
		pc = op.execute(pc, p)
	}

	return toString(p), nil
}

func ParseProgram(programStr string) (program, error) {
	words := strings.Split(programStr, ",")

	var result program
	for i, wordStr := range words {
		word, err := strconv.Atoi(wordStr)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing programStr at position %v", i))
		}
		result = append(result, word)
	}

	return result, nil
}

func toString(p program) string {
	var words []string

	for _, word := range p {
		words = append(words, strconv.Itoa(word))
	}

	return strings.Join(words, ",")
}
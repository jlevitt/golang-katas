package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Program = []int
type programCounter = int

const InstructionAdd = 1
const InstructionMul = 2
const InstructionHalt = 99

const SignalHalt = -1

type opcode interface {
	execute(programCounter, Program) programCounter
}

type binaryOpCode struct {
	leftAddr int
	rightAddr int
	outputAddr int
}

func (op binaryOpCode) args(p Program) (int, int) {
	return p[op.leftAddr], p[op.rightAddr]
}

func (op binaryOpCode) advance(pc programCounter) int {
	return pc + 4
}

func (op binaryOpCode) store(p Program, addr int, value int) {
	p[addr] = value
}

type add struct {
	binaryOpCode
}

func (op add) execute(pc programCounter, p Program) programCounter {
	left, right := op.args(p)
	result := left + right
	op.store(p, op.outputAddr, result)

	return op.advance(pc)
}


type multiply struct {
	binaryOpCode
}

func (op multiply) execute(pc programCounter, p Program) programCounter {
	left, right := op.args(p)
	result := left * right
	op.store(p, op.outputAddr, result)

	return op.advance(pc)
}

type halt struct {
}

func (op halt) execute(pc programCounter, p Program) programCounter {
	return SignalHalt
}

func parseOpcode(pc programCounter, p Program) (opcode, error) {
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

func SetNounVerb(p Program, noun int, verb int) {
	p[1] = noun
	p[2] = verb
}

func SetProgramAlarm(p Program) {
	SetNounVerb(p, 12, 2)
}

func RunProgram(p Program) (string, error) {
	pc := 0

	for pc < len(p) {
		op, err := parseOpcode(pc, p)

		if err != nil {
			return "", err
		}
		pc = op.execute(pc, p)

		if pc == SignalHalt {
			break
		}
	}

	return toString(p), nil
}

func ParseProgram(programStr string) (Program, error) {
	words := strings.Split(programStr, ",")

	var result Program
	for i, wordStr := range words {
		word, err := strconv.Atoi(wordStr)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing programStr at position %v", i))
		}
		result = append(result, word)
	}

	return result, nil
}

func toString(p Program) string {
	var words []string

	for _, word := range p {
		words = append(words, strconv.Itoa(word))
	}

	return strings.Join(words, ",")
}
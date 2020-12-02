package day2

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jlevitt/katas/advent/input"
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
	leftAddr   int
	rightAddr  int
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

func PartOne(path string) error {
	p := loadProgram(path)

	SetProgramAlarm(p)

	program, err := RunProgram(p)

	if err != nil {
		return err
	}

	fmt.Print(program)

	return nil
}

func PartTwo(path string) error {
	initialProgram := loadProgram(path)

	const expectedOutput = 19690720

	noun, verb, err := searchForOutput(initialProgram, expectedOutput)
	if err != nil {
		return err
	}

	fmt.Printf("Found noun=%v, verb=%v\n", noun, verb)

	return nil
}

func loadProgram(path string) Program {
	lines := input.ReadInputLines(path)
	p, err := ParseProgram(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func searchForOutput(initialProgram Program, expectedOutput int) (int, int, error) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			p := make(Program, len(initialProgram))
			copy(p, initialProgram)
			SetNounVerb(p, noun, verb)
			_, err := RunProgram(p)
			if err != nil {
				return -1, -1, err
			}
			output := p[0]

			if output == expectedOutput {
				return noun, verb, nil
			}
		}
	}

	return -1, -1, errors.New("expected output not found")
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
				leftAddr:   p[pc+1],
				rightAddr:  p[pc+2],
				outputAddr: p[pc+3],
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

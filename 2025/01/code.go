package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/samber/lo"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	instructions := parse(input)
	var currentIndex int64 = 50
	var zeroCounter int64 = 0
	if part2 {
		for _, instruction := range instructions {
			newIndex, zeroCount := instruction.getIndexAndZeroCount(100, currentIndex, false)
			currentIndex = newIndex
			zeroCounter += zeroCount
		}
	} else {
		for _, instruction := range instructions {
			newIndex, zeroCount := instruction.getIndexAndZeroCount(100, currentIndex, true)
			currentIndex = newIndex
			zeroCounter += zeroCount
		}
	}
	
	return zeroCounter
}

func parse(input string) []Instruction {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := lo.Map(lines, func(x string, index int) Instruction {
		instructionPieces := []rune(x)
		direction := instructionPieces[0]
		values := string(instructionPieces[1:])
		instruction_value, _ := strconv.ParseInt(values, 10, 64)
		return NewInstruction(direction, instruction_value)
	})
	return instructions
}

type Direction int

const (
	Left Direction = iota
	Right
)

func DirectionFromChar(c rune) (Direction, bool) {
	switch c {
	case 'L':
		return Left, true
	case 'R':
		return Right, true
	default:
		return 0, false
	}
}

type Instruction struct {
	direction Direction
	value int64
}

func NewInstruction(c rune, value int64) Instruction {
	dir, ok := DirectionFromChar(c)
	if !ok {
		panic(fmt.Sprintf("invalid direction character: %c", c))
	}
	return Instruction{
		direction: dir,
		value: value,
	}
}

func (instruction Instruction) getIndexAndZeroCount(safeDialEnd int64, currentIndex int64, part1 bool) (int64, int64) {
	var passes int64
	switch instruction.direction {
	case Left:
		newIndex :=
			(currentIndex + safeDialEnd - (instruction.value % safeDialEnd)) % safeDialEnd;
		if part1 {
			var passes int64
			passes = 0
			if newIndex == 0 {
                passes = 1
            }
			return newIndex, passes
		} else {
			passes := (instruction.value + safeDialEnd - currentIndex) / safeDialEnd;
			if currentIndex == 0 && passes > 0 {
				// If we start at zero and move left, we count an extra pass
				return newIndex, passes - 1
			} else {
				return newIndex, passes
			}
		}
	case Right:
		newIndex := (currentIndex + instruction.value) % safeDialEnd;
		if part1 {
			passes = 0
            if newIndex == 0 {
                passes = 1
            }
			return newIndex, passes
		} else {
			passes := (currentIndex + instruction.value) / safeDialEnd;
			return newIndex, passes
		}
	}
	return 0, 0
}

package main

import (
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
	id_ranges := parse(input)

	invalid_ids := lo.Map(id_ranges, func(id_range InclusiveRange, index int) uint64 {
		var sum uint64 = 0
		for id := id_range.start; id <= id_range.end; id++ {
			sum += checkIfValid(id, part2)
		}
		return sum
	})

	return lo.Sum(invalid_ids)
}

type InclusiveRange struct {
	start uint64
	end uint64
}

func parse(input string) []InclusiveRange {
	ranges := strings.Split(strings.TrimSpace(input), "\n")
	ranges = strings.Split(ranges[0], ",")
	inclusive_ranges := lo.Map(ranges, func(inclusive_range string, index int) InclusiveRange {
		pieces := strings.Split(inclusive_range, "-")
		start, _ := strconv.ParseUint(pieces[0], 10, 64)
		end, _ := strconv.ParseUint(pieces[1], 10, 64)
		return InclusiveRange{start, end}
	})
	return inclusive_ranges
}

func checkIfValid(id uint64, part2 bool) uint64 {
	if part2 {
		return validCheckPart2(id)
	}
	return validCheckPart1(id)
}

func validCheckPart1(id uint64) uint64 {
	stringifiedId := strconv.FormatUint(id, 10)
	runeifiedID := []rune(stringifiedId)
	idLength := len(runeifiedID)
	halfLength := idLength / 2
	if idLength % 2 != 0 {
		return 0
	}
	
	firstHalf := string(runeifiedID[:halfLength])
	secondHalf := string(runeifiedID[halfLength:])

	if firstHalf == secondHalf {
		return id
	} else {
		return 0
	}
}

func validCheckPart2(id uint64) uint64 {
	stringifiedId := strconv.FormatUint(id, 10)
	runeifiedID := []rune(stringifiedId)
	idLength := len(runeifiedID)
	halfLength := idLength / 2
	
	if idLength >= 2 {
		firstChar := runeifiedID[0]
		allSame := lo.EveryBy(runeifiedID, func(digit rune) bool {
			return digit == firstChar
		})
		if allSame {
            return id;
        }
	}



	for chunkCount := 2; chunkCount <= halfLength; chunkCount++ {
		var charChunks []string
		for i := 0; i < len(runeifiedID); i += chunkCount {
			end := i + chunkCount
			if end > len(runeifiedID) {
				end = len(runeifiedID)
			}
			chunk := string(runeifiedID[i:end])
			charChunks = append(charChunks, chunk)
		}

		// Check if all chunks equal the first
		if len(charChunks) > 0 {
			firstChunk := charChunks[0]
			allSame := lo.EveryBy(charChunks[1:], func(chunk string) bool {
				return chunk == firstChunk
			})
			if allSame {
				return id
			}
		}
	}

	return 0;
}

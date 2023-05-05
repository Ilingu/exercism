package wordy

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var Operation = map[string]rune{"plus": '+', "minus": '_', "multiplied": '*', "divided": '/'}

func SearchAndExecOperation(normalizedCalculus string) (int, bool) {
	IntNumbers, CurrentNumber, RemainingOps := []int{}, "", normalizedCalculus
	AddToNums := func() error {
		IntNum, err := strconv.Atoi(CurrentNumber)
		if err != nil {
			return errors.New("")
		}

		IntNumbers = append(IntNumbers, IntNum)
		RemainingOps = strings.Replace(RemainingOps, CurrentNumber, "", 1)
		CurrentNumber = ""
		return nil
	}

	for _, ch := range normalizedCalculus {
		if unicode.IsDigit(ch) || ch == '-' {
			CurrentNumber += string(ch)
			continue
		}

		err := AddToNums()
		if err != nil {
			return 0, false
		}
	}
	err := AddToNums()
	if err != nil {
		return 0, false
	}

	TotalCalculus := IntNumbers[0]
	for i, ops := range RemainingOps {
		if i+1 > len(IntNumbers)-1 {
			break
		}

		switch ops {
		case '+':
			TotalCalculus += IntNumbers[i+1]
		case '_':
			TotalCalculus -= IntNumbers[i+1]
		case '*':
			TotalCalculus *= IntNumbers[i+1]
		case '/':
			TotalCalculus /= IntNumbers[i+1]
		default:
			return 0, false // Not a Operation
		}
	}

	return TotalCalculus, true
}

func Answer(question string) (int, bool) {
	InvalidCalculus := regexp.MustCompile("[0-9]+ +[0-9]+")
	if InvalidCalculus.MatchString(question) {
		return 0, false
	}

	// Normalize Question
	question = strings.ReplaceAll(strings.TrimSuffix(strings.TrimPrefix(question, "What is"), "?"), " ", "")
	question = strings.ReplaceAll(question, "by", "")
	for _, op := range []string{"plus", "minus", "multiplied", "divided"} {
		question = strings.ReplaceAll(question, op, string(Operation[op]))
	}

	if len(question) == 1 {
		num, err := strconv.Atoi(string(question[0]))
		if err != nil {
			return 0, false
		}

		return num, true
	}

	return SearchAndExecOperation(question)
}

// JustNumRef, QuestionReg := regexp.MustCompile(`(?mi)What is ([0-9]+)\?`), regexp.MustCompile(`(?mi)What is ((?:-)?[0-9]+) ([a-zA-Z]+)(?: by)? ((?:-)?[0-9]+)`)
// ((?:-)?[0-9]+ ((?:plus|minus|multiplied(?: by)?|divided(?: by)?) ((?:-)?[0-9]+)+)+)

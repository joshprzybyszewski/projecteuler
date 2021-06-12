package easy

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem17() string {
	/*
		If the numbers 1 to 5 are written out in words:
		one, two, three, four, five, then there
		are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

		If all the numbers from 1 to 1000 (one thousand) inclusive
		were written out in words, how many letters would be used?

		NOTE: Do not count spaces or hyphens. For example, 342 (three
			hundred and forty-two) contains 23 letters and 115 (one
			hundred and fifteen) contains 20 letters. The use of
			"and" when writing out numbers is in compliance with
			British usage.
	*/
	ans := totalLettersFromOneToAThousand()
	return fmt.Sprintf("%v", ans)
}

func totalLettersFromOneToAThousand() int {
	total := 0
	for i := 1; i <= 1000; i++ {
		total += numLettersInWord(i)
	}
	return total
}

func numLettersInWord(n int) int {
	return len(numberToWord(n, false))
}

func numberToWord(
	n int,
	pretty bool,
) string {
	if n > 1000 {
		panic(`unsupported number`)
	}

	var sb strings.Builder

	if n >= 100 {
		d := n / 100
		if d == 10 {
			if pretty {
				return `one thousand`
			}
			return `onethousand`
		}
		sb.WriteString(singleDigitToWord(d))
		if pretty {
			sb.WriteString(` hundred`)
		} else {
			sb.WriteString(`hundred`)
		}
		if n%100 > 0 {
			if pretty {
				sb.WriteString(` and `)
			} else {
				sb.WriteString(`and`)
			}
		}
	}

	doubleDigitToWord(n%100, &sb, pretty)

	return sb.String()
}

func singleDigitToWord(n int) string {
	switch n {
	case 1:
		return `one`
	case 2:
		return `two`
	case 3:
		return `three`
	case 4:
		return `four`
	case 5:
		return `five`
	case 6:
		return `six`
	case 7:
		return `seven`
	case 8:
		return `eight`
	case 9:
		return `nine`
	default:
		panic(`unsupported number: ` + strconv.Itoa(n))
	}
}

func doubleDigitToWord(
	n int,
	sb *strings.Builder,
	pretty bool,
) {
	if n >= 100 {
		panic(`unsupported number`)
	}

	if n >= 20 {
		switch n / 10 {
		case 2:
			sb.WriteString(`twenty`)
		case 3:
			sb.WriteString(`thirty`)
		case 4:
			sb.WriteString(`forty`)
		case 5:
			sb.WriteString(`fifty`)
		case 6:
			sb.WriteString(`sixty`)
		case 7:
			sb.WriteString(`seventy`)
		case 8:
			sb.WriteString(`eighty`)
		case 9:
			sb.WriteString(`ninety`)
		}
		if n%10 > 0 && pretty {
			sb.WriteString(`-`)
		}
	} else if n >= 10 {
		switch n {
		case 10:
			sb.WriteString(`ten`)
		case 11:
			sb.WriteString(`eleven`)
		case 12:
			sb.WriteString(`twelve`)
		case 13:
			sb.WriteString(`thirteen`)
		case 14:
			sb.WriteString(`fourteen`)
		case 15:
			sb.WriteString(`fifteen`)
		case 16:
			sb.WriteString(`sixteen`)
		case 17:
			sb.WriteString(`seventeen`)
		case 18:
			sb.WriteString(`eighteen`)
		case 19:
			sb.WriteString(`nineteen`)
		}
		return
	}

	if n <= 0 || n%10 == 0 {
		return
	}

	sb.WriteString(singleDigitToWord(n % 10))
}

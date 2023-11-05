package yandex

type romanSymbol struct {
	Int    int
	Symbol string
}

func IntToRoman(num int) string {
	res := ""

	arr := []romanSymbol{{
		Int:    1000,
		Symbol: "M",
	}, {
		Int:    900,
		Symbol: "CM",
	}, {
		Int:    500,
		Symbol: "D",
	}, {
		Int:    400,
		Symbol: "CD",
	}, {
		Int:    100,
		Symbol: "C",
	}, {
		Int:    90,
		Symbol: "XC",
	}, {
		Int:    50,
		Symbol: "L",
	}, {
		Int:    40,
		Symbol: "XL",
	}, {
		Int:    10,
		Symbol: "X",
	}, {
		Int:    9,
		Symbol: "IX",
	}, {
		Int:    5,
		Symbol: "V",
	}, {
		Int:    4,
		Symbol: "IV",
	}, {
		Int:    1,
		Symbol: "I",
	}}

	for i := range arr {
		dev := num / arr[i].Int
		if dev == 0 {
			continue
		}
		j := 0
		for j < dev {
			res += arr[i].Symbol
			j++
		}

		num = num % arr[i].Int
	}

	return res
}

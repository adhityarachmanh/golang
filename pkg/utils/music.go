package utils

import "strings"

func SongMapping(song string) []interface{} {
	keyMap := map[string]int{
		"1": 24,
		"!": 25,
		"2": 26,
		"@": 27,
		"3": 28,
		"4": 29,
		"$": 30,
		"5": 31,
		"%": 32,
		"6": 33,
		"^": 34,
		"7": 35,
		"8": 36,
		"*": 37,
		"9": 38,
		"(": 39,
		"0": 40,
		"q": 41,
		"Q": 42,
		"w": 43,
		"W": 44,
		"e": 45,
		"E": 46,
		"r": 47,
		"t": 48,
		"T": 49,
		"y": 50,
		"Y": 51,
		"u": 52,
		"i": 53,
		"I": 54,
		"o": 55,
		"O": 56,
		"p": 57,
		"P": 58,
		"a": 59,
		"s": 60,
		"S": 61,
		"d": 62,
		"D": 63,
		"f": 64,
		"g": 65,
		"G": 66,
		"h": 67,
		"H": 68,
		"j": 69,
		"J": 70,
		"k": 71,
		"l": 72,
		"L": 73,
		"z": 74,
		"Z": 75,
		"x": 76,
		"c": 77,
		"C": 78,
		"v": 79,
		"V": 80,
		"b": 81,
		"B": 82,
		"n": 83,
		"m": 84,
		"_": 0,
		"|": -1,
	}
	var newArrResult []interface{}

	var isArray bool = false
	var arrayTemp []int
	arrSplit := strings.Split(song, "")

	for i := 0; i < len(arrSplit); i++ {
		d := arrSplit[i]

		if d == "[" {
			isArray = true
			// return
		} else if d == "]" {
			isArray = false

			newArrResult = append(newArrResult, arrayTemp)
			arrayTemp = nil
			// return
		} else {
			if isArray {
				c := keyMap[d]
				arrayTemp = append(arrayTemp, c)
			} else {
				c := keyMap[d]
				newArrResult = append(newArrResult, c)
			}
		}

	}
	return newArrResult
}

package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
    Подготовка строки. Заменяем заглавные буквы на маленькие с ! знаком (M -> !m)
	Превращаем строку в бинарную последовательность. В массив нулей и едениц ('101001011010010111011010')
	Разбиваем массив символов на чанки по 8 едениц в каждом ('10100101 10100101 11011010'). P.s. Делим массив битов на байты
	Представляем полученные байты в виде шестнадцатеричных чисел
	Возвращаем последовательность полученных чисел в виде строки
*/

type encodingTable map[rune]string

type BinaryChunks []BinaryChunk
type BinaryChunk string

const chunksSize = 8

type HexChunk string
type HexChunks []HexChunk

func Encode(str string) string {
	str = prepareText(str)
	chunks := splitByChunks(encodeBin(str), chunksSize)
	return chunks.ToHex().ToString()
}

//
func (hcs HexChunks) ToString() string {
	const sep = " "

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(hc))
	}

	return buf.String()
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexchunk := chunk.ToHex()

		res = append(res, hexchunk)
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

// prepareText Заменяет заглавные буквы на маленькие с восклицательным знаком (M -> !m)
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknow character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// func splitByChunks(bStr string, chunkSize int) BinaryChunks {
// 	strLen := utf8.RuneCountInString(bStr)
// 	chunksCount := strLen / chunkSize

// 	if strLen/chunkSize != 0 {
// 		chunksCount++
// 	}

// 	res := make(BinaryChunks, 0, chunksCount)

// 	var buf strings.Builder

// 	for i, ch := range bStr {
// 		buf.WriteString(string(ch))
// 		if (i+1) % chunkSize == 0 {
// 			res = append(res, BinaryChunk(buf.String()))
// 			buf.Reset()
// 		}
// 	}

// 	if buf.Len() != 0 {
// 		lastChunk := buf.String()
// 		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

// 		res = append(res, BinaryChunk(lastChunk))
// 	}

// 	return res
// }

func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}
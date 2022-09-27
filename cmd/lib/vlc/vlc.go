package vlc

import (
	"strings"
	"unicode"
)

/*
Этапы исполнения команды (функции) pack:
   	1) Подготовка строки. Заменяем заглавные буквы на маленькие с ! знаком (M -> !m)
	2) Превращаем строку в бинарную последовательность. В массив нулей и едениц ('101001011010010111011010')
	3) Разбиваем массив символов на чанки по 8 едениц в каждом ('10100101 10100101 11011010'). P.s. Делим массив битов на байты
	4) Представляем полученные байты в виде шестнадцатеричных чисел
	5) Возвращаем последовательность полученных чисел в виде строки
*/

func Encode(str string) []byte {
	str = prepareText(str)                              // подготавливаем строки. M -> !m
	chunks := splitByChunks(encodeBin(str), chunksSize) // представляем строку в виде двоичной последовательности и разбиваем ее на части (chunks)
	return chunks.Bytes()                               // конвертируем чанки в шестнадцатиричную систему и приводим к строке
}

/*
Шестнадцатеричные чанки приводим к двоичной системе счисления
Соединяем получанные чанки в одну бинарную строку
Строим дерево декодирования
С помощью ДД переводим закодированную строку в исходный текст
*/
func Decode(encodedData []byte) string {
	bString := NewBinChunks(encodedData).Join()
	dTree := getEncodingTable().DecodingTree()
	return exporText(dTree.Decode(bString))
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

// exportText Выполняет обратный prepareText функционал. Если в строке !m -> M
func exporText(str string) string {
	var buf strings.Builder
	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false

			continue
		}

		if ch == '!' {
			isCapital = true

			continue
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

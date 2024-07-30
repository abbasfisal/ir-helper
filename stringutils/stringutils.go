package stringutils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func HelloWorld() {
	fmt.Println("سلام دنیا:)")
}

// GenerateStrongRandom generates a strong random string of given length.
// If includeAlpha is true, it includes alphabets in the string.
func GenerateStrongRandom(length int, includeAlpha bool) (string, error) {
	var codeAlphabet string
	if includeAlpha {
		codeAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	} else {
		codeAlphabet = "0123456789"
	}

	max := big.NewInt(int64(len(codeAlphabet)))
	token := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		token[i] = codeAlphabet[num.Int64()]
	}

	return string(token), nil
}

// GenerateRandomNumber generates a random number string of given length.
func GenerateRandomNumber(length int) (string, error) {
	return GenerateStrongRandom(length, false)
}

// ConvertPersianNumbersToEnglish converts Persian numbers in the input string to English numbers.
func ConvertPersianNumbersToEnglish(input string) string {
	persianNumbers := []rune("۰۱۲۳۴۵۶۷۸۹")
	englishNumbers := []rune("0123456789")

	return replaceNumbers(input, persianNumbers, englishNumbers)
}

// ConvertArabicNumbersToEnglish converts Arabic numbers in the input string to English numbers.
func ConvertArabicNumbersToEnglish(input string) string {
	arabicNumbers := []rune("۰۱۲٣٤٥٦٧٨٩")
	englishNumbers := []rune("0123456789")

	return replaceNumbers(input, arabicNumbers, englishNumbers)
}

// ConvertNumbersToEnglish converts both Persian and Arabic numbers in the input string to English numbers.
func ConvertNumbersToEnglish(input string) string {
	input = ConvertPersianNumbersToEnglish(input)
	input = ConvertArabicNumbersToEnglish(input)
	return input
}

func replaceNumbers(input string, oldNumbers, newNumbers []rune) string {
	output := []rune(input)
	for i, char := range output {
		for j, oldChar := range oldNumbers {
			if char == oldChar {
				output[i] = newNumbers[j]
				break
			}
		}
	}
	return string(output)
}

// StripEnter removes all newline characters from the input string.
func StripEnter(txt string) string {
	txt = strings.ReplaceAll(txt, "\r\n", "")
	txt = strings.ReplaceAll(txt, "\r", "")
	txt = strings.ReplaceAll(txt, "\n", "")
	return txt
}

// StripTags removes HTML tags from the input string.
func StripTags(txt string) string {
	txt = UnEscape(txt)
	return stripHTMLTags(txt)
}

// UnEscape decodes HTML entities in the input string.
func UnEscape(txt string) string {
	return htmlUnescape(txt)
}

// CorrectMobileNo normalizes an Iranian mobile number.
func CorrectMobileNo(no string) string {
	if len(no) == 0 {
		return no
	}
	no = strings.TrimSpace(no)
	no = strings.TrimPrefix(no, "+")
	no = strings.TrimPrefix(no, "00")
	no = strings.TrimPrefix(no, "0")
	if !strings.HasPrefix(no, "98") && len(no) <= 10 {
		no = "98" + no[len(no)-10:]
	}
	if strings.HasPrefix(no, "980") {
		no = "98" + no[3:]
	}
	return no
}

// GenerateShortId generates a short unique ID of 8 characters.
func GenerateShortId() (string, error) {
	return GenerateStrongRandom(8, true)
}

// IsRTL detects if the input string is mostly written in a Right-to-Left language.
func IsRTL(s string) bool {
	rtlCount := 0
	totalCount := 0
	for _, r := range s {
		if isRTLRune(r) {
			rtlCount++
		}
		totalCount++
	}
	return rtlCount > totalCount/2
}

func isRTLRune(r rune) bool {
	// This is a simplified check for RTL characters
	return r >= 0x590 && r <= 0x8FF // Hebrew, Arabic, and other RTL ranges
}

// Reverse reverses the input string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func stripHTMLTags(txt string) string {
	// Implement a basic HTML tag stripper
	var output strings.Builder
	inTag := false
	for _, char := range txt {
		if char == '<' {
			inTag = true
			continue
		}
		if char == '>' {
			inTag = false
			continue
		}
		if !inTag {
			output.WriteRune(char)
		}
	}
	return output.String()
}

func htmlUnescape(txt string) string {
	// Implement an HTML entity decoder
	replacer := strings.NewReplacer(
		"&amp;", "&",
		"&lt;", "<",
		"&gt;", ">",
		"&quot;", "\"",
		"&#39;", "'",
	)
	return replacer.Replace(txt)
}

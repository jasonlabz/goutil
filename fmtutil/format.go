package fmtutil

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode"

	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/strutil"
)

// data size
const (
	OneKByte = 1024
	OneMByte = 1024 * 1024
	OneGByte = 1024 * 1024 * 1024
)

// DataSize format bytes number friendly.
//
// Usage:
//
//	file, err := os.Open(path)
//	fl, err := file.Stat()
//	fmtSize := DataSize(fl.Size())
func DataSize(size uint64) string {
	return basefn.DataSize(size)
}

// SizeToString alias of the DataSize
func SizeToString(size uint64) string { return DataSize(size) }

// StringToByte alias of the ParseByte
func StringToByte(sizeStr string) uint64 { return ParseByte(sizeStr) }

// ParseByte converts size string like 1GB/1g or 12mb/12M into an unsigned integer number of bytes
func ParseByte(sizeStr string) uint64 {
	sizeStr = strings.TrimSpace(sizeStr)
	lastPos := len(sizeStr) - 1
	if lastPos < 1 {
		return 0
	}

	if sizeStr[lastPos] == 'b' || sizeStr[lastPos] == 'B' {
		// last second char is k,m,g
		lastSec := sizeStr[lastPos-1]
		if lastSec > 'A' {
			lastPos--
		}
	}

	multiplier := float64(1)
	switch unicode.ToLower(rune(sizeStr[lastPos])) {
	case 'k':
		multiplier = 1 << 10
		sizeStr = strings.TrimSpace(sizeStr[:lastPos])
	case 'm':
		multiplier = 1 << 20
		sizeStr = strings.TrimSpace(sizeStr[:lastPos])
	case 'g':
		multiplier = 1 << 30
		sizeStr = strings.TrimSpace(sizeStr[:lastPos])
	default: // b
		multiplier = 1
		sizeStr = strings.TrimSpace(sizeStr[:lastPos])
	}

	size, _ := strconv.ParseFloat(sizeStr, 64)
	if size < 0 {
		return 0
	}

	return uint64(size * multiplier)
}

// PrettyJSON get pretty Json string
func PrettyJSON(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// StringsToInts string slice to int slice.
// Deprecated: please use the arrutil.StringsToInts()
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

// ArgsWithSpaces it like Println, will add spaces for each argument
func ArgsWithSpaces(vs []any) (message string) {
	if ln := len(vs); ln == 0 {
		return ""
	} else if ln == 1 {
		return strutil.SafeString(vs[0])
	} else {
		bs := make([]byte, 0, ln*8)
		for i := range vs {
			if i > 0 { // add space
				bs = append(bs, ' ')
			}
			bs = byteutil.AppendAny(bs, vs[i])
		}
		return string(bs)
	}
}

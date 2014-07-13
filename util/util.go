package util

import (
	"os"
	"path/filepath"
	"unicode/utf8"
)

var txtExt = map[string]bool{
	".js":    true,
	".json":  true,
	".html":  true,
	".md":    true,
	".rst":   true,
	".php":   true,
	".conf":  true,
	".go":    true,
	".css":   true,
	".py":    true,
	".log":   true,
	".pl":    true,
	".cofee": true,
	".dart":  true,
	".sql":   true,
}

// Take from here
// https://code.google.com/p/go/source/browse/godoc/util/util.go?repo=tools

func IsTextFile(filename string) bool {

	if istxt, found := txtExt[filepath.Ext(filename)]; found {
		return istxt
	}

	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer f.Close()

	var buf [1024]byte
	n, err := f.Read(buf[0:])
	if err != nil {
		return false
	}

	return IsText(buf[0:n])

}

// IsText reports whether a significant prefix of s looks like correct UTF-8;
// that is, if it is likely that s is human-readable text.
func IsText(s []byte) bool {
	const max = 1024 // at least utf8.UTFMax
	if len(s) > max {
		s = s[0:max]
	}
	for i, c := range string(s) {
		if i+utf8.UTFMax > len(s) {
			// last char may be incomplete - ignore
			break
		}
		if c == 0xFFFD || c < ' ' && c != '\n' && c != '\t' && c != '\f' {
			// decoding error or control character - not a text file
			return false
		}
	}
	return true
}

// Package langmap contains a list of display names of well known languages.
//
// This is a Go port of github.com/mozilla/language-mapping-list
package langmap

// Name holds English and Native names of a language / locale.
type Name struct {
	Native  string
	English string
}

// NativeName returns the native name of a language / locale code. It returns
// empty string if the language / locale is unknown or, very unlikely, it
// has no native name.
func NativeName(loc string) string {
	if n, ok := Names[loc]; ok {
		return n.Native
	}
	return ""
}

// EnglishName returs the English name of a given language / locale code.
func EnglishName(loc string) string {
	if n, ok := Names[loc]; ok {
		return n.English
	}
	return ""
}

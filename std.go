package i18n

import (
	"io/fs"

	"github.com/gookit/ini/v2"
)

/************************************************************
 * default instance
 ************************************************************/

// default instance
var std = NewEmpty()

// Std get default i18n instance
func Std() *I18n { return std }

// Default get default i18n instance
func Default() *I18n { return std }

// Reset std instance
func Reset() { std = NewEmpty() }

// Init the default language instance
func Init(langDir, defLang string, languages map[string]string) *I18n {
	std.langDir = langDir
	std.languages = languages
	std.DefaultLang = defLang

	return std.Init()
}

// InitFS init the default language instance, loading language files from fsys (eg: embed.FS).
// langDir is a slash-separated path inside fsys.
func InitFS(fsys fs.FS, langDir, defLang string, languages map[string]string) *I18n {
	std.fsys = fsys
	return Init(langDir, defLang, languages)
}

// Config the default instance
func Config(fn func(l *I18n)) {
	fn(std)
}

// T translate language key to value string
func T(lang, key string, args ...interface{}) string {
	return std.T(lang, key, args...)
}

// Tr translate language key to value string
func Tr(lang, key string, args ...interface{}) string {
	return std.Tr(lang, key, args...)
}

// Dt translate language key from default language
func Dt(key string, args ...interface{}) string {
	return std.DefTr(key, args...)
}

// Dtr translate language key from default language
func Dtr(key string, args ...interface{}) string {
	return std.DefTr(key, args...)
}

// DefTr translate language key from default language
func DefTr(key string, args ...interface{}) string {
	return std.DefTr(key, args...)
}

// AddLang register and init new language. alias of NewLang()
func AddLang(lang string, name string) {
	std.NewLang(lang, name)
}

// LangData get language data instance
func LangData(lang string) *ini.Ini {
	return std.Lang(lang)
}

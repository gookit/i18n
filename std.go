package i18n

import "github.com/gookit/ini/v2"

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

// DefTr translate language key from default language
func DefTr(key string, args ...interface{}) string {
	return std.DefTr(key, args...)
}

// Init the default language instance
func Init(langDir, defLang string, languages map[string]string) *I18n {
	std.langDir = langDir
	std.languages = languages
	std.DefaultLang = defLang

	return std.Init()
}

// AddLang register and init new language. alias of NewLang()
func AddLang(lang string, name string) {
	std.NewLang(lang, name)
}

// LangData get language data instance
func LangData(lang string) *ini.Ini {
	return std.Lang(lang)
}

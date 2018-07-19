/*
Simple i18n manage, use INI format file

Source code and other details for the project are available at GitHub:

   https://github.com/gookit/i18n

lang files:

	conf/
	    lang/
	        en.ini
	        zh-CN.ini

init:

    import "github/gookit/i18n"

    languages := map[string]string{
        "en": "English",
        "zh-CN": "简体中文",
        "zh-TW": "繁体中文",
    }

    i18n.Init("conf/lang", "en", languages)

usage:

    // translate from special language
    val := i18n.Tr("en", "key")

    // translate from default language
    val := i18n.DefTr("key")

 */
package i18n

import (
	"fmt"
	"bytes"
	"log"
	"github.com/gookit/ini"
	"errors"
)

// I18n language manager
type I18n struct {
	// languages data
	data map[string]*ini.Ini

	// language files directory
	langDir string
	// language list {en:English, zh-CN:简体中文}
	languages map[string]string

	// default language name. eg. "en"
	DefaultLang string
	// spare(fallback) language name. eg. "en"
	FallbackLang string
}

/************************************************************
 * default instance
 ************************************************************/

// default instance
var defI18n = NewEmpty()

// DefI18n get default i18n instance
func DefI18n() *I18n {
	return defI18n
}

// Tr
func Tr(lang string, key string, args ...interface{}) string {
	return defI18n.Tr(lang, key, args...)
}

// DefTr
func DefTr(key string, args ...interface{}) string {
	return defI18n.DefTr(key, args...)
}

// Init
func Init(langDir string, defLang string, languages map[string]string) *I18n {
	defI18n.langDir = langDir
	defI18n.languages = languages

	defI18n.DefaultLang = defLang

	return defI18n.Init()
}

/************************************************************
 * create instance
 ************************************************************/

// New
func New(langDir string, defLang string, languages map[string]string) *I18n {
	return &I18n{
		data:    make(map[string]*ini.Ini, 0),
		langDir: langDir,

		languages: languages,

		DefaultLang: defLang,
	}
}

// NewEmpty
func NewEmpty() *I18n {
	return &I18n{
		data: make(map[string]*ini.Ini, 0),

		languages: make(map[string]string, 0),
	}
}

// NewEmpty
func NewWithInit(langDir string, defLang string, languages map[string]string) *I18n {
	m := New(langDir, defLang, languages)

	return m.Init()
}

/************************************************************
 * translate
 ************************************************************/

// DefTr translate from default lang
func (l *I18n) DefTr(key string, args ...interface{}) string {
	return l.Tr(l.DefaultLang, key, args...)
}

// Tr translate from a lang by key
// site.name => [site]
//  			name = my blog
func (l *I18n) Tr(lang, key string, args ...interface{}) string {
	if !l.HasLang(lang) {
		// find from fallback lang
		val := l.transFromFallback(key)
		if val == "" {
			return key
		}

		// if has args
		if len(args) > 0 {
			val = fmt.Sprintf(val, args...)
		}

		return val
	}

	val, ok := l.data[lang].Get(key)
	if !ok {
		// find from fallback lang
		val = l.transFromFallback(key)
		if val == "" {
			return key
		}
	}

	// if has args
	if len(args) > 0 {
		val = fmt.Sprintf(val, args...)
	}

	return val
}

// translate from fallback language
func (l *I18n) transFromFallback(key string) (val string) {
	fb := l.FallbackLang
	if !l.HasLang(fb) {
		return
	}

	return l.data[fb].MustString(key)
}

/************************************************************
 * data manage
 ************************************************************/

// Init load add language files
func (l *I18n) Init() *I18n {
	for lang := range l.languages {
		lData, err := ini.LoadFiles(l.langDir + "/" + lang + ".ini")
		if err != nil {
			log.Fatalf("Fail to load language: %s, error %s", lang, err.Error())
		}

		l.data[lang] = lData
	}

	return l
}

// Add
func (l *I18n) Add(lang string, name string) {
	l.NewLang(lang, name)
}

// NewLang create/add a new language
// usage:
// 	i18n.NewLang("zh-CN", "简体中文")
func (l *I18n) NewLang(lang string, name string) {
	l.data[lang] = ini.New()
	l.languages[lang] = name
}

// Add new language or append lang data
// usage:
// 	i18n.LoadFile("zh-CN", "path/to/zh-CN.ini")
func (l *I18n) LoadFile(lang string, file string) (err error) {
	// append data
	if ld, ok := l.data[lang]; ok {
		err = ld.LoadFiles(file)
		if err != nil {
			return
		}
	} else {
		err = errors.New("language" + lang + " not exist, please create it before load data")
	}

	return
}

// LoadString load language data form a string
// usage:
// i18n.Set("zh-CN", "name = blog")
func (l *I18n) LoadString(lang string, data string) (err error) {
	// append data
	if ld, ok := l.data[lang]; ok {
		err = ld.LoadStrings(data)
		if err != nil {
			return
		}
	} else {
		err = errors.New("language" + lang + " not exist, please create it before load data")
	}

	return
}

// Lang get language data instance
func (l *I18n) Lang(lang string) *ini.Ini {
	if _, ok := l.languages[lang]; ok {
		return l.data[lang]
	}

	return nil
}

// ToString
func (l *I18n) ToString(lang string) string {
	if _, ok := l.languages[lang]; !ok {
		return ""
	}

	var buf bytes.Buffer
	l.data[lang].WriteTo(&buf)

	return buf.String()
}

// HasLang
func (l *I18n) HasLang(lang string) bool {
	_, ok := l.languages[lang]
	return ok
}

// DelLang
func (l *I18n) DelLang(lang string) bool {
	_, ok := l.languages[lang]
	if ok {
		delete(l.data, lang)
		delete(l.languages, lang)
	}

	return ok
}

// Languages
func (l *I18n) Languages() map[string]string {
	return l.languages
}

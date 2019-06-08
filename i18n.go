/*
Package i18n is a simple language manager, use INI format file.

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

Usage:

    // translate from special language
    val := i18n.Tr("en", "key")

    // translate from default language
    val := i18n.DefTr("key")

*/
package i18n

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gookit/ini/v2"
)

// language file load mode
const (
	// language name is file name. "en" -> "lang/en.ini"
	FileMode uint8 = 0
	// language name is dir name, will load all file in the dir. "en" -> "lang/en/*.ini"
	DirMode uint8 = 1
)

// I18n language manager
type I18n struct {
	// languages data
	data map[string]*ini.Ini

	// language files directory
	langDir string
	// language list {en:English, zh-CN:简体中文}
	languages map[string]string
	// loaded lang files
	// loadedFiles []string

	// mode for the load language files. mode: 0 single file, 1 multi dir
	LoadMode uint8
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

// Default get default i18n instance
func Default() *I18n {
	return defI18n
}

// Tr translate language key to value string
func Tr(lang string, key string, args ...interface{}) string {
	return defI18n.Tr(lang, key, args...)
}

// DefTr translate language key from default language
func DefTr(key string, args ...interface{}) string {
	return defI18n.DefTr(key, args...)
}

// Init the default language instance
func Init(langDir string, defLang string, languages map[string]string) *I18n {
	defI18n.langDir = langDir
	defI18n.languages = languages

	defI18n.DefaultLang = defLang

	return defI18n.Init()
}

/************************************************************
 * create instance
 ************************************************************/

// New a i18n instance
func New(langDir string, defLang string, languages map[string]string) *I18n {
	return &I18n{
		data: make(map[string]*ini.Ini, 0),

		langDir:   langDir,
		languages: languages,

		DefaultLang: defLang,
	}
}

// NewEmpty a empty i18n instance
func NewEmpty() *I18n {
	return &I18n{
		data: make(map[string]*ini.Ini, 0),

		languages: make(map[string]string, 0),
	}
}

// NewWithInit a i18n instance and call init
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

		if len(args) > 0 { // if has args
			val = fmt.Sprintf(val, args...)
		}

		return val
	}

	val, ok := l.data[lang].GetValue(key)
	if !ok {
		// find from fallback lang
		val = l.transFromFallback(key)
		if val == "" {
			return key
		}
	}

	if len(args) > 0 { // if has args
		val = fmt.Sprintf(val, args...)
	}

	return val
}

// HasKey in the language data
func (l *I18n) HasKey(lang, key string) (ok bool) {
	if !l.HasLang(lang) {
		return
	}

	_, ok = l.data[lang].GetValue(key)
	return
}

// translate from fallback language
func (l *I18n) transFromFallback(key string) (val string) {
	fb := l.FallbackLang
	if !l.HasLang(fb) {
		return
	}

	return l.data[fb].String(key)
}

/************************************************************
 * data manage
 ************************************************************/

// Init load add language files
func (l *I18n) Init() *I18n {
	if l.LoadMode == FileMode {
		l.loadSingleFiles()
	} else if l.LoadMode == DirMode {
		l.loadDirFiles()
	} else {
		panic("invalid load mode setting. only allow 0, 1")
	}

	return l
}

// load language files when LoadMode is 0
func (l *I18n) loadSingleFiles() {
	pathSep := string(os.PathSeparator)

	for lang := range l.languages {
		lData := ini.New()
		err   := lData.LoadFiles(l.langDir + pathSep + lang + ".ini")
		if err != nil {
			panic("fail to load language: " + lang + ", error " + err.Error())
		}

		l.data[lang] = lData
	}
}

// load language files when LoadMode is 1
func (l *I18n) loadDirFiles() {
	pathSep := string(os.PathSeparator)

	for lang := range l.languages {
		dirPath := l.langDir + pathSep + lang
		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			panic("read dir fail: " + dirPath + ", error " + err.Error())
		}

		sl := l.data[lang]

		for _, fi := range files {
			// skip dir and filter the specified format
			if fi.IsDir() || !strings.HasSuffix(fi.Name(), ".ini") {
				continue
			}

			var err error
			if sl != nil {
				err = sl.LoadFiles(dirPath + pathSep + fi.Name())
			} else { // create new language instance
				sl = ini.New()
				err = sl.LoadFiles(dirPath + pathSep + fi.Name())
				l.data[lang] = sl
			}

			if err != nil {
				panic("fail to load language file: " + lang + ", error " + err.Error())
			}
		}
	}
}

// Add new language
func (l *I18n) Add(lang string, name string) {
	l.NewLang(lang, name)
}

// NewLang create/add a new language
// Usage:
// 	i18n.NewLang("zh-CN", "简体中文")
func (l *I18n) NewLang(lang string, name string) {
	// lang exist
	if _, ok := l.languages[lang]; ok {
		return
	}

	l.data[lang] = ini.New()
	l.languages[lang] = name
}

// LoadFile append data to a exist language
// Usage:
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
// Usage:
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

// Export a language data as INI string
func (l *I18n) Export(lang string) string {
	if _, ok := l.languages[lang]; !ok {
		return ""
	}

	var buf bytes.Buffer

	_, err := l.data[lang].WriteTo(&buf)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

// HasLang in the manager
func (l *I18n) HasLang(lang string) bool {
	_, ok := l.languages[lang]
	return ok
}

// DelLang from the i18n manager
func (l *I18n) DelLang(lang string) bool {
	_, ok := l.languages[lang]
	if ok {
		delete(l.data, lang)
		delete(l.languages, lang)
	}

	return ok
}

// Languages get all languages
func (l *I18n) Languages() map[string]string {
	return l.languages
}

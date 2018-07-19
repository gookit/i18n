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
)

// I18n
type I18n struct {
    // languages data
    data map[string]*ini.Ini
    // default language name. eg. "en"
    defLang string
    // spare(fallback) language name. eg. "en"
    // spareLang string
    // language files directory
    langDir string
    // language list {en:English, zh-CN:简体中文}
    languages map[string]string
}

/************************************/
/********** default instance ********/
/************************************/

// default instance
var defI18n = &I18n{data: make(map[string]*ini.Ini, 0)}

// Tr
func Tr(lang string, key string, args ...interface{}) string {
    return defI18n.Tr(lang, key, args...)
}

// DefTr
func DefTr(key string, args ...interface{}) string {
    return defI18n.DefTr(key, args...)
}

// Init
func Init(langDir string, defLang string, languages map[string]string)  {
    defI18n.langDir = langDir
    defI18n.defLang = defLang
    defI18n.languages = languages

    defI18n.Init()
}

// DefI18n get default i18n instance
func DefI18n() *I18n {
	return defI18n
}

/************************************/
/************   creator  ************/
/************************************/

// NewEmpty
func NewEmpty() *I18n {
    return &I18n{data: make(map[string]*ini.Ini, 0)}
}

// New
func New(langDir string, defLang string, languages map[string]string) *I18n {
    return &I18n{
        langDir:   langDir,
        defLang:   defLang,
        languages: languages,

        data: make(map[string]*ini.Ini, 0),
    }
}

/************************************/
/************ main logic ************/
/************************************/

// Init load add language files
func (l *I18n) Init() {
    for lang := range l.languages {
        lData, err := ini.LoadFiles(l.langDir + "/" + lang + ".ini")
        if err != nil {
            log.Fatalf("Fail to load language: %s, error %s", lang, err.Error())
        }

        l.data[lang] = lData
    }
}

// Add new language or append lang data
// usage:
// 	i18n.Add("zh-CN", "path/to/zh-CN.ini", "简体中文")
func (l *I18n) Add(lang string, file string, name string) {
    langFile := l.langDir + "/" + lang + ".ini"

    // append data
    if ld, ok := l.data[lang]; ok {
        err := ld.LoadFiles(langFile)
        if err != nil {
            log.Fatalf("Fail to load language: %s, error %s", lang, err.Error())
        }
    } else {
        l.Set(lang, file, name)
    }
}

// Set new language
// usage:
// i18n.Set("zh-CN", "path/to/zh-CN.ini", "简体中文")
func (l *I18n) Set(lang string, file string, name string) {
    langFile := l.langDir + "/" + lang + ".ini"

    lData, err := ini.LoadFiles(langFile)
    if err != nil {
        log.Fatalf("Fail to load language: %s, error %s", lang, err.Error())
    }

    l.data[lang] = lData
    l.languages[lang] = name
}

// AddData
func (l *I18n) AddData(lang string, dataSource map[string]ini.Section) {
    // append data
    if ld, ok := l.data[lang]; ok {
    	ld.LoadData(dataSource)
    }
}

// DefTr translate from default lang
func (l *I18n) DefTr(key string, args ...interface{}) string {
    return l.Tr(l.defLang, key, args...)
}

// Tr translate from a lang by key
// site.name => [site]
//  			name = my blog
func (l *I18n) Tr(lang string, key string, args ...interface{}) string {
    if !l.HasLang(lang) {
        return ""
    }

    val, ok := l.data[lang].Get(key)
	if !ok {
		return key
	}

    // if has args
    if len(args) > 0 {
        val = fmt.Sprintf(val, args...)
    }

    return val
}

// LangSource
func (l *I18n) LangSource(lang string) string {
    if _, ok := l.languages[lang]; ok {
        return ""
    }

    var buf bytes.Buffer
    l.data[lang].WriteTo(&buf)

    return buf.String()
}

// HasLang
func (l *I18n) HasLang(lang string) bool {
    if _, ok := l.languages[lang]; ok {
        return true
    }

    return false
}

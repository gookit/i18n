/**
Simple i18n implement by "gopkg.in/ini.v1"
Source code and other details for the project are available at GitHub:

   https://github.com/gookit/i18n

usage:

    import "github/gookit/i18n"

    languages := map[string]string{
        "en": "English",
        "zh-CN": "简体中文",
        "zh-TW": "繁体中文",
    }

    i18n.Init("conf/lang", "en", languages)

ok, now:

    // translate from special language
    val := i18n.Tr("en", "key")

    // translate from default language
    val := i18n.DefTr("key")

 */
package i18n

import (
    "gopkg.in/ini.v1"
    "fmt"
    "strings"
    "log"
)

// I18n
type I18n struct {
    // languages data
    data map[string]*ini.File
    // default language name. eg. "en"
    defLang string
    // spare language name. eg. "en"
    spareLang string
    // language files directory
    langDir string
    // language list {en:English, zh-CN:简体中文}
    languages map[string]string
}

/************************************/
/********** default instance ********/
/************************************/

// default instance
var defI18n = &I18n{data: make(map[string]*ini.File, 0)}

func Tr(lang string, key string, args ...interface{}) string {
    return defI18n.Tr(lang, key, args...)
}

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
    return &I18n{data: make(map[string]*ini.File, 0)}
}

// New
func New(langDir string, defLang string, languages map[string]string) *I18n {
    return &I18n{
        langDir:   langDir,
        defLang:   defLang,
        languages: languages,

        data: make(map[string]*ini.File, 0),
    }
}

/************************************/
/************ main logic ************/
/************************************/

// Init load add language files
func (l *I18n) Init() {
    for lang := range l.languages {
        lData, err := ini.Load(l.langDir + "/" + lang + ".ini")

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
        err := ld.Append(langFile)

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

    lData, err := ini.Load(langFile)

    if err != nil {
        log.Fatalf("Fail to load language: %s, error %s", lang, err.Error())
    }

    l.data[lang] = lData
    l.languages[lang] = name
}

// AddData
func (l *I18n) AddData(lang string, dataSource interface{}) {
    // append data
    if ld, ok := l.data[lang]; ok {
    	ld.Append(dataSource)
    }
}

// DefTr translate from default lang
func (l *I18n) DefTr(key string, args ...interface{}) string {
    return l.Tr(l.defLang, key, args...)
}

// Tr translate from a lang by key
func (l *I18n) Tr(lang string, key string, args ...interface{}) string {
    if !l.HasLang(lang) {
        return ""
    }

    var val string

    // site.name => [site]
    //  			name = my blog
    if strings.Contains(key, ".") {
        nodes := strings.SplitN(key, ".", 2)
        val = l.data[lang].Section(nodes[0]).Key(nodes[1]).String()
    } else {
        val = l.data[lang].Section("").Key(key).String()
    }

    // if has args
    if len(args) > 0 {
        val = fmt.Sprintf(val, args...)
    }

    return val
}

// HasLang
func (l *I18n) HasLang(lang string) bool {
    if _, ok := l.languages[lang]; ok {
        return true
    }

    return false
}

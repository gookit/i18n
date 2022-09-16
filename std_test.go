package i18n_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/i18n"
)

func TestDefault(t *testing.T) {
	is := assert.New(t)
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	defer i18n.Reset()
	i18n.Init("testdata", "en", languages)

	m := i18n.Default()
	is.IsType(new(i18n.I18n), m)

	is.Eq("Blog", i18n.T("en", "name"))

	is.Eq("Blog", i18n.Tr("en", "name"))
	is.Eq("Blog", i18n.Dt("name"))
	is.Eq("Blog", i18n.Dtr("name"))
	is.Eq("Blog", i18n.DefTr("name"))
	is.Eq("博客", i18n.T("zh-CN", "name"))
	is.Eq("博客", i18n.Tr("zh-CN", "name"))
}

func TestAddLang(t *testing.T) {
	defer i18n.Reset()

	lang := "custom"
	i18n.AddLang(lang, "")
	i18n.Config(func(l *i18n.I18n) {
		l.DefaultLang = "en"
	})

	assert.True(t, i18n.Default().HasLang(lang))
	assert.NotEmpty(t, i18n.Default().Languages())

	err := i18n.Std().SetValues(lang, "", map[string]string{
		"name": "inhere",
		"age":  "234",
	})
	assert.NoErr(t, err)

	assert.NotNil(t, i18n.LangData(lang))
	assert.Eq(t, "inhere", i18n.LangData(lang).String("name"))
	assert.Eq(t, 234, i18n.LangData(lang).Int("age"))

	assert.Eq(t, "inhere", i18n.Tr(lang, "name"))
	assert.Eq(t, "234", i18n.T(lang, "age"))
}

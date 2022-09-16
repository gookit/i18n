package i18n_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/i18n"
)

func TestAddLang(t *testing.T) {
	lang := "custom"
	i18n.AddLang(lang, "")

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

package i18n

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Example() {
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	l := New("testdata", "en", languages)
	l.Init()

	fmt.Printf("name %s\n", l.Tr("en", "name"))
	fmt.Printf("name %s\n", l.DefTr("name"))
	fmt.Printf("name %s\n", l.Tr("zh-CN", "name"))
	fmt.Printf("use args: %s\n", l.DefTr("argMsg", "inhere"))

	// Output:
	// name Blog
	// name Blog
	// name 博客
	// use args: hello inhere, welcome
}

func TestDefI18n(t *testing.T) {
	st := assert.New(t)
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	Init("testdata", "en", languages)

	st.Equal("Blog", DefI18n().Tr("en", "name"))

	st.Equal("Blog", Tr("en", "name"))
	st.Equal("Blog", DefTr("name"))
	st.Equal("博客", Tr("zh-CN", "name"))
}

func TestI18n(t *testing.T) {
	st := assert.New(t)

	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	m := NewWithInit("testdata", "en", languages)
	st.True(m.HasLang("zh-CN"))
	st.False(m.HasLang("zh-TW"))

	str := m.Tr("zh-TW", "key")
	st.Equal("key", str)

	st.True(m.HasKey("en", "onlyInEn"))
	st.False(m.HasKey("zh-CN", "onlyInEn"))

	ls := m.Languages()
	st.Equal("English", ls["en"])

	// use args
	str = m.DefTr("argMsg", "inhere")
	st.Contains(str, "inhere")

	// fallback lang
	m.FallbackLang = "en"
	str = m.Tr("zh-CN", "onlyInEn")
	st.Equal("val0", str)

	l := m.Lang("en")
	st.NotNil(l)
	st.Equal("Blog", l.MustString("name"))

	ok := m.DelLang("zh-CN")
	st.True(ok)
	st.False(m.HasLang("zh-CN"))
}

func TestI18n_NewLang(t *testing.T) {
	st := assert.New(t)

	l := NewEmpty()
	l.NewLang("en", "English")

	err := l.LoadFile("en", "testdata/en.ini")
	st.Nil(err)

	st.Equal("Blog", l.Tr("en", "name"))
	st.Equal("name", l.DefTr("name"))

	// set default lang
	l.DefaultLang = "en"

	st.Equal("Blog", l.DefTr("name"))
}

func TestI18n_Export(t *testing.T) {
	st := assert.New(t)

	m := NewEmpty()
	m.NewLang("en", "English")

	err := m.LoadString("en", "name = Blog")
	st.Nil(err)

	st.Contains(m.Export("en"), "name = Blog")
}

func TestI18n_LoadString(t *testing.T) {
	st := assert.New(t)

	m := NewEmpty()
	m.NewLang("en", "English")
	st.True(m.HasLang("en"))

	err := m.LoadString("en", "name = Blog")
	st.Nil(err)
	st.Equal("Blog", m.Tr("en", "name"))
}

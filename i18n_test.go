package i18n

// test cover details: https://gocover.io/github.com/gookit/i18n
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

func TestInstance(t *testing.T) {
	st := assert.New(t)
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	Init("testdata", "en", languages)

	m := Instance()
	st.IsType(new(I18n), m)

	st.Equal("Blog", Tr("en", "name"))

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
	st.Equal(SingleFile, m.LoadMode)

	str := m.Tr("zh-TW", "key")
	st.Equal("key", str)

	st.True(m.HasKey("en", "onlyInEn"))
	st.False(m.HasKey("zh-CN", "onlyInEn"))
	st.False(m.HasKey("no-lang", "key"))

	ls := m.Languages()
	st.Equal("English", ls["en"])

	// use args
	str = m.DefTr("argMsg", "inhere")
	st.Contains(str, "inhere")

	// fallback lang
	m.FallbackLang = "en"
	str = m.Tr("zh-CN", "onlyInEn")
	st.Equal("val0", str)

	str = m.Tr("zh-CN", "noKey")
	st.Equal("noKey", str)

	str = m.Tr("no-lang", "argMsg", "inhere")
	st.Contains(str, "inhere")

	str = m.Tr("no-lang", "no-key")
	st.Equal("no-key", str)

	// get lang
	l := m.Lang("no-lang")
	st.Nil(l)

	l = m.Lang("en")
	st.NotNil(l)
	st.Equal("Blog", l.MustString("name"))

	ok := m.DelLang("zh-CN")
	st.True(ok)
	st.False(m.HasLang("zh-CN"))

	st.Panics(func() {
		languages["not-exist"] = "not-Exist"
		NewWithInit("testdata", "en", languages)
	})
}

func TestMultiFile(t *testing.T) {
	st := assert.New(t)
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	m := New("testdata", "en", languages)
	// setting
	m.LoadMode = MultiFile
	m.Init()

	st.True(m.HasLang("zh-CN"))
	st.False(m.HasLang("zh-TW"))
	st.Equal(MultiFile, m.LoadMode)

	st.Equal("inhere", m.DefTr("name"))
	st.Equal("语言管理", m.Tr("zh-CN", "use-for"))

	st.Panics(func() {
		m := New("testdata", "en", languages)
		// setting invalid mode
		m.LoadMode = 3
		m.Init()
	})

	st.Panics(func() {
		// invalid lang
		languages["not-exist"] = "not-Exist"

		m := New("testdata", "en", languages)
		m.LoadMode = MultiFile
		m.Init()
	})
}

func TestI18n_NewLang(t *testing.T) {
	st := assert.New(t)

	l := NewEmpty()
	l.Add("en", "English")

	err := l.LoadFile("en", "testdata/en.ini")
	st.Nil(err)

	// invalid file path
	err = l.LoadFile("en", "not-exist.ini")
	st.Error(err)

	// invalid data string
	err = l.LoadString("en", "invalid string")
	st.Error(err)

	st.Equal("Blog", l.Tr("en", "name"))
	st.Equal("name", l.DefTr("name"))

	// not exist lang
	err = l.LoadFile("zh-CN", "testdata/zh-CN.ini")
	st.Error(err)
	err = l.LoadString("zh-CN", "name = 博客")
	st.Error(err)
	l.NewLang("zh-CN", "简体中文")
	err = l.LoadString("zh-CN", "name = 博客")
	st.Nil(err)

	// set default lang
	l.DefaultLang = "en"

	st.Equal("Blog", l.DefTr("name"))
}

func TestI18n_Export(t *testing.T) {
	st := assert.New(t)

	m := NewEmpty()
	m.Add("en", "English")
	// repeat
	m.NewLang("en", "English")

	err := m.LoadString("en", "name = Blog")
	st.Nil(err)

	st.Contains(m.Export("en"), "name = Blog")
	st.Equal("", m.Export("no-lang"))
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

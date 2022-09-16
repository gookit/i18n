package i18n

// test cover details: https://gocover.io/github.com/gookit/i18n
import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/ini/v2"
)

func Example() {
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	l := New("testdata", "en", languages)
	l.Init()

	fmt.Printf("name: %s\n", l.T("en", "name"))
	fmt.Printf("name: %s\n", l.Dt("name"))
	fmt.Printf("name: %s\n", l.Tr("zh-CN", "name"))
	fmt.Printf("use args: %s\n", l.DefTr("argMsg", "inhere"))

	// Output:
	// name: Blog
	// name: Blog
	// name: 博客
	// use args: hello inhere, welcome
}

func TestNewWithInit(t *testing.T) {
	is := assert.New(t)

	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	m := NewWithInit("testdata", "en", languages)
	is.True(m.HasLang("zh-CN"))
	is.False(m.HasLang("zh-TW"))
	is.Eq(FileMode, m.LoadMode)

	str := m.Tr("zh-TW", "key")
	is.Eq("key", str)

	is.True(m.HasKey("en", "onlyInEn"))
	is.False(m.HasKey("zh-CN", "onlyInEn"))
	is.False(m.HasKey("no-lang", "key"))

	ls := m.Languages()
	is.Eq("English", ls["en"])

	// use args
	str = m.DefTr("argMsg", "inhere")
	is.Contains(str, "inhere")

	// fallback lang
	m.FallbackLang = "en"
	str = m.Tr("zh-CN", "onlyInEn")
	is.Eq("val0", str)

	str = m.Tr("zh-CN", "noKey")
	is.Eq("noKey", str)

	str = m.Tr("no-lang", "argMsg", "inhere")
	is.Contains(str, "inhere")

	str = m.Tr("no-lang", "no-key")
	is.Eq("no-key", str)

	// get lang
	l := m.Lang("no-lang")
	is.Nil(l)

	l = m.Lang("en")
	is.NotNil(l)
	is.Eq("Blog", l.String("name"))

	ok := m.DelLang("zh-CN")
	is.True(ok)
	is.False(m.HasLang("zh-CN"))

	is.Panics(func() {
		languages["not-exist"] = "not-Exist"
		NewWithInit("testdata", "en", languages)
	})
}

func TestDirMode(t *testing.T) {
	is := assert.New(t)
	languages := map[string]string{
		"en":    "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	m := New("testdata", "en", languages)
	// setting
	m.LoadMode = DirMode
	m.Init()

	is.True(m.HasLang("zh-CN"))
	is.False(m.HasLang("zh-TW"))
	is.Eq(DirMode, m.LoadMode)

	is.Eq("inhere", m.Dt("name"))
	is.Eq("inhere", m.DefTr("name"))
	is.Eq("语言管理", m.Tr("zh-CN", "use-for"))

	fmt.Println(m.Lang("zh-CN").Data())

	is.Panics(func() {
		m := New("testdata", "en", languages)
		// setting invalid mode
		m.LoadMode = 3
		m.Init()
	})

	is.Panics(func() {
		// invalid lang
		languages["not-exist"] = "not-Exist"

		m := New("testdata", "en", languages)
		m.LoadMode = DirMode
		m.Init()
	})
}

func TestI18n_TransMode(t *testing.T) {
	is := assert.New(t)

	m := NewEmpty()
	m.TransMode = ReplaceMode
	m.Add("en", "English")

	err := m.LoadString("en", "desc = i am {name}, age is {age}")
	is.NoErr(err)

	is.Eq("i am tom, age is 22", m.Tr("en", "desc", "name", "tom", "age", 22))
	is.Eq("i am tom, age is 22", m.Tr("en", "desc", map[string]interface{}{
		"name": "tom",
		"age":  22,
	}))
	is.Eq(
		"i am tom, age is CANNOT-TO-STRING",
		m.Tr("en", "desc", map[string]interface{}{"name": "tom", "age": []int{2}}),
	)
}

func TestI18n_NewLang(t *testing.T) {
	is := assert.New(t)

	l := NewEmpty()
	l.AddLang("en", "English")

	err := l.LoadFile("en", "testdata/en.ini")
	is.Nil(err)

	// invalid file path
	err = l.LoadFile("en", "not-exiis.ini")
	is.Err(err)

	// invalid data string
	err = l.LoadString("en", "invalid string")
	is.Err(err)

	is.Eq("Blog", l.Tr("en", "name"))
	is.Eq("name", l.DefTr("name"))

	// not exist lang
	err = l.LoadFile("zh-CN", "testdata/zh-CN.ini")
	is.Err(err)
	err = l.LoadString("zh-CN", "name = 博客")
	is.Err(err)
	l.NewLang("zh-CN", "简体中文")
	err = l.LoadString("zh-CN", "name = 博客")
	is.Nil(err)

	// set default lang
	l.Config(func(l *I18n) {
		l.DefaultLang = "en"
	})
	is.Eq("Blog", l.Dtr("name"))
}

func TestI18n_Export(t *testing.T) {
	is := assert.New(t)

	m := NewEmpty()
	m.Add("en", "English")
	// repeat
	m.NewLang("en", "English")

	err := m.LoadString("en", "name = Blog")
	is.Nil(err)

	is.Contains(m.Export("en"), "name = Blog")
	is.Eq("", m.Export("no-lang"))
}

func TestI18n_SetValues(t *testing.T) {
	is := assert.New(t)
	l := NewEmpty().WithLang("en", "English")

	err := l.SetValues("notExist", "", nil)
	is.Err(err)

	err = l.SetValues("en", "", nil)
	is.NoErr(err)

	err = l.SetValues("en", "site", ini.Section{
		"name": "my site",
	})
	is.NoErr(err)
	is.Eq("my site", l.Tr("en", "site.name"))
}

func TestI18n_LoadString(t *testing.T) {
	is := assert.New(t)

	m := NewEmpty()
	m.NewLang("en", "English")
	is.True(m.HasLang("en"))

	err := m.LoadString("en", "name = Blog")
	is.Nil(err)
	is.Eq("Blog", m.Tr("en", "name"))

	err = m.LoadString("en", "name = Blog")
	is.Nil(err)
	is.Eq("Blog", m.Tr("en", "name"))
}

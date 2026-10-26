package i18n

import (
	"embed"
	"testing"
	"testing/fstest"

	"github.com/gookit/goutil/x/assert"
)

//go:embed testdata
var testdataFS embed.FS

var fsLanguages = map[string]string{
	"en":    "English",
	"zh-CN": "简体中文",
}

func TestNewWithFS_embed(t *testing.T) {
	t.Run("file mode", func(t *testing.T) {
		l := NewWithFS(testdataFS, "testdata", "en", fsLanguages).Init()
		assert.Eq(t, "Blog", l.Dt("name"))
		assert.Eq(t, "博客", l.Tr("zh-CN", "name"))
		assert.Eq(t, "hello inhere, welcome", l.DefTr("argMsg", "inhere"))
	})

	t.Run("dir mode", func(t *testing.T) {
		l := NewWithFS(testdataFS, "testdata", "en", fsLanguages)
		l.LoadMode = DirMode
		l.Init()
		// same values as the OS file system loader (TestDirMode)
		assert.Eq(t, "inhere", l.Dt("name"))
		assert.Eq(t, "语言管理", l.Tr("zh-CN", "use-for"))
	})

	t.Run("missing language panics", func(t *testing.T) {
		assert.Panics(t, func() {
			NewWithFS(testdataFS, "testdata", "en", map[string]string{"not-exist": "Not"}).Init()
		})
		assert.Panics(t, func() {
			l := NewWithFS(testdataFS, "testdata", "en", map[string]string{"not-exist": "Not"})
			l.LoadMode = DirMode
			l.Init()
		})
	})
}

func TestI18n_LoadFileFS(t *testing.T) {
	mfs := fstest.MapFS{
		"lang/en.ini":       {Data: []byte("name = Blog\n")},
		"lang/en-extra.ini": {Data: []byte("extra = more\n")},
	}
	l := NewWithFS(mfs, "lang", "en", map[string]string{"en": "English"}).Init()
	assert.Eq(t, "Blog", l.Dt("name"))
	assert.NoErr(t, l.LoadFileFS(mfs, "en", "lang/en-extra.ini"))
	assert.Eq(t, "more", l.Dt("extra"))

	assert.Err(t, l.LoadFileFS(mfs, "fr", "lang/en-extra.ini"))
	assert.Err(t, l.LoadFileFS(mfs, "en", "lang/missing.ini"))
}

func TestInitFS(t *testing.T) {
	defer Reset()
	InitFS(testdataFS, "testdata", "en", fsLanguages)
	assert.Eq(t, "Blog", Dt("name"))
	assert.Eq(t, "博客", Tr("zh-CN", "name"))
}

func TestWithFS_nilRestoresOSFiles(t *testing.T) {
	l := NewWithFS(testdataFS, "testdata", "en", fsLanguages).WithFS(nil).Init()
	assert.Eq(t, "Blog", l.Dt("name"))
}

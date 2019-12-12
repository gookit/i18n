# I18n

[![GoDoc](https://godoc.org/github.com/gookit/i18n?status.svg)](https://godoc.org/github.com/gookit/i18n)
[![Build Status](https://travis-ci.org/gookit/i18n.svg?branch=master)](https://travis-ci.org/gookit/i18n)
[![Coverage Status](https://coveralls.io/repos/github/gookit/i18n/badge.svg?branch=master)](https://coveralls.io/github/gookit/i18n?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/i18n)](https://goreportcard.com/report/github.com/gookit/i18n)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/i18n)](https://github.com/gookit/i18n)

Use `INI` files, simple i18n manager implement.

> **[中文说明](README.zh-CN.md)**

## Features

- Easy to use, supports loading multiple languages, multiple files
- Two data loading modes: single file `FileMode`, folder` DirMode`; default is folder mode
- Support to set the default language and fallback language; when the default language data is not found, it will automatically try to find the fallback language
- Support parameter replacement, there are two modes: `SprintfMode` replaces parameters via` fmt.Sprintf`, `ReplaceMode` uses` strings.Replacer`

## Install

```bash
go get github.com/gookit/i18n
```

## Godoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/i18n.v1)
- [godoc for github](https://godoc.org/github.com/gookit/i18n)

## Usage

```text
lang/
    en/
        default.ini
        ...
    zh-CN/
        default.ini
        ...
```

### Init i18n

```go
    import "github/gookit/i18n"

    languages := map[string]string{
        "en": "English",
        "zh-CN": "简体中文",
        // "zh-TW": "繁体中文",
    }

    // The default instance initialized directly here
    i18n.Init("conf/lang", "en", languages)
    
    // Create a custom new instance
    // i18n.New(langDir string, defLang string, languages)
    // i18n.NewEmpty()
```

### Translate

```go
    // translate from special language
    msg := i18n.Tr("en", "key")

    // translate from default language
    msg = i18n.DefTr("key")
    // with arguments. 
    msg = i18n.DefTr("key1", "arg1", "arg2")
```

### Parameter replacement mode

Use `SprintfMode`(**defaults**):

```ini
# en.ini
desc = I am %s, age is %d
```

Usage with parameters:

```go
msg := i18n.Tr("en", "desc", "name", "tom", "age", 22)
// Output: "I am tom, age is 22"
```

Use `ReplaceMode`:

```ini
# en.ini
desc = I am {name}, age is {age}
```

Usage with `map[string]interface{}` params:

```go
i18n.TransMode = i18n.ReplaceMode

msg := i18n.Tr("en", "desc", "desc", map[string]interface{}{
    "name": "tom",
    "age": 22,
})
// Output: "I am tom, age is 22"
```

## Tests

```bash
go test -cover
```

## Dep packages

- [gookit/ini](https://github.com/gookit/ini) is an INI config file/data manage implement

## License

**[MIT](LICENSE)**

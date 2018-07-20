# i18n

[![GoDoc](https://godoc.org/github.com/gookit/i18n?status.svg)](https://godoc.org/github.com/gookit/i18n)
[![Build Status](https://travis-ci.org/gookit/i18n.svg?branch=master)](https://travis-ci.org/gookit/i18n)
[![Coverage Status](https://coveralls.io/repos/github/gookit/i18n/badge.svg?branch=master)](https://coveralls.io/github/gookit/i18n?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/i18n)](https://goreportcard.com/report/github.com/gookit/i18n)

Use `ini` files, simple i18n manager implement.

- easy to use，can load multi language, multi files
- Support set default language, fallback language
- Support parameter replacement

> **[中文说明](README_cn.md)**

## Godoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/i18n.v1)
- [godoc for github](https://godoc.org/github.com/gookit/i18n)

## Usage

```text
conf/
    lang/
        en.ini
        zh-CN.ini
        ...
```

- init

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

- usage

```go
    // translate from special language
    msg := i18n.Tr("en", "key")

    // translate from default language
    msg = i18n.DefTr("key")
```

## dep packages

- [gookit/ini](https://github.com/gookit/ini) ini 解析管理

## License

**MIT**

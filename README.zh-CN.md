# I18n

[![GoDoc](https://godoc.org/github.com/gookit/i18n?status.svg)](https://godoc.org/github.com/gookit/i18n)
[![Build Status](https://travis-ci.org/gookit/i18n.svg?branch=master)](https://travis-ci.org/gookit/i18n)
[![Coverage Status](https://coveralls.io/repos/github/gookit/i18n/badge.svg?branch=master)](https://coveralls.io/github/gookit/i18n?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/i18n)](https://goreportcard.com/report/github.com/gookit/i18n)

使用INI文件实现的语言数据管理使用。

> **[EN README](README.md)**

## 功能简介

- 使用简单，可加载多个语言，多个文件
- 两种模式：单文件、文件夹，默认是文件夹模式
- 支持设置默认语言，备用语言
- 支持参数替换

## Godoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/i18n.v1)
- [godoc for github](https://godoc.org/github.com/gookit/i18n)

## 快速使用

```text
lang/
    en/
        default.ini
        ...
    zh-CN/
        default.ini
        ...
```

### 初始化

```go
    import "github/gookit/i18n"

    languages := map[string]string{
        "en": "English",
        "zh-CN": "简体中文",
        // "zh-TW": "繁体中文",
    }

    // 这里直接初始化的默认实例
    i18n.Init("conf/lang", "en", languages)

    // 或者创建自定义的新实例
    // i18n.New(langDir string, defLang string, languages)
    // i18n.NewEmpty()
```

### 翻译数据

```go
    // 从指定的语言翻译
    msg := i18n.Tr("en", "key")

    // 从默认语言翻译
    msg = i18n.DefTr("key")
    // with arguments. 
    msg = i18n.DefTr("key1", "arg1", "arg2")
```

## 测试

```bash
go test -cover
```

## 依赖包

- [gookit/ini](https://github.com/gookit/ini) 功能强大的 INI 解析管理

## License

**[MIT](LICENSE)**
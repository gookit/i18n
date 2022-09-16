# I18n

[![Actions Status](https://github.com/gookit/i18n/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/i18n/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/i18n/badge.svg?branch=master)](https://coveralls.io/github/gookit/i18n?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/i18n)](https://goreportcard.com/report/github.com/gookit/i18n)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/i18n)](https://github.com/gookit/i18n)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/i18n?style=flat-square)
[![Go Reference](https://pkg.go.dev/badge/github.com/gookit/i18n.svg)](https://pkg.go.dev/github.com/gookit/i18n)

Management and use of multilingual data using `INI` files.

> **[中文说明](README.zh-CN.md)**

## Features

- Easy to use, supports loading multiple languages, multiple files
- Two data loading modes: single file `FileMode`(default), folder `DirMode`
- Support to set the default language and fallback language
  - when the default language data is not found, it will automatically try to find the fallback language
- Support parameter replacement, there are two modes
  - `SprintfMode` replaces parameters via `fmt.Sprintf`
  - `ReplaceMode` uses func `strings.Replacer`

## Install

```shell
go get github.com/gookit/i18n
```

## Godoc

- [go doc](https://pkg.go.dev/github.com/gookit/i18n)

## Usage

**Structs on use single `FileMode` mode**:

```text
lang/
    en.ini
    ru.ini
    zh-CN.ini
    zh-TW.ini
    ... ...
```

**Structs on use folder `DirMode` mode**:

```text
lang/
    en/
        default.ini
        ... ...
    zh-CN/
        default.ini
        ... ...
    zh-TW/
        default.ini
        ... ...
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

### Translate message

```go
// Translate from default language
msg = i18n.Dt("key")
// with arguments. 
msg = i18n.DefTr("key1", "arg1", "arg2")

// Translate from the specified language
msg := i18n.Tr("en", "key")
```

**Function list**:

```go
// Translate from default language
func Dt(key string, args ...interface{}) string
func Dtr(key string, args ...interface{}) string
func DefTr(key string, args ...interface{}) string

// Translate from the specified language
func T(lang, key string, args ...interface{}) string
func Tr(lang, key string, args ...interface{}) string
```

## Parameters replacement mode

### Use sprintf mode

> TIP: default mode is `SprintfMode`

```ini
# en.ini
desc = I am %s, age is %d
```

Usage with parameters like sprintf:

```go
msg := i18n.Tr("en", "desc", "tom", 22)
// Output: "I am tom, age is 22"
```

### Use replace mode

Enable replace mode:

```go
// set mode
i18n.Std().TransMode = i18n.ReplaceMode

// OR
i18n.Config(func(l *i18n.I18n) {
    l.TransMode = i18n.ReplaceMode
})
```

Examples for language data:

```ini
# en.ini
desc = I am {name}, age is {age}
```

**Usage with parameters**:

```go
// args is {"name": "tom", "age": 22}
msg := i18n.Tr("en", "desc", "name", "tom", "age", 22)
// Output: "I am tom, age is 22"
```

**Usage with kv-map parameters**:

```go
msg := i18n.Tr("en", "desc", map[string]interface{}{
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

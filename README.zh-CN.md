# I18n

[![Actions Status](https://github.com/gookit/i18n/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/i18n/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/i18n/badge.svg?branch=master)](https://coveralls.io/github/gookit/i18n?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/i18n)](https://goreportcard.com/report/github.com/gookit/i18n)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/i18n)](https://github.com/gookit/i18n)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/i18n?style=flat-square)
[![Go Reference](https://pkg.go.dev/badge/github.com/gookit/i18n.svg)](https://pkg.go.dev/github.com/gookit/i18n)

使用 `INI` 文件实现的多语言数据的管理和使用。

> **[EN README](README.md)**

## 功能简介

- 使用简单，支持加载多个语言，多个配置文件
- 两种数据加载模式：单文件 `FileMode`(默认) 、文件夹 `DirMode`
- 支持设置默认语言，备用语言；当在默认语言数据没找到时，自动尝试到备用语言查找
- 支持参数替换，有两种模式
  - `SprintfMode` 通过 `fmt.Sprintf` 替换参数
  - `ReplaceMode` 则使用 `strings.Replacer` 替换

## 安装

```bash
go get github.com/gookit/i18n
```

## 文档

- [Go doc](https://pkg.go.dev/github.com/gookit/i18n)

## 开始使用

**使用 `FileMode` 模式的语言文件结构**:

```text
lang/
    en.ini
    ru.ini
    zh-CN.ini
    zh-TW.ini
    ... ...
```

**使用 `DirMode` 模式的语言文件结构**:

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

### 初始化

```go
import "github/gookit/i18n"

defaultLang := "en"
languages := map[string]string{
    "en": "English",
    "zh-CN": "简体中文",
    // "zh-TW": "繁体中文",
}

// 这里直接初始化的默认实例
i18n.Init("conf/lang", defaultLang, languages)
```

**或者创建自定义的新实例**

```go
myI18n := i18n.New(langDir string, defLang string, languages)

myI18n := i18n.NewEmpty()
```

## 翻译数据

```go
// 从默认语言翻译
msg = i18n.Dtr("key")
// with arguments. 
msg = i18n.DefTr("key1", "arg1", "arg2")

// 从指定的语言翻译
msg := i18n.Tr("en", "key")
```

**方法列表**:

```go
// 从默认语言翻译
func Dt(key string, args ...interface{}) string
func Dtr(key string, args ...interface{}) string
func DefTr(key string, args ...interface{}) string

// 从指定的语言翻译
func T(lang, key string, args ...interface{}) string
func Tr(lang, key string, args ...interface{}) string
```

## 参数替换模式

### 使用 `SprintfMode` 模式

默认就是 `SprintfMode` 模式, 内部使用 `fmt.Spritf()` 进行参数的替换处理

```ini
# en.ini
desc = I am %s, age is %d
```

**按顺序传入参数使用**：

```go
msg := i18n.Tr("en", "desc", "tom", 22)
// Output: "I am tom, age is 22"
```

### 使用 `ReplaceMode` 模式

使用 `ReplaceMode` 模式, 内部使用 `strings.Replacer` 进行参数的替换处理

**启用 `ReplaceMode` 替换模式**:

```go
// set mode
i18n.Config(func(l *i18n.I18n) {
    l.TransMode = i18n.ReplaceMode
})

// OR
i18n.Std().TransMode = i18n.ReplaceMode
```

语言配置示例:

```ini
# en.ini
desc = I am {name}, age is {age}
```

**按kv顺序传入参数使用**：

```go
msg := i18n.Tr("en", "desc", "name", "tom", "age", 22)
// Output: "I am tom, age is 22"
```

**传入 kv-map 参数使用**：

```go
msg := i18n.Tr("en", "desc", map[string]interface{}{
    "name": "tom",
    "age": 22,
})
// Output: "I am tom, age is 22"
```

## 测试

```bash
go test -cover
```

## 依赖包

- [gookit/ini](https://github.com/gookit/ini) 功能强大的 INI 解析管理

## License

**[MIT](LICENSE)**

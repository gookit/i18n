package main

import (
	"github.com/gookit/i18n"
	"fmt"
)

// go run ./examples/demo.go
func main()  {
	languages := map[string]string{
		"en": "English",
		"zh-CN": "简体中文",
		// "zh-TW": "繁体中文",
	}

	i18n.Init("testdata", "en", languages)

	fmt.Printf("name %s\n", i18n.Tr("en", "name"))
	fmt.Printf("name %s\n", i18n.Tr("zh-CN", "name"))

	// Output:
	// name Blog
	// name 博客
}

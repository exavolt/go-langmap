# langmap

A Go port of [Mozilla's `langmap` JavaScript package](https://github.com/mozilla/language-mapping-list).

Please note that there's already package for getting display names of a
language which is located at `golang.org/x/text/language/display`. It's
maintained by Go team. You should give it a look before using this package.

Use this package if you want or need consistency with Mozilla's language
names.

## Usage

```
package main

import (
	"fmt"
	"github.com/exavolt/go-langmap"
)

func main() {
	fmt.Printf("Native en-US => %s\n", langmap.NativeName("en-US"))
	fmt.Printf("Native th => %s\n", langmap.NativeName("th"))
	fmt.Printf("English th => %s\n", langmap.EnglishName("th"))
}
```

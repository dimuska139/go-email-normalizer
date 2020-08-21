# go-email-normalizer

This is Golang library for providing a canonical representation of email address. It allows
to prevent multiple signups. `go-email-normalizer` contains some popular providers but you can easy append others.

[![Build Status](https://travis-ci.org/dimuska139/go-email-normalizer.svg?branch=master)](https://travis-ci.org/dimuska139/go-email-normalizer)
[![codecov](https://codecov.io/gh/dimuska139/go-email-normalizer/branch/master/graph/badge.svg)](https://codecov.io/gh/dimuska139/go-email-normalizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/go-email-normalizer)](https://goreportcard.com/report/github.com/dimuska139/go-email-normalizer)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/go-email-normalizer/blob/master/LICENSE)

### Download

```shell
go get -u github.com/dimuska139/go-email-normalizer
```

## Example

```go
package main

import (
	"fmt"
	"strings"
    normalizer "github.com/dimuska139/go-email-normalizer"
)

type zohoRule struct {}

func (rule *zohoRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *zohoRule) processDomain(domain string) string {
	return domain
}


func main() {
	n := normalizer.NewNormalizer()
	fmt.Println(n.Normalize("test@gm+.ail.com"))
	
	n.AddRule("zoho.com", &zohoRule{})
	fmt.Println(n.Normalize("test@zo+ho.com"))
}
```


## Supported providers

* Google
* Fastmail
* Microsoft
* Rambler

Also you can integrate another rules using `AddRule` function (see example above)
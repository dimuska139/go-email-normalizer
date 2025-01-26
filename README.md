# go-email-normalizer - email normalization for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/dimuska139/go-email-normalizer.svg)](https://pkg.go.dev/github.com/dimuska139/tilda-go)
[![codecov](https://codecov.io/github/dimuska139/go-email-normalizer/graph/badge.svg?token=E1TagDCBXc)](https://codecov.io/gh/dimuska139/go-email-normalizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/go-email-normalizer)](https://goreportcard.com/report/github.com/dimuska139/go-email-normalizer)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/go-email-normalizer/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

This is Golang library for providing a canonical representation of email address. It allows
to prevent multiple signups. `go-email-normalizer` contains some popular providers but you can easily append others.

## Download

```shell
go get -u github.com/dimuska139/go-email-normalizer/v5
```

## Usage

```go
package main

import (
	"fmt"
	"strings"
	normalizer "github.com/dimuska139/go-email-normalizer/v5"
)

type customRule struct {}

func (rule *customRule) ProcessUsername(username string) string {
	return strings.Replace(username, "-", "", -1)
}

func (rule *customRule) ProcessDomain(domain string) string {
	return domain
}

func main() {
	n := normalizer.NewNormalizer()
	fmt.Println(n.Normalize("vasya+pupkin@gmail.com")) // vasya@gmail.com
	fmt.Println(n.Normalize("t.e-St+vasya@gmail.com")) // te-st@gmail.com
	fmt.Println(n.Normalize("John+Brown@yahoo.com"))   // john+brown@yahoo.com
	fmt.Println(n.Normalize("John-Brown@yahoo.com"))   // john@yahoo.com
	fmt.Println(n.Normalize("t.e-St+@googlemail.com")) // te-st@gmail.com
	fmt.Println(n.Normalize("t.e-St+@google.com"))     // te-st@gmail.com
	
	n.AddRule("customrules.com", &customRule{})
	fmt.Println(n.Normalize(" tE-S-t@CustomRules.com.")) // tESt@customrules.com
}
```


## Supported providers

* Apple
* Fastmail
* Google
* Microsoft
* Protonmail
* Rackspace
* Rambler
* Yahoo
* Yandex
* Zoho

Also you can integrate other rules using `AddRule` function (see an example above)

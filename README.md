# go-email-normalizer - email normalization for Go

[![Build Status](https://travis-ci.org/dimuska139/go-email-normalizer.svg?branch=master)](https://travis-ci.org/dimuska139/go-email-normalizer)
[![codecov](https://codecov.io/gh/dimuska139/go-email-normalizer/branch/master/graph/badge.svg)](https://codecov.io/gh/dimuska139/go-email-normalizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/go-email-normalizer)](https://goreportcard.com/report/github.com/dimuska139/go-email-normalizer)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/go-email-normalizer/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

This is Golang library for providing a canonical representation of email address. It allows
to prevent multiple signups. `go-email-normalizer` contains some popular providers but you can easy append others.

## Download

```shell
go get github.com/dimuska139/go-email-normalizer
```

## Usage

```go
package main

import (
	"fmt"
	"strings"
	normalizer "github.com/dimuska139/go-email-normalizer"
)

type customRule struct {}

func (rule *customRule) processUsername(username string) string {
	return strings.Replace(username, "-", "", -1)
}

func (rule *customRule) processDomain(domain string) string {
	return domain
}

func main() {
	n := normalizer.NewNormalizer()
	fmt.Println(n.Normalize("vasya+pupkin@gmail.com")) // vasya@gmail.com
	fmt.Println(n.Normalize("t.e-St+@gmail.com")) // te-st@gmail.com
	fmt.Println(n.Normalize("t.e-St+@googlemail.com")) // te-st@gmail.com
	fmt.Println(n.Normalize("t.e-St+@google.com")) // te-st@gmail.com
	
	n.AddRule("customrules.com", &customRule{})
	fmt.Println(n.Normalize(" te-s-t@CustomRules.com.")) // test@customrules.com
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

Also you can integrate another rules using `AddRule` function (see an example above)

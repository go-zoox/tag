# Tag - Struct Tag Parser And Decoder

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/tag)](https://pkg.go.dev/github.com/go-zoox/tag)
[![Build Status](https://github.com/go-zoox/tag/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/tag/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/tag)](https://goreportcard.com/report/github.com/go-zoox/tag)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/tag/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/tag?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/tag.svg)](https://github.com/go-zoox/tag/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/tag.svg?label=Release)](https://github.com/go-zoox/tag/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/tag
```

## Features
* [x] Type Validation
  * [x] `omitempty`, such as `tag:"app_name,omitempty"`
  * [x] `required`, such as `tag:"app_name,required"`
  * [x] `default`, such as `tag:"app_name,default=my_app"`
  * [x] `enum`, such as `tag:"app_name,enum=my_app|my_app2"`
  * [x] `regexp`, such as `tag:"app_name,regexp=/^[a-zA-Z0-9_]+$/`
  * [x] `min`, such as `tag:"app_name,min=1`
    * if type is `string`, means the length of string
    * if type is `int64`, means the minimum value of int
  * [x] `max`, such as `tag:"app_name,max=10`
    * if type is `string`, means the length of string
    * if type is `int64`, means the maximum value of int
* [x] Auto Type Transform


## Getting Started

```go
// see test cases in tag_test.go
```

## License
GoZoox is released under the [MIT License](./LICENSE).
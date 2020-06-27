# Godash - A port of [lodash](http://lodash.com/) utility functions for Go language

[![Build Status](https://travis-ci.org/iamtabrezkhan/godash.svg?branch=master)](https://travis-ci.org/iamtabrezkhan/godash)
[![Coverage Status](https://coveralls.io/repos/github/iamtabrezkhan/godash/badge.svg?branch=master)](https://coveralls.io/github/iamtabrezkhan/godash?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamtabrezkhan/godash)](https://goreportcard.com/report/github.com/iamtabrezkhan/godash)

This library contains different packages for different collection types that you can import. For instance, to import functions related to slices you can do something like:

```go
import "github.com/iamtabrezkhan/godash/slice"
```

> Note: This library heavily relies on reflection, and it checks for types on runtime, so make sure you have good test suites. Although, for most of the known types we are avoiding the use of reflect and only using when there are unknown custom types.

## Install

```sh
go get github.com/iamtabrezkhan/godash
```

## Usage

```go
// import functions related to slice
import "github.com/iamtabrezkhan/godash/slice"

slice.Chunk([]int{1, 2, 3, 4, 5}, 2) // ==> [[1, 2], [3, 4], [5]]
```

## Packages

- [slice](slice/README.md) [![Documentation](https://godoc.org/github.com/iamtabrezkhan/godash?status.svg)](http://godoc.org/github.com/iamtabrezkhan/godash/slice)
- more soon...

## Author

👤 **Tabrez Khan**

- Website: https://iamtabrezkhan.netlify.app
- Twitter: [@TabrezX](https://twitter.com/TabrezX)
- Github: [@iamtabrezkhan](https://github.com/iamtabrezkhan)
- LinkedIn: [@iamtabrezkhan](https://linkedin.com/in/iamtabrezkhan)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/iamtabrezkhan/godash/issues). You can also take a look at the [contributing guide](https://github.com/iamtabrezkhan/godash/blob/master/CONTRIBUTING.md).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Tabrez Khan](https://github.com/iamtabrezkhan).<br />
This project is [MIT](LICENSE) licensed.

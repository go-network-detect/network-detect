# Network Detector for Go

[![GoDoc](https://godoc.org/github.com/go-network-detect/network-detect?status.svg)](https://godoc.org/github.com/go-network-detect/network-detect)

This package performs crude network detection by examining the
state of the machine on which it is executing (as opposed to
attempting to access the network). As such the detection should be
both very fast in general and run in approximately the same time
regardless of whether the network is up.

Currently only Linux operating systems are supported.

## Documentation

Documentation is available on [godoc.org](https://godoc.org/github.com/go-network-detect/network-detect).

## License

Released under the MIT license.

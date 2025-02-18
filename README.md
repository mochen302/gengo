# gengo

[![Travis Widget]][Travis] [![GoDoc Widget]][GoDoc]  [![GoReport]][GoReportStatus]

[Travis]: https://travis-ci.org/kubernetes/gengo
[Travis Widget]: https://travis-ci.org/kubernetes/gengo.svg?branch=master
[GoDoc]: https://godoc.org/github.com/mochen302/gengox
[GoDoc Widget]: https://godoc.org/github.com/mochen302/gengox?status.svg
[GoReport]: https://goreportcard.com/badge/github.com/kubernetes/gengo
[GoReportStatus]: https://goreportcard.com/report/github.com/kubernetes/gengo

A package for generating things based on go files. This mechanism was first used
in [Kubernetes code-generator](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/code-generator) and is split out here for ease of reuse and maintainability.

`go get github.com/mochen302/gengox/...`

## Examples

A set generator, deep-copy generator, defaulter generator and go-to-protobuf
generator are included here. Also, import-boss will enforce arbitrary rules about
import trees.

## args/

Package args defines common arguments for a generator binary.

## generator/

Package generator defines interfaces for code generators to implement, and
machinery that will execute those code generators.

## types/

Package types contains the type system definition. It is modeled after Go's type
system, but it's intended that you could produce these types by parsing
something else, if you want to write the parser/converter.

We don't directly use the go types in the go typecheck library because they are
based on implementing differing interfaces. A struct-based format is more
convenient input for template driven output.

## parser/

Package parser parses go source files.

## namer/

Package namer defines a naming system, for:
* helping you reference go objects in a syntactically correct way
* keeping track of what you reference, for importing the right packages
* and defining parallel tracks of names, for making public interfaces and
  private implementations.

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.

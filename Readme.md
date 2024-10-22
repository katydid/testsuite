# Cross Language Test Suite for Katydid

[![Build Status](https://github.com/katydid/testsuite/actions/workflows/build.yml/badge.svg)](https://github.com/katydid/testsuite/actions)

The test suite is a language agnostic test suite.

The idea is that katydid can be implemented in multiple programming languages.
Currently an implementations are available in:

  - [Go](https://github.com/katydid/validator-go) and
  - [Haskell](https://github.com/katydid/katydid-haskell).

Having one set of tests that can be used by multiple implementations helps to keep these implementations consistent
and is a great starting point when creating a new implementation.

Technically each of the two implementations consist of three implementations of validator's core algorithm,
so currently there are six users of the test suite.

The test suite is simply a bunch of folders and files that can be traversed and opened in any programming language.

The test suite contains some Go code to help to generate tests for multiple serialization formats.
The output is just a bunch of files and folders.
There should be no need to run this, except when adding more tests,
since all the output files and folders are already committed in the repository.
The reason for generating files and folders is that katydid is encoding agnostic,
so to help make sure that the parsers are implemented consistently,
we might want to use the same validator test for multiple serialization formats.

> This Go code for test suite generation should not be confused with the Go tests for the Go implementations of the validator.
The Go implementations of the validator all expect the test suite to located adjacent to katydid in the `GOPATH` and the files are read like in any other language's implementation, without using any Go code from the test suite repository.

## Tests

The validator tests are located in the `./validator/tests` folder.

Tests are grouped by codec: json, xml, etc.

Inside each codec folder is a list of testcase folders, each for a descriptive name.

Files in the codec folder are schemas that might be required by a testcase.
For example, marshaled protocol buffer file descriptor sets.

Inside each testcase folder is the validator grammar in several formats: txt, json and xml.
This for cases where a new language has not had time to fully implement a validator syntax parser.

Also found in the testcase folder is a filename starting with `valid` or `invalid` depending on whether the contents of the file is valid with respect to the provide grammar of invalid.
The file extension of this file again describes the codec of the contents.
In the case where this codec requires a schema, this filename is also provided as a suffix.

## Benchmarks

The validator benchmarks are located in the `./validator/benches` folder.

This folder is not checked in, because of its size.
Instead this folder can be generated, by running `make regenerate`.
This will require `go` to be installed and this folder to checked out to
`$GOPATH/src/github.com/katydid/testsuite`.

Benchmarks are grouped by codec: pbnum (no others are provided yet)

Inside each codec folder is a list of benchcase folders, each for a descriptive name.

Files in the codec folder are schemas that might be required by a benchcase.
For example, marshaled protocol buffer file descriptor sets.

Inside each benchcase folder is the validator grammar in several formats: txt, json and xml.
This for cases where a new language has not had time to fully implement a validator syntax parser.

Also found in the benchcase folder is a 1000 generated files each starting with a number and following by codec and schema information.

## Adding Tests

As mentioned before tests are generated from Go code.
Adding a test requires adding some more Go code in `gen-validator-tests`.

### Adding an XML Test

Create a new file `./gen-validator-tests/my_tests.go` with the following content:

```go
package main

import (
  "github.com/katydid/validator-go/validator/ast"
  . "github.com/katydid/validator-go/validator/combinator"
)

func init() {
  ValidateXMLString(
    "test_name", // test name
      G{"main": ast.NewTreeNode(ast.NewStringName("test"), ast.NewEmpty())}, // validator grammar
    "<test/>", // input xml
    true, // valid?
  )
}
```

run:
```
$ make regenerate
```

This will then generate files in the `./tests/xml/test_name` folder:

  - `validator.txt` containing the validator expression as a string.
  - `validator.xml` containing an xml serialized validator abstract syntax tree.  This is for new implementations of the validator which have not implemented a parser for the grammar yet.
  - `validator.json` containing a json serialized validator abstract syntax tree.  This is for new implementations of the validator which have not implemented a parser for the grammar yet.
  - `valid.xml` containing the valid xml input.

### Adding a JSON Test

Create a new file `./gen-validator-tests/my_tests.go` with the following content:

```go
package main

import (
  "github.com/katydid/validator-go/validator/ast"
  "github.com/katydid/validator-go/validator/parser"
)

func init() {
  str := `test=="abc"`
  grammar, err := parser.ParseGrammar(str)
  if err != nil {
    panic(err)
  }
  ValidateJsonString(
    "test_name",
    G(ast.NewRefLookup(grammar)),
    `{"test":"abc"}`,
    true,
  )
}
```

run:
```
$ make regenerate
```

This will then generate files in the `./tests/json/test_name` folder:

  - `validator.txt` containing the validator expression as a string.
  - `validator.xml` containing an xml serialized validator abstract syntax tree.  This is for new implementations of the validator which have not implemented a parser for the grammar yet.
  - `validator.json` containing a json serialized validator abstract syntax tree.  This is for new implementations of the validator which have not implemented a parser for the grammar yet.
  - `valid.json` containing the valid json input.

### Adding a Protocol Buffer Test

Create a new proto file `./gen-validator-tests/my.proto` with your protobuf definition.  This should use [gogoprotobuf](https://github.com/gogo/protobuf), since we require a generated Description method that returns a FileDescriptorSet.

Here is an example:

```proto
syntax = "proto2";

package main;

import "github.com/gogo/protobuf/gogoproto/gogo.proto";

option (gogoproto.gostring_all) = true;
option (gogoproto.description_all) = true;

message MyMessage {
	optional string MyField = 1;
}
```

In the `./gen-validator-tests/Makefile` we need to add a command to generate some code for this protocol buffer.

```
(protoc --gogo_out=. -I=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf my.proto)
```

Finally we can create our test file `./gen-validator-tests/my_tests.go` with the following content:

```go
package main

import (
  "github.com/katydid/validator-go/validator/ast"
  "github.com/katydid/validator-go/validator/parser"
  . "github.com/katydid/validator-go/validator/combinator"
  . "github.com/katydid/validator-go/validator/funcs"
)

func init() {
  str := `MyField == "abc"`
  grammar, err := parser.ParseGrammar(largeValidator)
  if err != nil {
    panic(err)
  }
  g := G{"main": grammar.GetTopPattern()}
  msg := &MyMessage{
    MyField: "abc",
  }
  ValidateProtoName(
    "test_name",
    g,
    msg,
    true,
  )
}
```

run:
```
$ make regenerate
```

This will then generate files in the `./tests/pbname/test_name` folder:

  - `validator.txt` containing the validator expression as a string.
  - `validator.xml` containing an xml serialized validator abstract syntax tree.  This is for new implementations of the validator which have not implemented a parser for the grammar yet.
  - `validator.json` containing a json serialized validator abstract syntax tree.  This is for new implementations of the validator which have not implemented a parser for the grammar yet.
  - `valid.main.MyMessage.desc.pbname` containing the valid marshaled protocol buffer input.

It also generates the file `./tests/pbname/main.MyMessage.desc` containing the marshaled FileDescriptorSet for MyMessage.

`ValidateProtoNum` is another function that could have been used to validate the protocol buffer.  This validating translates the validator expression into an expression that relies on field numbers instead of field names.  This is useful for optimization purposes. `ValidateProtoNum` works exactly like `ValidateProtoName` except that it translates the expression and outputs the files in the `pbnum` folder.  This lets the test runner know that it should use the protobuf parser that returns field numbers instead of field names.

`ValidateProtoNumNoRewrite` is another function that could be used to validate a protocol buffer, but it relies on the input expression to have already been translated to field numberes.

### Adding More Tests

`ValidateProtoEtc` and `ValidateProtoNumEtc` are validation functions that not only generates a test for protocol buffers, but also generates tests for json and goreflect. `ValidateProtoNumEtc` additionally generates a test for `pbnum`.

The `./tests/goreflect` folder contains tests for reflected go structures.  These tests are for `Go` implementations only.  Instead of the generated folder containing a `valid.json` file, it contains a `valid.goreflect` file, which also contains json.  This json will be unmarshaled into a go structure that will be passed to the goreflect parser.

# Cross Language Test Suite for Katydid

The test suite is a language agnostic test suite.

The idea is that katydid can be implemented in multiple programming languages. 
Currently an implementations are available in:

  - [Go](https://github.com/katydid/katydid) and 
  - [Haskell](https://github.com/katydid/katydid-haskell).

Having one set of tests that can be used by multiple implementations helps to keep these implementations consistent
and is a great starting point when creating a new implementation.

Technically each of the two implementations consist of three implementations of Relapse's core algorithm, 
so currently there are six users of the test suite.

The test suite is simply a bunch of folders and files that can be traversed and opened in any programming language.

The test suite contains some Go code to help to generate tests for multiple serialization formats. 
The output is just a bunch of files and folders.
There should be no need to run this, except when adding more tests, 
since all the output files and folders are already committed in the repository.
The reason for generating files and folders is that katydid is encoding agnostic, 
so to help make sure that the parsers are implemented consistently, 
we might want to use the same relapse test for multiple serialization formats.

> This Go code for test suite generation should not be confused with the Go tests for the Go implementations of Relapse.
The Go implementations of Relapse all expect the test suite to located adjacent to katydid in the `GOPATH` and the files are read like in any other language's implementation, without using any Go code from the test suite repository.

## Tests

Relapse tests are located in the `./relapse/tests` folder.

Tests are grouped by codec: json, xml, etc.

Inside each codec folder is a list of testcase folders, each for a descriptive name.

Files in the codec folder are schemas that might be required by a testcase.
For example, marshaled protocol buffer file descriptor sets.

Inside each testcase folder is the relapse grammar in several formats: txt, json and xml.
This for cases where a new language has not had time to fully implement a relapse syntax parser.

Also found in the testcase folder is a filename starting with `valid` or `invalid` depending on whether the contents of the file is valid with respect to the provide grammar of invalid.
The file extension of this file again describes the codec of the contents.
In the case where this codec requires a schema, this filename is also provided as a suffix.

## Benchmarks

Relapse benchmarks are located in the `./relapse/benches` folder.

This folder is not checked in, because of its size.
Instead this folder can be generated, by running `make regenerate`.
This will require `go` to be installed and this folder to checked out to 
`$GOPATH/src/github.com/katydid/testsuite`.

Benchmarks are grouped by codec: pbnum (no others are provided yet)

Inside each codec folder is a list of benchcase folders, each for a descriptive name.

Files in the codec folder are schemas that might be required by a benchcase.
For example, marshaled protocol buffer file descriptor sets.

Inside each benchcase folder is the relapse grammar in several formats: txt, json and xml.
This for cases where a new language has not had time to fully implement a relapse syntax parser.

Also found in the benchcase folder is a 1000 generated files each starting with a number and following by codec and schema information.

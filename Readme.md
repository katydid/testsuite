# Cross Language Test Suite for Katydid

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

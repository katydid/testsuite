# 2025-01-16

Benchmark ran using:
* validator-gogo-proto: 50570cdd8aa7506e2336a3d7bd8d707011e3983b
* testsuite: d9643280a7f6d8e19357fffb876718526e1b6982

```
% make bench
TESTSUITE=MUST go test -test.v -test.run=XXX -test.bench=. ./...
?       github.com/katydid/validator-gogo-proto/testsuite       [no test files]
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             9457218               128.3 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                 1000000              1027 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1302956               922.3 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                7456564               161.3 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              13927239                85.71 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  13579813                89.56 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            15955916                74.28 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     15422751                78.68 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3705471               322.7 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     20990870                56.79 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                             10556816               111.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3672676               324.6 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    2535153               468.7 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1437870               834.5 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                2159280               555.1 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2808925               424.6 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2419929               501.7 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/auto  26.381s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14              181797              6728 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   77623             14834 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                  68383             17266 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                 191707              5740 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                247668              4609 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                    239368              4697 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              370128              3101 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       343328              3380 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                              110335             10564 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       771426              1462 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                               188320              6035 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               98546             11992 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                     129967              8832 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       92706             12583 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                 109831             10385 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                   151879              7761 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                   142870              7974 ns/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/intern        21.858s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/interp
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14               49978             23892 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   12303             96912 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                   8070            146790 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                  39349             30183 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                 74841             16064 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                     65000             18001 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              128154              9000 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                        94285             12413 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                               14888             80564 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       323229              3321 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                                49590             23846 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               30903             38098 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                      25640             46789 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       16131             73885 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                  18859             63474 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                    27391             43434 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                    33355             35678 ns/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/interp        28.144s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             7933693               146.9 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                 1000000              1132 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1018 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                6315476               189.9 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              12227974                97.81 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  11872567               100.7 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            14115397                84.21 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     13132005                90.75 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3252654               364.3 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     19227187                63.15 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                              9441136               127.7 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3272506               365.7 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    2280838               522.5 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1313686               903.0 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1960850               611.8 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2506599               476.4 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2208940               542.2 ns/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/mem   25.673s
```

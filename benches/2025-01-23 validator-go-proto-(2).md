# 2025-01-23

Benchmark ran using:
* validator-go-proto: 748c5d232401e016c5d7841eddb44178f2a871c1
* testsuite: 4df6d7feece8ac98b6152687907b209480f944c9

```
TESTSUITE=MUST go test -test.v -test.run=XXX -test.bench=. ./...
?       github.com/katydid/validator-go-proto/testsuite [no test files]
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             8720832               138.2 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                  873488              1252 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1030 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                7412846               161.8 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              13062639                91.60 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  11939258               101.4 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            14561346                81.05 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     13763046                86.56 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3541036               334.4 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     19555651                61.96 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                              9850436               120.8 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3964698               300.0 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    2129930               559.2 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1254628               954.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1836501               650.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2592295               464.1 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2250192               534.9 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/auto    25.622s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14              180561              6708 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   73137             15866 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                  63866             18428 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                 189142              5923 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                235650              4674 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                    238458              4839 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              360606              3136 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       331058              3429 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                              102663             11033 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       757964              1508 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                               183836              6207 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               95895             11958 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                     117745              9723 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       83640             13569 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                 100562             11668 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                   139460              8320 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                   138888              8471 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/intern  21.883s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/interp
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14               62755             19170 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   15396             77758 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                   9450            112162 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                  55032             21292 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                 86536             13509 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                     76711             14768 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              156549              7186 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       123696              9010 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                               20415             58544 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       319575              3507 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                                61040             19206 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               34551             34278 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                      29018             40928 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       19998             59855 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                  22216             53809 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                    34708             34203 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                    39247             29954 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/interp  25.685s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             7564747               158.6 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                  898556              1292 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1148 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                6240261               194.0 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              11319944               106.6 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  10397610               114.8 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            12926431                92.50 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     12249808                97.36 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3050943               394.3 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     17614279                67.67 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                              8733997               138.1 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3174222               375.0 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    1887158               621.5 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1000000              1057 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1662080               734.0 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2326566               523.9 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2016978               587.5 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/mem     25.247s
```
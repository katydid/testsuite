# 2025-01-23

Benchmark ran using:
* validator-gogo-proto: 84ad204c5ff2c510c3ff16a915a8016b358378f0
* testsuite: bf8f2eaf95248ff9aba772430020987c59396529


```
TESTSUITE=MUST go test -test.v -test.run=XXX -test.bench=. ./...
?       github.com/katydid/validator-gogo-proto/testsuite       [no test files]
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             8852842               137.1 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                  995028              1174 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1037 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                7034703               162.1 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              13287202                90.60 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  12129583                98.02 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            14969584                80.23 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     14612376                82.75 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3686316               325.6 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     20349099                60.16 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                             10149818               118.3 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3951202               303.0 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    2138224               561.7 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1251451               962.0 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1844425               651.8 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2576907               466.2 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2269212               528.0 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/auto  26.351s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14              182071              6667 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   72495             16102 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                  62145             18755 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                 185827              5928 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                250282              4667 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                    235912              4827 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              366414              3143 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       330259              3430 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                              103784             10817 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       763026              1538 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                               181912              6258 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               96663             11975 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                     119066              9685 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       86149             13563 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                 100332             11539 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                   140148              8311 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                   137276              8528 ns/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/intern        21.954s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/interp
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14               61538             19389 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   15558             77239 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                   9084            112065 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                  55180             21407 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                 83268             13794 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                     77498             15023 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              155305              7297 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       126861              9083 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                               20635             57420 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       302846              3524 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                                60798             19277 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               34089             34727 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                      29109             41038 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       19849             59866 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                  22338             53893 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                    34610             34458 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                    39033             30351 ns/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/interp        25.653s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-gogo-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             7717353               155.4 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                  918756              1281 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1118 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                6299912               187.5 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              11601138               104.3 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  10904646               110.0 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            13212220                89.64 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     12590229                94.88 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3241988               373.0 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     18282672                65.72 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                              8874475               134.0 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3214746               373.5 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    1964583               608.2 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1000000              1033 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1681545               712.1 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2356833               507.9 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2082982               575.6 ns/op
PASS
ok      github.com/katydid/validator-gogo-proto/testsuite/mem   25.074s
```
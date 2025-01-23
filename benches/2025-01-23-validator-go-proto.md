# 2025-01-23

Benchmark ran using:
* validator-go-proto: aaadb2af4484f900e792a0a147d8f21942fc8210
* testsuite: bf8f2eaf95248ff9aba772430020987c59396529


```
TESTSUITE=MUST go test -test.v -test.run=XXX -test.bench=. ./...
?       github.com/katydid/validator-go-proto/testsuite [no test files]
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             8907015               137.8 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                  903616              1181 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1025 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                7416223               161.8 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              13149153                91.24 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  12106078                98.20 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            14960968                80.01 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     14219462                84.50 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3676624               325.0 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     20276534                59.76 ns/op            0 B/op          0 allocs/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                             10022068               118.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             4006885               298.2 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    2145066               558.8 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1276610               943.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1846922               648.0 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2586146               463.9 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2275358               529.3 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/auto    25.590s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14              181255              6736 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   72355             16030 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                  62769             18516 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                 194490              5905 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                245774              4649 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                    238098              4847 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              360687              3107 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       335379              3455 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                              101809             10899 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       744244              1521 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                               189211              6202 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               95934             12011 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                     117982              9695 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       85920             13594 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                  98583             11573 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                   142544              8301 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                   137259              8557 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/intern  21.968s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/interp
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14               61588             18963 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                   15578             76722 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                   9022            111339 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                  55430             21261 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                                 85513             13489 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                     78214             14921 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                              155172              7147 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                       127951              9027 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                               20865             57120 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                       305689              3485 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                                60480             19225 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                               34363             34120 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                      29143             40741 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                       20098             59474 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                  22352             53602 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                    34380             34012 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                    39783             30037 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/interp  25.539s
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/AndNameTelephonePb
BenchmarkSuite/AndNameTelephonePb-14             7679644               155.8 ns/op
BenchmarkSuite/BridgePepperPb
BenchmarkSuite/BridgePepperPb-14                  906013              1288 ns/op
BenchmarkSuite/BridgePepperAndFountainTargetPb
BenchmarkSuite/BridgePepperAndFountainTargetPb-14                1000000              1118 ns/op
BenchmarkSuite/ContextPersonPb
BenchmarkSuite/ContextPersonPb-14                                6468506               185.7 ns/op
BenchmarkSuite/CorrectNotNamePb
BenchmarkSuite/CorrectNotNamePb-14                              11689753               103.1 ns/op
BenchmarkSuite/EmptyOrNilPb
BenchmarkSuite/EmptyOrNilPb-14                                  10723312               112.3 ns/op
BenchmarkSuite/IncorrectNotNamePb
BenchmarkSuite/IncorrectNotNamePb-14                            13289478                90.17 ns/op
BenchmarkSuite/LenNamePb
BenchmarkSuite/LenNamePb-14                                     12710175                94.85 ns/op
BenchmarkSuite/ListIndexAddressPb
BenchmarkSuite/ListIndexAddressPb-14                             3262567               366.8 ns/op
BenchmarkSuite/NilNamePb
BenchmarkSuite/NilNamePb-14                                     18163081                65.36 ns/op
BenchmarkSuite/OrNameTelephonePb
BenchmarkSuite/OrNameTelephonePb-14                              8933236               134.2 ns/op
BenchmarkSuite/RecursiveSrcTreePb
BenchmarkSuite/RecursiveSrcTreePb-14                             3268129               369.6 ns/op
BenchmarkSuite/TypewriterPrisonDaisySledPb
BenchmarkSuite/TypewriterPrisonDaisySledPb-14                    1955425               615.8 ns/op
BenchmarkSuite/TypewriterPrisonMapSharkPb
BenchmarkSuite/TypewriterPrisonMapSharkPb-14                     1000000              1035 ns/op
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb
BenchmarkSuite/TypewriterPrisonMenuPaperclipPb-14                1681809               709.6 ns/op
BenchmarkSuite/TypewriterPrisonScarBusStopPb
BenchmarkSuite/TypewriterPrisonScarBusStopPb-14                  2358129               507.8 ns/op
BenchmarkSuite/TypewriterPrisonSmileLetterPb
BenchmarkSuite/TypewriterPrisonSmileLetterPb-14                  2080090               576.4 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/mem     25.108s
```
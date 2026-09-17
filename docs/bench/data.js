window.BENCHMARK_DATA = {
  "lastUpdate": 1789649880732,
  "repoUrl": "https://github.com/moov-io/wire20022",
  "entries": {
    "moov-io/wire20022": [
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "49063b65846bbc5f4f4aaddf8efac466ca337848",
          "message": "ci: run ISO 20022 parse/write Go benchmarks in this repository (#69)",
          "timestamp": "2026-09-14T18:43:57Z",
          "url": "https://github.com/moov-io/wire20022/commit/49063b65846bbc5f4f4aaddf8efac466ca337848"
        },
        "date": 1789411540325,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 75566,
            "unit": "ns/op\t   72200 B/op\t     842 allocs/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 75566,
            "unit": "ns/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72200,
            "unit": "B/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 84677,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 84677,
            "unit": "ns/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 473893,
            "unit": "ns/op\t  467238 B/op\t    5580 allocs/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 473893,
            "unit": "ns/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 467238,
            "unit": "B/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5580,
            "unit": "allocs/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1445631,
            "unit": "ns/op\t 1423936 B/op\t   17565 allocs/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1445631,
            "unit": "ns/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423936,
            "unit": "B/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2298,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2298,
            "unit": "ns/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 402341,
            "unit": "ns/op\t  363145 B/op\t    4626 allocs/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 402341,
            "unit": "ns/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363145,
            "unit": "B/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3767,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3767,
            "unit": "ns/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 752855,
            "unit": "ns/op\t  690739 B/op\t    8636 allocs/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 752855,
            "unit": "ns/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690739,
            "unit": "B/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 176090,
            "unit": "ns/op\t  150967 B/op\t    1864 allocs/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 176090,
            "unit": "ns/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 150967,
            "unit": "B/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1864,
            "unit": "allocs/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 410230,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 410230,
            "unit": "ns/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 435634,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 435634,
            "unit": "ns/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 820884,
            "unit": "ns/op\t  780081 B/op\t    9776 allocs/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 820884,
            "unit": "ns/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780081,
            "unit": "B/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1441 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "49063b65846bbc5f4f4aaddf8efac466ca337848",
          "message": "ci: run ISO 20022 parse/write Go benchmarks in this repository (#69)",
          "timestamp": "2026-09-14T18:43:57Z",
          "url": "https://github.com/moov-io/wire20022/commit/49063b65846bbc5f4f4aaddf8efac466ca337848"
        },
        "date": 1789477447042,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 56632,
            "unit": "ns/op\t   72413 B/op\t     842 allocs/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 56632,
            "unit": "ns/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72413,
            "unit": "B/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 63835,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 63835,
            "unit": "ns/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 370984,
            "unit": "ns/op\t  469415 B/op\t    5608 allocs/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 370984,
            "unit": "ns/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 469415,
            "unit": "B/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5608,
            "unit": "allocs/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1128888,
            "unit": "ns/op\t 1423955 B/op\t   17565 allocs/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1128888,
            "unit": "ns/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423955,
            "unit": "B/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1759,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1759,
            "unit": "ns/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 307155,
            "unit": "ns/op\t  363146 B/op\t    4626 allocs/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 307155,
            "unit": "ns/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363146,
            "unit": "B/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 2988,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 2988,
            "unit": "ns/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 566505,
            "unit": "ns/op\t  690749 B/op\t    8636 allocs/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 566505,
            "unit": "ns/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690749,
            "unit": "B/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 124115,
            "unit": "ns/op\t  154436 B/op\t    1908 allocs/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 124115,
            "unit": "ns/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 154436,
            "unit": "B/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1908,
            "unit": "allocs/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 312409,
            "unit": "ns/op\t  367481 B/op\t    4713 allocs/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 312409,
            "unit": "ns/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367481,
            "unit": "B/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 334374,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 334374,
            "unit": "ns/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 665176,
            "unit": "ns/op\t  780069 B/op\t    9776 allocs/op",
            "extra": "1864 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 665176,
            "unit": "ns/op",
            "extra": "1864 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780069,
            "unit": "B/op",
            "extra": "1864 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1864 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "49063b65846bbc5f4f4aaddf8efac466ca337848",
          "message": "ci: run ISO 20022 parse/write Go benchmarks in this repository (#69)",
          "timestamp": "2026-09-14T18:43:57Z",
          "url": "https://github.com/moov-io/wire20022/commit/49063b65846bbc5f4f4aaddf8efac466ca337848"
        },
        "date": 1789563684065,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 75857,
            "unit": "ns/op\t   72330 B/op\t     842 allocs/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 75857,
            "unit": "ns/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72330,
            "unit": "B/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 83286,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 83286,
            "unit": "ns/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 464944,
            "unit": "ns/op\t  460897 B/op\t    5506 allocs/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 464944,
            "unit": "ns/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 460897,
            "unit": "B/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5506,
            "unit": "allocs/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1455242,
            "unit": "ns/op\t 1423968 B/op\t   17565 allocs/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1455242,
            "unit": "ns/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423968,
            "unit": "B/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2381,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2381,
            "unit": "ns/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 396707,
            "unit": "ns/op\t  363142 B/op\t    4626 allocs/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 396707,
            "unit": "ns/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363142,
            "unit": "B/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3787,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3787,
            "unit": "ns/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 743448,
            "unit": "ns/op\t  690733 B/op\t    8636 allocs/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 743448,
            "unit": "ns/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690733,
            "unit": "B/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 163984,
            "unit": "ns/op\t  152132 B/op\t    1878 allocs/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 163984,
            "unit": "ns/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 152132,
            "unit": "B/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1878,
            "unit": "allocs/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 407033,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 407033,
            "unit": "ns/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 433858,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 433858,
            "unit": "ns/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 812007,
            "unit": "ns/op\t  780043 B/op\t    9776 allocs/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 812007,
            "unit": "ns/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780043,
            "unit": "B/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1450 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "49063b65846bbc5f4f4aaddf8efac466ca337848",
          "message": "ci: run ISO 20022 parse/write Go benchmarks in this repository (#69)",
          "timestamp": "2026-09-14T18:43:57Z",
          "url": "https://github.com/moov-io/wire20022/commit/49063b65846bbc5f4f4aaddf8efac466ca337848"
        },
        "date": 1789649880288,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 75784,
            "unit": "ns/op\t   72255 B/op\t     842 allocs/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 75784,
            "unit": "ns/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72255,
            "unit": "B/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 83660,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 83660,
            "unit": "ns/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 472656,
            "unit": "ns/op\t  467909 B/op\t    5596 allocs/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 472656,
            "unit": "ns/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 467909,
            "unit": "B/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5596,
            "unit": "allocs/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1437365,
            "unit": "ns/op\t 1423914 B/op\t   17565 allocs/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1437365,
            "unit": "ns/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423914,
            "unit": "B/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2278,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2278,
            "unit": "ns/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 398266,
            "unit": "ns/op\t  363135 B/op\t    4625 allocs/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 398266,
            "unit": "ns/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363135,
            "unit": "B/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3733,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3733,
            "unit": "ns/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 747728,
            "unit": "ns/op\t  690735 B/op\t    8636 allocs/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 747728,
            "unit": "ns/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690735,
            "unit": "B/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 168132,
            "unit": "ns/op\t  156331 B/op\t    1930 allocs/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 168132,
            "unit": "ns/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 156331,
            "unit": "B/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1930,
            "unit": "allocs/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 409419,
            "unit": "ns/op\t  367475 B/op\t    4713 allocs/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 409419,
            "unit": "ns/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367475,
            "unit": "B/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 439123,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 439123,
            "unit": "ns/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 812789,
            "unit": "ns/op\t  780037 B/op\t    9776 allocs/op",
            "extra": "1438 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 812789,
            "unit": "ns/op",
            "extra": "1438 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780037,
            "unit": "B/op",
            "extra": "1438 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1438 times\n4 procs"
          }
        ]
      }
    ]
  }
}
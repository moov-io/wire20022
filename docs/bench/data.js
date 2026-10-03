window.BENCHMARK_DATA = {
  "lastUpdate": 1791032242069,
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
        "date": 1789734995567,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 72379,
            "unit": "ns/op\t   72349 B/op\t     842 allocs/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 72379,
            "unit": "ns/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72349,
            "unit": "B/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 77953,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 77953,
            "unit": "ns/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 449531,
            "unit": "ns/op\t  469219 B/op\t    5609 allocs/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 449531,
            "unit": "ns/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 469219,
            "unit": "B/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5609,
            "unit": "allocs/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1356680,
            "unit": "ns/op\t 1423955 B/op\t   17565 allocs/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1356680,
            "unit": "ns/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423955,
            "unit": "B/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2217,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2217,
            "unit": "ns/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 373287,
            "unit": "ns/op\t  363148 B/op\t    4626 allocs/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 373287,
            "unit": "ns/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363148,
            "unit": "B/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3488,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3488,
            "unit": "ns/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 741881,
            "unit": "ns/op\t  690753 B/op\t    8636 allocs/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 741881,
            "unit": "ns/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690753,
            "unit": "B/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 160252,
            "unit": "ns/op\t  158836 B/op\t    1961 allocs/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 160252,
            "unit": "ns/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 158836,
            "unit": "B/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1961,
            "unit": "allocs/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 378908,
            "unit": "ns/op\t  367479 B/op\t    4713 allocs/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 378908,
            "unit": "ns/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367479,
            "unit": "B/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 402110,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 402110,
            "unit": "ns/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 765148,
            "unit": "ns/op\t  780053 B/op\t    9776 allocs/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 765148,
            "unit": "ns/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780053,
            "unit": "B/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1504 times\n4 procs"
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
        "date": 1789819989019,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 76658,
            "unit": "ns/op\t   72340 B/op\t     842 allocs/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 76658,
            "unit": "ns/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72340,
            "unit": "B/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 84277,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 84277,
            "unit": "ns/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 466899,
            "unit": "ns/op\t  460963 B/op\t    5507 allocs/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 466899,
            "unit": "ns/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 460963,
            "unit": "B/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5507,
            "unit": "allocs/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1442848,
            "unit": "ns/op\t 1423951 B/op\t   17565 allocs/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1442848,
            "unit": "ns/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423951,
            "unit": "B/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2289,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2289,
            "unit": "ns/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 435242,
            "unit": "ns/op\t  363143 B/op\t    4626 allocs/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 435242,
            "unit": "ns/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363143,
            "unit": "B/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3749,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3749,
            "unit": "ns/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 745594,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 745594,
            "unit": "ns/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 166993,
            "unit": "ns/op\t  155542 B/op\t    1921 allocs/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 166993,
            "unit": "ns/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 155542,
            "unit": "B/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1921,
            "unit": "allocs/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 406212,
            "unit": "ns/op\t  367476 B/op\t    4713 allocs/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 406212,
            "unit": "ns/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367476,
            "unit": "B/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 436375,
            "unit": "ns/op\t  414440 B/op\t    4981 allocs/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 436375,
            "unit": "ns/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414440,
            "unit": "B/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 811315,
            "unit": "ns/op\t  780045 B/op\t    9776 allocs/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 811315,
            "unit": "ns/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780045,
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
        "date": 1789908536280,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 84654,
            "unit": "ns/op\t   72324 B/op\t     842 allocs/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 84654,
            "unit": "ns/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72324,
            "unit": "B/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 90183,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 90183,
            "unit": "ns/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 548187,
            "unit": "ns/op\t  467572 B/op\t    5582 allocs/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 548187,
            "unit": "ns/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 467572,
            "unit": "B/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5582,
            "unit": "allocs/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1519323,
            "unit": "ns/op\t 1423921 B/op\t   17565 allocs/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1519323,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423921,
            "unit": "B/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2343,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2343,
            "unit": "ns/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 421546,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 421546,
            "unit": "ns/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3886,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3886,
            "unit": "ns/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 792977,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 792977,
            "unit": "ns/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 184938,
            "unit": "ns/op\t  158848 B/op\t    1962 allocs/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 184938,
            "unit": "ns/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 158848,
            "unit": "B/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1962,
            "unit": "allocs/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 440406,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 440406,
            "unit": "ns/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 468311,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 468311,
            "unit": "ns/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 878713,
            "unit": "ns/op\t  780054 B/op\t    9776 allocs/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 878713,
            "unit": "ns/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780054,
            "unit": "B/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1329 times\n4 procs"
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
        "date": 1790001991609,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 79018,
            "unit": "ns/op\t   72273 B/op\t     842 allocs/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 79018,
            "unit": "ns/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72273,
            "unit": "B/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 84215,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 84215,
            "unit": "ns/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 458830,
            "unit": "ns/op\t  460819 B/op\t    5515 allocs/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 458830,
            "unit": "ns/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 460819,
            "unit": "B/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5515,
            "unit": "allocs/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1432276,
            "unit": "ns/op\t 1423908 B/op\t   17565 allocs/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1432276,
            "unit": "ns/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423908,
            "unit": "B/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2265,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2265,
            "unit": "ns/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 393074,
            "unit": "ns/op\t  363133 B/op\t    4625 allocs/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 393074,
            "unit": "ns/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363133,
            "unit": "B/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3629,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3629,
            "unit": "ns/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 731541,
            "unit": "ns/op\t  690736 B/op\t    8636 allocs/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 731541,
            "unit": "ns/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690736,
            "unit": "B/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 167722,
            "unit": "ns/op\t  157166 B/op\t    1941 allocs/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 167722,
            "unit": "ns/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 157166,
            "unit": "B/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1941,
            "unit": "allocs/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 401169,
            "unit": "ns/op\t  367468 B/op\t    4712 allocs/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 401169,
            "unit": "ns/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367468,
            "unit": "B/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 426929,
            "unit": "ns/op\t  414436 B/op\t    4981 allocs/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 426929,
            "unit": "ns/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414436,
            "unit": "B/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 878857,
            "unit": "ns/op\t  780061 B/op\t    9776 allocs/op",
            "extra": "1419 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 878857,
            "unit": "ns/op",
            "extra": "1419 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780061,
            "unit": "B/op",
            "extra": "1419 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1419 times\n4 procs"
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
        "date": 1790082051021,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 49704,
            "unit": "ns/op\t   72432 B/op\t     842 allocs/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 49704,
            "unit": "ns/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72432,
            "unit": "B/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 57437,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 57437,
            "unit": "ns/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 311418,
            "unit": "ns/op\t  463201 B/op\t    5529 allocs/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 311418,
            "unit": "ns/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 463201,
            "unit": "B/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5529,
            "unit": "allocs/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1066296,
            "unit": "ns/op\t 1423961 B/op\t   17565 allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1066296,
            "unit": "ns/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423961,
            "unit": "B/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1484,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1484,
            "unit": "ns/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 258837,
            "unit": "ns/op\t  363150 B/op\t    4626 allocs/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 258837,
            "unit": "ns/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363150,
            "unit": "B/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 2633,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 2633,
            "unit": "ns/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 483067,
            "unit": "ns/op\t  690758 B/op\t    8636 allocs/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 483067,
            "unit": "ns/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690758,
            "unit": "B/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 113752,
            "unit": "ns/op\t  163545 B/op\t    2019 allocs/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 113752,
            "unit": "ns/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 163545,
            "unit": "B/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 2019,
            "unit": "allocs/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 265344,
            "unit": "ns/op\t  367485 B/op\t    4713 allocs/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 265344,
            "unit": "ns/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367485,
            "unit": "B/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 279992,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 279992,
            "unit": "ns/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 531199,
            "unit": "ns/op\t  780060 B/op\t    9776 allocs/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 531199,
            "unit": "ns/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780060,
            "unit": "B/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "2202 times\n4 procs"
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
        "date": 1790169303849,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 49649,
            "unit": "ns/op\t   72446 B/op\t     842 allocs/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 49649,
            "unit": "ns/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72446,
            "unit": "B/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 53647,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 53647,
            "unit": "ns/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 332033,
            "unit": "ns/op\t  472409 B/op\t    5643 allocs/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 332033,
            "unit": "ns/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 472409,
            "unit": "B/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5643,
            "unit": "allocs/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1045585,
            "unit": "ns/op\t 1423949 B/op\t   17565 allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1045585,
            "unit": "ns/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423949,
            "unit": "B/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1670,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1670,
            "unit": "ns/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 275893,
            "unit": "ns/op\t  363146 B/op\t    4626 allocs/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 275893,
            "unit": "ns/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363146,
            "unit": "B/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 2655,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 2655,
            "unit": "ns/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 487443,
            "unit": "ns/op\t  690756 B/op\t    8636 allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 487443,
            "unit": "ns/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690756,
            "unit": "B/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 111865,
            "unit": "ns/op\t  156119 B/op\t    1928 allocs/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 111865,
            "unit": "ns/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 156119,
            "unit": "B/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1928,
            "unit": "allocs/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 294202,
            "unit": "ns/op\t  367486 B/op\t    4713 allocs/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 294202,
            "unit": "ns/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367486,
            "unit": "B/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 297468,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 297468,
            "unit": "ns/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 556957,
            "unit": "ns/op\t  780067 B/op\t    9776 allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 556957,
            "unit": "ns/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780067,
            "unit": "B/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "2168 times\n4 procs"
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
        "date": 1790255165126,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 76068,
            "unit": "ns/op\t   72361 B/op\t     842 allocs/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 76068,
            "unit": "ns/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72361,
            "unit": "B/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 83765,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 83765,
            "unit": "ns/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 467408,
            "unit": "ns/op\t  459893 B/op\t    5495 allocs/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 467408,
            "unit": "ns/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 459893,
            "unit": "B/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5495,
            "unit": "allocs/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1455827,
            "unit": "ns/op\t 1423944 B/op\t   17565 allocs/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1455827,
            "unit": "ns/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423944,
            "unit": "B/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2284,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2284,
            "unit": "ns/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 405434,
            "unit": "ns/op\t  363145 B/op\t    4626 allocs/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 405434,
            "unit": "ns/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363145,
            "unit": "B/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 4035,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 4035,
            "unit": "ns/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 745650,
            "unit": "ns/op\t  690736 B/op\t    8636 allocs/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 745650,
            "unit": "ns/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690736,
            "unit": "B/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 162383,
            "unit": "ns/op\t  151037 B/op\t    1865 allocs/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 162383,
            "unit": "ns/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 151037,
            "unit": "B/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1865,
            "unit": "allocs/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 409411,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 409411,
            "unit": "ns/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 437127,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 437127,
            "unit": "ns/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 816331,
            "unit": "ns/op\t  780050 B/op\t    9776 allocs/op",
            "extra": "1460 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 816331,
            "unit": "ns/op",
            "extra": "1460 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780050,
            "unit": "B/op",
            "extra": "1460 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1460 times\n4 procs"
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
        "date": 1790341841087,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 53086,
            "unit": "ns/op\t   72371 B/op\t     842 allocs/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 53086,
            "unit": "ns/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72371,
            "unit": "B/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 60895,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 60895,
            "unit": "ns/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 374523,
            "unit": "ns/op\t  474551 B/op\t    5669 allocs/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 374523,
            "unit": "ns/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 474551,
            "unit": "B/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5669,
            "unit": "allocs/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1088041,
            "unit": "ns/op\t 1423944 B/op\t   17565 allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1088041,
            "unit": "ns/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423944,
            "unit": "B/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1798,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1798,
            "unit": "ns/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 292306,
            "unit": "ns/op\t  363148 B/op\t    4626 allocs/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 292306,
            "unit": "ns/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363148,
            "unit": "B/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 2728,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 2728,
            "unit": "ns/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 536868,
            "unit": "ns/op\t  690743 B/op\t    8636 allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 536868,
            "unit": "ns/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690743,
            "unit": "B/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 123419,
            "unit": "ns/op\t  161595 B/op\t    1995 allocs/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 123419,
            "unit": "ns/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 161595,
            "unit": "B/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1995,
            "unit": "allocs/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 292650,
            "unit": "ns/op\t  367480 B/op\t    4713 allocs/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 292650,
            "unit": "ns/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367480,
            "unit": "B/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 329700,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 329700,
            "unit": "ns/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 588664,
            "unit": "ns/op\t  780052 B/op\t    9776 allocs/op",
            "extra": "1968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 588664,
            "unit": "ns/op",
            "extra": "1968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780052,
            "unit": "B/op",
            "extra": "1968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1968 times\n4 procs"
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
        "date": 1790426249339,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 78890,
            "unit": "ns/op\t   72306 B/op\t     842 allocs/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 78890,
            "unit": "ns/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72306,
            "unit": "B/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 82662,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 82662,
            "unit": "ns/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 462641,
            "unit": "ns/op\t  463815 B/op\t    5547 allocs/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 462641,
            "unit": "ns/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 463815,
            "unit": "B/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5547,
            "unit": "allocs/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1433584,
            "unit": "ns/op\t 1423927 B/op\t   17565 allocs/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1433584,
            "unit": "ns/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423927,
            "unit": "B/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2287,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2287,
            "unit": "ns/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 404251,
            "unit": "ns/op\t  363144 B/op\t    4626 allocs/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 404251,
            "unit": "ns/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363144,
            "unit": "B/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3768,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3768,
            "unit": "ns/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 740256,
            "unit": "ns/op\t  690729 B/op\t    8636 allocs/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 740256,
            "unit": "ns/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690729,
            "unit": "B/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 174198,
            "unit": "ns/op\t  158691 B/op\t    1960 allocs/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 174198,
            "unit": "ns/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 158691,
            "unit": "B/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1960,
            "unit": "allocs/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 412544,
            "unit": "ns/op\t  367482 B/op\t    4713 allocs/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 412544,
            "unit": "ns/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367482,
            "unit": "B/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 435909,
            "unit": "ns/op\t  414436 B/op\t    4981 allocs/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 435909,
            "unit": "ns/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414436,
            "unit": "B/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 808986,
            "unit": "ns/op\t  780043 B/op\t    9776 allocs/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 808986,
            "unit": "ns/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780043,
            "unit": "B/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1467 times\n4 procs"
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
        "date": 1790515730127,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 66452,
            "unit": "ns/op\t   72408 B/op\t     842 allocs/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 66452,
            "unit": "ns/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72408,
            "unit": "B/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 73128,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 73128,
            "unit": "ns/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 420033,
            "unit": "ns/op\t  460798 B/op\t    5498 allocs/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 420033,
            "unit": "ns/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 460798,
            "unit": "B/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5498,
            "unit": "allocs/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1303760,
            "unit": "ns/op\t 1423935 B/op\t   17565 allocs/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1303760,
            "unit": "ns/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423935,
            "unit": "B/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2193,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2193,
            "unit": "ns/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 361633,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 361633,
            "unit": "ns/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3540,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3540,
            "unit": "ns/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 659975,
            "unit": "ns/op\t  690744 B/op\t    8636 allocs/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 659975,
            "unit": "ns/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690744,
            "unit": "B/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 155014,
            "unit": "ns/op\t  161567 B/op\t    1995 allocs/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 155014,
            "unit": "ns/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 161567,
            "unit": "B/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1995,
            "unit": "allocs/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 377254,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 377254,
            "unit": "ns/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 387549,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 387549,
            "unit": "ns/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 737703,
            "unit": "ns/op\t  780061 B/op\t    9776 allocs/op",
            "extra": "1639 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 737703,
            "unit": "ns/op",
            "extra": "1639 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780061,
            "unit": "B/op",
            "extra": "1639 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1639 times\n4 procs"
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
        "date": 1790612526938,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 42546,
            "unit": "ns/op\t   72335 B/op\t     842 allocs/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 42546,
            "unit": "ns/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72335,
            "unit": "B/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 46540,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 46540,
            "unit": "ns/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 260679,
            "unit": "ns/op\t  461625 B/op\t    5517 allocs/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 260679,
            "unit": "ns/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 461625,
            "unit": "B/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5517,
            "unit": "allocs/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 812418,
            "unit": "ns/op\t 1423943 B/op\t   17565 allocs/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 812418,
            "unit": "ns/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423943,
            "unit": "B/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1299,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1299,
            "unit": "ns/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 215825,
            "unit": "ns/op\t  363144 B/op\t    4626 allocs/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 215825,
            "unit": "ns/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363144,
            "unit": "B/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 2137,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 2137,
            "unit": "ns/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 432650,
            "unit": "ns/op\t  690749 B/op\t    8636 allocs/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 432650,
            "unit": "ns/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690749,
            "unit": "B/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 94910,
            "unit": "ns/op\t  160874 B/op\t    1987 allocs/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 94910,
            "unit": "ns/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 160874,
            "unit": "B/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1987,
            "unit": "allocs/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 225199,
            "unit": "ns/op\t  367475 B/op\t    4713 allocs/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 225199,
            "unit": "ns/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367475,
            "unit": "B/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 237693,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 237693,
            "unit": "ns/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 446768,
            "unit": "ns/op\t  780051 B/op\t    9776 allocs/op",
            "extra": "2629 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 446768,
            "unit": "ns/op",
            "extra": "2629 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780051,
            "unit": "B/op",
            "extra": "2629 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "2629 times\n4 procs"
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
        "date": 1790692346277,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 77453,
            "unit": "ns/op\t   72374 B/op\t     842 allocs/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 77453,
            "unit": "ns/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72374,
            "unit": "B/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 84182,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 84182,
            "unit": "ns/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 494330,
            "unit": "ns/op\t  460234 B/op\t    5493 allocs/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 494330,
            "unit": "ns/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 460234,
            "unit": "B/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5493,
            "unit": "allocs/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1528126,
            "unit": "ns/op\t 1423950 B/op\t   17565 allocs/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1528126,
            "unit": "ns/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423950,
            "unit": "B/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2335,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2335,
            "unit": "ns/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 427791,
            "unit": "ns/op\t  363150 B/op\t    4626 allocs/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 427791,
            "unit": "ns/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363150,
            "unit": "B/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3946,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3946,
            "unit": "ns/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 744008,
            "unit": "ns/op\t  690747 B/op\t    8636 allocs/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 744008,
            "unit": "ns/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690747,
            "unit": "B/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 162812,
            "unit": "ns/op\t  149924 B/op\t    1852 allocs/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 162812,
            "unit": "ns/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 149924,
            "unit": "B/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1852,
            "unit": "allocs/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 412524,
            "unit": "ns/op\t  367478 B/op\t    4713 allocs/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 412524,
            "unit": "ns/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367478,
            "unit": "B/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 441059,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 441059,
            "unit": "ns/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 826624,
            "unit": "ns/op\t  780047 B/op\t    9776 allocs/op",
            "extra": "1395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 826624,
            "unit": "ns/op",
            "extra": "1395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780047,
            "unit": "B/op",
            "extra": "1395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1395 times\n4 procs"
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
        "date": 1790778663550,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 39395,
            "unit": "ns/op\t   72329 B/op\t     842 allocs/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 39395,
            "unit": "ns/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72329,
            "unit": "B/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 41033,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 41033,
            "unit": "ns/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 244237,
            "unit": "ns/op\t  461206 B/op\t    5506 allocs/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 244237,
            "unit": "ns/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 461206,
            "unit": "B/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5506,
            "unit": "allocs/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 747987,
            "unit": "ns/op\t 1423946 B/op\t   17565 allocs/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 747987,
            "unit": "ns/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423946,
            "unit": "B/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1267,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1267,
            "unit": "ns/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 205333,
            "unit": "ns/op\t  363142 B/op\t    4626 allocs/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 205333,
            "unit": "ns/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363142,
            "unit": "B/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 1968,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 1968,
            "unit": "ns/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 401589,
            "unit": "ns/op\t  690749 B/op\t    8636 allocs/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 401589,
            "unit": "ns/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690749,
            "unit": "B/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 85666,
            "unit": "ns/op\t  156698 B/op\t    1935 allocs/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 85666,
            "unit": "ns/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 156698,
            "unit": "B/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1935,
            "unit": "allocs/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 217564,
            "unit": "ns/op\t  367472 B/op\t    4712 allocs/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 217564,
            "unit": "ns/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367472,
            "unit": "B/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 228722,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 228722,
            "unit": "ns/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 424649,
            "unit": "ns/op\t  780062 B/op\t    9776 allocs/op",
            "extra": "2835 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 424649,
            "unit": "ns/op",
            "extra": "2835 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780062,
            "unit": "B/op",
            "extra": "2835 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "2835 times\n4 procs"
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
        "date": 1790866850604,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 75272,
            "unit": "ns/op\t   72336 B/op\t     842 allocs/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 75272,
            "unit": "ns/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72336,
            "unit": "B/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 84463,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 84463,
            "unit": "ns/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 498004,
            "unit": "ns/op\t  461341 B/op\t    5510 allocs/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 498004,
            "unit": "ns/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 461341,
            "unit": "B/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5510,
            "unit": "allocs/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1456165,
            "unit": "ns/op\t 1423949 B/op\t   17565 allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1456165,
            "unit": "ns/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423949,
            "unit": "B/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2272,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2272,
            "unit": "ns/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 404189,
            "unit": "ns/op\t  363143 B/op\t    4626 allocs/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 404189,
            "unit": "ns/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363143,
            "unit": "B/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3600,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3600,
            "unit": "ns/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 731502,
            "unit": "ns/op\t  690736 B/op\t    8636 allocs/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 731502,
            "unit": "ns/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690736,
            "unit": "B/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 168327,
            "unit": "ns/op\t  160011 B/op\t    1976 allocs/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 168327,
            "unit": "ns/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 160011,
            "unit": "B/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1976,
            "unit": "allocs/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 408192,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 408192,
            "unit": "ns/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 426897,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 426897,
            "unit": "ns/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 806943,
            "unit": "ns/op\t  780039 B/op\t    9776 allocs/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 806943,
            "unit": "ns/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780039,
            "unit": "B/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1429 times\n4 procs"
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
        "date": 1790951012006,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 38275,
            "unit": "ns/op\t   72299 B/op\t     842 allocs/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 38275,
            "unit": "ns/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72299,
            "unit": "B/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 44137,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 44137,
            "unit": "ns/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 245548,
            "unit": "ns/op\t  461746 B/op\t    5517 allocs/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 245548,
            "unit": "ns/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 461746,
            "unit": "B/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5517,
            "unit": "allocs/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 775285,
            "unit": "ns/op\t 1423944 B/op\t   17565 allocs/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 775285,
            "unit": "ns/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423944,
            "unit": "B/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 1235,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 1235,
            "unit": "ns/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 208941,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 208941,
            "unit": "ns/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 2032,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 2032,
            "unit": "ns/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 422210,
            "unit": "ns/op\t  690753 B/op\t    8636 allocs/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 422210,
            "unit": "ns/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690753,
            "unit": "B/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 88340,
            "unit": "ns/op\t  156623 B/op\t    1934 allocs/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 88340,
            "unit": "ns/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 156623,
            "unit": "B/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1934,
            "unit": "allocs/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 223057,
            "unit": "ns/op\t  367481 B/op\t    4713 allocs/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 223057,
            "unit": "ns/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367481,
            "unit": "B/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 233107,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 233107,
            "unit": "ns/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 457010,
            "unit": "ns/op\t  780052 B/op\t    9776 allocs/op",
            "extra": "2487 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 457010,
            "unit": "ns/op",
            "extra": "2487 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780052,
            "unit": "B/op",
            "extra": "2487 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "2487 times\n4 procs"
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
        "date": 1791032241690,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 78303,
            "unit": "ns/op\t   72325 B/op\t     842 allocs/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 78303,
            "unit": "ns/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72325,
            "unit": "B/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck)",
            "value": 85513,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - ns/op",
            "value": 85513,
            "unit": "ns/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ConnectionCheck) - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 495778,
            "unit": "ns/op\t  480207 B/op\t    5745 allocs/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 495778,
            "unit": "ns/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 480207,
            "unit": "B/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 5745,
            "unit": "allocs/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest)",
            "value": 1489553,
            "unit": "ns/op\t 1423922 B/op\t   17565 allocs/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - ns/op",
            "value": 1489553,
            "unit": "ns/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - B/op",
            "value": 1423922,
            "unit": "B/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/DrawdownRequest) - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 2320,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 2320,
            "unit": "ns/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master)",
            "value": 408165,
            "unit": "ns/op\t  363141 B/op\t    4626 allocs/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - ns/op",
            "value": 408165,
            "unit": "ns/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - B/op",
            "value": 363141,
            "unit": "B/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/Master) - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 3806,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 3806,
            "unit": "ns/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn)",
            "value": 762826,
            "unit": "ns/op\t  690742 B/op\t    8636 allocs/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - ns/op",
            "value": 762826,
            "unit": "ns/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - B/op",
            "value": 690742,
            "unit": "B/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentReturn) - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 172445,
            "unit": "ns/op\t  151076 B/op\t    1866 allocs/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 172445,
            "unit": "ns/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 151076,
            "unit": "B/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 1866,
            "unit": "allocs/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest)",
            "value": 418549,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - ns/op",
            "value": 418549,
            "unit": "ns/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest) - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 443130,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 443130,
            "unit": "ns/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse)",
            "value": 833942,
            "unit": "ns/op\t  780048 B/op\t    9776 allocs/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - ns/op",
            "value": 833942,
            "unit": "ns/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - B/op",
            "value": 780048,
            "unit": "B/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML (github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse) - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1400 times\n4 procs"
          }
        ]
      }
    ]
  }
}
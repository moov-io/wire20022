window.BENCHMARK_DATA = {
  "lastUpdate": 1792760541352,
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 75566,
            "unit": "ns/op\t   72200 B/op\t     842 allocs/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 75566,
            "unit": "ns/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72200,
            "unit": "B/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "13948 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 84677,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 84677,
            "unit": "ns/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13654 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 473893,
            "unit": "ns/op\t  467238 B/op\t    5580 allocs/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 473893,
            "unit": "ns/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 467238,
            "unit": "B/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5580,
            "unit": "allocs/op",
            "extra": "2271 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1445631,
            "unit": "ns/op\t 1423936 B/op\t   17565 allocs/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1445631,
            "unit": "ns/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423936,
            "unit": "B/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2298,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2298,
            "unit": "ns/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "487891 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 402341,
            "unit": "ns/op\t  363145 B/op\t    4626 allocs/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 402341,
            "unit": "ns/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363145,
            "unit": "B/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2912 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3767,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3767,
            "unit": "ns/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "322882 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 752855,
            "unit": "ns/op\t  690739 B/op\t    8636 allocs/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 752855,
            "unit": "ns/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690739,
            "unit": "B/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1598 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 176090,
            "unit": "ns/op\t  150967 B/op\t    1864 allocs/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 176090,
            "unit": "ns/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 150967,
            "unit": "B/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1864,
            "unit": "allocs/op",
            "extra": "7298 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 410230,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 410230,
            "unit": "ns/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2817 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 435634,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 435634,
            "unit": "ns/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2667 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 820884,
            "unit": "ns/op\t  780081 B/op\t    9776 allocs/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 820884,
            "unit": "ns/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780081,
            "unit": "B/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 56632,
            "unit": "ns/op\t   72413 B/op\t     842 allocs/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 56632,
            "unit": "ns/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72413,
            "unit": "B/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "20652 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 63835,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 63835,
            "unit": "ns/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "18822 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 370984,
            "unit": "ns/op\t  469415 B/op\t    5608 allocs/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 370984,
            "unit": "ns/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 469415,
            "unit": "B/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5608,
            "unit": "allocs/op",
            "extra": "3249 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1128888,
            "unit": "ns/op\t 1423955 B/op\t   17565 allocs/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1128888,
            "unit": "ns/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423955,
            "unit": "B/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1759,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1759,
            "unit": "ns/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "668275 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 307155,
            "unit": "ns/op\t  363146 B/op\t    4626 allocs/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 307155,
            "unit": "ns/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363146,
            "unit": "B/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3811 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2988,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2988,
            "unit": "ns/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "389994 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 566505,
            "unit": "ns/op\t  690749 B/op\t    8636 allocs/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 566505,
            "unit": "ns/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690749,
            "unit": "B/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2114 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 124115,
            "unit": "ns/op\t  154436 B/op\t    1908 allocs/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 124115,
            "unit": "ns/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 154436,
            "unit": "B/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1908,
            "unit": "allocs/op",
            "extra": "9632 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 312409,
            "unit": "ns/op\t  367481 B/op\t    4713 allocs/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 312409,
            "unit": "ns/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367481,
            "unit": "B/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3657 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 334374,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 334374,
            "unit": "ns/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3445 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 665176,
            "unit": "ns/op\t  780069 B/op\t    9776 allocs/op",
            "extra": "1864 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 665176,
            "unit": "ns/op",
            "extra": "1864 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780069,
            "unit": "B/op",
            "extra": "1864 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 75857,
            "unit": "ns/op\t   72330 B/op\t     842 allocs/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 75857,
            "unit": "ns/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72330,
            "unit": "B/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15723 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 83286,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 83286,
            "unit": "ns/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14348 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 464944,
            "unit": "ns/op\t  460897 B/op\t    5506 allocs/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 464944,
            "unit": "ns/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 460897,
            "unit": "B/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5506,
            "unit": "allocs/op",
            "extra": "2714 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1455242,
            "unit": "ns/op\t 1423968 B/op\t   17565 allocs/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1455242,
            "unit": "ns/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423968,
            "unit": "B/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "762 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2381,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2381,
            "unit": "ns/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "528789 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 396707,
            "unit": "ns/op\t  363142 B/op\t    4626 allocs/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 396707,
            "unit": "ns/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363142,
            "unit": "B/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2950 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3787,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3787,
            "unit": "ns/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "307534 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 743448,
            "unit": "ns/op\t  690733 B/op\t    8636 allocs/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 743448,
            "unit": "ns/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690733,
            "unit": "B/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1556 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 163984,
            "unit": "ns/op\t  152132 B/op\t    1878 allocs/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 163984,
            "unit": "ns/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 152132,
            "unit": "B/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1878,
            "unit": "allocs/op",
            "extra": "6570 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 407033,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 407033,
            "unit": "ns/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2794 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 433858,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 433858,
            "unit": "ns/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2664 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 812007,
            "unit": "ns/op\t  780043 B/op\t    9776 allocs/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 812007,
            "unit": "ns/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780043,
            "unit": "B/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 75784,
            "unit": "ns/op\t   72255 B/op\t     842 allocs/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 75784,
            "unit": "ns/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72255,
            "unit": "B/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15549 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 83660,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 83660,
            "unit": "ns/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14313 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 472656,
            "unit": "ns/op\t  467909 B/op\t    5596 allocs/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 472656,
            "unit": "ns/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 467909,
            "unit": "B/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5596,
            "unit": "allocs/op",
            "extra": "2305 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1437365,
            "unit": "ns/op\t 1423914 B/op\t   17565 allocs/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1437365,
            "unit": "ns/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423914,
            "unit": "B/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2278,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2278,
            "unit": "ns/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "520951 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 398266,
            "unit": "ns/op\t  363135 B/op\t    4625 allocs/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 398266,
            "unit": "ns/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363135,
            "unit": "B/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "2804 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3733,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3733,
            "unit": "ns/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "321158 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 747728,
            "unit": "ns/op\t  690735 B/op\t    8636 allocs/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 747728,
            "unit": "ns/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690735,
            "unit": "B/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1575 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 168132,
            "unit": "ns/op\t  156331 B/op\t    1930 allocs/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 168132,
            "unit": "ns/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 156331,
            "unit": "B/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1930,
            "unit": "allocs/op",
            "extra": "6567 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 409419,
            "unit": "ns/op\t  367475 B/op\t    4713 allocs/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 409419,
            "unit": "ns/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367475,
            "unit": "B/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2766 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 439123,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 439123,
            "unit": "ns/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2590 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 812789,
            "unit": "ns/op\t  780037 B/op\t    9776 allocs/op",
            "extra": "1438 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 812789,
            "unit": "ns/op",
            "extra": "1438 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780037,
            "unit": "B/op",
            "extra": "1438 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 72379,
            "unit": "ns/op\t   72349 B/op\t     842 allocs/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 72379,
            "unit": "ns/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72349,
            "unit": "B/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 77953,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 77953,
            "unit": "ns/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14356 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 449531,
            "unit": "ns/op\t  469219 B/op\t    5609 allocs/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 449531,
            "unit": "ns/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 469219,
            "unit": "B/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5609,
            "unit": "allocs/op",
            "extra": "2437 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1356680,
            "unit": "ns/op\t 1423955 B/op\t   17565 allocs/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1356680,
            "unit": "ns/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423955,
            "unit": "B/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2217,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2217,
            "unit": "ns/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "537618 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 373287,
            "unit": "ns/op\t  363148 B/op\t    4626 allocs/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 373287,
            "unit": "ns/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363148,
            "unit": "B/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2799 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3488,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3488,
            "unit": "ns/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "330159 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 741881,
            "unit": "ns/op\t  690753 B/op\t    8636 allocs/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 741881,
            "unit": "ns/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690753,
            "unit": "B/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 160252,
            "unit": "ns/op\t  158836 B/op\t    1961 allocs/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 160252,
            "unit": "ns/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 158836,
            "unit": "B/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1961,
            "unit": "allocs/op",
            "extra": "7670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 378908,
            "unit": "ns/op\t  367479 B/op\t    4713 allocs/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 378908,
            "unit": "ns/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367479,
            "unit": "B/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3085 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 402110,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 402110,
            "unit": "ns/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2934 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 765148,
            "unit": "ns/op\t  780053 B/op\t    9776 allocs/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 765148,
            "unit": "ns/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780053,
            "unit": "B/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 76658,
            "unit": "ns/op\t   72340 B/op\t     842 allocs/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 76658,
            "unit": "ns/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72340,
            "unit": "B/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 84277,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 84277,
            "unit": "ns/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "12901 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 466899,
            "unit": "ns/op\t  460963 B/op\t    5507 allocs/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 466899,
            "unit": "ns/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 460963,
            "unit": "B/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5507,
            "unit": "allocs/op",
            "extra": "2485 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1442848,
            "unit": "ns/op\t 1423951 B/op\t   17565 allocs/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1442848,
            "unit": "ns/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423951,
            "unit": "B/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "816 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2289,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2289,
            "unit": "ns/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "522640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 435242,
            "unit": "ns/op\t  363143 B/op\t    4626 allocs/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 435242,
            "unit": "ns/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363143,
            "unit": "B/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2866 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3749,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3749,
            "unit": "ns/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "308682 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 745594,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 745594,
            "unit": "ns/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1492 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 166993,
            "unit": "ns/op\t  155542 B/op\t    1921 allocs/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 166993,
            "unit": "ns/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 155542,
            "unit": "B/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1921,
            "unit": "allocs/op",
            "extra": "6082 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 406212,
            "unit": "ns/op\t  367476 B/op\t    4713 allocs/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 406212,
            "unit": "ns/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367476,
            "unit": "B/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2847 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 436375,
            "unit": "ns/op\t  414440 B/op\t    4981 allocs/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 436375,
            "unit": "ns/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414440,
            "unit": "B/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2652 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 811315,
            "unit": "ns/op\t  780045 B/op\t    9776 allocs/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 811315,
            "unit": "ns/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780045,
            "unit": "B/op",
            "extra": "1450 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 84654,
            "unit": "ns/op\t   72324 B/op\t     842 allocs/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 84654,
            "unit": "ns/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72324,
            "unit": "B/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "14674 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 90183,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 90183,
            "unit": "ns/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "12999 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 548187,
            "unit": "ns/op\t  467572 B/op\t    5582 allocs/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 548187,
            "unit": "ns/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 467572,
            "unit": "B/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5582,
            "unit": "allocs/op",
            "extra": "2245 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1519323,
            "unit": "ns/op\t 1423921 B/op\t   17565 allocs/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1519323,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423921,
            "unit": "B/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2343,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2343,
            "unit": "ns/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "479109 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 421546,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 421546,
            "unit": "ns/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2712 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3886,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3886,
            "unit": "ns/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "307851 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 792977,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 792977,
            "unit": "ns/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1478 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 184938,
            "unit": "ns/op\t  158848 B/op\t    1962 allocs/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 184938,
            "unit": "ns/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 158848,
            "unit": "B/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1962,
            "unit": "allocs/op",
            "extra": "5733 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 440406,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 440406,
            "unit": "ns/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2720 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 468311,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 468311,
            "unit": "ns/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2488 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 878713,
            "unit": "ns/op\t  780054 B/op\t    9776 allocs/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 878713,
            "unit": "ns/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780054,
            "unit": "B/op",
            "extra": "1329 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 79018,
            "unit": "ns/op\t   72273 B/op\t     842 allocs/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 79018,
            "unit": "ns/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72273,
            "unit": "B/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 84215,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 84215,
            "unit": "ns/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13142 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 458830,
            "unit": "ns/op\t  460819 B/op\t    5515 allocs/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 458830,
            "unit": "ns/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 460819,
            "unit": "B/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5515,
            "unit": "allocs/op",
            "extra": "2470 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1432276,
            "unit": "ns/op\t 1423908 B/op\t   17565 allocs/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1432276,
            "unit": "ns/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423908,
            "unit": "B/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2265,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2265,
            "unit": "ns/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "531958 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 393074,
            "unit": "ns/op\t  363133 B/op\t    4625 allocs/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 393074,
            "unit": "ns/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363133,
            "unit": "B/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "2923 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3629,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3629,
            "unit": "ns/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "331550 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 731541,
            "unit": "ns/op\t  690736 B/op\t    8636 allocs/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 731541,
            "unit": "ns/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690736,
            "unit": "B/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1566 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 167722,
            "unit": "ns/op\t  157166 B/op\t    1941 allocs/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 167722,
            "unit": "ns/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 157166,
            "unit": "B/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1941,
            "unit": "allocs/op",
            "extra": "6649 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 401169,
            "unit": "ns/op\t  367468 B/op\t    4712 allocs/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 401169,
            "unit": "ns/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367468,
            "unit": "B/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "2836 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 426929,
            "unit": "ns/op\t  414436 B/op\t    4981 allocs/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 426929,
            "unit": "ns/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414436,
            "unit": "B/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2671 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 878857,
            "unit": "ns/op\t  780061 B/op\t    9776 allocs/op",
            "extra": "1419 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 878857,
            "unit": "ns/op",
            "extra": "1419 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780061,
            "unit": "B/op",
            "extra": "1419 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 49704,
            "unit": "ns/op\t   72432 B/op\t     842 allocs/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 49704,
            "unit": "ns/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72432,
            "unit": "B/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "24484 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 57437,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 57437,
            "unit": "ns/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "19268 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 311418,
            "unit": "ns/op\t  463201 B/op\t    5529 allocs/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 311418,
            "unit": "ns/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 463201,
            "unit": "B/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5529,
            "unit": "allocs/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1066296,
            "unit": "ns/op\t 1423961 B/op\t   17565 allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1066296,
            "unit": "ns/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423961,
            "unit": "B/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1484,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1484,
            "unit": "ns/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "698611 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 258837,
            "unit": "ns/op\t  363150 B/op\t    4626 allocs/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 258837,
            "unit": "ns/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363150,
            "unit": "B/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "4179 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2633,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2633,
            "unit": "ns/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "441661 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 483067,
            "unit": "ns/op\t  690758 B/op\t    8636 allocs/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 483067,
            "unit": "ns/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690758,
            "unit": "B/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2370 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 113752,
            "unit": "ns/op\t  163545 B/op\t    2019 allocs/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 113752,
            "unit": "ns/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 163545,
            "unit": "B/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 2019,
            "unit": "allocs/op",
            "extra": "9637 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 265344,
            "unit": "ns/op\t  367485 B/op\t    4713 allocs/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 265344,
            "unit": "ns/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367485,
            "unit": "B/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "4503 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 279992,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 279992,
            "unit": "ns/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 531199,
            "unit": "ns/op\t  780060 B/op\t    9776 allocs/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 531199,
            "unit": "ns/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780060,
            "unit": "B/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 49649,
            "unit": "ns/op\t   72446 B/op\t     842 allocs/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 49649,
            "unit": "ns/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72446,
            "unit": "B/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "24147 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 53647,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 53647,
            "unit": "ns/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "22368 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 332033,
            "unit": "ns/op\t  472409 B/op\t    5643 allocs/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 332033,
            "unit": "ns/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 472409,
            "unit": "B/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5643,
            "unit": "allocs/op",
            "extra": "3516 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1045585,
            "unit": "ns/op\t 1423949 B/op\t   17565 allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1045585,
            "unit": "ns/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423949,
            "unit": "B/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1670,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1670,
            "unit": "ns/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "711745 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 275893,
            "unit": "ns/op\t  363146 B/op\t    4626 allocs/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 275893,
            "unit": "ns/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363146,
            "unit": "B/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3879 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2655,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2655,
            "unit": "ns/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "413275 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 487443,
            "unit": "ns/op\t  690756 B/op\t    8636 allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 487443,
            "unit": "ns/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690756,
            "unit": "B/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 111865,
            "unit": "ns/op\t  156119 B/op\t    1928 allocs/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 111865,
            "unit": "ns/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 156119,
            "unit": "B/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1928,
            "unit": "allocs/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 294202,
            "unit": "ns/op\t  367486 B/op\t    4713 allocs/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 294202,
            "unit": "ns/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367486,
            "unit": "B/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "4152 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 297468,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 297468,
            "unit": "ns/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3554 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 556957,
            "unit": "ns/op\t  780067 B/op\t    9776 allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 556957,
            "unit": "ns/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780067,
            "unit": "B/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 76068,
            "unit": "ns/op\t   72361 B/op\t     842 allocs/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 76068,
            "unit": "ns/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72361,
            "unit": "B/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15464 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 83765,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 83765,
            "unit": "ns/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14364 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 467408,
            "unit": "ns/op\t  459893 B/op\t    5495 allocs/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 467408,
            "unit": "ns/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 459893,
            "unit": "B/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5495,
            "unit": "allocs/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1455827,
            "unit": "ns/op\t 1423944 B/op\t   17565 allocs/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1455827,
            "unit": "ns/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423944,
            "unit": "B/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "804 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2284,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2284,
            "unit": "ns/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "518713 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 405434,
            "unit": "ns/op\t  363145 B/op\t    4626 allocs/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 405434,
            "unit": "ns/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363145,
            "unit": "B/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2802 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 4035,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 4035,
            "unit": "ns/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "302955 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 745650,
            "unit": "ns/op\t  690736 B/op\t    8636 allocs/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 745650,
            "unit": "ns/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690736,
            "unit": "B/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 162383,
            "unit": "ns/op\t  151037 B/op\t    1865 allocs/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 162383,
            "unit": "ns/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 151037,
            "unit": "B/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1865,
            "unit": "allocs/op",
            "extra": "6933 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 409411,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 409411,
            "unit": "ns/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2725 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 437127,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 437127,
            "unit": "ns/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2631 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 816331,
            "unit": "ns/op\t  780050 B/op\t    9776 allocs/op",
            "extra": "1460 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 816331,
            "unit": "ns/op",
            "extra": "1460 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780050,
            "unit": "B/op",
            "extra": "1460 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 53086,
            "unit": "ns/op\t   72371 B/op\t     842 allocs/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 53086,
            "unit": "ns/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72371,
            "unit": "B/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "21873 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 60895,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 60895,
            "unit": "ns/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "19848 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 374523,
            "unit": "ns/op\t  474551 B/op\t    5669 allocs/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 374523,
            "unit": "ns/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 474551,
            "unit": "B/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5669,
            "unit": "allocs/op",
            "extra": "2924 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1088041,
            "unit": "ns/op\t 1423944 B/op\t   17565 allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1088041,
            "unit": "ns/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423944,
            "unit": "B/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1798,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1798,
            "unit": "ns/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "662224 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 292306,
            "unit": "ns/op\t  363148 B/op\t    4626 allocs/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 292306,
            "unit": "ns/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363148,
            "unit": "B/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3636 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2728,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2728,
            "unit": "ns/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "440022 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 536868,
            "unit": "ns/op\t  690743 B/op\t    8636 allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 536868,
            "unit": "ns/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690743,
            "unit": "B/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 123419,
            "unit": "ns/op\t  161595 B/op\t    1995 allocs/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 123419,
            "unit": "ns/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 161595,
            "unit": "B/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1995,
            "unit": "allocs/op",
            "extra": "9660 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 292650,
            "unit": "ns/op\t  367480 B/op\t    4713 allocs/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 292650,
            "unit": "ns/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367480,
            "unit": "B/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3919 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 329700,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 329700,
            "unit": "ns/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3674 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 588664,
            "unit": "ns/op\t  780052 B/op\t    9776 allocs/op",
            "extra": "1968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 588664,
            "unit": "ns/op",
            "extra": "1968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780052,
            "unit": "B/op",
            "extra": "1968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 78890,
            "unit": "ns/op\t   72306 B/op\t     842 allocs/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 78890,
            "unit": "ns/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72306,
            "unit": "B/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 82662,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 82662,
            "unit": "ns/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14476 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 462641,
            "unit": "ns/op\t  463815 B/op\t    5547 allocs/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 462641,
            "unit": "ns/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 463815,
            "unit": "B/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5547,
            "unit": "allocs/op",
            "extra": "2610 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1433584,
            "unit": "ns/op\t 1423927 B/op\t   17565 allocs/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1433584,
            "unit": "ns/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423927,
            "unit": "B/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "834 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2287,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2287,
            "unit": "ns/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "529988 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 404251,
            "unit": "ns/op\t  363144 B/op\t    4626 allocs/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 404251,
            "unit": "ns/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363144,
            "unit": "B/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2977 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3768,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3768,
            "unit": "ns/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "306123 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 740256,
            "unit": "ns/op\t  690729 B/op\t    8636 allocs/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 740256,
            "unit": "ns/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690729,
            "unit": "B/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1576 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 174198,
            "unit": "ns/op\t  158691 B/op\t    1960 allocs/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 174198,
            "unit": "ns/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 158691,
            "unit": "B/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1960,
            "unit": "allocs/op",
            "extra": "6948 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 412544,
            "unit": "ns/op\t  367482 B/op\t    4713 allocs/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 412544,
            "unit": "ns/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367482,
            "unit": "B/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2895 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 435909,
            "unit": "ns/op\t  414436 B/op\t    4981 allocs/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 435909,
            "unit": "ns/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414436,
            "unit": "B/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2600 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 808986,
            "unit": "ns/op\t  780043 B/op\t    9776 allocs/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 808986,
            "unit": "ns/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780043,
            "unit": "B/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 66452,
            "unit": "ns/op\t   72408 B/op\t     842 allocs/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 66452,
            "unit": "ns/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72408,
            "unit": "B/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 73128,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 73128,
            "unit": "ns/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "16287 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 420033,
            "unit": "ns/op\t  460798 B/op\t    5498 allocs/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 420033,
            "unit": "ns/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 460798,
            "unit": "B/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5498,
            "unit": "allocs/op",
            "extra": "2580 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1303760,
            "unit": "ns/op\t 1423935 B/op\t   17565 allocs/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1303760,
            "unit": "ns/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423935,
            "unit": "B/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2193,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2193,
            "unit": "ns/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "534564 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 361633,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 361633,
            "unit": "ns/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3243 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3540,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3540,
            "unit": "ns/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "322434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 659975,
            "unit": "ns/op\t  690744 B/op\t    8636 allocs/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 659975,
            "unit": "ns/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690744,
            "unit": "B/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1759 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 155014,
            "unit": "ns/op\t  161567 B/op\t    1995 allocs/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 155014,
            "unit": "ns/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 161567,
            "unit": "B/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1995,
            "unit": "allocs/op",
            "extra": "6685 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 377254,
            "unit": "ns/op\t  367477 B/op\t    4713 allocs/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 377254,
            "unit": "ns/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367477,
            "unit": "B/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3259 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 387549,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 387549,
            "unit": "ns/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3045 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 737703,
            "unit": "ns/op\t  780061 B/op\t    9776 allocs/op",
            "extra": "1639 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 737703,
            "unit": "ns/op",
            "extra": "1639 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780061,
            "unit": "B/op",
            "extra": "1639 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 42546,
            "unit": "ns/op\t   72335 B/op\t     842 allocs/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 42546,
            "unit": "ns/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72335,
            "unit": "B/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "27465 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 46540,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 46540,
            "unit": "ns/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "26030 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 260679,
            "unit": "ns/op\t  461625 B/op\t    5517 allocs/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 260679,
            "unit": "ns/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 461625,
            "unit": "B/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5517,
            "unit": "allocs/op",
            "extra": "4362 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 812418,
            "unit": "ns/op\t 1423943 B/op\t   17565 allocs/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 812418,
            "unit": "ns/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423943,
            "unit": "B/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1299,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1299,
            "unit": "ns/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "943954 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 215825,
            "unit": "ns/op\t  363144 B/op\t    4626 allocs/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 215825,
            "unit": "ns/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363144,
            "unit": "B/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "4984 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2137,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2137,
            "unit": "ns/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "540741 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 432650,
            "unit": "ns/op\t  690749 B/op\t    8636 allocs/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 432650,
            "unit": "ns/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690749,
            "unit": "B/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2874 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 94910,
            "unit": "ns/op\t  160874 B/op\t    1987 allocs/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 94910,
            "unit": "ns/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 160874,
            "unit": "B/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1987,
            "unit": "allocs/op",
            "extra": "12349 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 225199,
            "unit": "ns/op\t  367475 B/op\t    4713 allocs/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 225199,
            "unit": "ns/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367475,
            "unit": "B/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "5257 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 237693,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 237693,
            "unit": "ns/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "4732 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 446768,
            "unit": "ns/op\t  780051 B/op\t    9776 allocs/op",
            "extra": "2629 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 446768,
            "unit": "ns/op",
            "extra": "2629 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780051,
            "unit": "B/op",
            "extra": "2629 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 77453,
            "unit": "ns/op\t   72374 B/op\t     842 allocs/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 77453,
            "unit": "ns/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72374,
            "unit": "B/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15321 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 84182,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 84182,
            "unit": "ns/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13966 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 494330,
            "unit": "ns/op\t  460234 B/op\t    5493 allocs/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 494330,
            "unit": "ns/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 460234,
            "unit": "B/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5493,
            "unit": "allocs/op",
            "extra": "2554 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1528126,
            "unit": "ns/op\t 1423950 B/op\t   17565 allocs/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1528126,
            "unit": "ns/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423950,
            "unit": "B/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2335,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2335,
            "unit": "ns/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "506782 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 427791,
            "unit": "ns/op\t  363150 B/op\t    4626 allocs/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 427791,
            "unit": "ns/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363150,
            "unit": "B/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2840 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3946,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3946,
            "unit": "ns/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "280096 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 744008,
            "unit": "ns/op\t  690747 B/op\t    8636 allocs/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 744008,
            "unit": "ns/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690747,
            "unit": "B/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 162812,
            "unit": "ns/op\t  149924 B/op\t    1852 allocs/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 162812,
            "unit": "ns/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 149924,
            "unit": "B/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1852,
            "unit": "allocs/op",
            "extra": "8230 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 412524,
            "unit": "ns/op\t  367478 B/op\t    4713 allocs/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 412524,
            "unit": "ns/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367478,
            "unit": "B/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2670 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 441059,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 441059,
            "unit": "ns/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2641 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 826624,
            "unit": "ns/op\t  780047 B/op\t    9776 allocs/op",
            "extra": "1395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 826624,
            "unit": "ns/op",
            "extra": "1395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780047,
            "unit": "B/op",
            "extra": "1395 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 39395,
            "unit": "ns/op\t   72329 B/op\t     842 allocs/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 39395,
            "unit": "ns/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72329,
            "unit": "B/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "30172 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 41033,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 41033,
            "unit": "ns/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "28858 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 244237,
            "unit": "ns/op\t  461206 B/op\t    5506 allocs/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 244237,
            "unit": "ns/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 461206,
            "unit": "B/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5506,
            "unit": "allocs/op",
            "extra": "4447 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 747987,
            "unit": "ns/op\t 1423946 B/op\t   17565 allocs/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 747987,
            "unit": "ns/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423946,
            "unit": "B/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1267,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1267,
            "unit": "ns/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "960627 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 205333,
            "unit": "ns/op\t  363142 B/op\t    4626 allocs/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 205333,
            "unit": "ns/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363142,
            "unit": "B/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "5493 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 1968,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 1968,
            "unit": "ns/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "612830 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 401589,
            "unit": "ns/op\t  690749 B/op\t    8636 allocs/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 401589,
            "unit": "ns/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690749,
            "unit": "B/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "3069 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 85666,
            "unit": "ns/op\t  156698 B/op\t    1935 allocs/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 85666,
            "unit": "ns/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 156698,
            "unit": "B/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1935,
            "unit": "allocs/op",
            "extra": "13543 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 217564,
            "unit": "ns/op\t  367472 B/op\t    4712 allocs/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 217564,
            "unit": "ns/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367472,
            "unit": "B/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "5503 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 228722,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 228722,
            "unit": "ns/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "5187 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 424649,
            "unit": "ns/op\t  780062 B/op\t    9776 allocs/op",
            "extra": "2835 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 424649,
            "unit": "ns/op",
            "extra": "2835 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780062,
            "unit": "B/op",
            "extra": "2835 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 75272,
            "unit": "ns/op\t   72336 B/op\t     842 allocs/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 75272,
            "unit": "ns/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72336,
            "unit": "B/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "16178 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 84463,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 84463,
            "unit": "ns/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13032 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 498004,
            "unit": "ns/op\t  461341 B/op\t    5510 allocs/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 498004,
            "unit": "ns/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 461341,
            "unit": "B/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5510,
            "unit": "allocs/op",
            "extra": "2464 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1456165,
            "unit": "ns/op\t 1423949 B/op\t   17565 allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1456165,
            "unit": "ns/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423949,
            "unit": "B/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2272,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2272,
            "unit": "ns/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "528171 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 404189,
            "unit": "ns/op\t  363143 B/op\t    4626 allocs/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 404189,
            "unit": "ns/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363143,
            "unit": "B/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2926 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3600,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3600,
            "unit": "ns/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "324565 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 731502,
            "unit": "ns/op\t  690736 B/op\t    8636 allocs/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 731502,
            "unit": "ns/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690736,
            "unit": "B/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1573 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 168327,
            "unit": "ns/op\t  160011 B/op\t    1976 allocs/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 168327,
            "unit": "ns/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 160011,
            "unit": "B/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1976,
            "unit": "allocs/op",
            "extra": "7047 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 408192,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 408192,
            "unit": "ns/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2806 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 426897,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 426897,
            "unit": "ns/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2689 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 806943,
            "unit": "ns/op\t  780039 B/op\t    9776 allocs/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 806943,
            "unit": "ns/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780039,
            "unit": "B/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 38275,
            "unit": "ns/op\t   72299 B/op\t     842 allocs/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 38275,
            "unit": "ns/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72299,
            "unit": "B/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "30590 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 44137,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 44137,
            "unit": "ns/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "27606 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 245548,
            "unit": "ns/op\t  461746 B/op\t    5517 allocs/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 245548,
            "unit": "ns/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 461746,
            "unit": "B/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5517,
            "unit": "allocs/op",
            "extra": "4911 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 775285,
            "unit": "ns/op\t 1423944 B/op\t   17565 allocs/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 775285,
            "unit": "ns/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423944,
            "unit": "B/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1557 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1235,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1235,
            "unit": "ns/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "983042 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 208941,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 208941,
            "unit": "ns/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "5640 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2032,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2032,
            "unit": "ns/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "604741 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 422210,
            "unit": "ns/op\t  690753 B/op\t    8636 allocs/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 422210,
            "unit": "ns/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690753,
            "unit": "B/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 88340,
            "unit": "ns/op\t  156623 B/op\t    1934 allocs/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 88340,
            "unit": "ns/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 156623,
            "unit": "B/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1934,
            "unit": "allocs/op",
            "extra": "13683 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 223057,
            "unit": "ns/op\t  367481 B/op\t    4713 allocs/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 223057,
            "unit": "ns/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367481,
            "unit": "B/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "5184 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 233107,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 233107,
            "unit": "ns/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "4861 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 457010,
            "unit": "ns/op\t  780052 B/op\t    9776 allocs/op",
            "extra": "2487 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 457010,
            "unit": "ns/op",
            "extra": "2487 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780052,
            "unit": "B/op",
            "extra": "2487 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 78303,
            "unit": "ns/op\t   72325 B/op\t     842 allocs/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 78303,
            "unit": "ns/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72325,
            "unit": "B/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "14960 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 85513,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 85513,
            "unit": "ns/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13790 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 495778,
            "unit": "ns/op\t  480207 B/op\t    5745 allocs/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 495778,
            "unit": "ns/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 480207,
            "unit": "B/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5745,
            "unit": "allocs/op",
            "extra": "2514 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1489553,
            "unit": "ns/op\t 1423922 B/op\t   17565 allocs/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1489553,
            "unit": "ns/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423922,
            "unit": "B/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "810 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2320,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2320,
            "unit": "ns/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "508480 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 408165,
            "unit": "ns/op\t  363141 B/op\t    4626 allocs/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 408165,
            "unit": "ns/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363141,
            "unit": "B/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2701 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3806,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3806,
            "unit": "ns/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "312322 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 762826,
            "unit": "ns/op\t  690742 B/op\t    8636 allocs/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 762826,
            "unit": "ns/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690742,
            "unit": "B/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1504 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 172445,
            "unit": "ns/op\t  151076 B/op\t    1866 allocs/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 172445,
            "unit": "ns/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 151076,
            "unit": "B/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1866,
            "unit": "allocs/op",
            "extra": "7016 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 418549,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 418549,
            "unit": "ns/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2749 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 443130,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 443130,
            "unit": "ns/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2538 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 833942,
            "unit": "ns/op\t  780048 B/op\t    9776 allocs/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 833942,
            "unit": "ns/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780048,
            "unit": "B/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1400 times\n4 procs"
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
        "date": 1791121040635,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 76659,
            "unit": "ns/op\t   72305 B/op\t     842 allocs/op",
            "extra": "15295 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 76659,
            "unit": "ns/op",
            "extra": "15295 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72305,
            "unit": "B/op",
            "extra": "15295 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15295 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 83819,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "14322 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 83819,
            "unit": "ns/op",
            "extra": "14322 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "14322 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14322 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 488707,
            "unit": "ns/op\t  473256 B/op\t    5655 allocs/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 488707,
            "unit": "ns/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 473256,
            "unit": "B/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5655,
            "unit": "allocs/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1449757,
            "unit": "ns/op\t 1423901 B/op\t   17565 allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1449757,
            "unit": "ns/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423901,
            "unit": "B/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2308,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "499287 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2308,
            "unit": "ns/op",
            "extra": "499287 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "499287 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "499287 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 436759,
            "unit": "ns/op\t  363150 B/op\t    4626 allocs/op",
            "extra": "2910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 436759,
            "unit": "ns/op",
            "extra": "2910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363150,
            "unit": "B/op",
            "extra": "2910 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2910 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3788,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "315051 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3788,
            "unit": "ns/op",
            "extra": "315051 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "315051 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "315051 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 747846,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1524 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 747846,
            "unit": "ns/op",
            "extra": "1524 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1524 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1524 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 163773,
            "unit": "ns/op\t  153164 B/op\t    1891 allocs/op",
            "extra": "6270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 163773,
            "unit": "ns/op",
            "extra": "6270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 153164,
            "unit": "B/op",
            "extra": "6270 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1891,
            "unit": "allocs/op",
            "extra": "6270 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 408843,
            "unit": "ns/op\t  367478 B/op\t    4713 allocs/op",
            "extra": "2786 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 408843,
            "unit": "ns/op",
            "extra": "2786 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367478,
            "unit": "B/op",
            "extra": "2786 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2786 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 436720,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "2746 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 436720,
            "unit": "ns/op",
            "extra": "2746 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "2746 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2746 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 812100,
            "unit": "ns/op\t  780045 B/op\t    9776 allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 812100,
            "unit": "ns/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780045,
            "unit": "B/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1434 times\n4 procs"
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
        "date": 1791218558076,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 74484,
            "unit": "ns/op\t   72235 B/op\t     842 allocs/op",
            "extra": "15703 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 74484,
            "unit": "ns/op",
            "extra": "15703 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72235,
            "unit": "B/op",
            "extra": "15703 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15703 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 83167,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "13738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 83167,
            "unit": "ns/op",
            "extra": "13738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "13738 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "13738 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 474231,
            "unit": "ns/op\t  473097 B/op\t    5649 allocs/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 474231,
            "unit": "ns/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 473097,
            "unit": "B/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5649,
            "unit": "allocs/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1436385,
            "unit": "ns/op\t 1423981 B/op\t   17565 allocs/op",
            "extra": "842 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1436385,
            "unit": "ns/op",
            "extra": "842 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423981,
            "unit": "B/op",
            "extra": "842 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "842 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2252,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "528193 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2252,
            "unit": "ns/op",
            "extra": "528193 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "528193 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "528193 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 393475,
            "unit": "ns/op\t  363142 B/op\t    4626 allocs/op",
            "extra": "2928 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 393475,
            "unit": "ns/op",
            "extra": "2928 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363142,
            "unit": "B/op",
            "extra": "2928 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2928 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3760,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "312513 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3760,
            "unit": "ns/op",
            "extra": "312513 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "312513 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "312513 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 735782,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1605 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 735782,
            "unit": "ns/op",
            "extra": "1605 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1605 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1605 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 169018,
            "unit": "ns/op\t  159472 B/op\t    1969 allocs/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 169018,
            "unit": "ns/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 159472,
            "unit": "B/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1969,
            "unit": "allocs/op",
            "extra": "6961 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 403353,
            "unit": "ns/op\t  367481 B/op\t    4713 allocs/op",
            "extra": "2955 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 403353,
            "unit": "ns/op",
            "extra": "2955 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367481,
            "unit": "B/op",
            "extra": "2955 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2955 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 495431,
            "unit": "ns/op\t  414441 B/op\t    4981 allocs/op",
            "extra": "2750 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 495431,
            "unit": "ns/op",
            "extra": "2750 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414441,
            "unit": "B/op",
            "extra": "2750 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2750 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 809456,
            "unit": "ns/op\t  780066 B/op\t    9776 allocs/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 809456,
            "unit": "ns/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780066,
            "unit": "B/op",
            "extra": "1431 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1431 times\n4 procs"
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
        "date": 1791297599406,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 68291,
            "unit": "ns/op\t   72363 B/op\t     842 allocs/op",
            "extra": "17161 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 68291,
            "unit": "ns/op",
            "extra": "17161 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72363,
            "unit": "B/op",
            "extra": "17161 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17161 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 75939,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14946 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 75939,
            "unit": "ns/op",
            "extra": "14946 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14946 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14946 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 438818,
            "unit": "ns/op\t  462219 B/op\t    5523 allocs/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 438818,
            "unit": "ns/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 462219,
            "unit": "B/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5523,
            "unit": "allocs/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1368203,
            "unit": "ns/op\t 1423956 B/op\t   17565 allocs/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1368203,
            "unit": "ns/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423956,
            "unit": "B/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2217,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "544522 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2217,
            "unit": "ns/op",
            "extra": "544522 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "544522 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "544522 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 372721,
            "unit": "ns/op\t  363147 B/op\t    4626 allocs/op",
            "extra": "3036 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 372721,
            "unit": "ns/op",
            "extra": "3036 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363147,
            "unit": "B/op",
            "extra": "3036 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3036 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3483,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "331882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3483,
            "unit": "ns/op",
            "extra": "331882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "331882 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "331882 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 689559,
            "unit": "ns/op\t  690741 B/op\t    8636 allocs/op",
            "extra": "1726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 689559,
            "unit": "ns/op",
            "extra": "1726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690741,
            "unit": "B/op",
            "extra": "1726 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1726 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 160020,
            "unit": "ns/op\t  158283 B/op\t    1955 allocs/op",
            "extra": "6586 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 160020,
            "unit": "ns/op",
            "extra": "6586 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 158283,
            "unit": "B/op",
            "extra": "6586 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1955,
            "unit": "allocs/op",
            "extra": "6586 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 380859,
            "unit": "ns/op\t  367481 B/op\t    4713 allocs/op",
            "extra": "3092 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 380859,
            "unit": "ns/op",
            "extra": "3092 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367481,
            "unit": "B/op",
            "extra": "3092 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3092 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 404935,
            "unit": "ns/op\t  414439 B/op\t    4981 allocs/op",
            "extra": "2857 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 404935,
            "unit": "ns/op",
            "extra": "2857 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414439,
            "unit": "B/op",
            "extra": "2857 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2857 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 760240,
            "unit": "ns/op\t  780060 B/op\t    9776 allocs/op",
            "extra": "1510 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 760240,
            "unit": "ns/op",
            "extra": "1510 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780060,
            "unit": "B/op",
            "extra": "1510 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1510 times\n4 procs"
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
        "date": 1791385198690,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 40098,
            "unit": "ns/op\t   72255 B/op\t     842 allocs/op",
            "extra": "30181 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 40098,
            "unit": "ns/op",
            "extra": "30181 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72255,
            "unit": "B/op",
            "extra": "30181 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "30181 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 46791,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "25446 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 46791,
            "unit": "ns/op",
            "extra": "25446 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "25446 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "25446 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 252830,
            "unit": "ns/op\t  471218 B/op\t    5625 allocs/op",
            "extra": "4315 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 252830,
            "unit": "ns/op",
            "extra": "4315 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 471218,
            "unit": "B/op",
            "extra": "4315 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5625,
            "unit": "allocs/op",
            "extra": "4315 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 797362,
            "unit": "ns/op\t 1423978 B/op\t   17565 allocs/op",
            "extra": "1567 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 797362,
            "unit": "ns/op",
            "extra": "1567 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423978,
            "unit": "B/op",
            "extra": "1567 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1567 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1226,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "956476 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1226,
            "unit": "ns/op",
            "extra": "956476 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "956476 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "956476 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 210750,
            "unit": "ns/op\t  363150 B/op\t    4626 allocs/op",
            "extra": "5668 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 210750,
            "unit": "ns/op",
            "extra": "5668 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363150,
            "unit": "B/op",
            "extra": "5668 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "5668 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2063,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "541088 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2063,
            "unit": "ns/op",
            "extra": "541088 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "541088 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "541088 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 412237,
            "unit": "ns/op\t  690750 B/op\t    8636 allocs/op",
            "extra": "2925 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 412237,
            "unit": "ns/op",
            "extra": "2925 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690750,
            "unit": "B/op",
            "extra": "2925 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2925 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 83321,
            "unit": "ns/op\t  154045 B/op\t    1903 allocs/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 83321,
            "unit": "ns/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 154045,
            "unit": "B/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1903,
            "unit": "allocs/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 207207,
            "unit": "ns/op\t  367474 B/op\t    4713 allocs/op",
            "extra": "5595 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 207207,
            "unit": "ns/op",
            "extra": "5595 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367474,
            "unit": "B/op",
            "extra": "5595 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "5595 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 229225,
            "unit": "ns/op\t  414438 B/op\t    4981 allocs/op",
            "extra": "5107 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 229225,
            "unit": "ns/op",
            "extra": "5107 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414438,
            "unit": "B/op",
            "extra": "5107 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "5107 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 446202,
            "unit": "ns/op\t  780059 B/op\t    9776 allocs/op",
            "extra": "2676 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 446202,
            "unit": "ns/op",
            "extra": "2676 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780059,
            "unit": "B/op",
            "extra": "2676 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "2676 times\n4 procs"
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
        "date": 1791472107201,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 73253,
            "unit": "ns/op\t   72460 B/op\t     842 allocs/op",
            "extra": "17070 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 73253,
            "unit": "ns/op",
            "extra": "17070 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72460,
            "unit": "B/op",
            "extra": "17070 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17070 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 76963,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "15597 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 76963,
            "unit": "ns/op",
            "extra": "15597 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "15597 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "15597 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 443573,
            "unit": "ns/op\t  463153 B/op\t    5532 allocs/op",
            "extra": "2769 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 443573,
            "unit": "ns/op",
            "extra": "2769 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 463153,
            "unit": "B/op",
            "extra": "2769 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5532,
            "unit": "allocs/op",
            "extra": "2769 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1367398,
            "unit": "ns/op\t 1423950 B/op\t   17565 allocs/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1367398,
            "unit": "ns/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423950,
            "unit": "B/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2224,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "539029 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2224,
            "unit": "ns/op",
            "extra": "539029 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "539029 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "539029 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 369866,
            "unit": "ns/op\t  363143 B/op\t    4626 allocs/op",
            "extra": "2960 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 369866,
            "unit": "ns/op",
            "extra": "2960 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363143,
            "unit": "B/op",
            "extra": "2960 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2960 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3671,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "335326 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3671,
            "unit": "ns/op",
            "extra": "335326 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "335326 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "335326 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 681338,
            "unit": "ns/op\t  690742 B/op\t    8636 allocs/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 681338,
            "unit": "ns/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690742,
            "unit": "B/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 156077,
            "unit": "ns/op\t  157121 B/op\t    1941 allocs/op",
            "extra": "7588 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 156077,
            "unit": "ns/op",
            "extra": "7588 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 157121,
            "unit": "B/op",
            "extra": "7588 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1941,
            "unit": "allocs/op",
            "extra": "7588 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 379021,
            "unit": "ns/op\t  367475 B/op\t    4713 allocs/op",
            "extra": "2980 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 379021,
            "unit": "ns/op",
            "extra": "2980 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367475,
            "unit": "B/op",
            "extra": "2980 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2980 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 404664,
            "unit": "ns/op\t  414440 B/op\t    4981 allocs/op",
            "extra": "2913 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 404664,
            "unit": "ns/op",
            "extra": "2913 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414440,
            "unit": "B/op",
            "extra": "2913 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2913 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 762042,
            "unit": "ns/op\t  780069 B/op\t    9776 allocs/op",
            "extra": "1519 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 762042,
            "unit": "ns/op",
            "extra": "1519 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780069,
            "unit": "B/op",
            "extra": "1519 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1519 times\n4 procs"
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
        "date": 1791557636989,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 53781,
            "unit": "ns/op\t   72394 B/op\t     842 allocs/op",
            "extra": "21826 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 53781,
            "unit": "ns/op",
            "extra": "21826 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72394,
            "unit": "B/op",
            "extra": "21826 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "21826 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 62415,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "19657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 62415,
            "unit": "ns/op",
            "extra": "19657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "19657 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "19657 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 346508,
            "unit": "ns/op\t  457334 B/op\t    5463 allocs/op",
            "extra": "3003 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 346508,
            "unit": "ns/op",
            "extra": "3003 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 457334,
            "unit": "B/op",
            "extra": "3003 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5463,
            "unit": "allocs/op",
            "extra": "3003 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1084515,
            "unit": "ns/op\t 1423942 B/op\t   17565 allocs/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1084515,
            "unit": "ns/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423942,
            "unit": "B/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "1074 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1710,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "672493 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1710,
            "unit": "ns/op",
            "extra": "672493 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "672493 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "672493 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 287244,
            "unit": "ns/op\t  363141 B/op\t    4626 allocs/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 287244,
            "unit": "ns/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363141,
            "unit": "B/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "3944 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2751,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "423066 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2751,
            "unit": "ns/op",
            "extra": "423066 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "423066 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "423066 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 549550,
            "unit": "ns/op\t  690745 B/op\t    8636 allocs/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 549550,
            "unit": "ns/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690745,
            "unit": "B/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 124152,
            "unit": "ns/op\t  161609 B/op\t    1996 allocs/op",
            "extra": "9261 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 124152,
            "unit": "ns/op",
            "extra": "9261 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 161609,
            "unit": "B/op",
            "extra": "9261 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1996,
            "unit": "allocs/op",
            "extra": "9261 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 298530,
            "unit": "ns/op\t  367479 B/op\t    4713 allocs/op",
            "extra": "3874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 298530,
            "unit": "ns/op",
            "extra": "3874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367479,
            "unit": "B/op",
            "extra": "3874 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "3874 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 321350,
            "unit": "ns/op\t  414437 B/op\t    4981 allocs/op",
            "extra": "3530 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 321350,
            "unit": "ns/op",
            "extra": "3530 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414437,
            "unit": "B/op",
            "extra": "3530 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3530 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 604262,
            "unit": "ns/op\t  780063 B/op\t    9776 allocs/op",
            "extra": "1972 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 604262,
            "unit": "ns/op",
            "extra": "1972 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780063,
            "unit": "B/op",
            "extra": "1972 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1972 times\n4 procs"
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
        "date": 1791641364453,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 73939,
            "unit": "ns/op\t   72279 B/op\t     842 allocs/op",
            "extra": "15878 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 73939,
            "unit": "ns/op",
            "extra": "15878 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72279,
            "unit": "B/op",
            "extra": "15878 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "15878 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 81464,
            "unit": "ns/op\t   72150 B/op\t     955 allocs/op",
            "extra": "14696 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 81464,
            "unit": "ns/op",
            "extra": "14696 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72150,
            "unit": "B/op",
            "extra": "14696 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14696 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 482634,
            "unit": "ns/op\t  475845 B/op\t    5689 allocs/op",
            "extra": "2266 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 482634,
            "unit": "ns/op",
            "extra": "2266 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 475845,
            "unit": "B/op",
            "extra": "2266 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5689,
            "unit": "allocs/op",
            "extra": "2266 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1447614,
            "unit": "ns/op\t 1423973 B/op\t   17565 allocs/op",
            "extra": "825 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1447614,
            "unit": "ns/op",
            "extra": "825 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423973,
            "unit": "B/op",
            "extra": "825 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "825 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2264,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "521043 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2264,
            "unit": "ns/op",
            "extra": "521043 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "521043 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "521043 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 400244,
            "unit": "ns/op\t  363148 B/op\t    4626 allocs/op",
            "extra": "2791 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 400244,
            "unit": "ns/op",
            "extra": "2791 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363148,
            "unit": "B/op",
            "extra": "2791 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2791 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3760,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "320988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3760,
            "unit": "ns/op",
            "extra": "320988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "320988 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "320988 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 737821,
            "unit": "ns/op\t  690738 B/op\t    8636 allocs/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 737821,
            "unit": "ns/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690738,
            "unit": "B/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1572 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 163448,
            "unit": "ns/op\t  153589 B/op\t    1897 allocs/op",
            "extra": "6225 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 163448,
            "unit": "ns/op",
            "extra": "6225 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 153589,
            "unit": "B/op",
            "extra": "6225 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1897,
            "unit": "allocs/op",
            "extra": "6225 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 404285,
            "unit": "ns/op\t  367476 B/op\t    4713 allocs/op",
            "extra": "2946 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 404285,
            "unit": "ns/op",
            "extra": "2946 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367476,
            "unit": "B/op",
            "extra": "2946 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2946 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 437614,
            "unit": "ns/op\t  414440 B/op\t    4981 allocs/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 437614,
            "unit": "ns/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414440,
            "unit": "B/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 810419,
            "unit": "ns/op\t  780049 B/op\t    9776 allocs/op",
            "extra": "1442 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 810419,
            "unit": "ns/op",
            "extra": "1442 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780049,
            "unit": "B/op",
            "extra": "1442 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9776,
            "unit": "allocs/op",
            "extra": "1442 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "ceda5fa82d6f6633e773811308aef524080b3076",
          "message": "fix(deps): restore compatible fedwire20022 revision\n\nRenovate #70 set go.mod to fedwire20022 v0.3.5 without go.sum hashes.\nTagged releases dropped gen/<MessageType>/<version> packages this\nmodule still imports, so tidy cannot record v0.3.5. Pin the last\nrevision that builds and ignore further fedwire20022 updates.",
          "timestamp": "2026-10-14T14:17:35Z",
          "url": "https://github.com/moov-io/wire20022/commit/ceda5fa82d6f6633e773811308aef524080b3076"
        },
        "date": 1791987543099,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 75507,
            "unit": "ns/op\t   72285 B/op\t     842 allocs/op",
            "extra": "13262 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 75507,
            "unit": "ns/op",
            "extra": "13262 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72285,
            "unit": "B/op",
            "extra": "13262 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "13262 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 83559,
            "unit": "ns/op\t   72149 B/op\t     955 allocs/op",
            "extra": "14286 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 83559,
            "unit": "ns/op",
            "extra": "14286 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72149,
            "unit": "B/op",
            "extra": "14286 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14286 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 477221,
            "unit": "ns/op\t  463633 B/op\t    5546 allocs/op",
            "extra": "2545 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 477221,
            "unit": "ns/op",
            "extra": "2545 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 463633,
            "unit": "B/op",
            "extra": "2545 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5546,
            "unit": "allocs/op",
            "extra": "2545 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1475646,
            "unit": "ns/op\t 1423921 B/op\t   17565 allocs/op",
            "extra": "814 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1475646,
            "unit": "ns/op",
            "extra": "814 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423921,
            "unit": "B/op",
            "extra": "814 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17565,
            "unit": "allocs/op",
            "extra": "814 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2313,
            "unit": "ns/op\t    1024 B/op\t      18 allocs/op",
            "extra": "521278 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2313,
            "unit": "ns/op",
            "extra": "521278 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "521278 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "521278 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 411127,
            "unit": "ns/op\t  363142 B/op\t    4626 allocs/op",
            "extra": "2793 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 411127,
            "unit": "ns/op",
            "extra": "2793 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363142,
            "unit": "B/op",
            "extra": "2793 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4626,
            "unit": "allocs/op",
            "extra": "2793 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3810,
            "unit": "ns/op\t    3745 B/op\t      41 allocs/op",
            "extra": "311805 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3810,
            "unit": "ns/op",
            "extra": "311805 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3745,
            "unit": "B/op",
            "extra": "311805 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "311805 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 756459,
            "unit": "ns/op\t  690734 B/op\t    8636 allocs/op",
            "extra": "1539 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 756459,
            "unit": "ns/op",
            "extra": "1539 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690734,
            "unit": "B/op",
            "extra": "1539 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8636,
            "unit": "allocs/op",
            "extra": "1539 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 172654,
            "unit": "ns/op\t  155295 B/op\t    1918 allocs/op",
            "extra": "6612 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 172654,
            "unit": "ns/op",
            "extra": "6612 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 155295,
            "unit": "B/op",
            "extra": "6612 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1918,
            "unit": "allocs/op",
            "extra": "6612 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 442364,
            "unit": "ns/op\t  367476 B/op\t    4713 allocs/op",
            "extra": "2679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 442364,
            "unit": "ns/op",
            "extra": "2679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367476,
            "unit": "B/op",
            "extra": "2679 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4713,
            "unit": "allocs/op",
            "extra": "2679 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 443271,
            "unit": "ns/op\t  414435 B/op\t    4981 allocs/op",
            "extra": "2552 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 443271,
            "unit": "ns/op",
            "extra": "2552 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414435,
            "unit": "B/op",
            "extra": "2552 times\n4 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2552 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 829056,
            "unit": "ns/op\t  780056 B/op\t    9776 allocs/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 829056,
            "unit": "ns/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 780056,
            "unit": "B/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "a14ce1c7051454cb503ef6526c48e1ef53dd363c",
          "message": "ci: store median of 6 bench samples at 2 CPUs\n\nRun go test -count=6 -cpu=2 and collapse each benchmark to a median\nline before github-action-benchmark records it. Raise alert-threshold\nto 150% and pin ubuntu-24.04.\n\nDrop BenchmarkCustomerCreditTransfer from the suite; CreateDocument\nstill fatals on UniqueEndToEndTransactionRef.",
          "timestamp": "2026-10-14T14:53:19Z",
          "url": "https://github.com/moov-io/wire20022/commit/a14ce1c7051454cb503ef6526c48e1ef53dd363c"
        },
        "date": 1791989845075,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 66922,
            "unit": "ns/op\t72074 B/op\t842 allocs/op",
            "extra": "17768 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 66922,
            "unit": "ns/op",
            "extra": "17768 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72074,
            "unit": "B/op",
            "extra": "17768 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "17768 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 74605,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "15955 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 74605,
            "unit": "ns/op",
            "extra": "15955 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "15955 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "15955 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 429557.5,
            "unit": "ns/op\t461886 B/op\t5532 allocs/op",
            "extra": "2692 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 429557.5,
            "unit": "ns/op",
            "extra": "2692 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 461886,
            "unit": "B/op",
            "extra": "2692 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5532,
            "unit": "allocs/op",
            "extra": "2692 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1318103.5,
            "unit": "ns/op\t1423571 B/op\t17563 allocs/op",
            "extra": "904 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1318103.5,
            "unit": "ns/op",
            "extra": "904 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423571,
            "unit": "B/op",
            "extra": "904 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "904 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2186,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "511616 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2186,
            "unit": "ns/op",
            "extra": "511616 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "511616 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "511616 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 355698.5,
            "unit": "ns/op\t363058 B/op\t4625 allocs/op",
            "extra": "3225 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 355698.5,
            "unit": "ns/op",
            "extra": "3225 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363058,
            "unit": "B/op",
            "extra": "3225 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "3225 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3409.5,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "337798 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3409.5,
            "unit": "ns/op",
            "extra": "337798 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "337798 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "337798 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 675666.5,
            "unit": "ns/op\t690568 B/op\t8635 allocs/op",
            "extra": "1753 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 675666.5,
            "unit": "ns/op",
            "extra": "1753 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690568,
            "unit": "B/op",
            "extra": "1753 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "1753 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 156523,
            "unit": "ns/op\t163120 B/op\t2014 allocs/op",
            "extra": "7130 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 156523,
            "unit": "ns/op",
            "extra": "7130 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 163120,
            "unit": "B/op",
            "extra": "7130 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 2014,
            "unit": "allocs/op",
            "extra": "7130 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 366632.5,
            "unit": "ns/op\t367396 B/op\t4712 allocs/op",
            "extra": "3135 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 366632.5,
            "unit": "ns/op",
            "extra": "3135 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367396,
            "unit": "B/op",
            "extra": "3135 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "3135 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 384517.5,
            "unit": "ns/op\t414391 B/op\t4981 allocs/op",
            "extra": "3034 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 384517.5,
            "unit": "ns/op",
            "extra": "3034 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414391,
            "unit": "B/op",
            "extra": "3034 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3034 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 724791.5,
            "unit": "ns/op\t779858 B/op\t9775 allocs/op",
            "extra": "1629 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 724791.5,
            "unit": "ns/op",
            "extra": "1629 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779858,
            "unit": "B/op",
            "extra": "1629 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1629 times\n2 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "473ad635ed3a68beac7e7c095104676c1b4699ec",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:56Z",
          "url": "https://github.com/moov-io/wire20022/commit/473ad635ed3a68beac7e7c095104676c1b4699ec"
        },
        "date": 1791991173156,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 73381.5,
            "unit": "ns/op\t72068 B/op\t842 allocs/op",
            "extra": "16206 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 73381.5,
            "unit": "ns/op",
            "extra": "16206 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72068,
            "unit": "B/op",
            "extra": "16206 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "16206 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 80564,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "14757 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 80564,
            "unit": "ns/op",
            "extra": "14757 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "14757 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14757 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 456026.5,
            "unit": "ns/op\t467054 B/op\t5584 allocs/op",
            "extra": "2550 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 456026.5,
            "unit": "ns/op",
            "extra": "2550 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 467054,
            "unit": "B/op",
            "extra": "2550 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5584,
            "unit": "allocs/op",
            "extra": "2550 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1381240.5,
            "unit": "ns/op\t1423576 B/op\t17563 allocs/op",
            "extra": "864 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1381240.5,
            "unit": "ns/op",
            "extra": "864 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423576,
            "unit": "B/op",
            "extra": "864 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "864 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2251,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "500338 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2251,
            "unit": "ns/op",
            "extra": "500338 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "500338 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "500338 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 383572,
            "unit": "ns/op\t363061 B/op\t4625 allocs/op",
            "extra": "3037 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 383572,
            "unit": "ns/op",
            "extra": "3037 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363061,
            "unit": "B/op",
            "extra": "3037 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "3037 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3638,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "324103 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3638,
            "unit": "ns/op",
            "extra": "324103 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "324103 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "324103 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 704230.5,
            "unit": "ns/op\t690556 B/op\t8635 allocs/op",
            "extra": "1676 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 704230.5,
            "unit": "ns/op",
            "extra": "1676 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690556,
            "unit": "B/op",
            "extra": "1676 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "1676 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 161474,
            "unit": "ns/op\t155242 B/op\t1918 allocs/op",
            "extra": "7038 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 161474,
            "unit": "ns/op",
            "extra": "7038 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 155242,
            "unit": "B/op",
            "extra": "7038 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1918,
            "unit": "allocs/op",
            "extra": "7038 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 392566.5,
            "unit": "ns/op\t367396 B/op\t4712 allocs/op",
            "extra": "2968 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 392566.5,
            "unit": "ns/op",
            "extra": "2968 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367396,
            "unit": "B/op",
            "extra": "2968 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "2968 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 420096,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "2853 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 420096,
            "unit": "ns/op",
            "extra": "2853 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "2853 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2853 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 786832.5,
            "unit": "ns/op\t779853 B/op\t9775 allocs/op",
            "extra": "1509 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 786832.5,
            "unit": "ns/op",
            "extra": "1509 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779853,
            "unit": "B/op",
            "extra": "1509 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1509 times\n2 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "473ad635ed3a68beac7e7c095104676c1b4699ec",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:56Z",
          "url": "https://github.com/moov-io/wire20022/commit/473ad635ed3a68beac7e7c095104676c1b4699ec"
        },
        "date": 1791991799190,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 74736,
            "unit": "ns/op\t72112 B/op\t842 allocs/op",
            "extra": "16197 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 74736,
            "unit": "ns/op",
            "extra": "16197 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72112,
            "unit": "B/op",
            "extra": "16197 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "16197 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 82087,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "14600 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 82087,
            "unit": "ns/op",
            "extra": "14600 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "14600 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "14600 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 465212,
            "unit": "ns/op\t468666 B/op\t5606 allocs/op",
            "extra": "2580 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 465212,
            "unit": "ns/op",
            "extra": "2580 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 468666,
            "unit": "B/op",
            "extra": "2580 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5606,
            "unit": "allocs/op",
            "extra": "2580 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1420297.5,
            "unit": "ns/op\t1423580 B/op\t17563 allocs/op",
            "extra": "849 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1420297.5,
            "unit": "ns/op",
            "extra": "849 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423580,
            "unit": "B/op",
            "extra": "849 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "849 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2314.5,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "492876 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2314.5,
            "unit": "ns/op",
            "extra": "492876 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "492876 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "492876 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 393725,
            "unit": "ns/op\t363059 B/op\t4625 allocs/op",
            "extra": "2946 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 393725,
            "unit": "ns/op",
            "extra": "2946 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363059,
            "unit": "B/op",
            "extra": "2946 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "2946 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3705,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "313021 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3705,
            "unit": "ns/op",
            "extra": "313021 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "313021 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "313021 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 727123,
            "unit": "ns/op\t690564 B/op\t8635 allocs/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 727123,
            "unit": "ns/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690564,
            "unit": "B/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 163058.5,
            "unit": "ns/op\t151731 B/op\t1874 allocs/op",
            "extra": "6878 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 163058.5,
            "unit": "ns/op",
            "extra": "6878 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 151731,
            "unit": "B/op",
            "extra": "6878 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1874,
            "unit": "allocs/op",
            "extra": "6878 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 400972.5,
            "unit": "ns/op\t367394 B/op\t4712 allocs/op",
            "extra": "2836 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 400972.5,
            "unit": "ns/op",
            "extra": "2836 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367394,
            "unit": "B/op",
            "extra": "2836 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "2836 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 428864,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "2750 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 428864,
            "unit": "ns/op",
            "extra": "2750 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "2750 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2750 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 801463,
            "unit": "ns/op\t779847 B/op\t9775 allocs/op",
            "extra": "1478 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 801463,
            "unit": "ns/op",
            "extra": "1478 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779847,
            "unit": "B/op",
            "extra": "1478 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1478 times\n2 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "473ad635ed3a68beac7e7c095104676c1b4699ec",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:56Z",
          "url": "https://github.com/moov-io/wire20022/commit/473ad635ed3a68beac7e7c095104676c1b4699ec"
        },
        "date": 1792069352041,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 52482.5,
            "unit": "ns/op\t72050 B/op\t842 allocs/op",
            "extra": "22744 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 52482.5,
            "unit": "ns/op",
            "extra": "22744 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72050,
            "unit": "B/op",
            "extra": "22744 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "22744 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 58372,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "20641 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 58372,
            "unit": "ns/op",
            "extra": "20641 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "20641 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "20641 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 331937,
            "unit": "ns/op\t466106 B/op\t5584 allocs/op",
            "extra": "3432 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 331937,
            "unit": "ns/op",
            "extra": "3432 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 466106,
            "unit": "B/op",
            "extra": "3432 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5584,
            "unit": "allocs/op",
            "extra": "3432 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1017298,
            "unit": "ns/op\t1423574 B/op\t17563 allocs/op",
            "extra": "1164 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1017298,
            "unit": "ns/op",
            "extra": "1164 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423574,
            "unit": "B/op",
            "extra": "1164 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "1164 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1683.5,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "666667 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1683.5,
            "unit": "ns/op",
            "extra": "666667 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "666667 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "666667 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 281758.5,
            "unit": "ns/op\t363060 B/op\t4625 allocs/op",
            "extra": "4053 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 281758.5,
            "unit": "ns/op",
            "extra": "4053 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363060,
            "unit": "B/op",
            "extra": "4053 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "4053 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2652.5,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "435456 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2652.5,
            "unit": "ns/op",
            "extra": "435456 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "435456 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "435456 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 517680,
            "unit": "ns/op\t690566 B/op\t8635 allocs/op",
            "extra": "2274 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 517680,
            "unit": "ns/op",
            "extra": "2274 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690566,
            "unit": "B/op",
            "extra": "2274 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "2274 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 114759.5,
            "unit": "ns/op\t154946 B/op\t1914 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 114759.5,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 154946,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1914,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 282003.5,
            "unit": "ns/op\t367396 B/op\t4712 allocs/op",
            "extra": "4137 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 282003.5,
            "unit": "ns/op",
            "extra": "4137 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367396,
            "unit": "B/op",
            "extra": "4137 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "4137 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 305687.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "3801 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 305687.5,
            "unit": "ns/op",
            "extra": "3801 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "3801 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3801 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 577467,
            "unit": "ns/op\t779856 B/op\t9775 allocs/op",
            "extra": "2044 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 577467,
            "unit": "ns/op",
            "extra": "2044 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779856,
            "unit": "B/op",
            "extra": "2044 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "2044 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#72)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T15:12:31Z",
          "url": "https://github.com/moov-io/wire20022/commit/dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd"
        },
        "date": 1792155520313,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 72774.5,
            "unit": "ns/op\t72124 B/op\t842 allocs/op",
            "extra": "16525 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 72774.5,
            "unit": "ns/op",
            "extra": "16525 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72124,
            "unit": "B/op",
            "extra": "16525 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "16525 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 79460,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "15096 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 79460,
            "unit": "ns/op",
            "extra": "15096 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "15096 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "15096 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 451668,
            "unit": "ns/op\t466461 B/op\t5579 allocs/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 451668,
            "unit": "ns/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 466461,
            "unit": "B/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5579,
            "unit": "allocs/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1368385.5,
            "unit": "ns/op\t1423582 B/op\t17563 allocs/op",
            "extra": "878 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1368385.5,
            "unit": "ns/op",
            "extra": "878 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423582,
            "unit": "B/op",
            "extra": "878 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "878 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2216.5,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "516050 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2216.5,
            "unit": "ns/op",
            "extra": "516050 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "516050 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "516050 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 381141,
            "unit": "ns/op\t363062 B/op\t4625 allocs/op",
            "extra": "3075 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 381141,
            "unit": "ns/op",
            "extra": "3075 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363062,
            "unit": "B/op",
            "extra": "3075 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "3075 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3611,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "324046 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3611,
            "unit": "ns/op",
            "extra": "324046 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "324046 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "324046 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 695515,
            "unit": "ns/op\t690563 B/op\t8635 allocs/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 695515,
            "unit": "ns/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690563,
            "unit": "B/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "1718 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 152704.5,
            "unit": "ns/op\t148804 B/op\t1838 allocs/op",
            "extra": "7144 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 152704.5,
            "unit": "ns/op",
            "extra": "7144 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 148804,
            "unit": "B/op",
            "extra": "7144 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1838,
            "unit": "allocs/op",
            "extra": "7144 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 390629,
            "unit": "ns/op\t367397 B/op\t4712 allocs/op",
            "extra": "2983 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 390629,
            "unit": "ns/op",
            "extra": "2983 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367397,
            "unit": "B/op",
            "extra": "2983 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "2983 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 412906.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 412906.5,
            "unit": "ns/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 773135.5,
            "unit": "ns/op\t779858 B/op\t9775 allocs/op",
            "extra": "1550 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 773135.5,
            "unit": "ns/op",
            "extra": "1550 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779858,
            "unit": "B/op",
            "extra": "1550 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1550 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#72)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T15:12:31Z",
          "url": "https://github.com/moov-io/wire20022/commit/dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd"
        },
        "date": 1792238225245,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 73774,
            "unit": "ns/op\t72058 B/op\t842 allocs/op",
            "extra": "16424 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 73774,
            "unit": "ns/op",
            "extra": "16424 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72058,
            "unit": "B/op",
            "extra": "16424 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "16424 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 80186,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "15038 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 80186,
            "unit": "ns/op",
            "extra": "15038 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "15038 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "15038 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 458461,
            "unit": "ns/op\t467068 B/op\t5592 allocs/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 458461,
            "unit": "ns/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 467068,
            "unit": "B/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5592,
            "unit": "allocs/op",
            "extra": "2473 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1389198.5,
            "unit": "ns/op\t1423584 B/op\t17563 allocs/op",
            "extra": "865 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1389198.5,
            "unit": "ns/op",
            "extra": "865 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423584,
            "unit": "B/op",
            "extra": "865 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "865 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2247,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "506571 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2247,
            "unit": "ns/op",
            "extra": "506571 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "506571 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "506571 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 385830.5,
            "unit": "ns/op\t363064 B/op\t4625 allocs/op",
            "extra": "3043 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 385830.5,
            "unit": "ns/op",
            "extra": "3043 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363064,
            "unit": "B/op",
            "extra": "3043 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "3043 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3665,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "317797 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3665,
            "unit": "ns/op",
            "extra": "317797 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "317797 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "317797 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 718073.5,
            "unit": "ns/op\t690564 B/op\t8635 allocs/op",
            "extra": "1660 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 718073.5,
            "unit": "ns/op",
            "extra": "1660 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690564,
            "unit": "B/op",
            "extra": "1660 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "1660 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 163053.5,
            "unit": "ns/op\t159737 B/op\t1973 allocs/op",
            "extra": "7195 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 163053.5,
            "unit": "ns/op",
            "extra": "7195 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 159737,
            "unit": "B/op",
            "extra": "7195 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1973,
            "unit": "allocs/op",
            "extra": "7195 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 390875,
            "unit": "ns/op\t367396 B/op\t4712 allocs/op",
            "extra": "2949 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 390875,
            "unit": "ns/op",
            "extra": "2949 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367396,
            "unit": "B/op",
            "extra": "2949 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "2949 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 419575.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "2858 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 419575.5,
            "unit": "ns/op",
            "extra": "2858 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "2858 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "2858 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 781889,
            "unit": "ns/op\t779860 B/op\t9775 allocs/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 781889,
            "unit": "ns/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779860,
            "unit": "B/op",
            "extra": "1530 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1530 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#72)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T15:12:31Z",
          "url": "https://github.com/moov-io/wire20022/commit/dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd"
        },
        "date": 1792326305884,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 38242,
            "unit": "ns/op\t72144 B/op\t842 allocs/op",
            "extra": "31436 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 38242,
            "unit": "ns/op",
            "extra": "31436 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72144,
            "unit": "B/op",
            "extra": "31436 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "31436 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 43192.5,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "28102 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 43192.5,
            "unit": "ns/op",
            "extra": "28102 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "28102 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "28102 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 251094,
            "unit": "ns/op\t468532 B/op\t5602 allocs/op",
            "extra": "4754 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 251094,
            "unit": "ns/op",
            "extra": "4754 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 468532,
            "unit": "B/op",
            "extra": "4754 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5602,
            "unit": "allocs/op",
            "extra": "4754 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 757268.5,
            "unit": "ns/op\t1423570 B/op\t17563 allocs/op",
            "extra": "1569 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 757268.5,
            "unit": "ns/op",
            "extra": "1569 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423570,
            "unit": "B/op",
            "extra": "1569 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "1569 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1237,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "909918 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1237,
            "unit": "ns/op",
            "extra": "909918 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "909918 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "909918 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 202783,
            "unit": "ns/op\t363061 B/op\t4625 allocs/op",
            "extra": "5803 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 202783,
            "unit": "ns/op",
            "extra": "5803 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363061,
            "unit": "B/op",
            "extra": "5803 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "5803 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2045.5,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "558824 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2045.5,
            "unit": "ns/op",
            "extra": "558824 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "558824 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "558824 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 395698.5,
            "unit": "ns/op\t690555 B/op\t8635 allocs/op",
            "extra": "2936 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 395698.5,
            "unit": "ns/op",
            "extra": "2936 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690555,
            "unit": "B/op",
            "extra": "2936 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "2936 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 85889,
            "unit": "ns/op\t153822 B/op\t1900 allocs/op",
            "extra": "13432 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 85889,
            "unit": "ns/op",
            "extra": "13432 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 153822,
            "unit": "B/op",
            "extra": "13432 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1900,
            "unit": "allocs/op",
            "extra": "13432 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 215209,
            "unit": "ns/op\t367398 B/op\t4712 allocs/op",
            "extra": "5354 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 215209,
            "unit": "ns/op",
            "extra": "5354 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367398,
            "unit": "B/op",
            "extra": "5354 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "5354 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 230886.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "5107 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 230886.5,
            "unit": "ns/op",
            "extra": "5107 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "5107 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "5107 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 431062,
            "unit": "ns/op\t779864 B/op\t9775 allocs/op",
            "extra": "2683 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 431062,
            "unit": "ns/op",
            "extra": "2683 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779864,
            "unit": "B/op",
            "extra": "2683 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "2683 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#72)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T15:12:31Z",
          "url": "https://github.com/moov-io/wire20022/commit/dcdeeeaffbbc2bb369dcc2f41c636b799eaea4dd"
        },
        "date": 1792422857352,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 46746.5,
            "unit": "ns/op\t72092 B/op\t842 allocs/op",
            "extra": "25632 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 46746.5,
            "unit": "ns/op",
            "extra": "25632 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72092,
            "unit": "B/op",
            "extra": "25632 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "25632 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 51799.5,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "23173 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 51799.5,
            "unit": "ns/op",
            "extra": "23173 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "23173 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "23173 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 299290,
            "unit": "ns/op\t464868 B/op\t5562 allocs/op",
            "extra": "3927 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 299290,
            "unit": "ns/op",
            "extra": "3927 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 464868,
            "unit": "B/op",
            "extra": "3927 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5562,
            "unit": "allocs/op",
            "extra": "3927 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 966967.5,
            "unit": "ns/op\t1423576 B/op\t17563 allocs/op",
            "extra": "1228 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 966967.5,
            "unit": "ns/op",
            "extra": "1228 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423576,
            "unit": "B/op",
            "extra": "1228 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "1228 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1458,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "783950 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1458,
            "unit": "ns/op",
            "extra": "783950 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "783950 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "783950 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 254144,
            "unit": "ns/op\t363058 B/op\t4625 allocs/op",
            "extra": "4671 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 254144,
            "unit": "ns/op",
            "extra": "4671 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363058,
            "unit": "B/op",
            "extra": "4671 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "4671 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2491.5,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "460976 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2491.5,
            "unit": "ns/op",
            "extra": "460976 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "460976 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "460976 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 459781,
            "unit": "ns/op\t690558 B/op\t8635 allocs/op",
            "extra": "2568 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 459781,
            "unit": "ns/op",
            "extra": "2568 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690558,
            "unit": "B/op",
            "extra": "2568 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "2568 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 100661.5,
            "unit": "ns/op\t152684 B/op\t1886 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 100661.5,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 152684,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1886,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 253656.5,
            "unit": "ns/op\t367399 B/op\t4712 allocs/op",
            "extra": "4513 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 253656.5,
            "unit": "ns/op",
            "extra": "4513 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367399,
            "unit": "B/op",
            "extra": "4513 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "4513 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 267358.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "4342 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 267358.5,
            "unit": "ns/op",
            "extra": "4342 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "4342 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "4342 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 501064,
            "unit": "ns/op\t779854 B/op\t9775 allocs/op",
            "extra": "2367 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 501064,
            "unit": "ns/op",
            "extra": "2367 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779854,
            "unit": "B/op",
            "extra": "2367 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "2367 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "9b716076ef57e617839b50b66c575c3fb6cbfd05",
          "message": "chore(deps): update go toolchain directive to v1.27.3 (#73)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-20T12:58:46Z",
          "url": "https://github.com/moov-io/wire20022/commit/9b716076ef57e617839b50b66c575c3fb6cbfd05"
        },
        "date": 1792502093118,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 52802.5,
            "unit": "ns/op\t72040 B/op\t842 allocs/op",
            "extra": "22917 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 52802.5,
            "unit": "ns/op",
            "extra": "22917 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72040,
            "unit": "B/op",
            "extra": "22917 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "22917 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 58086.5,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "20778 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 58086.5,
            "unit": "ns/op",
            "extra": "20778 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "20778 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "20778 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 334271,
            "unit": "ns/op\t472237 B/op\t5654 allocs/op",
            "extra": "3464 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 334271,
            "unit": "ns/op",
            "extra": "3464 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 472237,
            "unit": "B/op",
            "extra": "3464 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5654,
            "unit": "allocs/op",
            "extra": "3464 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1019783,
            "unit": "ns/op\t1423572 B/op\t17563 allocs/op",
            "extra": "1173 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1019783,
            "unit": "ns/op",
            "extra": "1173 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423572,
            "unit": "B/op",
            "extra": "1173 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "1173 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1690.5,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "667495 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1690.5,
            "unit": "ns/op",
            "extra": "667495 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "667495 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "667495 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 278770.5,
            "unit": "ns/op\t363062 B/op\t4625 allocs/op",
            "extra": "4075 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 278770.5,
            "unit": "ns/op",
            "extra": "4075 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363062,
            "unit": "B/op",
            "extra": "4075 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "4075 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2641,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "435517 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2641,
            "unit": "ns/op",
            "extra": "435517 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "435517 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "435517 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 517879,
            "unit": "ns/op\t690568 B/op\t8635 allocs/op",
            "extra": "2284 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 517879,
            "unit": "ns/op",
            "extra": "2284 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690568,
            "unit": "B/op",
            "extra": "2284 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "2284 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 121047,
            "unit": "ns/op\t162937 B/op\t2013 allocs/op",
            "extra": "9040 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 121047,
            "unit": "ns/op",
            "extra": "9040 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 162937,
            "unit": "B/op",
            "extra": "9040 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 2013,
            "unit": "allocs/op",
            "extra": "9040 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 285150,
            "unit": "ns/op\t367394 B/op\t4712 allocs/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 285150,
            "unit": "ns/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367394,
            "unit": "B/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 304790.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "3824 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 304790.5,
            "unit": "ns/op",
            "extra": "3824 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "3824 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3824 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 574025,
            "unit": "ns/op\t779856 B/op\t9775 allocs/op",
            "extra": "2068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 574025,
            "unit": "ns/op",
            "extra": "2068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779856,
            "unit": "B/op",
            "extra": "2068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "2068 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "9b716076ef57e617839b50b66c575c3fb6cbfd05",
          "message": "chore(deps): update go toolchain directive to v1.27.3 (#73)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-20T12:58:46Z",
          "url": "https://github.com/moov-io/wire20022/commit/9b716076ef57e617839b50b66c575c3fb6cbfd05"
        },
        "date": 1792589035365,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 52015,
            "unit": "ns/op\t72068 B/op\t842 allocs/op",
            "extra": "23192 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 52015,
            "unit": "ns/op",
            "extra": "23192 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72068,
            "unit": "B/op",
            "extra": "23192 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "23192 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 57946,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "20722 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 57946,
            "unit": "ns/op",
            "extra": "20722 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "20722 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "20722 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 334186,
            "unit": "ns/op\t466195 B/op\t5581 allocs/op",
            "extra": "3476 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 334186,
            "unit": "ns/op",
            "extra": "3476 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 466195,
            "unit": "B/op",
            "extra": "3476 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5581,
            "unit": "allocs/op",
            "extra": "3476 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1022810.5,
            "unit": "ns/op\t1423572 B/op\t17563 allocs/op",
            "extra": "1166 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1022810.5,
            "unit": "ns/op",
            "extra": "1166 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423572,
            "unit": "B/op",
            "extra": "1166 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "1166 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1687,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "666970 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1687,
            "unit": "ns/op",
            "extra": "666970 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "666970 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "666970 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 275559.5,
            "unit": "ns/op\t363058 B/op\t4625 allocs/op",
            "extra": "4093 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 275559.5,
            "unit": "ns/op",
            "extra": "4093 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363058,
            "unit": "B/op",
            "extra": "4093 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "4093 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2616.5,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "432672 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2616.5,
            "unit": "ns/op",
            "extra": "432672 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "432672 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "432672 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 518921,
            "unit": "ns/op\t690568 B/op\t8635 allocs/op",
            "extra": "2275 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 518921,
            "unit": "ns/op",
            "extra": "2275 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690568,
            "unit": "B/op",
            "extra": "2275 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "2275 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 112729.5,
            "unit": "ns/op\t151324 B/op\t1869 allocs/op",
            "extra": "9362 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 112729.5,
            "unit": "ns/op",
            "extra": "9362 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 151324,
            "unit": "B/op",
            "extra": "9362 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1869,
            "unit": "allocs/op",
            "extra": "9362 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 283121,
            "unit": "ns/op\t367398 B/op\t4712 allocs/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 283121,
            "unit": "ns/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367398,
            "unit": "B/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "4068 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 302025.5,
            "unit": "ns/op\t414390 B/op\t4981 allocs/op",
            "extra": "3849 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 302025.5,
            "unit": "ns/op",
            "extra": "3849 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414390,
            "unit": "B/op",
            "extra": "3849 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3849 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 574116.5,
            "unit": "ns/op\t779859 B/op\t9775 allocs/op",
            "extra": "2064 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 574116.5,
            "unit": "ns/op",
            "extra": "2064 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779859,
            "unit": "B/op",
            "extra": "2064 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "2064 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "a1d13ca65a92687dcead406817f74c50690e111d",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#74)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T18:15:16Z",
          "url": "https://github.com/moov-io/wire20022/commit/a1d13ca65a92687dcead406817f74c50690e111d"
        },
        "date": 1792675442032,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 65742,
            "unit": "ns/op\t72088 B/op\t842 allocs/op",
            "extra": "18417 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 65742,
            "unit": "ns/op",
            "extra": "18417 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72088,
            "unit": "B/op",
            "extra": "18417 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "18417 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 73793.5,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "16179 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 73793.5,
            "unit": "ns/op",
            "extra": "16179 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "16179 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "16179 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 423936.5,
            "unit": "ns/op\t469310 B/op\t5622 allocs/op",
            "extra": "2680 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 423936.5,
            "unit": "ns/op",
            "extra": "2680 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 469310,
            "unit": "B/op",
            "extra": "2680 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5622,
            "unit": "allocs/op",
            "extra": "2680 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1298424.5,
            "unit": "ns/op\t1423573 B/op\t17563 allocs/op",
            "extra": "927 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1298424.5,
            "unit": "ns/op",
            "extra": "927 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423573,
            "unit": "B/op",
            "extra": "927 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "927 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 2056.5,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "561610 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 2056.5,
            "unit": "ns/op",
            "extra": "561610 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "561610 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "561610 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 355029,
            "unit": "ns/op\t363063 B/op\t4625 allocs/op",
            "extra": "3381 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 355029,
            "unit": "ns/op",
            "extra": "3381 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363063,
            "unit": "B/op",
            "extra": "3381 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "3381 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 3498.5,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "335131 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 3498.5,
            "unit": "ns/op",
            "extra": "335131 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "335131 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "335131 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 664494,
            "unit": "ns/op\t690564 B/op\t8635 allocs/op",
            "extra": "1798 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 664494,
            "unit": "ns/op",
            "extra": "1798 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690564,
            "unit": "B/op",
            "extra": "1798 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "1798 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 149899.5,
            "unit": "ns/op\t158960 B/op\t1963 allocs/op",
            "extra": "7718 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 149899.5,
            "unit": "ns/op",
            "extra": "7718 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 158960,
            "unit": "B/op",
            "extra": "7718 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1963,
            "unit": "allocs/op",
            "extra": "7718 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 364550.5,
            "unit": "ns/op\t367399 B/op\t4712 allocs/op",
            "extra": "3205 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 364550.5,
            "unit": "ns/op",
            "extra": "3205 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367399,
            "unit": "B/op",
            "extra": "3205 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "3205 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 378636,
            "unit": "ns/op\t414391 B/op\t4981 allocs/op",
            "extra": "3042 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 378636,
            "unit": "ns/op",
            "extra": "3042 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414391,
            "unit": "B/op",
            "extra": "3042 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3042 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 722157.5,
            "unit": "ns/op\t779859 B/op\t9775 allocs/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 722157.5,
            "unit": "ns/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779859,
            "unit": "B/op",
            "extra": "1671 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1671 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "a1d13ca65a92687dcead406817f74c50690e111d",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#74)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T18:15:16Z",
          "url": "https://github.com/moov-io/wire20022/commit/a1d13ca65a92687dcead406817f74c50690e111d"
        },
        "date": 1792760540800,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWriteXML/ConnectionCheck",
            "value": 53345.5,
            "unit": "ns/op\t72041 B/op\t842 allocs/op",
            "extra": "22429 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - ns/op",
            "value": 53345.5,
            "unit": "ns/op",
            "extra": "22429 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - B/op",
            "value": 72041,
            "unit": "B/op",
            "extra": "22429 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ConnectionCheck - allocs/op",
            "value": 842,
            "unit": "allocs/op",
            "extra": "22429 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck",
            "value": 60775.5,
            "unit": "ns/op\t72142 B/op\t955 allocs/op",
            "extra": "19351 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - ns/op",
            "value": 60775.5,
            "unit": "ns/op",
            "extra": "19351 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - B/op",
            "value": 72142,
            "unit": "B/op",
            "extra": "19351 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ConnectionCheck - allocs/op",
            "value": 955,
            "unit": "allocs/op",
            "extra": "19351 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest",
            "value": 337383.5,
            "unit": "ns/op\t464770 B/op\t5568 allocs/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - ns/op",
            "value": 337383.5,
            "unit": "ns/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - B/op",
            "value": 464770,
            "unit": "B/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/DrawdownRequest - allocs/op",
            "value": 5568,
            "unit": "allocs/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest",
            "value": 1037756,
            "unit": "ns/op\t1423564 B/op\t17563 allocs/op",
            "extra": "1156 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - ns/op",
            "value": 1037756,
            "unit": "ns/op",
            "extra": "1156 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - B/op",
            "value": 1423564,
            "unit": "B/op",
            "extra": "1156 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/DrawdownRequest - allocs/op",
            "value": 17563,
            "unit": "allocs/op",
            "extra": "1156 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master",
            "value": 1752.5,
            "unit": "ns/op\t1024 B/op\t18 allocs/op",
            "extra": "650364 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - ns/op",
            "value": 1752.5,
            "unit": "ns/op",
            "extra": "650364 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - B/op",
            "value": 1024,
            "unit": "B/op",
            "extra": "650364 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/Master - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "650364 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master",
            "value": 290067,
            "unit": "ns/op\t363060 B/op\t4625 allocs/op",
            "extra": "4048 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - ns/op",
            "value": 290067,
            "unit": "ns/op",
            "extra": "4048 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - B/op",
            "value": 363060,
            "unit": "B/op",
            "extra": "4048 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/Master - allocs/op",
            "value": 4625,
            "unit": "allocs/op",
            "extra": "4048 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn",
            "value": 2778,
            "unit": "ns/op\t3744 B/op\t41 allocs/op",
            "extra": "423538 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - ns/op",
            "value": 2778,
            "unit": "ns/op",
            "extra": "423538 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - B/op",
            "value": 3744,
            "unit": "B/op",
            "extra": "423538 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentReturn - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "423538 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn",
            "value": 540469,
            "unit": "ns/op\t690563 B/op\t8635 allocs/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - ns/op",
            "value": 540469,
            "unit": "ns/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - B/op",
            "value": 690563,
            "unit": "B/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentReturn - allocs/op",
            "value": 8635,
            "unit": "allocs/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest",
            "value": 117415.5,
            "unit": "ns/op\t155346 B/op\t1918 allocs/op",
            "extra": "9510 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - ns/op",
            "value": 117415.5,
            "unit": "ns/op",
            "extra": "9510 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - B/op",
            "value": 155346,
            "unit": "B/op",
            "extra": "9510 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/PaymentStatusRequest - allocs/op",
            "value": 1918,
            "unit": "allocs/op",
            "extra": "9510 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest",
            "value": 291260.5,
            "unit": "ns/op\t367390 B/op\t4712 allocs/op",
            "extra": "3936 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - ns/op",
            "value": 291260.5,
            "unit": "ns/op",
            "extra": "3936 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - B/op",
            "value": 367390,
            "unit": "B/op",
            "extra": "3936 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/PaymentStatusRequest - allocs/op",
            "value": 4712,
            "unit": "allocs/op",
            "extra": "3936 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse",
            "value": 319611.5,
            "unit": "ns/op\t414391 B/op\t4981 allocs/op",
            "extra": "3615 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - ns/op",
            "value": 319611.5,
            "unit": "ns/op",
            "extra": "3615 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - B/op",
            "value": 414391,
            "unit": "B/op",
            "extra": "3615 times\n2 procs"
          },
          {
            "name": "BenchmarkWriteXML/ReturnRequestResponse - allocs/op",
            "value": 4981,
            "unit": "allocs/op",
            "extra": "3615 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse",
            "value": 603422,
            "unit": "ns/op\t779853 B/op\t9775 allocs/op",
            "extra": "1928 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - ns/op",
            "value": 603422,
            "unit": "ns/op",
            "extra": "1928 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - B/op",
            "value": 779853,
            "unit": "B/op",
            "extra": "1928 times\n2 procs"
          },
          {
            "name": "BenchmarkParseXML/ReturnRequestResponse - allocs/op",
            "value": 9775,
            "unit": "allocs/op",
            "extra": "1928 times\n2 procs"
          }
        ]
      }
    ]
  }
}
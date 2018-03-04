# pdfsort

pdfsort takes two single-sided scans and reorders the pages as if it was a double-sided scan.

When scanning a double-sided stack of papers on a single-sided scanner, first scan one side (producing pages 1, 3, ..., 2n-1), and then flip the entire stack and scan the other side (producing pages 2n, 2n-2, ..., 2).

pdfsort orders a pdf produced this way.

## Installation

    go get github.com/jbeder/pdfsort

## Execution

    go run pdfsort.go <input> <output>

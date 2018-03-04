// $ pdfsort <input> <output>
//
// pdfsort takes two single-sided scans and reorders the pages as if it was a double-sided scan.
//
// When scanning a double-sided stack of papers on a single-sided scanner, first scan one side (producing pages 1, 3, ..., 2n-1), and then flip the entire stack and scan the other side (producing pages 2n, 2n-2, ..., 2).
// pdfsort orders a pdf produced this way.
package main

import (
	"fmt"
	"log"
	"os"

	pdf "github.com/unidoc/unidoc/pdf/model"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: pdfsort <input> <output>")
	}
	if err := pdfsort(os.Args[1], os.Args[2]); err != nil {
		log.Fatalf("pdfsort: %v", err)
	}
}

func pages(n int) []int {
	p := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			p = append(p, i/2+1)
		} else {
			p = append(p, n-(i-1)/2)
		}
	}
	return p
}

func pdfsort(in, out string) error {
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}
	w := pdf.NewPdfWriter()

	n, err := r.GetNumPages()
	if err != nil {
		return err
	}
	if n%2 != 0 {
		return fmt.Errorf("expected even number of pages, found %d", n)
	}
	ocProps, err := r.GetOCProperties()
	if err != nil {
		return err
	}
	w.SetOCProperties(ocProps)

	for _, i := range pages(n) {
		page, err := r.GetPage(i)
		if err != nil {
			return err
		}
		if err = w.AddPage(page); err != nil {
			return err
		}
	}

	wf, err := os.Create(out)
	if err != nil {
		return err
	}
	defer wf.Close()
	return w.Write(wf)
}

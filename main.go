package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const page = `<html><head><title>Fib(%[1]d)</title></head><body><h1>Fib(%[1]d) = %[2]d</h1><a href="/?n=%[3]d">Fib(%[3]d)</a></body></html>`
const maxN = 64

func main() {
	handlerFunction := func(res http.ResponseWriter, req *http.Request) {
		n, err := getN(req)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(res, page, n, Fib(n), (n+1)%maxN)
	}

	http.ListenAndServe(":8080", http.HandlerFunc(handlerFunction))
}

// getN parses the "n" query param and ensures it is in the range [0, maxN]
func getN(req *http.Request) (int, error) {
	nStr := req.URL.Query().Get("n")
	if nStr == "" {
		return 1, nil
	}
	qn, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, err
	}
	if qn < 0 || qn > maxN {
		return 0, fmt.Errorf("n must be [0, %d]", maxN)
	}
	return qn, nil
}

type TwoByTwo [2][2]int

// Fib calculates the nth fib number using
func Fib(nth int) int {
	m1 := TwoByTwo{{1, 1}, {1, 0}}
	m2 := TwoByTwo{{1, 1}, {1, 0}}

	for ; nth > 0; nth = nth >> 1 {
		if (nth & 1) > 0 {
			m1 = matrixProduct(m1, m2)
		}
		m2 = matrixProduct(m2, m2)
	}

	return m1[1][0]
}

func matrixProduct(m1, m2 TwoByTwo) TwoByTwo {
	return TwoByTwo{
		{m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0], m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]},
		{m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]},
	}
}

package main

import "testing"

func TestCmplx_add(t *testing.T) {
    n1 := cmplx_num{1, 2}
    n2 := cmplx_num{3, 4}
    res := cmplx_add(n1, n2)
    if (res.real != 4 || res.img != 6) {
        t.Error()
    }
}


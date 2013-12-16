package main

/*
#include "svm.h"
*/
import "C"
import "unsafe"

func main() {

    godata := [][]C.struct_parts{
        []C.struct_parts{
            C.struct_parts{i: 1, j: 2},
            C.struct_parts{i: 2, j: 3},
            C.struct_parts{i: 3, j: 4},
        },
        []C.struct_parts{
            C.struct_parts{i: 4, j: 5},
            C.struct_parts{i: 5, j: 6},
            C.struct_parts{i: 6, j: 7},
        },
    }

    var _ = C.struct_checking_pointers{
        x: (**C.struct_parts)(unsafe.Pointer(&godata[0])),
    }

}

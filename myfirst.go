package main

/*
#include <svm.h>
*/

import "C"

func main() {
    a := C.struct_parts{i:C.int(1), j:C.int(2)}
    b := C.struct_parts{i:C.int(2), j:C.int(3)}
    c := C.struct_parts{i:C.int(3), j:C.int(4)}
    d := C.struct_parts{i:C.int(4), j:C.int(5)}
    e := C.struct_parts{i:C.int(5), j:C.int(6)}
    f := C.struct_parts{i:C.int(6), j:C.int(7)}

    //How do I then create checking_pointers?
}

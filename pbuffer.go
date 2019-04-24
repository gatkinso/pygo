package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
// int PyArg_ParseTuple_s(PyObject * args, char ** str);
import "C" 

import (
	"log"
	"fmt"
	"strings"
	"github.com/golang/protobuf/jsonpb"
)

//export setmsg
func setmsg(self, args *C.PyObject) *C.PyObject {
	pb := &Stencil{}

	var buffer *C.char

	if C.PyArg_ParseTuple_s(args, &buffer) == 0 {
		return nil
	}

	var go_buffer = C.GoString(buffer)

	//fmt.Println(go_buffer)

	if err := jsonpb.Unmarshal(strings.NewReader(go_buffer), pb); err != nil {
		log.Fatalln("Error converting JSON to proto:", err)
	}

    fmt.Println( pb )

	return C.PyLong_FromLongLong(0)
}

func main() {}
package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
// int PyArg_ParseTuple_LL(PyObject *, long long *, long long *);
// int PyArg_ParseTuple_s(PyObject * args, char ** str);
import "C" 

import (
	//"log"
	"fmt"
	//"strings"
	//"github.com/golang/protobuf/jsonpb"
)

//export add
func add(self, args *C.PyObject) *C.PyObject {
	var a, b C.longlong
	if C.PyArg_ParseTuple_LL(args, &a, &b) == 0 {
		return nil
	}
	return C.PyLong_FromLongLong(a + b)
}

//export setmsg
func setmsg(self, args *C.PyObject) *C.PyObject {
	pb := &Stencil{}

	var buffer *C.char

	if C.PyArg_ParseTuple_s(args, &buffer) == 0 {
		return nil
	}

	var go_buffer = C.GoString(buffer)

	fmt.Println("Hi")

	//if err := jsonpb.Unmarshal(strings.NewReader(go_buffer), pb); err != nil {
	//	log.Fatalln("Error converting JSON to proto:", err)
	//}

    //fmt.Println( pb.GetStringValues()["timestamp"] )

	return nil
}

func main() {}
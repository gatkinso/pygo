#define Py_LIMITED_API
#include <Python.h>

PyObject* add(PyObject *self, PyObject *args);

int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {
    return PyArg_ParseTuple(args, "LL", a, b);
}

int PyArg_ParseTuple_s(PyObject * args, char ** str) {
    return PyArg_ParseTuple(args, "s", str);
}

static PyMethodDef PygoMethods[] = {
	{"add", add, METH_VARARGS, "Add two numbers."},
	{NULL, NULL, 0, NULL}
};

static struct PyModuleDef pygomodule = {
	PyModuleDef_HEAD_INIT, "pygo", NULL, -1, PygoMethods
};

PyMODINIT_FUNC PyInit_pygo(void)
{
	return PyModule_Create(&pygomodule);
}

#define Py_LIMITED_API
#include <Python.h>

PyObject* setmsg(PyObject *self, PyObject *args);

int PyArg_ParseTuple_s(PyObject * args, char ** str) {
    return PyArg_ParseTuple(args, "s", str);
}

static PyMethodDef LibagentMethods[] = {
	{"setmsg", setmsg, METH_VARARGS, "Pass json to Go."},
	{NULL, NULL, 0, NULL}
};

static struct PyModuleDef libagentmodule = {
	PyModuleDef_HEAD_INIT, "libagent", NULL, -1, LibagentMethods
};

PyMODINIT_FUNC PyInit_libagent(void)
{
	return PyModule_Create(&libagentmodule);
}

package ast

import (
	"surf/code"
	"surf/logger"
	"surf/token"
)

// FileCode is a representation of the file source code
type FileCode struct {
	// functions holds the functions of the current file
	// and of all the imported files.
	functions map[string]map[string]*code.Function
	// modules holds the defined modules in the file
	modules map[string]*code.SurfMod
}

// NewFileCode creates a new FileCode object
func NewFileCode() FileCode {
	return FileCode{
		functions: make(map[string]map[string]*code.Function),
		modules:   make(map[string]*code.SurfMod),
	}
}

// GetFunctions returns the functions of the FileCode
func (fc *FileCode) GetFunctions() *map[string]map[string]*code.Function {
	return &fc.functions
}

// AddFunction adds a function to the FileCode
func (fc *FileCode) AddFunction(trace token.Token, file string, name string, function code.Function) {
	functions, ok := fc.functions[file]

	// Make sure the file exists in the map
	if !ok {
		fc.functions[file] = make(map[string]*code.Function)
		functions, _ = fc.functions[file]
	}

	for key := range functions {
		if key == name {
			logger.TokenError(
				trace,
				"Redefinition of function "+name,
				"The function "+name+" has already been defined in this file",
				"Change the name of the function",
			)
		}
	}

	functions[name] = &function
}

// GetFunction returns a function from the FileCode
func (fc *FileCode) GetFunction(file string, name string) (*code.Function, bool) {
	// Make sure the file exists in the map
	functions, ok := fc.functions[file]

	if !ok {
		return &code.Function{}, false
	}

	function, _ok := functions[name]
	return function, _ok
}

// LocateFunction locates a function in the FileCode and returns it
// along with a boolean indicating if the function was found
// and another boolean indicating if the function was found in the same file
func LocateFunction(
	functions map[string]map[string]*code.Function,
	file string,
	name string,
) (*code.Function, bool, bool) {
	fileFunctions, ok := functions[file]
	// Make sure the file exists in the map
	if !ok {
		return &code.Function{}, false, false
	}

	// Check if the function is in the current file
	if fun, ok := fileFunctions[name]; ok {
		return fun, true, true
	}

	// Check if the function is in an imported file
	for _, value := range functions {
		if fun, ok := value[name]; ok {
			return fun, true, false
		}
	}

	return &code.Function{}, false, false
}

// GetModules returns the modules of the FileCode
func (fc *FileCode) GetModules() *map[string]*code.SurfMod {
	return &fc.modules
}

// AddModule adds a new module to the FileCode
func (fc *FileCode) AddModule(name string, module *code.SurfMod) {
	fc.modules[name] = module
}

package main

/*
#include <stdlib.h>
#include "type.h"
const size_t FILE_SIZE = sizeof(struct File);
const size_t IDENT_SIZE = sizeof(struct Ident);
const size_t COBJECT_SIZE = sizeof(struct CObject);
const size_t GENDECL_SIZE = sizeof(struct GenDecl);
const size_t FUNCDECL_SIZE = sizeof(struct FuncDecl);
const size_t DECL_SIZE = sizeof(struct Decl);
*/
import "C"
import "unsafe"
import "go/ast"
import "fmt"
import "reflect"

func CreateCPointer(size C.size_t) unsafe.Pointer {
	return unsafe.Pointer(C.malloc(size))
}

func ConvertGenDecl(genDecl *ast.GenDecl) *C.struct_GenDecl {
	CGenDecl := (*C.struct_GenDecl)(CreateCPointer(C.GENDECL_SIZE))
	CGenDecl.Count = C.int(len(genDecl.Specs))
	return CGenDecl
}

func ConvertFuncDecl(funcDecl *ast.FuncDecl) *C.struct_FuncDecl {
	CFuncDecl := (*C.struct_FuncDecl)(CreateCPointer(C.FUNCDECL_SIZE))
	CFuncDecl.Name = C.CString(funcDecl.Name.Name)
	return CFuncDecl
}

func ConvertDecl(decl ast.Decl) *C.struct_Decl {
	CDecl := (*C.struct_Decl)(CreateCPointer(C.DECL_SIZE))
	switch d := interface{}(decl).(type) {
	case *ast.GenDecl:
		{
			CDecl.GenDecl = ConvertGenDecl(d)
			CDecl.Type = C.GenDeclType
		}
	case *ast.FuncDecl:
		{
			CDecl.FuncDecl = ConvertFuncDecl(d)
			CDecl.Type = C.FuncDeclType
		}
	default:
		fmt.Println(reflect.TypeOf(d))
	}
	return CDecl
}

func ConvertAstObject(obj *ast.Object) *C.struct_CObject {
	// TODO
	return (*C.struct_CObject)(CreateCPointer(C.COBJECT_SIZE))
}

func ConvertAstIdent(ident *ast.Ident) *C.struct_Ident {
	Cident := (*C.struct_Ident)(CreateCPointer(C.IDENT_SIZE))
	Cident.Obj = ConvertAstObject(ident.Obj)
	Cident.Name = C.CString(ident.Name)
	return Cident
}

func ConvertAstDecls(decls []ast.Decl) *C.struct_Decl {
	CDecls := (*C.struct_Decl)(CreateCPointer(C.DECL_SIZE * C.size_t(len(decls))))
	for pos, decl := range decls {
		CDecl := (*C.struct_Decl)(unsafe.Pointer(uintptr(unsafe.Pointer(CDecls)) + uintptr(C.size_t(pos)*C.IDENT_SIZE)))
		*CDecl = *ConvertDecl(decl)
	}
	return CDecls
}

func ConvertAstFile(file *ast.File) *C.struct_File {
	CFile := (*C.struct_File)(CreateCPointer(C.FILE_SIZE))
	CFile.Name = ConvertAstIdent(file.Name)
	CFile.Decls = ConvertAstDecls(file.Decls)
	CFile.DeclsCount = C.int(len(file.Decls))
	return CFile
}

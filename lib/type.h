#ifndef TYPE_H
#define TYPE_H
#include <stdlib.h>

enum ObjKind
{
    Bad,
    Pkg,
    Con,
    Typ,
    Var,
    Fun,
    Lbl
};

struct CObject
{
    enum ObjKind Kind;
    char* Name;
};
extern const size_t COBJECT_SIZE;

struct Ident
{
    char* Name;
    int NamePos;
    struct CObject* Obj;
};
extern const size_t IDENT_SIZE;

struct FuncDecl
{
    char* Name;
};
extern const size_t FUNCDECL_SIZE;

struct GenDecl
{
    int Count;
};
extern const size_t GENDECL_SIZE;


struct Decl
{
    struct FuncDecl* FuncDecl;
    struct GenDecl* GenDecl;

    enum {
        FuncDeclType,
        GenDeclType
    } Type;
};
extern const size_t DECL_SIZE;

struct File
{
    struct Ident* Name;
    struct File* Next;
    struct Decl* Decls;
    int DeclsCount;
};
extern const size_t FILE_SIZE;

struct File* CreateFile(char* name, struct File* next);
#endif
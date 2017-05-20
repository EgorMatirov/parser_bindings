#include "libparser_bindings.h"
#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <type.h>

int main()
{
    GoString str;
    str.p = "package main; var b int = 3 \nfunc main(){\n    a := b\n\n}\nfunc test(){\nif} func test2(){}";
    str.n = strlen(str.p);
    File *file = ParseFile(str);
    std::cout << "Package name: " << file->Name->Name << std::endl;
    std::cout << "Global declarations count: " << file->DeclsCount << std::endl;
    for(int i = 0; i<file->DeclsCount; ++i)
    {
        auto current = file->Decls+i;
        if(current->Type == Decl::FuncDeclType) {
            std::cout << (i+1) << ". Function " << current->FuncDecl->Name << std::endl;
        }
        else {
            std::cout << (i+1) << ". General declaration. Count: " << current->GenDecl->Count << std::endl;
        }
    }
    return 0;
}

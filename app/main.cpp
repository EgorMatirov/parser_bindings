#include "libparser_bindings.h"
#include <iostream>
#include <stdlib.h>
#include <string.h>
int main()
{
        GoString str;
        str.p = "package main; func main(){\n}";
        str.n = strlen(str.p);
        File *file = ParseFile(str);
        std::cout << "Package name: " << file->Name << std::endl;
        return 0;
}

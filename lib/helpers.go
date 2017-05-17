package main;

/*
#include <stdlib.h>
#include "type.h"
struct File* CreateFile(char* name, struct File* next)
{
    struct File *file = malloc(sizeof(struct File));
    file->Name = name;
    file->Next = next;
    return file;
}
*/
import "C"

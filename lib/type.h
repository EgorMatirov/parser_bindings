#ifndef TYPE_H
#define TYPE_H
struct File
{
    char* Name;
    struct File* Next;
};

struct File* CreateFile(char* name, struct File* next);
#endif
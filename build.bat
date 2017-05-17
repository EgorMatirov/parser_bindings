mkdir build
cd build
go build -o ./libparser_bindings.a -buildmode=c-archive parser_bindings/lib
g++ -I. -I../lib -L. -o main ../app/main.cpp -lparser_bindings -lmingw32 -lwinmm -lws2_32
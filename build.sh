mkdir build
cd build
go build -o ./libparser_bindings.a -buildmode=c-archive ../lib
g++ -I. -I../lib -L. -o main ../app/main.cpp -lparser_bindings -lpthread
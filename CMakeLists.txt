cmake_minimum_required(VERSION 3.5)
project(app)

set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES main.cpp)
add_executable(app ${SOURCE_FILES})

set(PARSER_BINDINGS_LIB_PATH ${CMAKE_BINARY_DIR}/go_lib/)

set(GO_LIB_FILES ${CMAKE_CURRENT_SOURCE_DIR}/lib/helpers.go ${CMAKE_CURRENT_SOURCE_DIR}/lib/main.go ${CMAKE_CURRENT_SOURCE_DIR}/lib/type.h)

add_custom_command(
        OUTPUT ${PARSER_BINDINGS_LIB_PATH}/libparser_bindings.a ${PARSER_BINDINGS_LIB_PATH}/libparser_bindings.h
        COMMAND go build -o ${PARSER_BINDINGS_LIB_PATH}/libparser_bindings.a -buildmode=c-archive ./lib
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        DEPENDS ${GO_LIB_FILES}
)

add_custom_target(go_library ALL DEPENDS ${PARSER_BINDINGS_LIB_PATH}/libparser_bindings.a)

add_library(parser_bindings STATIC IMPORTED)
set_target_properties(parser_bindings PROPERTIES IMPORTED_LOCATION ${PARSER_BINDINGS_LIB_PATH}/libparser_bindings.a)
add_dependencies(parser_bindings go_library)

target_link_libraries(app parser_bindings)
target_include_directories(app PUBLIC ${PARSER_BINDINGS_LIB_PATH})
target_include_directories(app PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/lib)

if(WIN32)
    target_link_libraries(app mingw32 winmm ws2_32)
endif(WIN32)

if (UNIX)
    target_link_libraries(app pthread)
endif ()

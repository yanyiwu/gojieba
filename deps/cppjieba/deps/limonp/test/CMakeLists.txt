SET(EXECUTABLE_OUTPUT_PATH ${PROJECT_BINARY_DIR}/test)
SET(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/test/lib)

INCLUDE_DIRECTORIES(${PROJECT_SOURCE_DIR}/include)

ADD_EXECUTABLE(demo demo.cpp)

#if(NOT MSVC)
#	ADD_EXECUTABLE(thread_pool_demo thread_pool_demo.cpp)
#	TARGET_LINK_LIBRARIES(thread_pool_demo pthread)
#endif()

ADD_SUBDIRECTORY(unittest)

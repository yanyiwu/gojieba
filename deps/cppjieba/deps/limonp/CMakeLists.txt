cmake_minimum_required(VERSION 3.14)

PROJECT(limonp
        LANGUAGES CXX)

################
# cmake config #
################

if(NOT DEFINED CMAKE_CXX_STANDARD)
    set(CMAKE_CXX_STANDARD 11)
endif()
message(STATUS "CMAKE_CXX_STANDARD is ${CMAKE_CXX_STANDARD}")

set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)


##############
# dependency #
##############
include(GNUInstallDirs)

##########
# target #
##########

add_library(${PROJECT_NAME} INTERFACE)
add_library(${PROJECT_NAME}::${PROJECT_NAME} ALIAS ${PROJECT_NAME})

target_include_directories(${PROJECT_NAME}
                           INTERFACE
                               $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include>
                               $<INSTALL_INTERFACE:${CMAKE_INSTALL_FULL_INCLUDEDIR}>)

########
# test #
########

ENABLE_TESTING()

ADD_SUBDIRECTORY(test)
ADD_TEST(NAME ./test/demo COMMAND ./test/demo)
ADD_TEST(NAME ./test/test.run COMMAND ./test/test.run)

#if(NOT MSVC)
#	ADD_TEST(NAME ./test/thread_pool_demo COMMAND ./test/thread_pool_demo)
#endif()

###########
# install #
###########

include(GNUInstallDirs)

install(TARGETS ${PROJECT_NAME}
        EXPORT ${PROJECT_NAME})

install(EXPORT ${PROJECT_NAME}
        DESTINATION ${CMAKE_INSTALL_DATADIR}/${PROJECT_NAME}/cmake/
        NAMESPACE ${PROJECT_NAME}::
        FILE ${PROJECT_NAME}-config.cmake)

install(DIRECTORY include/
        DESTINATION ${CMAKE_INSTALL_INCLUDEDIR})

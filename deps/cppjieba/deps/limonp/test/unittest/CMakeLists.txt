SET(GTEST_ROOT_DIR ${PROJECT_SOURCE_DIR}/test/gtest)

INCLUDE_DIRECTORIES(${GTEST_ROOT_DIR} ${GTEST_ROOT_DIR}/include ${PROJECT_SOURCE_DIR}/include)
ADD_EXECUTABLE(test.run 
    gtest_main.cpp 
    TLogging.cpp  
    TArgvContext.cpp  
    TConfig.cpp
    TMd5.cpp
    TStringUtil.cpp 
    TStdExtension.cpp 
    TMutexLock.cpp 
    TBlockingQueue.cpp 
    TBoundedQueue.cpp 
    TLocalVector.cpp 
    TClosure.cpp 
    TThread.cpp 
    TThreadPool.cpp 
    TColorPrint.cpp
    TFileLock.cpp
    )
TARGET_LINK_LIBRARIES(test.run gtest pthread)

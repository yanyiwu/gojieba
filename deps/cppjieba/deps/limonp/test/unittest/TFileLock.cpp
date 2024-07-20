#include "limonp/FileLock.hpp"
#include "gtest/gtest.h"

using namespace limonp;

static const char* test_file = "/tmp/limonp_test_filelock";

TEST(FileLockTest, Test1) {
  FileLock fileLock1;
  fileLock1.Open(test_file);
  ASSERT_TRUE(fileLock1.Ok());
  fileLock1.Lock();
  ASSERT_TRUE(fileLock1.Ok());
  fileLock1.UnLock();
  ASSERT_TRUE(fileLock1.Ok());
}

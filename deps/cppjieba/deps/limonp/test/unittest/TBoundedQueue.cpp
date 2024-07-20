#include "limonp/BoundedQueue.hpp"
#include "gtest/gtest.h"

using namespace limonp;

TEST(BoundedQueue, Test1) {
  const size_t size = 3;
  BoundedQueue<size_t> que(size);
  ASSERT_EQ(que.Capacity(), size);
  for(size_t i = 0; i < que.Capacity(); i++) {
    que.Push(i);
    ASSERT_EQ(que.Size(), i + 1);
  }
  ASSERT_TRUE(que.Full());
  for(size_t i = 0; que.Size(); i++) {
    ASSERT_EQ(que.Pop(), i);
  }
  ASSERT_TRUE(que.Empty());

  //second time
  for(size_t i = 0; i < que.Capacity(); i++) {
    que.Push(i);
    ASSERT_EQ(que.Size(), i + 1);
  }
  ASSERT_TRUE(que.Full());
  for(size_t i = 0; que.Size(); i++) {
    ASSERT_EQ(que.Pop(), i);
  }
  ASSERT_TRUE(que.Empty());
}


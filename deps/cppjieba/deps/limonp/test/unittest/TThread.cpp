#include "gtest/gtest.h"
#include "limonp/Thread.hpp"
using namespace limonp;

TEST(IThread, Test1) {
  class ThreadHandle: public IThread {
   public:
    ThreadHandle():num(1) {}
    virtual ~ThreadHandle() {}
   private:
   public:
    virtual void Run() {
      ASSERT_EQ(num, 1u);
      num = 2;
    }
   public:
    size_t num;
  };
  {
    ThreadHandle thr;
    thr.Start();
    thr.Join();
    ASSERT_EQ(thr.num, 2u);
  }
  {
    IThread* ptr = new ThreadHandle();
    ptr->Start();
    ptr->Join();
    delete ptr;
  }
}



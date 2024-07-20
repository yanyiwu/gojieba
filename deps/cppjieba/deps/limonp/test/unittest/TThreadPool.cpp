#include "gtest/gtest.h"
#include "limonp/ThreadPool.hpp"
#include <exception>
using namespace limonp;

//static void addOne(void * data)
//{
//    size_t * i = (size_t *) data;
//    (*i) ++;
//}

void Incr(size_t* x) {
  (*x)++;
}

class Exception: public exception {
 public:
  Exception(const string& error)
    : error_(error) {
  }
  virtual ~Exception() throw() {
  }
  virtual const char* what() const throw() {
    return "hello Exception";
  }
 private:
  string error_;
};

void IncrWithThrow(int* x) {
  throw Exception("hello exception!!!");
}

TEST(ThreadPool, Test1) {
  const size_t threadNum = 2;
  vector<size_t> numbers(threadNum);
  {
    ThreadPool threadPool(threadNum);
    threadPool.Start();
    for(size_t i = 0; i < numbers.size(); i ++) {
      numbers[i] = i;
      threadPool.Add(NewClosure(&Incr, &numbers[i]));
    }
  }
  for(size_t i = 0; i < numbers.size(); i++) {
    ASSERT_EQ(i + 1, numbers[i]);
  }
}

TEST(ThreadPool, Exception) {
  const size_t threadNum = 2;
  ThreadPool threadPool(threadNum);
  threadPool.Start();
  
  int x;
  threadPool.Add(NewClosure(&IncrWithThrow, &x));
}

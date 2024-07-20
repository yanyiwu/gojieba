#include "gtest/gtest.h"
#include "limonp/BlockingQueue.hpp"
#include "limonp/BoundedBlockingQueue.hpp"

using namespace limonp;
using namespace std;

class CBlockingQueueTest1 {
 private:
  vector<pthread_t> pthreads_;
 public:
  static void* workerLocked(void * arg) {
    BlockingQueue<size_t> * res = (BlockingQueue<size_t> *)arg;
    size_t t = res->Pop();
    res->Push(t);
    return NULL;
  }
 public:
  CBlockingQueueTest1(const size_t threadSum, void* arg): pthreads_(threadSum) {
    for(size_t i = 0; i < pthreads_.size(); i++) {
      XCHECK(!pthread_create(&pthreads_[i], NULL, workerLocked, arg));
    }
  }
  ~CBlockingQueueTest1() {
  }
 public:
  void Wait() {
    for(size_t i = 0; i < pthreads_.size(); i++) {
      XCHECK(!pthread_join(pthreads_[i], NULL));
    }
  }

};

class CBlockingQueueTest2 {
 public:
  static void * thread_Pop(void * arg) {
    BlockingQueue<size_t> * res = (BlockingQueue<size_t> *)arg;
    for(size_t i = 0; i < 10; i++) {
      res->Pop();
    }
    return NULL;
  }
  static void * thread_push(void * arg) {
    BlockingQueue<size_t> * res = (BlockingQueue<size_t> *)arg;
    for(size_t i = 0; i < 10; i++) {
      usleep(10);
      res->Push(i);
    }
    return NULL;
  }
};
class CBoundedBlockingQueueTest3 {
 public:
  static void * thread_Pop(void * arg) {
    BoundedBlockingQueue<size_t> * res = (BoundedBlockingQueue<size_t> *)arg;
    for(size_t i = 0; i < 10; i++) {
      res->Pop();
    }
    return NULL;
  }
  static void * thread_push(void * arg) {
    BoundedBlockingQueue<size_t> * res = (BoundedBlockingQueue<size_t> *)arg;
    for(size_t i = 0; i < 10; i++) {
      usleep(10);
      res->Push(i);
    }
    return NULL;
  }
};

TEST(BlockingQueue, Test1) {
  size_t threadnum = 3;
  BlockingQueue<size_t> res;
  CBlockingQueueTest1 obj(threadnum, &res);
  //sleep(1);
  res.Push(1);
  obj.Wait();
  ASSERT_EQ(1u, res.Size());
  ASSERT_EQ(1u, res.Pop());

}

TEST(BlockingQueue, Test2) {
  BlockingQueue<size_t> queue;
  pthread_t pth_push;
  pthread_t pth_pop;
  pthread_create(&pth_push, NULL, CBlockingQueueTest2::thread_push, &queue);
  pthread_create(&pth_pop, NULL, CBlockingQueueTest2::thread_Pop, &queue);
  pthread_join(pth_push, NULL);
  pthread_join(pth_pop, NULL);
  ASSERT_TRUE(queue.Empty());
}

TEST(BlockingQueue, Test3) {
  BoundedBlockingQueue<size_t> queue(3);
  pthread_t pth_push;
  pthread_t pth_pop;
  pthread_create(&pth_push, NULL, CBoundedBlockingQueueTest3::thread_push, &queue);
  pthread_create(&pth_pop, NULL, CBoundedBlockingQueueTest3::thread_Pop, &queue);
  pthread_join(pth_push, NULL);
  pthread_join(pth_pop, NULL);
  ASSERT_TRUE(queue.Empty());
}


#include "gtest/gtest.h"
#include "limonp/StdExtension.hpp"
#include "limonp/MutexLock.hpp"
#include <vector>

using namespace limonp;
using namespace std;

static const size_t THREAD_SUM = 3;
static const size_t FOR_SIZE = 3;

static vector<size_t> res;

struct PthreadInfo {
  size_t id;
  pthread_t pthread_id;
  MutexLock* ptMutexLock;
};


class ThreadsLocked {
 private:
  MutexLock mutex_;
  vector<PthreadInfo> pthreadInfos_;
 public:
  ThreadsLocked(size_t threadSum): pthreadInfos_(threadSum) {
  }
  ~ThreadsLocked() {}
 public:
  static void* workerLocked(void * arg) {
    PthreadInfo * ptInfo = (PthreadInfo *) arg;
    MutexLockGuard lock(*ptInfo->ptMutexLock);
    for(size_t i = 0; i < FOR_SIZE; i ++) {
      //cout << ptInfo->id << ':' << i << endl;
      res.push_back(i);
      usleep(100 * i);
    }
    return NULL;
  }
 public:
  void Start() {
    for(size_t i = 0; i < pthreadInfos_.size(); i++) {
      pthreadInfos_[i].id = i;
      pthreadInfos_[i].ptMutexLock = &mutex_;
      assert(!pthread_create(&pthreadInfos_[i].pthread_id, NULL, workerLocked, &pthreadInfos_[i]));
    }
    for(size_t i = 0; i < pthreadInfos_.size(); i++) {
      assert(!pthread_join(pthreadInfos_[i].pthread_id, NULL));
    }
  }
};

TEST(MutexLock, Test1) {
  string str;
  //ThreadsNoLocked noLock(THREAD_SUM);
  //res.clear();
  //noLock.Start();
  //ASSERT_EQ(str << res, "[\"0\", \"0\", \"0\", \"1\", \"1\", \"1\", \"2\", \"2\", \"2\"]");
  ThreadsLocked locked(THREAD_SUM);
  res.clear();
  locked.Start();
  ASSERT_EQ(str << res,  "[0, 1, 2, 0, 1, 2, 0, 1, 2]");
}



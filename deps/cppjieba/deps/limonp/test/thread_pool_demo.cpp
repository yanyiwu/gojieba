#include "limonp/ThreadPool.hpp"
#include "limonp/StdExtension.hpp"

using namespace std;

const size_t THREAD_NUM = 4;

void Increase(int* x) {
  (*x) ++;
}

void DemoNormalFunction() {
  vector<int> numbers(4, 0);
  cout << numbers << endl;
  {
    limonp::ThreadPool thread_pool(THREAD_NUM);
    thread_pool.Start();
    for (size_t i = 0; i < numbers.size(); i++) {
      thread_pool.Add(limonp::NewClosure(&Increase, &numbers[i]));
    }
  }
  cout << numbers << endl;
}

class Numbers {
 public:
  Numbers(size_t num): numbers_(num, 0) {
  }
  void Increase(size_t index) {
    assert(index < numbers_.size());
    numbers_[index]++;
  }

  vector<int> numbers_;
};

void DemoClassFunction() {
  Numbers numbers(4);
  cout << numbers.numbers_ << endl;
  limonp::ThreadPool thread_pool(THREAD_NUM);
  thread_pool.Start();
  for (size_t i = 0; i < numbers.numbers_.size(); i++) {
    thread_pool.Add(limonp::NewClosure(&numbers, &Numbers::Increase, i));
  }
  thread_pool.Stop();
  cout << numbers.numbers_ << endl;
}

int main() {
  DemoNormalFunction();
  DemoClassFunction();
  return 0;
}

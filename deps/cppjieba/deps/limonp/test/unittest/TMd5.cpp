#include "limonp/Md5.hpp"
#include "gtest/gtest.h"
#include <string>
using namespace std;
using namespace limonp;

const char* const DICT_FILE[] = {
  "../test/testdata/jieba.dict.0.utf8",
  "../test/testdata/jieba.dict.0.1.utf8",
  "../test/testdata/jieba.dict.1.utf8",
  "../test/testdata/jieba.dict.2.utf8"
};

const char* const DICT_FILE_MD5[] = {
  "5aef74a56b363d994095c407c4809d84",
  "5aef74a56b363d994095c407c4809d84",
  "55f1116c05c8051ab53171f0b7455197",
  "b123553a2418c4bda51abc64d705d5d4"
};

const char* const TEST_STR[] = {
  "0000"
};

const char* const TEST_STR_MD5[] = {
  "4a7d1ed414474e4033ac29ccb8653d9b"
};

TEST(Md5Test, Test1) {
  ASSERT_EQ(sizeof(DICT_FILE)/sizeof(DICT_FILE[0]), sizeof(DICT_FILE_MD5)/sizeof(DICT_FILE_MD5[0]));
  string tmp;
  for (uint i = 0; i < sizeof(DICT_FILE)/sizeof(DICT_FILE[0]); i++) {
    md5File(DICT_FILE[i], tmp);
    ASSERT_EQ(tmp, string(DICT_FILE_MD5[i]));
  }

  ASSERT_EQ(sizeof(TEST_STR)/sizeof(TEST_STR[0]), sizeof(TEST_STR_MD5)/sizeof(TEST_STR_MD5[0]));
  for (uint i = 0; i < sizeof(TEST_STR)/sizeof(TEST_STR[0]); i++) {
    md5String(TEST_STR[i], tmp);
    ASSERT_EQ(tmp, string(TEST_STR_MD5[i]));
  }
}


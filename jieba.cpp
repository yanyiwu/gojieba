extern "C" {
    #include "jieba.h"
}

#include "jieba/MixSegment.hpp"

using namespace CppJieba;

extern "C" {

Jieba NewJieba(const char* dict_path, const char* hmm_path, const char* user_dict) {
  Jieba handler = (Jieba)(new MixSegment(dict_path, hmm_path, user_dict));
  return handler;
}

void FreeJieba(Jieba handle) {
  MixSegment* x = (MixSegment*)handle;
  delete x;
}

char** Cut(Jieba handle, const char* sentence) {
  MixSegment* x = (MixSegment*)handle;
  vector<string> words;
  x->cut(sentence, words);
  char ** res = (char**)malloc(sizeof(char*) * (words.size() + 1));
  for (size_t i = 0; i < words.size(); i++) {
    res[i] = (char*)malloc(sizeof(char) * (words[i].length() + 1));
    strcpy(res[i], words[i].c_str());
  }
  res[words.size()] = '\0';
  return res;
}

void FreeWords(char** words) {
  char** x = words;
  while (x && *x) {
    free(*x);
    *x = NULL;
    x ++;
  }
  free(words);
}


}

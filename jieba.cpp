extern "C" {
    #include "jieba.h"
}

#include "jieba/Jieba.hpp"

static char** ConvertWords(const std::vector<std::string>& words) {
  char ** res = (char**)malloc(sizeof(char*) * (words.size() + 1));
  for (size_t i = 0; i < words.size(); i++) {
    res[i] = (char*)malloc(sizeof(char) * (words[i].length() + 1));
    strcpy(res[i], words[i].c_str());
  }
  res[words.size()] = '\0';
  return res;
}

Jieba NewJieba(const char* dict_path, const char* hmm_path, const char* user_dict) {
  Jieba x = (Jieba)(new CppJieba::Jieba(dict_path, hmm_path, user_dict));
  return x;
}

void FreeJieba(Jieba x) {
  delete (CppJieba::Jieba*)x;
}

char** Cut(Jieba x, const char* sentence, int is_hmm_used) {
  std::vector<std::string> words;
  ((CppJieba::Jieba*)x)->Cut(sentence, is_hmm_used, words);
  char** res = ConvertWords(words);
  return res;
}

char** CutAll(Jieba x, const char* sentence) {
  std::vector<std::string> words;
  ((CppJieba::Jieba*)x)->CutAll(sentence, words);
  char** res = ConvertWords(words);
  return res;
}

char** CutForSearch(Jieba x, const char* sentence) {
  std::vector<std::string> words;
  ((CppJieba::Jieba*)x)->CutForSearch(sentence, words);
  char** res = ConvertWords(words);
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

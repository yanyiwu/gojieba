extern "C" {
    #include "extractor.h"
}

#include "cppjieba/KeywordExtractor.hpp"

static char** ConvertWords(const std::vector<std::string>& words) {
  char ** res = (char**)malloc(sizeof(char*) * (words.size() + 1));
  for (size_t i = 0; i < words.size(); i++) {
    res[i] = (char*)malloc(sizeof(char) * (words[i].length() + 1));
    strcpy(res[i], words[i].c_str());
  }
  res[words.size()] = NULL;
  return res;
}


Extractor NewExtractor(const char* dict_path, const char* hmm_path, const char* user_dict, const char* idf_path, const char* stop_word_path) {
  return (Extractor)(new cppjieba::KeywordExtractor(dict_path, hmm_path, idf_path, stop_word_path, user_dict));
}

void FreeExtractor(Extractor x) {
  delete (cppjieba::KeywordExtractor*)x;
}

char** Extract(Extractor handle, const char* sentence, int top_k) {
  std::vector<std::string> words;
  ((cppjieba::KeywordExtractor*)handle)->Extract(sentence, words, top_k);
  char** res = ConvertWords(words);
  return res;
}

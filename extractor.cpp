extern "C" {
    #include "extractor.h"
}

#include "cppjieba/KeywordExtractor.hpp"
#include "cppjieba/Jieba.hpp"

Extractor NewExtractor(const Jieba j, const char* idf_path, const char* stop_word_path) {
  const cppjieba::Jieba* x = (const cppjieba::Jieba*)j;
  return (Extractor)(new cppjieba::KeywordExtractor(x->GetDictTrie(), x->GetHMMModel(), idf_path, stop_word_path));
}

void FreeExtractor(Extractor x) {
  delete (cppjieba::KeywordExtractor*)x;
}


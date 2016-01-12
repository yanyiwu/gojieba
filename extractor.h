#ifndef CJIEBA_EXTRACTOR_H
#define CJIEBA_EXTRACTOR_H

#include "jieba.h"

typedef void* Extractor;

Extractor NewExtractor(const Jieba j, const char* idf_path, const char* stop_word_path);
void FreeExtractor(Extractor);

#endif // CJIEBA_EXTRACTOR_H

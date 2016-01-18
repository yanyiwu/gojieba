#ifndef CJIEBA_EXTRACTOR_H
#define CJIEBA_EXTRACTOR_H

#include "util.h"

typedef void* Extractor;

Extractor NewExtractor(const char* dict_path, const char* hmm_path, const char* user_dict, const char* idf_path, const char* stop_word_path);
void FreeExtractor(Extractor);

char** Extract(Extractor handle, const char* sentence, int top_k);

#endif // CJIEBA_EXTRACTOR_H

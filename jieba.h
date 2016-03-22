#ifndef CJIEBA_JIEBA_H
#define CJIEBA_JIEBA_H

#include "util.h"

typedef void* Jieba;

Jieba NewJieba(const char* dict_path, const char* hmm_path, const char* user_dict);
void FreeJieba(Jieba);

char** Cut(Jieba handle, const char* sentence, int is_hmm_used);
char** CutAll(Jieba handle, const char* sentence);
char** CutForSearch(Jieba handle, const char* sentence, int is_hmm_used);
char** Tag(Jieba handle, const char* sentence);

#endif // CJIEBA_JIEBA_H

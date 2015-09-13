#ifndef CJIEBA_JIEBA_H
#define CJIEBA_JIEBA_H

typedef void* Jieba;

Jieba NewJieba(const char* dict_path, const char* hmm_path, const char* user_dict);
void FreeJieba(Jieba);

char** Cut(Jieba handle, const char* sentence, int is_hmm_used);
char** CutAll(Jieba handle, const char* sentence);
char** CutForSearch(Jieba handle, const char* sentence, int is_hmm_used);

void FreeWords(char** words);

#endif

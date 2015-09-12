#ifndef CJIEBA_C_API_H
#define CJIEBA_C_API_H

typedef void* CJieba;
CJieba NewCJieba(const char* dict_path, const char* hmm_path, const char* user_dict);
void FreeCJieba(CJieba);
char* Cut(CJieba handle, const char* sentence, const char* seperator);

#endif

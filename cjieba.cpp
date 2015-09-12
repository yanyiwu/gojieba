extern "C" {
    #include "cjieba.h"
}

#include "jieba/MixSegment.hpp"

using namespace CppJieba;

extern "C" {

CJieba NewCJieba(const char* dict_path, const char* hmm_path, const char* user_dict) {
    CJieba handler = (CJieba)(new MixSegment(dict_path, hmm_path, user_dict));
    return handler;
}

void FreeCJieba(CJieba handle) {
    MixSegment* x = (MixSegment*)handle;
    delete x;
}

char* Cut(CJieba handle, const char* sentence, const char* seperator) {
    MixSegment* x = (MixSegment*)handle;
    vector<string> words;
    x->cut(sentence, words);
    string str = join(words.begin(), words.end(), seperator);
    if(str.empty()) {
        return NULL;
    }
    size_t size = str.size() + 1;
    char* res = (char*)malloc(size);
    memcpy(res, str.c_str(), size);
    return res;
}

}

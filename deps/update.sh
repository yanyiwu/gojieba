echo $1
rm -rf ./cppjieba
#git clone --branch v5.1.2 --single-branch git@github.com:yanyiwu/cppjieba.git
git clone --recurse-submodules --branch $1 --single-branch git@github.com:yanyiwu/cppjieba.git
cd cppjieba
rm -rf .git
cd ..
git add cppjieba


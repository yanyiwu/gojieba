# example: bash update.sh v5.3.0
set -x
echo $1
rm -rf ./cppjieba
git clone --recurse-submodules --branch $1 --single-branch git@github.com:yanyiwu/cppjieba.git
cd cppjieba
rm -rf .git
cd ..
git add cppjieba


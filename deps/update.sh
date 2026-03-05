# example: bash update.sh v5.3.0
# This script updates the bundled cppjieba (and its limonp dependency) by cloning
# the specified tag and embedding all files directly into the repository so that
# no git submodule initialization is needed (compatible with `go mod vendor`).
set -x
echo $1
rm -rf ./cppjieba
git clone --recurse-submodules --branch $1 --single-branch https://github.com/yanyiwu/cppjieba.git
# Remove all .git directories and .gitmodules files so the bundled copy is
# a plain directory tree with no submodule references.
find cppjieba -name ".git" -exec rm -rf {} + 2>/dev/null || true
find cppjieba -name ".gitmodules" -exec rm -f {} + 2>/dev/null || true
cd ..
git add deps/cppjieba


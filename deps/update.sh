# example: bash update.sh v5.3.0
# This script updates the bundled cppjieba by cloning the specified tag or ref
# and embedding all files directly into the repository so that
# no git submodule initialization is needed (compatible with `go mod vendor`).
set -x
echo $1
rm -rf ./cppjieba
git clone --recurse-submodules --branch $1 --single-branch https://github.com/yanyiwu/cppjieba.git
# Remove all .git directories and .gitmodules files so the bundled copy is
# a plain directory tree with no submodule references.
find cppjieba -name ".git" -exec rm -rf {} + 2>/dev/null || true
find cppjieba -name ".gitmodules" -exec rm -f {} + 2>/dev/null || true
# Empty Go packages keep C++ headers and dictionary data in `go mod vendor`.
printf '%s\n' '// Package cppjieba contains the bundled C++ headers.' 'package cppjieba' > cppjieba/include/cppjieba/pkg.go
printf '%s\n' '// Package dict contains the bundled cppjieba dictionaries.' 'package dict' > cppjieba/dict/pkg.go
printf '%s\n' '// Package pos_dict contains the bundled POS dictionaries.' 'package pos_dict' > cppjieba/dict/pos_dict/pkg.go
git add cppjieba

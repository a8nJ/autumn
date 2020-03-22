#!/bin/dash
set -e -u
git clone --depth 1 git://github.com/curl/curl
cd curl
mingw32-make -C lib -f Makefile.m32 -j 5
mkdir -p /Path/include/curl /Path/lib
cp lib/libcurl.a /Path/lib
cp include/curl/*.h /Path/include/curl

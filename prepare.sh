#!/bin/bash

cd "$( dirname "${BASH_SOURCE[0]}" )"

test -d adaptagrams || git clone https://github.com/mjwybrow/adaptagrams.git

if [[ `uname` == 'Linux' ]]; then
	exit 0
fi

# prepare library
rm -f libvpsc.a
mkdir -p build
cd build
g++ -c -O3 -I ../adaptagrams/cola ../adaptagrams/cola/libvpsc/*.cpp
ar rcs libvpsc.a *.o

g++ -c -O3 -I ../adaptagrams/cola -DBUILDING_VPSC_DLL -o vpsc_dll.o ../windll/vpsc.cpp
g++ -shared -o vpsc.dll vpsc_dll.o -L . -l vpsc
rm libvpsc.a *.o
cd ..

cp build/vpsc.dll ../..

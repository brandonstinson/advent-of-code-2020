#!/bin/bash

if [ -f answers.txt ] ; then
	rm answers.txt
fi

for d in d* ; do
	cd ${d}
	go run ${d}.go >> ../answers.txt
	printf -- '-%.0s' {1..30} >> ../answers.txt
	echo "" >> ../answers.txt
	cd ..
done

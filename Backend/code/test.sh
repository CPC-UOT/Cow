#!/usr/bin/env bash

cd "/home/coll/Dev/Cow/Backend/code"

for i in $(seq 1 10);
do
	cp $"./data/$i.in" ./blist.in
	./bb
	diff -Z -B -b --strip-trailing-cr $"./data/$i.out" ./blist.out
	if [ "$?" -eq 1 ];
	then
		echo "$i"
		rm ./bb
		echo "your wrong output:"
		bat ./blist.out
		echo "expected:"
		bat $"./data/$i.out"
		rm ./blist.out
		cp $"./basic.in" ./blist.in
		cp $"./basic.out" ./blist.out
		exit 1
	fi
done
cp $"./basic.in" ./blist.in
rm ./bb

#!/bin/bash

for sc_awk in `ls *-awk.sh`
do
    sc_ago=`echo $sc_awk | sed -e 's/-awk.sh/-ago.sh/'`
    #if [[ ! -f $sc_awk ]];
    #then
	#echo "ERROR: $sc_awk does not exist"
	#continue;
    #fi
    if [[ ! -f $sc_ago ]];
    then
	#echo "ERROR: $sc_ago does not exist"
	#continue;
	sed -e 's/^awk/ago -g/' $sc_awk > $sc_ago
	chmod u+x $sc_ago
    fi
    diff <(./$sc_awk) <(./$sc_ago)
    result=$?
    if [ $result -ne 0 ];
    then
	echo "ERROR: $sc_awk $sc_ago"
	continue
    fi
    echo "Passed: $sc_awk $sc_ago"
done

FILE=countries
for sc_awk in `ls *.awk`
do
    sc_ago=`echo $sc_awk | sed -e 's/.awk/.ago/'`

    if [[ ! -f $sc_ago ]];
    then
	sc_ago=$sc_awk
    fi
    diff <(awk -f $sc_awk $FILE) <(ago -g -f $sc_ago $FILE)
    result=$?
    if [ $result -ne 0 ];
    then
	echo "ERROR: $sc_awk $sc_ago"
	continue
    fi
    echo "Passed: $sc_awk $sc_ago"
done

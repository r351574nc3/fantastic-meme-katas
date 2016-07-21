#!/bin/sh


count() {
    FILENAME=$1
    while read x; do 
        echo $(grep -caF "$x" $FILENAME) $x
    done < $FILENAME | sort -n | uniq
}

count $1 > "$1".sorted

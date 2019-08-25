#!/bin/bash
# build_plugin.sh builds one go plugin, given a year and day

# $# is the number of arguments
# If the number of arguments is not 2, then print usage instructions and exit
if [[ $# -ne 2 ]]; then
    echo "Usage: ./build_plugin.sh <year> <day>"
    exit 1
fi

YEAR=$1
DAY=$2

# If YEAR does not match the regex ^201[5-8]$, print error and exit
if ! [[ $YEAR =~ ^201[5-8]$ ]]; then
    echo "Invalid <year>. Must be between 2015 and 2018, inclusive."
    exit 1
fi

# If DAY is not an integer between 1 and 31, print error and exit
if ! [[ $DAY =~ ^[012]?[0-9]$ ]] && ! [[ $DAY =~ ^3[01]$ ]]; then
    echo "Invalid <day>. Must be between 1 and 31, inclusive."
    exit 1
fi

# If DAY is one digit, left-pad with a zero
if [[ $DAY =~ ^[0-9]$ ]]; then
    DAY="0$DAY"
fi

DIR="puzzles/$YEAR/$DAY"

# If DIR exists, cd into it and run build script
if [[ -d $DIR ]]; then
    cd $DIR
    go build -buildmode=plugin -o $YEAR$DAY.so
else
    echo "$DIR does not exist"
    exit 1
fi
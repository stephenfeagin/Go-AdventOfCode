#/bin/bash
# new_puzzle.sh creates a new folder and copies template files for a new day's puzzle

# If the number of arguments is not 2, then print usage instructions and exit
if [[ $# -ne 2 ]]; then
    echo "Usage: ./new_puzzle.sh <year> <day>"
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

if [[ -d $DIR ]]; then
    echo "Error: Directory \"$DIR already exists.\""
    exit 1
fi

mkdir -p $DIR

cp template/input.txt.templ $DIR/input.txt
cp template/main.go.templ $DIR/main.go

#!/bin/bash
# builds all go plugins in their year/day directories

# cd into puzzles directory
cd puzzles

# every subdir in puzzles is one year
for YEAR in *; do

    # cd into the year dir
    cd $YEAR
    echo "$YEAR"

    # every subdir in $YEAR is one day
    for DAY in *; do

        # cd into the day dir
        cd $DAY
        echo -e "\t$DAY"

        # if there's a main.go file, run go build
        if [[ -f main.go ]]; then
            go build -buildmode=plugin -o $YEAR$DAY.so
        fi

        # cd back up to the year dir
        cd ..
    done
    # cd back into the puzzles dir
    cd ..
done

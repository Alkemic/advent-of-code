#!/bin/bash -eu

input_file="${1:-/dev/stdin}"

length=0

i=0
declare -A letters
while read -r line; do
    for ((j = 0; j < ${#line}; j++)); do
        letters[$i,$j]="${line:j:1}"
    done
    length=${#line}
    ((i+=1))
done < "${input_file}"

height=$((${#letters[@]}/$length))

count=0
check-string() {
    # passing associative array as a reference
    local -n letters_ref=$1
    local i=$2
    local j=$3
    local searched=$4

#    local count=0
    set +u # lazy workaround, so we don't have to check bounds

    # left to right
    txt="${letters_ref[$i,$j]}${letters_ref[$i,$((j+1))]}${letters_ref[$i,$((j+2))]}${letters_ref[$i,$((j+3))]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi
    # right to right
    txt="${letters_ref[$i,$j]}${letters_ref[$i,$((j-1))]}${letters_ref[$i,$((j-2))]}${letters_ref[$i,$((j-3))]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi

    # top to bottom
    txt="${letters_ref[$i,$j]}${letters_ref[$((i+1)),$j]}${letters_ref[$((i+2)),$j]}${letters_ref[$((i+3)),$j]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi
    # bottom to top
    txt="${letters_ref[$i,$j]}${letters_ref[$((i-1)),$j]}${letters_ref[$((i-2)),$j]}${letters_ref[$((i-3)),$j]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi

    # left-top to right-bottom
    txt="${letters_ref[$i,$j]}${letters_ref[$((i+1)),$((j+1))]}${letters_ref[$((i+2)),$((j+2))]}${letters_ref[$((i+3)),$((j+3))]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi
    # left-bottom to right-top
    txt="${letters_ref[$i,$j]}${letters_ref[$((i-1)),$((j-1))]}${letters_ref[$((i-2)),$((j-2))]}${letters_ref[$((i-3)),$((j-3))]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi

    # left-top to right-bottom
    txt="${letters_ref[$i,$j]}${letters_ref[$((i+1)),$((j-1))]}${letters_ref[$((i+2)),$((j-2))]}${letters_ref[$((i+3)),$((j-3))]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi
    # left-bottom to right-top
    txt="${letters_ref[$i,$j]}${letters_ref[$((i-1)),$((j+1))]}${letters_ref[$((i-2)),$((j+2))]}${letters_ref[$((i-3)),$((j+3))]}"
    if [[ "$txt" == "$searched" ]]; then
        ((count+=1))
    fi

    set -u

    # doing via echo, and then adding in main loop makes it 20x times slower...
#    echo "$count"
}

#found=0
for (( i = 0; i < length; i++ )); do
    for (( j = 0; j < height; j++ )); do
        # limit brute-forcing to starting positions of 'X'
        if [[ "${letters[$i,$j]}" == "X" ]]; then
            check-string letters $i $j "XMAS"
#            f=$(check-string letters $i $j "XMAS")
#            found=$(($found + $f))
        fi
    done
done

#echo "found $found"
echo "found $count"

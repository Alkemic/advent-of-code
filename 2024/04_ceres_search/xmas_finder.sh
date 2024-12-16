#!/bin/bash -eux

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

x-mas-find() {
    local -n letters_ref=$1
    local i=$2
    local j=$3

    local count=0
    set +u # lazy workaround, so we don't have to check bounds

    # left to right
    txt="${letters_ref[$((i-1)),$((j-1))]}${letters_ref[$((i-1)),$((j+1))]}${letters_ref[$((i+1)),$((j-1))]}${letters_ref[$((i+1)),$((j+1))]}"
    if [[ "$txt" == "MMSS" || "$txt" == "SSMM" || "$txt" == "MSMS" || "$txt" == "SMSM" ]]; then
        exit 0
    fi
    set -u

    exit 1
}

found=0
for (( i = 0; i < length; i++ )); do
    for (( j = 0; j < height; j++ )); do
        # limit brute-forcing to starting positions of 'A', center of X-MAS
        if [[ "${letters[$i,$j]}" == "A" ]] && (x-mas-find letters $i $j); then
            ((found+=1))
        fi
    done
done

echo "found $found"

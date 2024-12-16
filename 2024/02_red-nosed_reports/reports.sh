#!/bin/bash -eu

input_file="${1:-/dev/stdin}"

input=$(< "${input_file}")

process() {
    data=($1)

    if ((data[0] < data[1])); then
        asc=true
    elif ((data[0] > data[1])); then
        asc=false
    else
        return 1
    fi

    for (( i = 1; i < ${#data[@]}; i++ )); do
        prev=${data[i-1]}
        cur=${data[i]}

        delta=$((cur-prev))
        delta=${delta#-}

        if ((delta > 3 || delta < 1)); then
            return 1
        fi

        if $asc && ((prev > cur)); then
            return 1
        fi

        if ! $asc && ((prev < cur)); then
            return 1
        fi
    done

    return 0
}

safe=0
while read -r line; do
    if process "$line"; then
        ((safe+=1))
    fi
done <<< $input

echo "$safe"

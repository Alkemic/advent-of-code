#!/bin/bash -eu

input_file="${1:-/dev/stdin}"

process() {
    local data=($@)
    local asc
    if ((data[0] < data[1])); then
        asc=true
    elif ((data[0] > data[1])); then
        asc=false
    else
        return 1
    fi

    local i # fuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu
    for (( i = 1; i < ${#data[@]}; i++ )); do
        local prev=${data[i-1]}
        local cur=${data[i]}

        local delta=$((cur-prev))
        local delta=${delta#-}

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
while read -ra line; do
    if process "${line[@]}"; then
        ((safe++))
    else
        len=${#line[@]}
        for (( i = 0; i < len; i++ )); do
            arr=("${line[@]:0:i}" "${line[@]:i+1}")
            if process "${arr[@]}"; then
                ((safe+=1))
#                ((safe++)) # todo: why this is no working?!
                break
            fi
        done
    fi
done < "${input_file}"

echo "$safe"

#!/bin/bash -eu

input_file="${1:-/dev/stdin}"

input=$(< "${input_file}")

nums=()
occurences=()
while IFS=" " read -r left right; do
    nums+=($left)
    ((occurences[$right]+=1))
done <<< $input

sum=0
for num in ${nums[*]}; do
    dist=$((${occurences[$num]:-0} * $num))
    # ((sum+=dist)) # why this breaks this script?
    sum=$(($sum+$dist))
done

echo $sum

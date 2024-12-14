#!/bin/bash -eu

input_file="${1:-/dev/stdin}"

input=$(< "${input_file}")

left_nums=()
right_nums=()
while IFS=" " read -r value1 value2; do
    left_nums+=("$value1")
    right_nums+=("$value2")
done <<< $input

left_nums=($(printf '%d\n' "${left_nums[@]}" | sort -n))
right_nums=($(printf '%d\n' "${right_nums[@]}" | sort -n))

sum=0
len=${#right_nums[@]}
for (( i=0; i < ${len}; i++ )); do
    dist=$((${left_nums[$i]}-${right_nums[$i]}))
    dist=${dist#-} # abs value
    ((sum+=dist))
done

echo $sum

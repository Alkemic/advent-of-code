#!/bin/bash -eu

regexp='mul\(([0-9]{1,3}),([0-9]{1,3})\)'

input_file="${1:-/dev/stdin}"

input=$(< "${input_file}")
input=${input//)/)$'\n'}

out=0
while read -r line; do
    if [[ $line =~ $regexp ]]; then
        ((out+=$((BASH_REMATCH[1]*BASH_REMATCH[2]))))
    fi
done <<< "$input"

echo $out

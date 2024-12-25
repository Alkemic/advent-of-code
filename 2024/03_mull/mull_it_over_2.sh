#!/bin/bash -eu

input_file="${1:-/dev/stdin}"
input=$(< "${input_file}")

input=${input//$'\n'} # rm newlines

# insert bnewline before do/don't, for easy grepping
input=${input//do()/$'\n'do()}
input=${input//don\'t()/$'\n'don\'t()}

# todo: find a pure bash way
input=${input##'t\(\).*'}
input=$(grep -v "^don't()" <<< $input)
input=${input//)/)$'\n'} # break after each closing bracket, for easier parsing


result=0
regexp='mul\(([0-9]{1,3}),([0-9]{1,3})\)'
while read -r line; do
    if [[ $line =~ $regexp ]]; then
        ((result+=$((BASH_REMATCH[1]*BASH_REMATCH[2]))))
    fi
done <<< "$input"

echo $result

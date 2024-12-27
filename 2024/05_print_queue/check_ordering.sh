#!/bin/bash -eu

input_file="${1:-/dev/stdin}"

# contains x [a,b,c]
contains() {
    echo "c $@"
    local needle=$1
    shift
    local hay
    for hay in "$@"; do
        [[ "$needle" == "$hay" ]] && return 0
    done

    return 1
}

# check if pages follow `ordering` rules
check() {
    local pages=("$@")
    for (( i = 0; i < ${#pages[@]}; i++ )); do
        local page=${pages[i]}

        local afters
        read -ra afters <<< "${ordering[page]-}"

        # for each page in rule, check if it's not before current page, if so
#        echo "page $page: rules ${rules[*]}"
        for after in "${afters[@]}"; do
            # check if page this must be `after`, current one is in previous pages, making this check false
            # if number that must be after is before, then we fail this check
            if contains "${after}" "${pages[@]:i+1}"; then
                return 1
            fi
        done
    done
    return 0
}



# num => list of pages that must after `num`
declare ordering=()
sum=0
updates=()
while read -r line; do
    # matching ordering rules lines
    if [[ $line == *'|'* ]]; then
        IFS='|' read -r prev next <<< "$line"
        ordering[$prev]+="$next "
#        ordering[$next]+="$prev "
    fi

    # matching pages list lines
    if [[ $line == *","* ]]; then
#        echo $line
        updates+=($line)
#        IFS=, read -ra pages <<< "$line"
#
#        if check "${pages[@]}"; then
#            len=${#pages[@]}
#            mid=$((len/2))
#            ((sum+=pages[mid]))
#        fi
    fi
done < "${input_file}"


for update in "${updates[@]}"; do
    IFS=, read -ra pages <<< "$update"

    if check "${pages[@]}"; then
        len=${#pages[@]}
        mid=$((len/2))
        ((sum+=pages[mid]))
    fi
done

echo "$sum"

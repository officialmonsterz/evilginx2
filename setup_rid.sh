#!/usr/bin/env bash

script_name="evilginx2 rid setup"

function print_good () {
    echo -e "[${script_name}] \x1B[01;32m[+]\x1B[0m $1"
}

function print_error () {
    echo -e "[${script_name}] \x1B[01;31m[-]\x1B[0m $1"
}

function print_info () {
    echo -e "[${script_name}] \x1B[01;34m[*]\x1B[0m $1"
}

if [[ $# -ne 1 ]]; then
    print_error "Usage: ./setup_rid.sh <rid replacement>"
    print_error "Example: ./setup_rid.sh user_id"
    print_error "Replaces 'client_id' with your custom string everywhere in .go/.yaml/.json files, then rebuilds."
    exit 2
fi

rid_replacement="${1}"

function replace_rid () {
    print_info "Replacing 'client_id' with '${rid_replacement}'..."
    find . -type f -name "*.go" -exec sed -i "s|client_id|${rid_replacement}|g" {} \;
    find . -type f -name "*.yaml" -exec sed -i "s|client_id|${rid_replacement}|g" {} \;
    find . -type f -name "*.json" -exec sed -i "s|client_id|${rid_replacement}|g" {} \;
    print_good "Replaced all occurrences!"

    print_info "Rebuilding project..."
    go build -o evilginx2 .
    print_good "Rebuild complete! RID is now: ${rid_replacement}"
}

function main () {
    replace_rid
    print_good "RID replacement complete!"
}

main

#!/usr/bin/env bash
script_name="replace rid"
function print_good () { echo -e "[${script_name}] \x1B[01;32m[+]\x1B[0m $1"; }
function print_error () { echo -e "[${script_name}] \x1B[01;31m[-]\x1B[0m $1"; }

if [[ $# -ne 2 ]]; then
    print_error "Usage: ./replace_rid.sh <previous rid> <new rid>"
    print_error "Example: ./replace_rid.sh user_id client_id"
    exit 2
fi

previous_rid="${1}"
new_rid="${2}"

function main () {
    find . -type f -name "*.go" -exec sed -i "s|${previous_rid}|${new_rid}|g" {} \;
    find . -type f -name "*.yaml" -exec sed -i "s|${previous_rid}|${new_rid}|g" {} \;
    find . -type f -name "*.json" -exec sed -i "s|${previous_rid}|${new_rid}|g" {} \;
    go build -o evilginx2 .
    print_good "Replaced rid and rebuilt successfully!"
}

main

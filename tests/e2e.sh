#!/bin/bash

set -e


check_output_matches_file() {
    # Usage: check_output_matches_file "command" file_path
    # Run the command and capture the output
    local output
    output=$(eval "$1")

    # Read the content of the file
    local expected_content
    expected_content=$(<"$2")

    # Compare the output and expected content
    if [[ "$output" == "$expected_content" ]]; then
        echo "Success"
    else
        echo "Failed"
        echo ""
        echo "Output:"
        echo "$output"
        echo ""
        echo "Expected:"
        echo "$expected_content"
        exit 1
    fi
}


echo -e "\n\n### Test 1"
check_output_matches_file "../bin/top10url -file ./testdata/1_in.txt -top 2" ./testdata/1_out.txt

echo -e "\n\n### Test 2"
check_output_matches_file "../bin/top10url -file ./testdata/2_in.txt -top 3" ./testdata/2_out.txt

echo -e "\n\n### Test 3"
check_output_matches_file "../bin/top10url -file ./testdata/3_in.txt -top 3" ./testdata/3_out.txt



#!/bin/zsh

# Function to process files
process_file() {
    local file=$1
    # Use sed to delete the empty line following "fmt"
    sed -i '/import (/{
        N
        /import (.*\n.*"fmt"/{
            N
            /import (.*\n.*"fmt".*\n^$/d
        }
    }' "$file"
}

# Iterate over all files in the current directory
for file in *; do
    if [[ -f $file ]]; then
        # Check if the file contains the specified pattern
        if grep -q 'import (\n\s*"fmt"\n\s*-\n\s*"piscine"\n)' "$file"; then
            echo "Processing file '$file'."
            process_file "$file"
        fi
    fi
done
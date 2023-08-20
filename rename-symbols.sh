#!/bin/bash

# Loop through files in the current directory
for filename in *; do
    #echo "$filename" == *"\:"*
    #echo "$filename" == *":"*
    #echo "$filename" | tr ':' '-'
    #echo "$filename" | sed 's/[:：-]/_/g'
    # Check if the filename contains double-quote characters
#    if [[ "$filename" =~ "[-]" ]]; then
        # Replace double-quote with hyphen and rename the file
        #new_filename=$(echo "$filename" | tr ':' '-')
        new_filename=$(echo "$filename" | sed 's/[,?|&:：]/- /g')
        mv "$filename" "$new_filename"
        echo -e "Renamed:\n $filename ->\n $new_filename "
 #   fi


done

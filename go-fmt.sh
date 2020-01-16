#!/usr/bin/env bash

function cd_go_fmt() { # go fmt .
    for cf in `pwd`/*; do
        if [ -d $cf ]; then
            echo $cf
            cd $cf
            file=$(ls)
            # echo $file
            if [ "${file##*.}"x = "go"x ]; then
                go fmt .
                # echo $cf
                echo ''
            else
                cd_go_fmt
            fi
            cd ..
        fi
    done
}

cd_go_fmt


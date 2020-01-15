#!/usr/bin/env bash

function cd_go_fmt() { # go fmt .
    for cf in `pwd`/*; do
        if [ -d $cf ]; then
            echo $cf
            cd $cf
            if [ -d ".go" ]; then
                go fmt .
                echo ''
            else
                cd_go_fmt
            fi
            cd ..
        fi
    done
}

cd_go_fmt


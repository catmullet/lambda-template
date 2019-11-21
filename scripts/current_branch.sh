#!/bin/bash

get_branch() {
    # not a good idea to use this to actually checkout git branches
    # this is mostly just good for use in getting the environment from get_environment()

    # TODO remember to turn this back on
    # if we are in a jenkins build then the BRANCH_NAME env var will be set
    if [ -z ${BRANCH_NAME} ]; then
    		# one way to get the current branch
    		git branch | grep "\*" | cut -d' ' -f 2
    else
    		echo ${BRANCH_NAME}
    fi
}

get_environment() {
    local BRANCH_NAME=$(get_branch)
    echo ${BRANCH_NAME} | grep -q master && echo production || echo $(get_branch)
}
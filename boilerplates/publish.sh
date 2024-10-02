#!/bin/bash

# If any command on pipe fails, return failure
set -o pipefail

function yesno {
    unset response
    while [ "$response" != "y" ] && [ "$response" != "n" ]
    do
        echo -n "$@"
        echo -n " [y/n] "
        read response
        response=`echo $response | tr '[:upper:]' '[:lower:]'`
    done
    [ "$response" == "y" ] && return 0 || return 1
}

yesno "This will publish current branch code to master branch, continue?" || exit 0

# Return: Data in the format dd/mm/yyy
currentDate=`date +%d/%m/%Y`
challenge=`basename "$PWD"`
main_branch="main"

current_branch=$(git rev-parse --abbrev-ref HEAD)

# Checking if current branch is different than master
if [ $current_branch != $main_branch ]; then
    echo "You should be at $main_branch. Please, move this branch's code to $main_branch. 😞"
    echo "  git checkout $main_branch"
    exit 1
fi

# Updating repository
echo "Updating repository"
git fetch --all
git pull --rebase origin $main_branch

# Commiting the changes made
echo "Commiting the day's dojo files"
git add .
git commit -m "Dojo do dia ${currentDate} - Desafio $challenge"

# Pushing the changes to master
git push origin $main_branch
echo "Your commits were pushed sucessfully ⛅"

exit 0

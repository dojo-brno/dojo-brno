#! /bin/bash -e

# Basic color values..
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
reset=`tput sgr0`


MY_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH="$MY_PATH/.."


FOLDER="$MY_PATH/../2019/2019-09-12/fizzbuzz"
cd $FOLDER && go fmt && go test -v | sed /PASS/s//$(printf "${green}PASS${reset}")/ | sed /FAIL/s//$(printf "${red}FAIL${reset}")/

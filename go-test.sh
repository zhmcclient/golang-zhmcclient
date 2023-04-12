#!/bin/sh
args=("$@")
package=${args[0]}
file=${args[1]}

coveroutput="cover.out"
tmpfile="tmp.out"


function check_env() {

  if [ "$TRAVIS" = true ]; then
   if [ -z "$GH_USERNAME" ]; then
         echo "github.ibm.com username missing, cannot pull private dependencies. Exiting ..."
         exit 1
    fi
    if [ -z "$GH_POA" ]; then
         echo "github.ibm.com token missing, cannot pull private dependencies. Exiting ..."
         exit 1
    fi
    set_env
  fi
}

# setup git credentials for github.ibm.com so we can pull private go modules
function set_env() {
  echo "Setting up git credentials for github.ibm.com..."
  GIT_CRED_FILE="$(pwd)/.git-credentials"
  echo "https://${GH_USERNAME}:${GH_POA}@github.ibm.com" > $GIT_CRED_FILE
  git config --global credential.helper "store --file=${GIT_CRED_FILE}"
}

check_env


if [[ $package =~ ^[a-z]{0,2}$ ]];
  then
    echo "Invalid package name, please enter valid package and retry command"
    exit 1
fi

test_coversummary=`go test -v -cover ./... | grep  -E -A15 --text "(^\-\-\-\-|Fail|Pass).*"`
test_cover_result=`echo "$test_coversummary" | grep --text "^--- "| cut -d " " -f2| sed "s/://g"`

#check if the test cases are passing before getting the coverage
if [[ $test_cover_result == "FAIL" ]];then
  echo -e "$test_coversummary"
  exit 1
fi
go_test_coverout=`go test ./$package -coverprofile=$coveroutput`
go tool cover -func=$coveroutput > $tmpfile
unlink $coveroutput
coveragesum=0
filecounter=0
if [[ $file != "" ]]; then
  while read line; do
    regex=".*"$file+"\.go.*"
    if [[ $line =~ $regex ]]; then
      echo $line
      coverage=`echo "$line" | awk '{print $3}' | sed -E "s/\..*%//g"`
      filecounter=$((filecounter + 1))
      coveragesum=$((coveragesum + coverage))
    fi
  done < $tmpfile;
  unlink $tmpfile
  if [[ $coveragesum != 0  ]];then
    echo ""
    echo "Coverage for $file.go: $(( $coveragesum / $filecounter))%"
  else
    echo "No coverage found for the $file.go"
  fi
else
  go_test_coverout=`go test ./$package -coverprofile=$coveroutput`
  go tool cover -func=$coveroutput
  unlink $coveroutput
fi

test_result=`echo $go_test_coverout | cut -d ":" -f2 | cut -d " " -f2 | awk '{printf("%d\n",$0+=$0<0?0:0.9)}'`
# If coverage is below 65 then fail the build
if [[ $test_result -lt 65 ]]; then
  exit 1
fi

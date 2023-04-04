#!/bin/sh
args=("$@")
package=${args[0]}
file=${args[1]}

coveroutput="cover.out"
tmpfile="tmp.out"


function check_env() {

   if [ -z "$GH_USERNAME" ]; then
         echo "github.ibm.com username missing, cannot pull private dependencies. Exiting ..."
         exit 1
    fi
    if [ -z "$GH_POA" ]; then
         echo "github.ibm.com token missing, cannot pull private dependencies. Exiting ..."
         exit 1
    fi

}

check_env

# setup git credentials for github.ibm.com so we can pull private go modules
echo "Setting up git credentials for github.ibm.com..."
GIT_CRED_FILE="$(pwd)/.git-credentials"
echo "https://${GH_USERNAME}:${GH_POA}@github.ibm.com" > $GIT_CRED_FILE
git config --global credential.helper "store --file=${GIT_CRED_FILE}"

if [[ $package =~ ^[a-z]{0,2}$ ]];
  then
    echo "Invalid package name, please enter valid package and retry command"
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

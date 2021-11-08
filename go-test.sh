#!/bin/sh
args=("$@")
package=${args[0]}
file=${args[1]}

coveroutput="cover.out"
tmpfile="tmp.out"

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

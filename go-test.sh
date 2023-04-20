#!/bin/sh
args=("$@")
v_flag=''
f_flag=''
p_flag=''

while getopts 'fpv:' flag; do
  case "${flag}" in
    f) 
        file=${OPTARG}
        f_flag='true'
        ;;
    v) 
        package=${OPTARG}
        v_flag='true'
        ;;
    p) 
        package=${OPTARG}
        p_flag='true'
        ;;
    *) exit 1
       ;;
  esac
done

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

function run_vul_check() {
  package=${args[1]}
  go install golang.org/x/vuln/cmd/govulncheck@latest
  scan_result=`govulncheck -json ./... | jq -r '.vulnerability | {cves: .osv.aliases, id: .osv.id, details: .osv.details, packages: .modules, url: .osv.database_specific.url}' | jq --sort-keys | jq -s .| jq '.[] | with_entries( select( .value != null and . !={} )) | del(..|select(. == {}))| select( . != null)' | jq -s .`
  echo "$scan_result"
  vul_count=`echo $scan_result | jq length`
  echo "$vul_count found in go package: $package"
}

function run_unit_tests() {

  file=${args[2]}
  package=${args[1]}

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
}

check_env

if [[  $f_flag == 'true'  ]]; then
  run_unit_tests $package $file
fi

if [[  $v_flag == 'true'  ]]; then
  run_vul_check $package
fi

if [[  $p_flag == 'true'  ]]; then
  run_unit_tests $package
fi
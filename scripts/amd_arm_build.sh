#!/bin/bash

# get go mod name
go_mod=$(head -n 1 go.mod | awk '{print $2}')
echo "go mod: $go_mod"

# get last commit hash
git_commit=$(git log --pretty=format:'%H' -1)   # commit hash
abbr_git_commit=$(git log --pretty=format:'%h' -1)   # abbreviated commit hash
echo "last commit hash: $git_commit[long]  $abbr_git_commit[short]"

# build module and version
while getopts m:v: opt
do 
	case "${opt}" in
		m) module=${OPTARG};;
		v) version=${OPTARG};;
	esac
done
version=${version:-$abbr_git_commit}
echo "build version: $version"
module=${module:-$(ls cmd | head -n 1)}
echo "build module: $module"

amd64_name="$go_mod-x86_64-$version"
arm64_name="$go_mod-arm64-$version"

build_time=$(date "+%Y-%m-%d %H:%M:%S")
go_version=$(go version)
remote_origin=$(git remote get-url origin)

build_infos="\
-X 'main.gitHash=$git_commit' \
-X 'main.buildTime=$build_time' \
-X 'main.goVersion=$go_version' \
-X 'main.remoteOrigin=$remote_origin'\
"

module_path="cmd/$module"
if [ ! -d "$module_path" ]; then
  echo "module dir does not exist: $module_path"
  exit 1
fi

# linux amd64 
cd $module_path
export GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && \
    go build -a -ldflags "-extldflags '-static' $build_infos" -o "$amd64_name"
cd ../../ && mv $module_path/$amd64_name ./
# linux arm64 
cd $module_path
export GOOS=linux GOARCH=arm64 CGO_ENABLED=0 && \
    go build -a -ldflags "$build_infos" -o "$arm64_name"
cd ../../ && mv $module_path/$arm64_name ./

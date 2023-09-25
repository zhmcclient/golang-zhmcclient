# golang-zhmcclient

zhmcclient - A golang client library for the IBM Z HMC Web Services API

## Generate Fake APIs

```bash
cd ./pkg/zhmcclient
go get github.com/maxbrunsfeld/counterfeiter/v6
export COUNTERFEITER_NO_GENERATE_WARNING=true
go generate ./...
```

## Build

```bash
make sample-build
make sample-build-mac
make sample-build-docker
```

## Unit Test

```bash
# 'FILE' corresponds to the filename w/o .go extentsion
# 'PKG' package to be tested
# If only PKG is provided all files (go modules) under the package will be tested
make unit-test PKG=pkg/zhmcclient FILE=lpar
```

## Sample Usage

```bash
make sample-build
export HMC_ENDPOINT="https://192.168.195.118:6794"
export HMC_USERNAME=${username}
export HMC_PASSWORD=${password}
./bin/sample
```

or

```bash
make sample-build-mac
export HMC_ENDPOINT="https://192.168.195.118:6794"
export HMC_USERNAME=${username}
export HMC_PASSWORD=${password}
./bin/sample-mac
```

## Contributing

Third party contributions to this project are welcome!

In order to contribute, create a [Git pull request](https://help.github.com/articles/using-pull-requests/), considering this:

* Test is required.
* Each commit should only contain one "logical" change.
* A "logical" change should be put into one commit, and not split over multiple
  commits.
* Large new features should be split into stages.
* The commit message should not only summarize what you have done, but explain
  why the change is useful.
* The commit message must follow the format explained below.

What comprises a "logical" change is subject to sound judgement. Sometimes, it
makes sense to produce a set of commits for a feature (even if not large).
For example, a first commit may introduce a (presumably) compatible API change
without exploitation of that feature. With only this commit applied, it should
be demonstrable that everything is still working as before. The next commit may
be the exploitation of the feature in other components.

For further discussion of good and bad practices regarding commits, see:

 - [OpenStack Git Commit Good Practice](https://wiki.openstack.org/wiki/GitCommitMessages)

 - [How to Get Your Change Into the Linux Kernel](https://www.kernel.org/doc/Documentation/process/submitting-patches.rst)


## License

The zhmcclient package is licensed under the [Apache 2.0 License](https://github.com/zhmcclient/golang-zhmcclient/blob/master/LICENSE).
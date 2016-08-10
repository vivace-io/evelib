# evelib

[ ![Codeship Status for vivace-io/evelib](https://codeship.com/projects/5d940200-40da-0134-fc42-4202afd6e25f/status?branch=master)](https://codeship.com/projects/167799)

Library for connecting to and using Eve Online and other related 3rd party services.

This collection of libraries is still in a very early stage of development, and APIs are subject to change without notice on master.

**Stable API Versions:** when a version is deemed stable

Stable API versions will be tagged on the repository, and released through gopkg.in!

## Version 0

`go get gopkg.in/vivace-io/evelib.v0`

## Usage

This repository contains a number of libraries. For documentation on usage and examples, view the README.md in each package directory. Proper source documentation through godoc is a WIP and expected to be out soon.

## Contributing

Contributions, suggestions and feedback are always welcome!

Releases are versioned following [Semantic Versioning 2.0](http://semver.org/spec/v2.0.0.html).

### Working on the Code

**Please read the Pull Requests Section before making a pull request.**

This library is versioned through gopkg.in, and some libraries are inter-dependent on one another (the `zkill` depends on the package `crest`, for example).

With that in mind, you will have to do a little more than run the typical `go get`:

 1. Create a fork of this repository on GitHub (lets assume its `github.com/example/evelib`)
 2. On your development machine, run `go get gopkg.in/vivace-io/evelib.v#` where `#` is either the latest version or the version you intend to work on.
 3. Your `GOPATH` should now look like `$GOPATH/src/gopkg.in/vivace-io/evelib.v#`
 4. In the folder you just cloned, you need to reset your origin URL:
    - `git remote set-url origin git@github.com:<example>/evelib.git`
    - `git fetch`
 5. You should work on your code and push to your fork from that directory/repository.

### Pull Requests
To keep everything running smoothly, there are a few things that I ask of you:

 1. Do your best to ensure all tests pass prior to creating your pull request. If you are unable to fix any failing tests, or do not know how to, make note of it in the pull request details and ask for help. Don't be ashamed about asking for help!
 2. If exported API functions are changed with your pull request, it should be noted in the pull request. This would require a version bump.
 3. For current todos, feature requests, bugs, etc. see the repository issue tracker.
 4. While I prefer to have an issue to link to a PR, its not necessary. Feel free to take the initiative.

The primary developer of this repository is still relatively new to programming and Go, with only a few years experience in each. I'm always willing to learn about better and different ways of implementations!

# df-savecheck
DF save file validator

## Download
* [Source code](https://github.com/lethosor/df-savecheck)
* [Binaries](https://github.com/lethosor/df-savecheck/releases) (this is the "releases" tab if you are at the GitHub repository from the link above)

## Building

Uses [goxc](https://github.com/laher/goxc)

* Compile:

    * For current platform: ``go build``
    * For all platforms: ``goxc compile``

* Make release packages: ``goxc package``
* Bump version:

    * ``goxc bump`` for patches
    * ``goxc bump -dot=1`` for minor releases
    * ``goxc bump -dot=0`` for major releases

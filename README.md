# df-savecheck
DF save file validator

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

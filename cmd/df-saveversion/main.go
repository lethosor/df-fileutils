package main

import (
    "flag"
    "fmt"
    "os"

    "../../dfversions"
    "../../util"
)

func main() {
    var version_only bool
    flag.BoolVar(&version_only, "o", false, "Display only version numbers")
    flag.Parse()
    util.ForEachFile(flag.Args(), func(file string, f *os.File) {
        version, err := util.ReadUInt32(f)
        if err != nil {
            fmt.Printf("%s: read failed: %v\n", file, err)
            return
        }
        version_str := dfversions.Describe(version)
        if (version_only) {
            fmt.Println(version_str)
        } else {
            fmt.Printf("%s: %s\n", file, version_str)
        }
    })
}

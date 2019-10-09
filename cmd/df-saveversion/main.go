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
    for _, file := range flag.Args() {
        f, err := os.Open(file)
        if err != nil {
            fmt.Printf("%s: %v\n", file, err)
            continue
        }
        defer f.Close()

        version, err := util.ReadUInt32(f)
        if err != nil {
            fmt.Printf("%s: read failed: %v\n", file, err)
            continue
        }
        version_str := dfversions.Describe(version)
        if (version_only) {
            fmt.Println(version_str)
        } else {
            fmt.Printf("%s: %s\n", file, version_str)
        }
    }
}

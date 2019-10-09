package util

import (
    "encoding/binary"
    "fmt"
    "os"
)

func ReadUInt32(f *os.File) (ret uint32, err error) {
    err = binary.Read(f, binary.LittleEndian, &ret)
    return
}

func BytesRemaining (f *os.File) (ret uint64) {
    orig, _ := f.Seek(0, 1)
    end, _ := f.Seek(0, 2)
    ret = uint64(end - orig)
    f.Seek(orig, 0)
    return
}

func ForEachFile (filenames []string, callback func (string, *os.File)) {
    for _, filename := range filenames {
        f, err := os.Open(filename)
        if err != nil {
            fmt.Printf("%s: %v\n", filename, err)
            continue
        }
        defer f.Close()
        callback(filename, f)
    }
}

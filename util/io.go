package util

import (
    "encoding/binary"
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


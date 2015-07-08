package main

import (
    "bytes"
    "compress/zlib"
    "encoding/binary"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func readInt32(f *os.File) (ret int32, err error) {
    buf := make([]byte, 4)
    n, err := f.Read(buf)
    if err != nil || n != 4 {
        ret = 0
        return
    }
    binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &ret)
    return
}

func main() {
    if len(os.Args) == 1 {
        fmt.Printf("%s: No save file(s) given\n", os.Args[0])
        return
    }
    for i, file := range os.Args {
        if i == 0 {
            continue
        }

        f, err := os.Open(file)
        if err != nil {
            fmt.Printf("%s: Not found\n", file)
            continue
        }

        save_version, err := readInt32(f)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("%s: Save version %d\n", file, save_version)

        compressed, err := readInt32(f)
        if err != nil {
            fmt.Println(err)
            continue
        }
        if compressed != 1 {
            fmt.Printf("%s: Not compressed\n", file)
            continue
        }

        // Attempt to decompress
        ok := false
        chunk := 1
        for ;; chunk++ {
            length, err := readInt32(f)
            if err == io.EOF {
                ok = true
                break
            }
            if err != nil {
                fmt.Printf("%s (chunk %d): %s\n", file, chunk, err)
                break
            }
            mbytes := length / (1024 * 1024)
            if mbytes > 10 {
                fmt.Printf("%s (chunk %d): Memory threshold exceeded: tried to read %d MB\n", file, chunk, mbytes)
                break
            }

            buf := make([]byte, length)
            n, err := f.Read(buf)
            if err != nil {
                fmt.Printf("%s (chunk %d): %s\n", file, chunk, err)
                break
            }
            if n != int(length) {
                fmt.Printf("%s (chunk %d): Expected %d bytes, got %d\n", file, chunk, length, n)
                break
            }

            reader, err := zlib.NewReader(bytes.NewReader(buf))
            if err != nil {
                fmt.Printf("%s (chunk %d): Read failed: %s\n", file, chunk, err)
                break
            }
            _, err = ioutil.ReadAll(reader)
            if err != nil {
                fmt.Printf("%s (chunk %d): Compression error: %s\n", file, chunk, err)
                break
            }
        }
        if ok {
            fmt.Printf("%s: No compression errors detected (%d chunks)\n", file, chunk)
        } else {
            remaining, _ := ioutil.ReadAll(f)
            fmt.Printf("%s: %d bytes unread\n", file, len(remaining))
        }
    }
}

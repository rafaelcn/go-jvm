//
//
// Rafael Campos Nunes <rafaelnunes@engineer.com>
//

package main

import (
    "flag"
    "io/ioutil"
    "log"
    "os"
)

var (
    filename = flag.String("filename", "", "Which .class to interpret")

    verbose = flag.Bool("verbose", false, "")
    debug   = flag.Bool("debug", false, "")
)

func init() {
    flag.Usage = usage
}

func main() {
    flag.Parse()

    if len(*filename) == 0 {
        flag.Usage()
        os.Exit(1)
    }

    // Read .class file
    bytes := read(*filename)

    if *verbose {
        log.Printf("%x", bytes)
    }
}

// read returns all the bytes inside a class file
func read(f string) []byte {
    b, err := ioutil.ReadFile(f)

    if err != nil {
        log.Fatalf("[!] Reading filename %s has failed. Reason %v", f, err)
        return nil
    }

    return b
}

func usage() {
    // TODO: Implement this function
}

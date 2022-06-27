# YottaGo-Gitpod

A Gitpod environment for development with Golang and YottaDB

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/RamSailopal/YottaGo-Gitpod)

Once the environment has provisioned fully, four windows will open at the bottom of the screen.

The first window can be ignored, the second window will give access to the YottaDB environment, the third will be access to a command line to run Go code.

![Alt text](YottaGo.JPG?raw=true "Gitpod")

A simple test code example is stored in the code folder:

    package main

    import ( "lang.yottadb.com/go/yottadb"
         "fmt"
    )
    func main() {
                var err error
                var errstr yottadb.BufferT
                var tptoken uint64
                err = yottadb.SetValE(tptoken, nil, "Testing", "^TEST", []string{"1"})
                resp, err := yottadb.ValE(tptoken, &errstr, "^TEST", []string{"1"})
                if err == nil {
                        fmt.Print(resp, "\n")
                } else {
                        fmt.Print(err, "\n")
                }
    }

    
This can be run in the third window with:

    go run test.go
    
The code performs the following equivalent in M:

    S ^TEST("1")="Testing"
    
This can further be seen by running the following in the second window:

    YDB>D ^%G

    Output device: <terminal>:

    List ^TEST
    ^TEST(1)="Testing"
     
 # Further Coding Reference
 
 https://docs.yottadb.com/MultiLangProgGuide/goprogram.html#

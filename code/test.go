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

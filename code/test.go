package main

import ( "lang.yottadb.com/go/yottadb"
         "fmt" 
)
func main() {
                var err error
                var errstr yottadb.BufferT
                var tptoken uint64
		err = yottadb.SetValE(tptoken, nil, "1", "^TEST", []string{"Testing"})
		resp, err := yottadb.ValE(tptoken, &errstr, "^TEST", []string{"1"})
                fmt.Print(resp)
}

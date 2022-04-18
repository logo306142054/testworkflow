package main

import (
	"fmt"
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

type TestNestedModel struct {
	X string
}

type TestModel struct {
	ID                string
	Name              string
	Count             int
	Bool              bool
	Numbers           []int64
	Strings           []string
	KeyValues         map[string]string
	Nested            TestNestedModel
	MapStructs        map[string]TestNestedModel
	MapPointerStructs map[string]*TestNestedModel
}

func main() {
	fdb.MustAPIVersion(630)
	db := fdb.MustOpenDefault()

	ret, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		tr.Set(fdb.Key("hello"), []byte("world"))
		return tr.Get(fdb.Key("hello")).MustGet(), nil
		// db.Transact automatically commits (and if necessary,
		// retries) the transaction
	})
	if e != nil {
		log.Fatalf("Unable to perform FDB transaction (%v)", e)
	}

	fmt.Printf("hello is now world, foo was: %s\n", string(ret.([]byte)))
}

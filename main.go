package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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

func handlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	fmt.Println("proto:", request.Proto)
	fmt.Println("request uri:", request.RequestURI)
	fmt.Println("request url unescape:", request.URL.Query())
	fmt.Println("useragent:", request.UserAgent())
	request.ParseForm()
	fmt.Println("form:", request.Form)
	fmt.Println("header:", request.Header)
	fmt.Print("body:")
	io.Copy(os.Stdout, request.Body)
	fmt.Println()
	fmt.Println("escapedpath:", request.URL.EscapedPath())
	fmt.Println("Fragment", request.URL.Fragment)
	fmt.Println("Opaque", request.URL.Opaque)
	fmt.Println("Path", request.URL.Path)
	fmt.Println("RawPath", request.URL.RawPath)
	fmt.Println("RawQuery", request.URL.RawQuery)
	fmt.Println("scheme:", request.URL.Scheme)
	fmt.Println("host:", request.Host)
	fmt.Println("url:", request.URL.String())

	headRange := request.Header["Range"]
	fmt.Println("headrange:", headRange, " ==", request.Header.Get("Range"))
	//fmt.Println("", request.URL.)
	writer.Write([]byte(fmt.Sprintf("hello world:%s", time.Now().Format(time.RFC3339))))
	fmt.Println("-----------------")
	// go func(writer http.ResponseWriter, request *http.Request) {

	// }()
}

func initServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunc)
	srv := http.Server{
		Addr:         "127.0.0.1:5041",
		IdleTimeout:  1 * time.Minute,
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
	}
	lerr := srv.ListenAndServe()
	if lerr != nil && lerr == http.ErrServerClosed {
		fmt.Println("server clsoesd")
		return
	}
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

	initServer()
}

package main // OMIT
// OMIT
import _ "net/http/pprof"
import "log"      // OMIT
import "net/http" // OMIT

func main() {
	log.Println(http.ListenAndServe("localhost:3999", nil))
}

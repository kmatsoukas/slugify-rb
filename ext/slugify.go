package main

import (
	"C"

	"github.com/gosimple/slug"
)

func init() {
	slug.CustomSub = map[string]string{
		"@": "",
		"_": "-",
	}
}

//export slugify
func slugify(s *C.char) *C.char {
	gos := C.GoString(s)
	slugString := slug.Make(gos)

	return C.CString(slugString)
}

func main() {

}

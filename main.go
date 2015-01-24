package main

/*
#cgo LDFLAGS: -ljq
#include <jq.h>
*/
import "C"
import "fmt"

var testData = `{"data":[{"id":"6003598240487","name":"la biblia","audience_size":7693470,"path":[],"description":null},{"id":"6011716527242","name":"soccer on facebook","audience_size":14450580,"path":[],"description":null},{"id":"6003512642864","name":"hanging out with friends","audience_size":6330690,"path":[],"description":null},{"id":"6002926351162","name":"Hockey","audience_size":10611850,"path":[],"description":null},{"id":"6003022269556","name":"Rugby football","audience_size":13935770,"path":[],"description":null},{"id":"6003146664949","name":"Netball","audience_size":4229890,"path":[],"description":null},{"id":"6003013291881","name":"Kaizer Chiefs F.C.","audience_size":1794420,"path":[],"description":null},{"id":"6003255367499","name":"Generations (U.S. TV series)","audience_size":2278880,"path":[],"description":null},{"id":"6010738923046","name":"movies on facebook","audience_size":14951700,"path":[],"description":null},{"id":"6004030252209","name":"rb","audience_size":5030410,"path":[],"description":null},{"id":"6003187324105","name":"kaizer chiefs","audience_size":1489320,"path":[],"description":null},{"id":"6003170242302","name":"watching tv","audience_size":6956980,"path":[],"description":null},{"id":"6003400886535","name":"espn sportscenter","audience_size":228910,"path":[],"description":null},{"id":"6002925969459","name":"watching movies","audience_size":4784650,"path":[],"description":null},{"id":"6003214125247","name":"lakers","audience_size":349740,"path":[],"description":null}]}`

// This is just a simple POC.
// I have not included any memory management yet.
func main() {
	var state *C.jq_state = C.jq_init()
	if state == nil {
		panic("state nil")
	}

	// Compiles a `program` into jq_state.
	C.jq_compile(state, C.CString(".data[].name"))

	parser := C.jv_parser_new(0)

	// Make a simple input then convert to CString.
	str := testData
	cstr := C.CString(str)
	C.jv_parser_set_buf(parser, cstr, C.int(len(str)), 0)

	// This should be done in a loop but we arent for the POC
	v := C.jv_parser_next(parser)

	// Check if v is valid.
	if C.jv_is_valid(v) == 1 {

		C.jq_start(state, v, C.int(0))

		// Should be in loop for same reason as above.
		res := C.jq_next(state)

		// Check if res is valid.
		if C.jv_is_valid(res) == 1 {
			// Dump it!
			dumped := C.jv_dump_string(res, 0)

			// Convert dump to string!
			strval := C.jv_string_value(dumped)
			fmt.Println(C.GoString(strval))
		}
	}
}

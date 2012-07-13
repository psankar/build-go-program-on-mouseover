/* 
 * Author: Sankar சங்கர்
 * License: Creative Commons Zero License 1.0
 */
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1024", nil)
}

var src = `<html>
		<head>
		<script>function mysubmit(a) { document.forms["myform"].progtobuild.value = a; document.forms["myform"].submit(); }</script>
		</head>
		<body>
		<table>
		<form id="myform" method="post">
		<ul>
			<li onmouseover=mysubmit("Print123.go")>Print123.go</li>
			<li onmouseover=mysubmit("Print456.go")>Print456.go</li>
			<li onmouseover=mysubmit("Printabc.go")>Printabc.go</li>
			<li onmouseover=mysubmit("HelloPuvi.go")>HelloPuvi.go</li>
			<li onmouseover=mysubmit("HelloLoka.go")>HelloLoka.go</li>
			<li onmouseover=mysubmit("HelloWorld.go")>HelloWorld.go</li>
		</ul>
		<input type="hidden" id="progtobuild" name="progtobuild">
		</form>
		</body>
		</html>`

func handler(w http.ResponseWriter, r *http.Request) {

	output := ""

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			fmt.Fprint(w, fmt.Sprintf("Error parsing the submitted HTML form:\n%s", err))
			return
		}

		cmd := exec.Command("go", "build", r.FormValue("progtobuild"))

		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		cmd = exec.Command(fmt.Sprintf("./%s", strings.TrimRight(r.FormValue("progtobuild"), ".go")))

		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		output = out.String()
	}

	fmt.Fprint(w, src)
	fmt.Fprint(w, output)
}

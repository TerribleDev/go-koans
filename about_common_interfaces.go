package go_koans

import "bytes"
import "strings"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/
		in.WriteTo(out)
		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")
		out := new(bytes.Buffer)
		str := in.String()
		out.WriteString(strings.Split(str, " ")[0])

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}

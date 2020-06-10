package main

import (
	"fmt"

	"github.com/nickwells/param.mod/v5/param"
)

// addExamples adds some examples of how gosh might be used to the standard
// help message
func addExamples(ps *param.PSet) error {
	ps.AddExample(`gosh -pln '"Hello, World!"'`, `This prints Hello, World!`)

	ps.AddExample("gosh -pln 'math.Pi'", `This prints the value of Pi`)

	ps.AddExample("gosh -pln '17*12.5'",
		`This prints the results of a simple calculation`)

	ps.AddExample(
		"gosh -n -b 'count := 0' -e 'count++' -a-pln 'count'",
		"This reads from the standard input and prints"+
			" the number of lines read"+
			"\n"+
			"\n-n sets up the loop reading from standard input"+
			"\n-b 'count := 0' declares and initialises the counter"+
			" before the loop"+
			"\n-e 'count++' increments the counter inside the loop"+
			"\n-a-pln 'count' prints the counter using fmt.Println"+
			" after the loop.")

	ps.AddExample("gosh -n -b-p '\"Radius: \"'"+
		" -e 'r, _ := strconv.ParseFloat(_l.Text(), 10)'"+
		" -pf '\"Area: %9.2f\\n\", r*r*math.Pi'"+
		" -p '\"Radius: \"'",
		"This repeatedly prompts the user for a Radius and prints"+
			" the Area of the corresponding circle"+
			"\n"+
			"\n-n sets up the loop reading from standard input"+
			"\n-b-p '\"Radius: \"' prints the first prompt"+
			" before the loop."+
			"\n-e 'r, _ := strconv.ParseFloat(_l.Text(), 10)' sets"+
			" the radius from the text read from standard input,"+
			" ignoring errors."+
			"\n-pf '\"Area: %9.2f\\n\", r*r*math.Pi' calculates and"+
			" prints the area using fmt.Printf."+
			"\n-p '\"Radius: \"' prints the next prompt.")

	ps.AddExample(
		`gosh -i -w-pln `+
			`'strings.ReplaceAll(string(_l.Text()), "mod/pkg", "mod/v2/pkg")'`+
			` -- abc.go xyz.go `,
		"This changes each line in the two files abc.go and xyz.go"+
			" replacing any reference to mod/pkg with mod/v2/pkg. You"+
			" might find this useful when you are upgrading a Go module"+
			" which has changed its major version number. The files will"+
			" be changed and the original contents will be left behind in"+
			" files called abc.go.orig and xyz.go.orig.")

	ps.AddExample(`gosh -http-handler 'http.FileServer(http.Dir("/tmp/xxx"))'`,
		"This runs a web server that serves files from /tmp/xxx."+
			" Remember that any relative paths (not starting with '/')"+
			" will be evaluated ralative to the temporary run directory"+
			" from where the gosh program is run and not relative to the"+
			" directory where you are running gosh.")

	ps.AddExample(`gosh -web-p '"Gosh!"'`,
		"This runs a web server (listening on port "+
			fmt.Sprint(dfltHTTPPort)+") that returns 'Gosh!' for every"+
			" request.")

	ps.AddExample(`gosh -n -e 'if l := len(_l.Text()); l > 80 { '`+
		` -pf '"%3d: %s\n", l, _l.Text()' -e '}'`,
		"This will read from standard input and print out each line that"+
			" is longer than 80 characters.")

	return nil
}

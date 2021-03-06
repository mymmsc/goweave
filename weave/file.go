package weave

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// reWriteFile rewrites curfile with out && adds any missing imports
func (w *Weave) reWriteFile(curfile string, out string, importsNeeded []string) {

	path := w.buildLocation + "/" + curfile

	f, err := os.Create(path)
	if err != nil {
		w.flog.Println(err)
	}

	defer f.Close()

	if len(importsNeeded) > 0 {
		out = w.addMissingImports(importsNeeded, out)
	}

	_, err = f.WriteString(out)
	if err != nil {
		w.flog.Println(err)
	}

	w.reWorkImports(path)

}

// fileAsStr reads path file and returns it as a string
func fileAsStr(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	return string(buf)
}

// returns a slice of lines from file path
func fileLines(path string) []string {
	stuff := []string{}

	file, err := os.Open(path)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		stuff = append(stuff, scanner.Text())
	}

	return stuff
}

// writeOut writes nlines to path
func (w *Weave) writeOut(path string, nlines string) {

	b := []byte(nlines)
	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		w.flog.Println(err)
	}
}

// writeAtLine inserts writes to fname lntxt @ iline
func (w *Weave) writeAtLine(fname string, iline int, lntxt string) string {
	flines := fileLines(fname)

	out := ""
	for i := 0; i < len(flines); i++ {
		if i == iline {
			out += lntxt + "\n"
		}

		if strings.TrimSpace(flines[i]) == "" {
			out += "\n"
		} else {
			out += flines[i] + "\n"
		}
	}

	w.writeOut(fname, out)

	return out
}

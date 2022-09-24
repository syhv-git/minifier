package minifier

import (
	"bytes"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"os"
	"path"
	"regexp"
)

// Minifier takes three parameters and creates a minified file given by fp.
// mime is the media type for the files located at dir.
// dir can be a directory or individual files.
//
// The function will skip file fp if it is located in dir.
func Minifier(fp, mime string, dir ...string) error {
	buf := bytes.NewBuffer(nil)
	n := path.Base(fp)

	c := make(chan []byte)
	e := make(chan error)
	go walkDir(dir, n, c, e)

wait:
	for {
		select {
		case b, ok := <-c:
			if !ok {
				break wait
			}
			if _, err := buf.Write(b); err != nil {
				return err
			}
			if err := buf.WriteByte(byte('\n')); err != nil {
				return err
			}
		case err := <-e:
			panic(err.Error())
		}
	}

	min := minify.New()
	min.AddFunc("text/css", css.Minify)
	min.AddFunc("text/html", html.Minify)
	min.AddFunc("image/svg+xml", svg.Minify)
	min.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	min.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	min.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	tmp, err := min.Bytes(mime, buf.Bytes())
	if err != nil {
		return err
	}
	if err = os.WriteFile(fp, tmp, 0777); err != nil {
		return err
	}
	return nil
}

func walkDir(dir []string, ex string, b chan []byte, e chan error) {
	defer close(b)
	if len(dir) < 1 {
		return
	}

	for _, d := range dir {
		info, err := os.Stat(d)
		if err != nil {
			e <- err
			return
		}
		if info.Name() == ex {
			continue
		}
		if !info.IsDir() {
			tmp, err := os.ReadFile(d)
			if err != nil {
				e <- err
			}
			b <- tmp
			continue
		}
		nd, _ := os.ReadDir(d)
		p := handleDir(d, ex, nd)

		c := make(chan []byte)
		e2 := make(chan error)
		go walkDir(p, ex, b, e)

	wait2:
		for {
			select {
			case x, ok := <-c:
				if !ok {
					break wait2
				}
				b <- x
			case err = <-e2:
				e <- err
				return
			}
		}
	}
}

func handleDir(d, ex string, c []os.DirEntry) (paths []string) {
	for _, de := range c {
		if de.Name() == ex {
			continue
		}
		p := d + "/" + de.Name()
		paths = append(paths, p)
	}
	return
}

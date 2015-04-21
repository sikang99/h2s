package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/xyproto/mime"
)

// When serving a file. The file must exist. Must be given a full filename.
func filePage(w http.ResponseWriter, filename string, mimereader *mime.MimeReader) {
	// Mimetypes
	ext := path.Ext(filename)
	// Markdown pages are handled differently
	if ext == ".md" {
		w.Header().Add("Content-Type", "text/html")
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(w, "Unable to read %s: %s", filename, err)
			return
		}
		markdownBody := string(blackfriday.MarkdownCommon(b))
		fmt.Fprint(w, markdownPage(filename, markdownBody))
		return
	}
	// Set the correct Content-Type
	mimereader.SetHeader(w, ext)
	// Write to the ResponseWriter, from the File
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(w, "Can't open %s: %s", filename, err)
	}
	// Serve the file
	io.Copy(w, file)
	return
}

// Directory listing
func directoryListing(w http.ResponseWriter, dirname string) {
	var buf bytes.Buffer
	sep := string(os.PathSeparator)
	for _, filename := range getFilenames(dirname) {

		// Find the full name
		full_filename := dirname
		if !strings.HasSuffix(full_filename, sep) {
			full_filename += sep
		}
		full_filename += filename

		// Output different entries for files and directories
		buf.WriteString(easyLink(filename, full_filename, isDir(full_filename)))
	}
	title := dirname
	// Strip the leading "./"
	if strings.HasPrefix(title, "./") {
		title = title[2:]
	}
	// Use the application title for the main page
	//if title == "" {
	//	title = version_string
	//}
	if buf.Len() > 0 {
		fmt.Fprint(w, easyPage(title, buf.String()))
	} else {
		fmt.Fprint(w, easyPage(title, "Empty directory"))
	}
}

// When serving a directory. The directory must exist. Must be given a full filename.
func dirPage(w http.ResponseWriter, dirname string, mimereader *mime.MimeReader) {
	// Handle the serving of index files, if needed
	for _, indexfile := range indexFilenames {
		filename := path.Join(dirname, indexfile)
		if exists(filename) {
			filePage(w, filename, mimereader)
			return
		}
	}
	// Serve a directory listing of no index file is found
	directoryListing(w, dirname)
}

// When a file is not found
func noPage(filename string) string {
	return easyPage("Not found", "File not found: "+filename)
}

// Serve all files in the current directory, or only a few select filetypes (html, css, js, png and txt)
func registerHandlers(mux *http.ServeMux, servedir string) {
	// Read in the mimetype information from the system. Set UTF-8 when setting Content-Type.
	mimereader := mime.New("/etc/mime.types", true)

	// Handle all requests with this function
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		urlpath := req.URL.Path
		filename := url2filename(servedir, urlpath)
		// Remove the trailing slash from the filename, if any
		noslash := filename
		if strings.HasSuffix(filename, "/") {
			noslash = filename[:len(filename)-1]
		}
		hasdir := exists(filename) && isDir(filename)
		hasfile := exists(noslash)
		// Share the directory or file
		if hasdir {
			dirPage(w, filename, mimereader)
			return
		} else if !hasdir && hasfile {
			// Share a single file instead of a directory
			filePage(w, noslash, mimereader)
			return
		}
		// Not found
		fmt.Fprint(w, noPage(filename))
	})
}

# h2s

Simple HTTP/2 test server, for serving the files in a given directory.


### Features and limitations

* Supports HTTP/2 and HTTPS.
* Supports Markdown when displaying files ending with `.md`.
* If `index.html`, `index.md` or `index.txt` is found, it will be used for the main page.
* Reasonably fast. Runs as a native executable.
* Uses UTF-8 whenever possible.
* Sets Content-Type for a whole range of file extensions if `/etc/mime.types` exists. If not, Content-Type is set for a few commonly used types, like png and css.
* Self-signed TLS certificates will make the browser complain, unless the certificates are imported somehow.

### Usage

`h2s [directory] [host:port] [certfile] [keyfile]`

`host:port` can be just `:port` for localhost.

### Examples

Share the current directory as https://localhost:3000/

`h2s . :3000`


Share a single file as the main page at https://localhost/. This will make h2s listen to port 443, which may require more permissions.

`./h2s README.md`


### General information

<img src="https://raw.githubusercontent.com/sikang99/h2s/master/img/spdy-to-http2.png">

Reference
---------
- [xyproto/snusnu](https://github.com/xyproto/snusnu) - HTTP/2 web server for static files
- [HTTPS and Go](https://www.kaihag.com/https-and-go/)


### License

MIT


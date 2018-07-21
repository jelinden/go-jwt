# go-jwt

A small project to test how jwt works on the server side.

`/authenticate` endpoint takes username and password as a POST request and json body.

`/protected` endpoint is protected by a token check

Username, password and the key to sign the token are in the code for clarity, but on a production server they should never be available publicly.

## Running

Get vgo `go get -u golang.org/x/vgo`

`vgo build && ./go-jwt`

More info on vgo: `https://godoc.org/golang.org/x/vgo`

## Testing

Run the simple test with `go test`.

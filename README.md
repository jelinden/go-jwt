# go-jwt

## Part one

A small project to test how jwt works on the server side.

`/authenticate` endpoint takes username and password as a POST request and json body.

`/protected` endpoint is protected by a token check

Username, password and the key to sign the token are in the code for clarity, but on a production server they should never be available publicly.

### Running

`cd part1`

Get vgo `go get -u golang.org/x/vgo`

`vgo build && ./part1`

More info on vgo: `https://godoc.org/golang.org/x/vgo`

### Testing

Run the simple test with `go test`.

## Part two

A Mithril.js web application to test how jwt token basically works. After successfull login the token is saved to local storage.

<img src="https://github.com/jelinden/go-jwt/raw/master/jwt-login.png" width="350">

### Running

`npm install && npm run build`

`vgo build && ./part2`

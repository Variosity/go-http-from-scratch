# Go HTTP Server — Stdlib Only
An HTTP Server made from scratch with JSON Content-Type handling.

## What it does
This project has two endpoints, it uses a 'helloHandler' for normal string handling and a 'jsonHandler' for JSON output handling.

## How to run it

```
git clone https://github.com/Variosity/go-http-from-scratch.git
cd go-http-from-scratch
go run http.go
```

## How it works

1. net/http.HandleFunc registers the handler functions registers a URL path to a handler function so that when a request hits that path, Go knows which function to call.

2. ResponseWriter is used to write a response for when you make a request to the server.

3. The JSON handler calls Header().Set() before writing the body so that the client knows what kind of data to expect so it can properly handle it.
   
4. encoding/json.Marshal parses the JSON content, it returns the JSON encoding of 'jsonr', it is converted into a string before writing it into the response so that io.WriteString can output it.

## What I learned

1. A method lives on a value: myValue.DoSomething(), a function lives in a package: packageName.DoSomething(myValue), there was a confusion while I was building this because I named my json struct 'json' which caused errors.
   
2. http.Error sends an error message and a status code in one call, the return exits the handler function and nothing runs.

## What's next

*Next up I will be building an SQLite Uptime Monitor which will build on my HTTP Server knowledge with clients and operational thinking.*

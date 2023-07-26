# Template on how api design

Logic pulled out of main. Have main call func that returns error.

Server struct to hold all shared dependencies for api. No need to keep values as global state
 - Serve Mux
 - Logger (Im not attached to zeroLog)
 - Database
 - Controller
 - Config (Port, Api base path, etc)

Since Server struct hold shared dependancy, can have different func load values from different source. ie dev, flags, fromFile, from enviromental vars.

Handlers attached to server stuct
```
func (s *server) handleSomething() http.HandlerFunc { ... }
```

Handlers return HandlerFunc. It is more versitle since it can be used as handler or handlerFunc. It can only do a one-time per handle initalization (prepareThing).
```
func (s *server) handleSomething() http.HandlerFunc {
    thing := prepareThing()
    return func(w http.ResponseWriter, r *http.Request) {
        // use thing        
    }
}
```

Use middleware to handle repeative tasks like check if authorized.

Response function to handle header content and formats data to be returned.

Don't use number when setting status code, use name. 
Don't do
```
w.WriteHeader(200)
r, err := http.NewRequest("GET", u, nil)
do
```
w.WriteHeader(http.StatusOK)
r, err := http.NewRequest(http.MethodGet, u, nil)
```
Testing
```go test -v ./...```

todo:
 - [xxx]
## Middleware (Advanced)
This example will show how to create a more advanced version of middleware in Go.
A middleware in itself simply takes a http.HandleFunc as one of its parameters, wraps it and returns a new http.HandlerFunc for the server to call.
Here we define a new type Middleware which makes it eventually easier to chain multiple middlewares toghether. This idea is inspired by Mat Ryers'talk about Building APIs. You can find a more detailed explaination including the talk here.
This snippet explains in detail how a new middleware is created. In the full example below, we reduce this version by some boilerplate code.
```
func createNewMiddleware() Middleware {
    // Create a new Middleware
    middleware := func(next http.HandlerFunc) htt.HanlderFunc {

        // Define the http.HandlerFunc which is called by the server eventually
        handler := func(w http.ResponseWriter, r *http.Request) {

            // ... do middleware things

            // Call the next middleware/handler in chain
            next(w, r)
        }
        
        // Return newly created handler
        return handler
    }

    // Return newly created middleware
    return middleware
}
```
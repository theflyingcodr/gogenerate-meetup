# API Gen

One of the main things developers use Go for is writting microservice and web apis. This is a simple domain driven api example using the standard http libs.

Take a look at the [airplanes](airplanes) domain package and you'll see the service definitions in the [airplanes.go](airplanes/airplanes.go) file. We then have [airplanes/http.go](http.go) containing controllers, [service.go](airplanes/service.go) contains buisness logic and [mem_store.go](airplanes/mem_store.go) is a really basic in memory store that returns a list of airplanes and can add new airplanes.

When developing APIs we often need to write the same bootstrapped code to get a new endpoint or service up and running. This is tedious.

Under cmd/ there are two packages [api](cmd/api) and [gen](cmd/gen) open api and run `go run main.go` and you'll have the web service running on port :8080. You can hit [http://localhost:8080/api/v1/airplanes](http://localhost:8080/api/v1/airplanes) and get a response.

The second package [gen](cmd/gen) contains a generator. This will generate a new domain, taking away the boring task of bootstrapping. You will then have a full new endpoint you can hit. The idea of this in the real world is you could then go and add in whatever business logic is required a lot faster than manually creating the same boring files and structure over and over. It also enforces consistency. You can ensure everything is commented in your template files, that you are using best practises and everything follows common guidelines that you can determine.

To run the generator cd into the cmd/gen package and run `go run main.go -domain=<whatever>`. You can then hit /api/v1/<whatever>s and get a basic response back, awesome!

The generator uses AST in go to firstly update the [routes.go](routes.go) file with a new constant. It then updates the code in [api/main.go](cmd/api/main.go) to wire up your new endpoint and add the require import. Pretty cool!

Feel free to use this code for your own generators!

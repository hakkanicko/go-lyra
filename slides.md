% title: Workshop GO
% subtitle: Workshop GO
% author: Emile Vauge
% thankyou: Thank you!
% contact: <a href="https://twitter.com/emilevauge">@emilevauge</a>

---
title: Setup

- On GitHub, fork [https://github.com/emilevauge/go-lyra](https://github.com/emilevauge/go-lyra)
- Go to [http://c9.io](http://c9.io)
- Sign in using your GitHub account
- In Cloud9, edit this repository
- Let's code !

---
title: GO in few words

- designed by Google (Robert Griesemer, Rob Pike, and Ken Thompson)
- open source
- commonly called **golang**
- simple, productive, readable
- a modern take on C

---
title: more words...

- fast (compile, run)
- garbage collected
- strongly typed
- statically linked binaries

---
title: Where is the catch?

- still early (1.0 in March 2012)
- dependencies management
- no generics (yet)
- not really OOP
- lack of powerful IDE

---
title: Here and there

- Google (obviously)
- Docker
- Twitter
- AirBnB
- SoundCloud
- Netflix

---
title: Tools

- `go build`: which builds Go binaries
- `go test`: for unit testing and microbenchmarks
- `go fmt`: for formatting code
- `go run`: a shortcut for building and executing code
- `godoc`: for displaying documentation or serving it via HTTP
- `go generate`: a standard way to invoke code generators


---
title: Hello, Playground!

[http://play.golang.org/](http://play.golang.org/)

<pre class="prettyprint" data-lang="go">
package main

import "fmt"

func main() {
  fmt.Println("Hello, playground")
}
</pre>

---
title: Main function

- Every go program has a main() function
- This is where execution starts


<pre class="prettyprint" data-lang="go">
package main

func main() {
  // Your program starts here
}
</pre>


---
title: Output

- fmt is a core package (means “format”)
- pronounced "fumpt" (weird)

<pre class="prettyprint" data-lang="go">
package main

import "fmt"

func main() {
  fmt.Println("Hello Lyra!")
}
</pre>


---
title: Exercise `1_output`

---
title: Variables

- Define

<pre class="prettyprint" data-lang="go">
var myVar string
</pre>

- Initialize

<pre class="prettyprint" data-lang="go">
var myVar string
</pre>

---
title: Type Inference

- Define and Initialize

<pre class="prettyprint" data-lang="go">
var myVar string
</pre>

- Initialize

<pre class="prettyprint" data-lang="go">
myVar := "Hello"
</pre>


---
title: Functions

<pre class="prettyprint" data-lang="go">
func withArgs(a int, b string) int {
    return 3
}

result := withArgs(3, "Hello")
</pre>

<pre class="prettyprint" data-lang="go">
func withMultipleReturns() (string, error) {
  return "Hello", nil
}


s, err := withMultipleReturns()
</pre>


---
title: Exercise `4_returns`


---
title: Structs
subtitle: Define

<pre class="prettyprint" data-lang="go">
type Person struct {
  Name string
  Age  int
}

p := Person{
  Name: "Greg",
  Age:  25,
}


fmt.Println(p.Name)
</pre>


---
title: Structs
subtitle: Access

- public Properties Are Uppercase

*accessible to anything*

- private properties are lowercase

*only accessible via methods*



---
title: Structs
subtitle: Instantiation

<pre class="prettyprint" data-lang="go">
type Person struct {
  Name string
  Age  int
}

p := Person{
  Name: "Greg",
  Age:  25,
}


fmt.Println(p.Name)
</pre>


---
title: Structs
subtitle: Methods

<pre class="prettyprint" data-lang="go">
type Person struct {
  Name   string
  Age    int
}

func (p *Person) CanEnterPub() bool {
	return p.Age >= 19
}


person.CanEnterPub()
</pre>


---
title: Exercise `6_methods`


---
title: GoRoutines

Goroutines are lightweight threads of execution

- executed concurrently
- supports > 100,000 goroutines
- automatically distributes load across CPUs
- Worry free!


---
title: GoRoutines
subtitle: How?


<pre class="prettyprint" data-lang="go">
doSomething()
</pre>

<pre class="prettyprint" data-lang="go">
go doSomething()
</pre>


---
title: Channels

Goroutines communicate via channels: one-way pipes
Operations:

- Push to channel
- Listen to channel

Channels have types too:

- var chan int
- var chan string
- var chan map[string]MyStruct

Use channels to share state, for control flow

---
title: Channels
subtitle: Pushing

- Use `make()` to create a channel
<pre class="prettyprint" data-lang="go">
var channel := make(chan int)
</pre>

- Push to a channel with `<-`
<pre class="prettyprint" data-lang="go">
channel <- 5
channel <- 3
</pre>

---
title: Channels
subtitle: Listening

- Use select to wait (block) for a single message
- Similar syntax to case
<pre class="prettyprint" data-lang="go">
select {
case i <- channel:
  fmt.Printf("Got message %d", i)
}
</pre>

- Watch for multiple messages using range
<pre class="prettyprint" data-lang="go">
for i := range channel {
  fmt.Printf("Got message %d", i)
}
</pre>

---
title: Exercise `9_channels`


---
title: Loops

<pre class="prettyprint" data-lang="go">
// for loop
for i := 0; i < 10; i++ {...}


// while loop
for i < 10 {...}


// infinite loop
for {...}
</pre>

---
title: Web

<pre class="prettyprint" data-lang="go">
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
</pre>

---
title: Exercise `12_web`

---
title: Gorilla
subtitle: Advanced routing

- URL host
-  path
- path prefix
- schemes
- header
- query values
- HTTP methods
- custom matchers

---
title: Gorilla
subtitle: Example


<pre class="prettyprint" data-lang="go">
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/products", ProductsHandler)
    r.HandleFunc("/articles", ArticlesHandler)
	http.ListenAndServe(":8080", r)
}
</pre>

---
title: Gorilla
subtitle: Variables

<pre class="prettyprint" data-lang="go">
r := mux.NewRouter()
r.HandleFunc("/products/{key}", ProductHandler)
r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
</pre>


<pre class="prettyprint" data-lang="go">
vars := mux.Vars(request)
category := vars["category"]
</pre>


---
title: Gorilla
subtitle: Domain, Dynamic sub-domain

<pre class="prettyprint" data-lang="go">
r := mux.NewRouter()
// Only matches if domain is "www.example.com".
r.Host("www.example.com")
// Matches a dynamic subdomain.
r.Host("{subdomain:[a-z]+}.domain.com")
</pre>


---
title: Exercise `13_rest`
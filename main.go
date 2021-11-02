package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my wow site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>To get in touch, email me at <a href=\"mailto:daniel7288@gmail.com\">daniel7288@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li><b>Q: Is there a free version?</b>
      <br>A: Yes!! we offer a free trial for 30 days on any paid plans.</br></li>
  <li><b>Q: What are your support hours?</b>
      <br>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</br></li>
  <li><b>Q: How do I contact support?</b><br>A: Email us - <a href="mailto:daniel7288@gmail.com">daniel7288@gmail.com</a></br></li>
</ul>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
//   switch r.URL.Path {
//   case "/":
//     homeHandler(w, r)
//   case "/contact":
//     contactHandler(w, r)
//   default:
//     // TODO: handle the page not found error
//     http.Error(w, "Page not found", http.StatusNotFound)
//     // w.WriteHeader(http.StatusNotFound)
//     // fmt.Fprint(w, "Page not found")
//   }
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		// TODO: handle the page not found error
		http.Error(w, "Page not found", http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "Page not found")
	}
}

func main() {
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	// var router http.HandlerFunc = pathHandler
	// var router Router
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

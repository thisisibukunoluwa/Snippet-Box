package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// pre existing variables 
	// type config struct {
	// 	addr string 
	// 	staticDir string 
	// }

	// var cfg config 

	// flag.StringVar(&cfg.addr, "addr",":4000","HTTP network address")
	// flag.StringVar(&cfg.staticDir, "static-dir","./ui/static","Path to static assets")
		
	addr := flag.String("addr", ":4000", "HTTP network")

	flag.Parse()

	// normally you log your output to standard streams and redirect the output to a file at runtime , but if you don;t want to do this , then you can always open a file in go and use it as your log destination 

	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE,0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorlog,
		infoLog: infoLog,
	}
	mux := http.NewServeMux()

	//serve a direcetory over http
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// way to coerce a normal funciton to conforming to the http.Handler interface
	mux.HandleFunc("/", http.HandlerFunc(app.home))
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	
	// make http.Server use our errorLog
	
	srv := &http.Server{
		// Addr: cfg.addr,
		Addr: *addr,
		ErrorLog: errorlog,
		Handler: mux,
	}
	
	infoLog.Printf("Starting server on %s", *addr)

	httperr := srv.ListenAndServe()
	errorlog.Fatal(httperr)
}

// we want to disable directory listen in "./ui/static, what we will do is we will create a blank index.thml file and put it there , or we create our own custom filesystem"

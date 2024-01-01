package main

import "net/http" // contains tools for web server



func main(){
	
}





	/* basic server:
						 w is where we write our response to  r contains info about request, e.g path
	http.HandleFunc("/daniel", func (w http.ResponseWriter, r *http.Request){

		w.Write([]byte("hello world"))
	})

	initialize server (default go server on desired port)
	http.ListenAndServe(":8080", nil) // nil not important for now
	*/
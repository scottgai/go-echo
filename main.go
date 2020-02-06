package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			buf := make([]byte, 4096)

			for {
				n, err := r.Body.Read(buf)
				fmt.Printf("Read %d bytes\n", n)

				if n > 0 {
					nn, err := w.Write(buf[:n])
					fmt.Printf("Wrote %d bytes\n", nn)

					if err != nil {
						fmt.Println("ERROR: ", err)
						break
					}
				}


				if err != nil {
					if err == io.EOF {
						fmt.Println("EOF")
						err = r.Body.Close()

						if err != nil {
							fmt.Println("ERROR (body.Close): ", err)
						}
					}
					break
				}
			}
		}
	})

	http.ListenAndServe(":8080", mux)


}

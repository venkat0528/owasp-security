package communsecurity

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("This is an example server.\n"))
	})

	// yourCert.pem - path to your server certificate in PEM format
	// yourKey.pem -  path to your server private key in PEM format
	log.Fatal(http.ListenAndServeTLS(":443", "yourCert.pem", "yourKey.pem", nil))
}

type Certificates struct {
	CertFile string
	KeyFile  string
}

func ServerStart() {
	httpsServer := &http.Server{
		Addr: ":8080",
	}

	var certs []Certificates
	certs = append(certs, Certificates{
		CertFile: "../etc/yourSite.pem", //Your site certificate key
		KeyFile:  "../etc/yourSite.key", //Your site private key
	})

	config := &tls.Config{}
	config.Certificates = make([]tls.Certificate, len(certs))

	for i, v := range certs {
		config.Certificates[i], _ = tls.LoadX509KeyPair(v.CertFile, v.KeyFile)
	}

	conn, err := net.Listen("tcp", ":8080")
	fmt.Println(err)

	tlsListener := tls.NewListener(conn, config)
	httpsServer.Serve(tlsListener)
	fmt.Println("Listening on port 8080...")
}

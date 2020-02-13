package ReversProxy

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
)

func newTlsLn(Port string, caCertPath string,CertPath string, KeyPath string) (net.Listener, error) {

	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}

	pool.AppendCertsFromPEM(caCrt)

	cfg := &tls.Config{
		ClientCAs:  pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
	}

	cert, err := tls.LoadX509KeyPair(CertPath, KeyPath)
	if err != nil {
		return nil, err
	}

	cfg.Certificates = append(cfg.Certificates, cert)

	cfg.BuildNameToCertificate()

	ln, err := net.Listen("tcp4", ":"+Port)

	if err != nil {
		panic(err)
	}

	lnTls := tls.NewListener(ln, cfg)

	return lnTls, nil
}

package middleware

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
	"github.com/eddique/okta_go_saml/pkg/core/configs"
)

func SamlMiddleware() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair("okta-app.cert", "okta-app.key")
	if err != nil {
		log.Fatalln("Fatal Error:", err)
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Fatalln("Fatal Error:", err)
	}

	idpMetadataURL, err := url.Parse(configs.IdpMetadataUrl())
	if err != nil {
		log.Fatalln("Fatal Error:", err)
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		log.Fatalln("Fatal Error:", err)
	}
	rootURL, err := url.Parse(configs.RootUrl())
	if err != nil {
		log.Fatalln("Fatal Error:", err)
	}

	samlSP, err := samlsp.New(samlsp.Options{
		URL:               *rootURL,
		Key:               keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:       keyPair.Leaf,
		IDPMetadata:       idpMetadata,
		AllowIDPInitiated: true,
	})
	return samlSP, err
}

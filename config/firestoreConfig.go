package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type firebaseCred struct {
	Type          string `json:"type"`
	ProjectID     string `json:"project_id"`
	PrivateKeyID  string `json:"private_key_id"`
	PrivateKey    string `json:"private_key"`
	Clientemail   string `json:"client_email"`
	ClientID      string `json:"client_id"`
	AuthURI       string `json:"auth_uri"`
	TokenURI      string `json:"token_uri"`
	AuthProvider  string `json:"auth_provider_x509_cert_url"`
	ClientCertURL string `json:"client_x509_cert_url"`
}

// FSClient : list of firestore clients
type FSClient struct {
	AuthClient *auth.Client
	DBClient   *firestore.Client
}

// SetupFirebaseConfig : Return list of firestore clients
func SetupFirebaseConfig() *FSClient {
	opt := option.WithCredentialsJSON(createCredentials())
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authclient, err2 := app.Auth(context.Background())
	if err2 != nil {
		log.Fatalf("error getting Auth client: %v\n", err2)
		panic(fmt.Sprintf("error getting Auth client: %v\n", err2))
	}

	dbClient, err3 := app.Firestore(context.Background())
	if err3 != nil {
		log.Fatalf("error getting Firestore client: %v\n", err3)
	}

	return &FSClient{
		AuthClient: authclient,
		DBClient:   dbClient,
	}
}

func createCredentials() []byte {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %q", err)
	}
	credentials := firebaseCred{
		os.Getenv("F_TYPE"),
		os.Getenv("F_PROJECT_ID"),
		os.Getenv("F_PRIVATE_KEY_ID"),
		createPrivateKey(),
		os.Getenv("F_CLIENT_EMAIL"),
		os.Getenv("F_CLIENT_ID"),
		os.Getenv("F_AUTH_URI"),
		os.Getenv("F_TOKEN_URI"),
		os.Getenv("F_AUTH_PROVIDER_x509_CERT_URL"),
		os.Getenv("F_CLIENT_x509_CERT_URL"),
	}
	res, err := json.Marshal(credentials)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func createPrivateKey() string {
	s := os.Getenv("F_PRIVATE_KEY")
	c := 0
	finalStr := ""
	for i, v := range s {
		if string(v) == `\` && string(s[i+1]) == "n" {
			finalStr = finalStr + s[c:i]
			i += 2
			c = i
			if i < len(s) {
				finalStr += "\n"
			}
		}
	}
	return "-----BEGIN PRIVATE KEY-----\n" + finalStr + os.Getenv("F_PRIVATE_KEY_END") + "\n-----END PRIVATE KEY-----\n"
}

package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/Seifbarouni/fast-utils/utils"
	"github.com/zemirco/keycloak"
	"golang.org/x/oauth2"
)


var K *keycloak.Keycloak

func InitKeycloakClient(){

	url,err:=utils.GetVariable("config", "KEYCLOAK_URL")

	if err!=nil{
		log.Fatal(err)
	}

	name,err:=utils.GetVariable("config", "KEYCLOAK_ADMIN_NAME")

	if err!=nil{
		log.Fatal(err)
	}

	password,err:=utils.GetVariable("secrets", "KEYCLOAK_PASSWORD")

	if err!=nil{
		log.Fatal(err)
	}

	config := oauth2.Config{
        ClientID: "admin-cli",
        Endpoint: oauth2.Endpoint{
            TokenURL: fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", url),
        },
    }
	
	ctx := context.Background()
	token, err := config.PasswordCredentialsToken(ctx, name, password)

	 if err != nil {
        log.Fatal(err)
    }

    // create a new http client that uses the token on every request
   client := config.Client(ctx, token)

    // create a new keycloak instance and provide the http client
    k, err := keycloak.NewKeycloak(client, fmt.Sprintf("%s/",url))
    if err != nil {
        log.Fatal(err)
    }

	K = k

}
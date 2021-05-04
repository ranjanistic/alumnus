package handler

import (
	"fmt"
	"log"
	"context"
	"net/http"
	"net/url"
	"github.com/coreos/go-oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/ranjanistic/alumnus/config"
	"github.com/ranjanistic/alumnus/app"
	"github.com/ranjanistic/alumnus/auth"
	// "crypto/rand"
	"encoding/base64"
)

func Root(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"title": config.Env.APPNAME,
	})
}

func Login(c *fiber.Ctx) error {
	ht,err := http.NewRequestWithContext(c.Context(),c.Method(),c.Context().URI().String(),c.Context().RequestBodyStream())
	b := make([]byte, 32)
	// _, err := rand.Read(b)
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := app.Store.Get(ht.Response.Request, "auth-session")
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}
	session.Values["state"] = state
	// err = session.Save(r, w)
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}
	return c.Redirect(authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func Signup(c *fiber.Ctx) error {
	return c.Render("signup",fiber.Map{})
}

func Logout(c *fiber.Ctx) error{
	domain := "alumnus.us.auth0.com"

	logoutUrl, err := url.Parse("https://" + domain)

	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	// var scheme string
	// if c.Context().IsTLS() {
	// 	scheme = "https"
	// } else {
	// 	scheme = "http"
	// }

	c.Context().Host()
	// returnTo, err := url.Parse(scheme + "://" + c.Context().Host())

	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}
	// parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", "jRiCVXAKYgxCBWBGMNFnS8RE1TfW4SyR")
	logoutUrl.RawQuery = parameters.Encode()
	return c.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}

func Dash(c *fiber.Ctx) error {
	// session, err := app.Store.Get(r, "auth-session")
	// if err != nil {
	// 	c.Context().Error(err.Error(), http.StatusInternalServerError)
	// 	return err
	// }
	return c.Render("dash",fiber.Map{})
}

func Profile(c *fiber.Ctx) error {
	fmt.Println(c.Params("username"))
	return c.Render("profile",fiber.Map{})
}

func Settings(c *fiber.Ctx) error {
	return c.Render("settings",fiber.Map{})
}

func CallbackHandler(c *fiber.Ctx) error{
	ht,err := http.NewRequestWithContext(c.Context(),c.Method(),c.Context().URI().String(),c.Context().RequestBodyStream())
	session, err := app.Store.Get(ht.Response.Request, "auth-session")
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}

	if c.Query("state") != session.Values["state"] {
		c.Context().Error("Invalid state parameter" , http.StatusBadRequest)
		return err
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}

	token, err := authenticator.Config.Exchange(context.TODO(), c.Query("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		c.SendStatus(http.StatusUnauthorized)
		return err
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		c.Context().Error("No id_token field in oauth2 token.",http.StatusInternalServerError)
		return err
	}

	oidcConfig := &oidc.Config{
		ClientID: "jRiCVXAKYgxCBWBGMNFnS8RE1TfW4SyR",
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		c.Context().Error("Failed to verify ID Token: " + err.Error(),http.StatusInternalServerError)
		return err
	}

	
	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		c.Context().Error(err.Error(),http.StatusInternalServerError)
		return err
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	c.Context()
	// err = session.Save(ht.Response.Request)
	if err != nil {
		c.Context().Error(err.Error(), http.StatusInternalServerError)
		return err
	}

	// Redirect to logged in page
	return c.Redirect("/alum",http.StatusSeeOther)
	// http.Redirect(w, r, "/user", http.StatusSeeOther)
}
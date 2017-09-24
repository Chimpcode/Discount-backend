package auth_api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"errors"
)

const (
	FACEBOOK_KEY = "1862002230782717"
	FACEBOOK_SECRET = "96755ea094756358108238f71ec9a050"
)

var SessionsManager *sessions.Sessions

func BeginAuthHandler(ctx iris.Context) {
	url, err := GetAuthURL(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Writef("%v", err)
		return
	}

	ctx.Redirect(url, iris.StatusTemporaryRedirect)
}

func GetAuthURL(ctx iris.Context) (string, error) {
	providerName := "facebook"

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return "", err
	}
	state := "state"
	sess, err := provider.BeginAuth(state)
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}
	session := SessionsManager.Start(ctx)
	session.Set(providerName, sess.Marshal())
	return url, nil
}

var CompleteUserAuth = func(ctx iris.Context) (goth.User, error) {
	providerName := "facebook"

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}
	session := SessionsManager.Start(ctx)
	value := session.GetString(providerName)
	if value == "" {
		return goth.User{}, errors.New("session value for " + providerName + " not found")
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, ctx.Request().URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	session.Set(providerName, sess.Marshal())
	return provider.FetchUser(sess)
}

func Logout(ctx iris.Context) error {
	providerName := "facebook"

	session := SessionsManager.Start(ctx)
	session.Delete(providerName)
	return nil
}

func SetAuthAPI(api iris.Party) {
	goth.UseProviders(facebook.New(FACEBOOK_KEY, FACEBOOK_SECRET, "http://localhost:3000/auth/facebook/callback"))

	// start of the router

	api.Get("/auth/facebook/callback", func(ctx iris.Context) {

		user, err := CompleteUserAuth(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("%v", err)
			return
		}
		ctx.ViewData("", user)
		if err := ctx.View("user.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	api.Get("/logout/facebook", func(ctx iris.Context) {
		Logout(ctx)
		ctx.Redirect("/", iris.StatusTemporaryRedirect)
	})

	api.Get("/auth/facebook", func(ctx iris.Context) {
		// try to get the user without re-authenticating
		if gothUser, err := CompleteUserAuth(ctx); err == nil {
			ctx.ViewData("", gothUser)
			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			BeginAuthHandler(ctx)
		}
	})

	api.Get("/", func(ctx iris.Context) {

		ctx.ViewData("", 3)

		if err := ctx.View("index.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})
}
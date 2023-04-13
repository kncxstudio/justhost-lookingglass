package main

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func GetCookie() (cookieStr string) {
	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		panic(err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		panic(err)
	}
	defer page.Close()

	_, err = page.Goto("https://justhost.asia/en/looking-glass")
	if err != nil {
		panic(err)
	}

	page.WaitForLoadState("domcontentloaded")

	cookies, err := page.Context().Cookies()
	if err != nil {
		panic(err)
	}

	for _, cookie := range cookies {
		cookieStr += fmt.Sprintf("%s=%s; ", cookie.Name, cookie.Value)
	}

	return
}

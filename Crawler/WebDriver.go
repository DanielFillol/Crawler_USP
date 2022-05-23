package Crawler

import (
	"errors"
	"github.com/tebeka/selenium"
)

func WebDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, errors.New("error on creating web driver")
	}

	return driver, nil
}

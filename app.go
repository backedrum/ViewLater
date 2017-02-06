package main

import (
	"github.com/dominikh/go-js-dom"
	"github.com/fabioberger/chrome"
	"strconv"
)

func main() {
	c := chrome.NewChrome()

	tabDetails := chrome.Object{
		"active": false,
	}

	c.Tabs.Create(tabDetails, func(tab chrome.Tab) {
		notification := "Tab with id: " + strconv.Itoa(tab.Id) + " created!"
		dom.GetWindow().Document().GetElementByID("notification").SetInnerHTML(notification)
	})

}

package main

import (
	"encoding/json"
	"github.com/fabioberger/chrome"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"strconv"
)

type Link struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type LinksList []Link

var storage = js.Global.Get("localStorage")

var urls = LinksList{}

func main() {

	c := chrome.NewChrome()
	d := dom.GetWindow().Document()

	table := d.GetElementByID("links")

	// get json from storage
	linksJson := storage.Get("readLaterLinks").String()

	if linksJson != "" {
		err := json.Unmarshal([]byte(linksJson), &urls)
		if err != nil {
			println("Error has been occurred ", err.Error())
		}
	}

	for _, link := range urls {
		row := d.CreateElement("tr").(*dom.HTMLTableRowElement)
		urlColumn := d.CreateElement("td").(*dom.HTMLTableCellElement)
		row.AppendChild(urlColumn)

		urlColumn.SetAttribute("id", strconv.Itoa(link.Id))
		urlColumn.SetInnerHTML("<a href=\"" + link.Url + "\">" + link.Title + "</a>")

		addRemoveButton(d, table, row, link.Id)
		table.AppendChild(row)
	}

	add := d.GetElementByID("addButton").(*dom.HTMLInputElement)
	add.Call("addEventListener", "click", func(event *js.Object) {
		row := d.CreateElement("tr").(*dom.HTMLTableRowElement)
		urlColumn := d.CreateElement("td").(*dom.HTMLTableCellElement)
		row.AppendChild(urlColumn)

		c.Windows.GetCurrent(chrome.Object{}, func(window chrome.Window) {
			id := window.Id

			c.Tabs.GetSelected(id, func(tab chrome.Tab) {
				currentUrl := tab.Url

				nextId := 0
				if len(urls) > 0 {
					nextId = urls[len(urls)-1].Id + 1
				}

				urlColumn.SetAttribute("id", strconv.Itoa(nextId))
				urlColumn.SetInnerHTML("<a href=\"" + currentUrl + "\">" + tab.Title + "</a>")
				addRemoveButton(d, table, row, nextId)
				table.AppendChild(row)

				urls = append(urls, Link{Id: nextId, Url: currentUrl, Title: tab.Title})
				marshalUrlsToStorage()
			})

		})
	})
}

func addRemoveButton(d dom.Document, table dom.Element, tableRow *dom.HTMLTableRowElement, urlId int) {
	remove := d.CreateElement("input").(*dom.HTMLInputElement)
	remove.SetAttribute("type", "button")
	remove.SetAttribute("value", "X")
	remove.SetClass("button removeButton")

	remove.Call("addEventListener", "click", func(event *js.Object) {
		removeLink(urlId)
		marshalUrlsToStorage()
		table.RemoveChild(tableRow)
	})
	tableRow.AppendChild(remove)

	return
}

func removeLink(id int) {
	modified := LinksList{}
	for _, link := range urls {
		if link.Id == id {
			continue
		}

		modified = append(modified, link)
	}

	urls = modified
}

func marshalUrlsToStorage() {
	bytes, err := json.Marshal(urls)
	if err != nil {
		println("Error has occurred during marshalling ", err.Error())
	}

	storage.Set("readLaterLinks", string(bytes))
}

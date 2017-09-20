# ViewLater is a Chrome extension to keep links you want (maybe) to visit later

ViewLater is inspired by ReadLater extension (https://github.com/napsternxg/ReadLater).
It allows you to pin the link you would like to visit later. 
ViewLater Chrome extension is written in Go with the usage of GopherJS.

# How to install?
1. Download the sources.
2. Init govendor:
```
govendor init
govendor fetch github.com/fabioberger/chrome
govendor fetch honnef.co/go/js/dom
```
3. Build the extension:
```
govendor fetch honnef.co/go/js/dom
```
As result viewlater/ folder will be created. 

4. Go to chrome://extensions/ and click on "Load unpacked extension..." button.
Navigate to viewlater/ folder and click "Select" button. Extension will be added to you Chrome browser.

# How to use?
When on the tab you want to view later click on ViewLater extension button and press "Add".
Want to visit previously pinned link? Click on ViewLater extension button and then click on the link to visit.

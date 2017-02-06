ASSEMBLY=./quicktip

.DEFAULT_GOAL: $(ASSEMBLY)

$(ASSEMBLY):
	$(GOPATH)/bin/gopherjs build app.go -o ${ASSEMBLY}/app.js
	cp manifest.json ${ASSEMBLY}
	cp info.html ${ASSEMBLY}
clean:
	if [ -d ${ASSEMBLY} ] ; then rm -rf ${ASSEMBLY} ; fi

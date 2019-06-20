GOPATH=$(CURDIR)/../../../../
GOPATHCMD=GOPATH=$(GOPATH)

COVERDIR=$(CURDIR)/.cover
COVERAGEFILE=$(COVERDIR)/cover.out
COVERAGEREPORT=$(COVERDIR)/report.html

test:
	@${GOPATHCMD} ginkgo --failFast ./...

test-watch:
	@${GOPATHCMD} ginkgo watch -cover -r ./...

coverage-ci:
	@mkdir -p $(COVERDIR)
	@${GOPATHCMD} ginkgo -r -covermode=count --cover --trace ./
	@echo "mode: count" > "${COVERAGEFILE}"
	@find ./* -type f -name *.coverprofile -exec grep -h -v "^mode:" {} >> "${COVERAGEFILE}" \; -exec rm -f {} \;

coverage: coverage-ci
	@sed -i -e "s|_$(CURDIR)/|./|g" "${COVERAGEFILE}"
	@cp "${COVERAGEFILE}" coverage.txt

coverage-html:
	@$(GOPATHCMD) go tool cover -html="${COVERAGEFILE}" -o $(COVERAGEREPORT)
	@xdg-open $(COVERAGEREPORT) 2> /dev/null > /dev/null

dep-ensure:
	@$(GOPATHCMD) dep ensure -v $(PACKAGE)

dep-update:
ifdef PACKAGE
	@$(GOPATHCMD) dep ensure -update -v $(PACKAGE)
else
	@echo "Usage: PACKAGE=<package url> make dep-update"
	@echo "The environment variable \`PACKAGE\` is not defined."
endif

vet:
	@$(GOPATHCMD) go vet ./...

fmt:
	@$(GOPATHCMD) go fmt ./...


.PHONY: test test-watch dep-ensure dep-update coverage coverage-ci coverage-html vet fmt
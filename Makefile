all: $(patsubst %.adoc,%.html,$(wildcard *.adoc))

%.pdf: %.xml
	pandoc -f docbook -S $? --latex-engine=xelatex -o $@

%.xml: %.adoc
	asciidoc -b docbook -d article -o $@ $?

%.html: %.adoc
	asciidoctor -b html5 \
		--failure-level=WARN \
		-d article \
		-o $@ \
		$<

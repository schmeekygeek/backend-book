PANDOCFLAGS = --pdf-engine=xelatex \
		--indented-code-classes=javascript \
		--highlight-style=monochrome \
		-V documentclass=report \
		-V papersize=A5 \
		-V geometry:margin=0.7in

intro:
	mkdir -p dist && \
		pandoc src/introduction/*.md \
		-o dist/introduction.pdf \
		metadata.txt \
		$(PANDOCFLAGS)


book:
	mkdir -p dist && \
		pandoc src/book/*.md \
		-o dist/book.pdf \
		--table-of-contents \
		--number-sections \
		$(PANDOCFLAGS)

merge:
	pdfunite \
		dist/introduction.pdf \
		dist/book.pdf \
		dist/result.pdf && \
	rm dist/introduction.pdf \
		dist/book.pdf

all: intro book merge

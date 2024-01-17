.DEFAULT_GOAL := build

build:
	mkdir -p dist && \
	pandoc src/*.md -o dist/book.pdf \
		--table-of-contents \
		--number-sections \
		metadata.txt \
		--pdf-engine=xelatex \
		--indented-code-classes=javascript \
		--highlight-style=monochrome \
		-V documentclass=report \
		-V papersize=A5 \
		-V geometry:margin=0.5in

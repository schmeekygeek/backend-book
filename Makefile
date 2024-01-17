.DEFAULT_GOAL := build

build:
	pandoc book.md -o book.pdf \
		--table-of-contents \
		--number-sections \
		--pdf-engine=xelatex \
		--indented-code-classes=javascript \
		--highlight-style=monochrome \
		-V documentclass=report \
		-V papersize=A5 \
		-V geometry:margin=1in

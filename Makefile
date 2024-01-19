PANDOCFLAGS = --pdf-engine=xelatex \
							--indented-code-classes=javascript \
							--highlight-style=monochrome \
							-V documentclass=book \
							-V papersize=A5 \
							--include-before-body src/book/cover.md \
							-V geometry:margin=0.7in

book:
	mkdir -p dist && \
		pandoc src/book/*.md \
		-o dist/book.pdf \
		--table-of-contents \
		--number-sections \
		-f markdown+implicit_figures \
		metadata.txt \
		$(PANDOCFLAGS)

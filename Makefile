PANDOCFLAGS = --pdf-engine=xelatex \
							--indented-code-classes=javascript \
							--highlight-style=kate \
							-V documentclass=book \
							-V papersize=A5 \
							-V monofont=Iosevka\ Nerd\ Font\ Complete.ttf \
							-V geometry:margin=0.7in \
							-o dist/book.pdf \
							--table-of-contents \
							--number-sections \
							-f markdown+implicit_figures \
							metadata.txt \

book:
	mkdir -p dist && \
		pandoc src/book/*.md \
		$(PANDOCFLAGS)

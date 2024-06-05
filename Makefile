Iosevka = Iosevka\ Nerd\ Font\ Complete.ttf
FONT = 
PANDOCFLAGS = --pdf-engine=xelatex \
							--indented-code-classes=javascript \
							--highlight-style=tango \
							-V documentclass=book \
							-V papersize=A5 \
							-V linkcolor:blue \
							-V monofont=$(FONT) \
							-V geometry:margin=0.7in \
							--table-of-contents \
							--number-sections \
							-f markdown+implicit_figures \
							metadata.txt \

pdf:
	mkdir -p dist && \
		pandoc src/book/*.md \
		-o dist/book.pdf \
		$(PANDOCFLAGS)


epub:
	mkdir -p dist && \
		pandoc src/book/*.md \
		-o dist/book.epub \
		$(PANDOCFLAGS)

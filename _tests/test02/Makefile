.PHONY: tex
tex:
	../../dist/dox_linux_amd64_v1/dox -i test.md -o test.tex -p preamble.tex

.PHONY: pdf
pdf:
	lualatex --shell-escape -output-directory=. test.tex

.PHONY: pdf2
pdf2: tex
	pdflatex --shell-escape test.tex

.PHONY: clean
clean:
	@rm -rf *.log *.pdf test.tex *.out *.aux _minted-* *.fls *.fdb_latexmk *.pyg

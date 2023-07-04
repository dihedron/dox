.PHONY: binary
binary:
	goreleaser build --single-target --snapshot --clean

.PHONY: clean
clean:
	rm -rf _test/*.log _test/*.pdf _test/test*.tex _test/*.out _test/*.aux _test/_minted-*

.PHONY: tex
tex: binary
	dist/dox_linux_amd64_v1/dox -i _test/test0.md -o _test/test0.tex -p _test/preamble.tex

.PHONY: pdf
pdf:
	lualatex -shell-escape -output-directory=_test _test/test0.tex

# .PHONY: test1
# test1: binary
# 	dist/dox_linux_amd64_v1/dox -i _test/test1.md -o _test/test1.tex -p _test/preamble.tex
# 	pdflatex -output-directory=_test _test/test1.tex

# .PHONY: test2
# test2: binary
# 	dist/dox_linux_amd64_v1/dox -i _test/test2.md -o _test/test2.tex
# 	pdflatex -output-directory=_test _test/test2.tex

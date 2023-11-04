.PHONY: binary
binary:
	goreleaser build --single-target --snapshot --clean

.PHONY: clean
clean: ./_tests/*/
	@for test in $^; do  \
		echo "this is my path: [$${test}]" && \
		cd $${test} && \
		make clean && \
		cd ../..; \
	done


# .PHONY: test1
# test1: binary
# 	dist/dox_linux_amd64_v1/dox -i _test/test1.md -o _test/test1.tex -p _test/preamble.tex
# 	pdflatex -output-directory=_test _test/test1.tex

# .PHONY: test2
# test2: binary
# 	dist/dox_linux_amd64_v1/dox -i _test/test2.md -o _test/test2.tex
# 	pdflatex -output-directory=_test _test/test2.tex

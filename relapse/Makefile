regenerate:
	rm -rf ./tests/
	(cd gen-relapse-tests && make regenerate && go install .)
	gen-relapse-tests ./tests/
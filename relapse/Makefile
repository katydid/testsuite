regenerate:
	rm -rf ./tests/ || true
	rm -rf ./benches/ || true
	(cd gen-relapse-tests && make regenerate && go install .)
	gen-relapse-tests -seed 12345 -benches .
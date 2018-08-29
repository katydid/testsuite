regenerate:
	make nuke
	(cd gen-relapse-tests && make regenerate && go install .)
	gen-relapse-tests -seed 12345 -benches .

regenerate-tests:
	make nuke-tests
	(cd gen-relapse-tests && make regenerate && go install .)
	gen-relapse-tests

nuke-tests:
	rm -rf ./tests/ || true

nuke-benches:
	rm -rf ./benches/ || true

nuke:
	make nuke-tests
	make nuke-benches
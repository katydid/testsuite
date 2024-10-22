regenerate-tests:
	make nuke-tests
	(cd gen-validator-tests && make regenerate && go install .)
	gen-validator-tests -seed 12345

regenerate-all:
	make nuke
	(cd gen-validator-tests && make regenerate && go install .)
	gen-validator-tests -seed 12345 -benches .

nuke:
	make nuke-tests
	make nuke-benches

nuke-tests:
	rm -rf ./tests/ || true

nuke-benches:
	rm -rf ./benches/ || true

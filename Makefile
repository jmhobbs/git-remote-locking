default: install

install: install-lock install-unlock manpages

manpages:
	cp ./man/* /usr/local/share/man/man1/

install-lock:
	cd git-lock-remote && go install

install-unlock:
	cd git-unlock-remote && go install

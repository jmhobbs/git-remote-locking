# Easily lock and and unlock the pushurl for a remote

Don't want to accidentally push to your upstream?  Lock it.

```
fatal: 'no_push' does not appear to be a git repository
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.
```
![Demo](https://s3.amazonaws.com/static.velvetcache.org/temp/longterm/git-lock-demo.gif)

# Installation

### Option 1
Clone the repo and run `make install`

### Option 2
`go get -u github.com/jmhobbs/git-remote-locking/...` (no man files)

### Option 3
Download from [releases](https://github.com/jmhobbs/git-remote-locking/releases) and put in your PATH

# Uh, this could just be some shell scripts

Yes. They could just be some shell scripts.

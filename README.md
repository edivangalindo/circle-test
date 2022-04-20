# circle-test

Sometimes, during the day as an offensive security engineer or bug hunter, we come across a leaked token and we need to test it. This tool was born out of my need to do this routinely with tokens coming from CircleCI.

Single URL:

```
echo asd98asd909890898908990890890890 | circle-test
```

Multiple tokens:

```
cat tokens.txt | circle-test
```

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile circle-test:
```
go install github.com/edivangalindo/circle-test@latest
```

You can now run `~/go/bin/circle-test`. If you'd like to just run `circle-test` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.

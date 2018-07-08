h12.io/cv
=========

An HTML CV generator in Go that suports:
* JSON Resume data format
* Go text/template

Usage
-----

```bash
go get h12.io/cv/cmd/mkcv                  # install mkcv command line tool
cd $GOPATH/src/h12.io/cv/tmpl/mnjul-resume # choose a template
$GOPATH/bin/mkcv                           # execute mkcv with default options
open resume.html                           # view generated resume.html

mkcv -h # to see available command line options
```

Credit
------

This work is inspired by:

* resumake.io
* github.com/mnjul/html-resume
* jsonresume.org

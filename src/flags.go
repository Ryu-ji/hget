package main

import "flag"

var (
	flagUseDigest = flag.Bool("d", false, "")

	flagOut    = flag.String("out", "", "")
	flagDigest = flag.String("digest", "", "")
	flagUri    = flag.String("uri", "", "")
)

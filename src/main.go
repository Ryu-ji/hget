package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"./fetch"
)

func about() {
	fmt.Print("*HGET Command 0.0.1 dev\n *Copyright Ryuji Furukawa(ryu-ji)\n  *https://ryu-ji.github.io\n   *https://github.com/Ryu-ji/hget\n    *MIT License\n     *HGET is a tool that downloads files using the http protocol.\n\n")
}

func version() { fmt.Print("0.0.1 dev\n\n") }

func help() {}

func cmdRun(cmd string) {

	switch strings.ToLower(cmd) {
	case cmdVersion:
		version()
		break

	case cmdAbout:
		about()
		break

	case cmdHelp:
		help()
		break

	default:
		cmdRun(cmdHelp)
	}

	os.Exit(0)
}

func download(uri string) (file string, data []byte) {

	var res fetch.Response

	var err error

	if getScheme(uri) == SchemeHTTPS {

		res, err = fetch.HTTPSFetch(uri)

	} else {
		res, err = fetch.HTTPFetch(uri)
	}

	if err != nil {
		log.Println(err)

		return file, data
	}

	if filepath.Ext(res.FileName) != res.Ext {
		res.FileName += res.Ext
	}

	file = res.FileName

	data = res.Body

	return file, data
}


func main() {

	//コマンドライン処理
	flag.Parse()

	if flag.Lookup("uri").Value.String() == "" && flag.NArg() == 1 {

		cmd := flag.Arg(0)

		if flag.NFlag() == 0 && isVaildScheme(cmd) == false {
			cmdRun(cmd)
		} else {
			flagUri = &cmd
		}
	}

	//出力先
	var outPaths []string
	var outPathsLen int = 0

	if flag.Lookup("out").Value.String() != "" {
		outPaths = strings.Split(*flagOut, " ")
		outPathsLen = len(outPaths)
	}

	//ダウンロード処理
	for i, uri := range strings.Split(*flagUri, " ") {

		//-oの指定がなければURLからのファイル名を使う
		var file string
		var data []byte

		file, data = download(uri)

		if i != outPathsLen {
			file = outPaths[i]
		}

		write(file, data)

		if *flagUseDigest == true {
			d := digest(data, choiceHashFunc(flagDigest))
			fmt.Println(uri, d)
		}

	}

}

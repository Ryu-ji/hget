package fetch

import (
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)

type Response struct {
	Body     []byte
	FileName string
	Ext      string
	/*実装予定Charset string*/
}

func arrangeTheResponse(resp http.Response) (response Response) {
	response.Body, _ = ioutil.ReadAll(resp.Body)

	cnttype := resp.Header.Get("Content-Type")

	exts, err := mime.ExtensionsByType(cnttype)

	if err != nil {
		response.Ext = ""
	}

	response.Ext = choiceExt(exts)

	response.FileName = filepath.Base(resp.Request.URL.String())

	//実装予定 response.Charset = strings.Split(cnttype, ";")[1]

	return response
}

func choiceExt(exts []string) string {

	var ext string

	extsLen := len(exts)

	for i := 0; i < extsLen; i++ {

		if i == extsLen || i == extsLen-1 {
			continue
		}

		if i == 0 {
			ext = exts[i]
		}

		x := len(ext)
		y := len(exts[i+1])

		if x < y {
			ext = exts[i+1]
		}

	}
	return ext
}

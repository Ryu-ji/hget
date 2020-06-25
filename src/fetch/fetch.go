package fetch

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

func HTTPFetch(uri string) (Response, error) {

	parsedURI, err := url.Parse(uri)

	if err != nil {
		return Response{}, err
	}

	res, err := makeReq(parsedURI.String(), nil, http.Client{})

	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()

	return arrangeTheResponse(*res), nil
}

func HTTPSFetch(uri string) (Response, error) {

	parsedURI, err := url.Parse(uri)

	if err != nil {
		return Response{}, err
	}

	tls := &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: parsedURI.Hostname()},
	}

	res, err := makeReq(parsedURI.String(), nil, http.Client{Transport: tls})

	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()

	return arrangeTheResponse(*res), nil

}

func makeReq(uri string, headers map[string]string, client http.Client) (*http.Response, error) {

	req, _ := http.NewRequest("GET", uri, nil)

	if headers != nil {

		for k, v := range headers {
			req.Header.Set(k, v)
		}

	}

	return client.Do(req)
}

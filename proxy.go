package main

import (
    "github.com/zeroxoneafour/html-rewriter"
    "log"
    "net/http"
    "os"
    "net/url"
    "io"
    "strings"
)

type proxyHandler struct {} // empty struct lol

func (h proxyHandler) ServeHTTP (resp http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/" {
        client := new(http.Client)

        req.URL, _ = url.Parse(req.URL.String()[1:])
        url := req.URL.String()
        req.RequestURI = "" // idk why i have to do this
        if req.URL.Scheme == "" { // return if it requests ex. favicon.ico
            return
        }
        log.Println("Req.URL:", url)

        req.Host = req.URL.Host // set host
        delete(req.Header, "Accept-Encoding") // we don't like encodings
        rawResp, err := client.Do(req)

        if err != nil {
            log.Panicln("Error encountered in making URL request:", err)
        } else {
            log.Println("Status:", rawResp.Status)
            buf := new(strings.Builder)
            io.Copy(buf, rawResp.Body)
            var newPage *strings.Reader
            log.Println(rawResp.Header["Content-Type"][0])
            if strings.Contains(rawResp.Header["Content-Type"][0], "text/html") {
                newPage = strings.NewReader(htmlrewriter.RewriteHTML(buf.String(), url))
            } else {
                newPage = strings.NewReader(buf.String()) // probably a more efficient way to do this but idk
            }
            for k, v := range(rawResp.Header) { // set headers
                resp.Header()[k] = v
            }
            io.Copy(resp, newPage)

        }
    } else {
        io.Copy(resp, strings.NewReader("<h1>go-rproxy by Vaughan Milliman</h1>\n<a href=\"https://github.com/zeroxoneafour/go-rproxy\">Source Code</a> licensed under the MIT License"))
    }
}

func main() {
    var port, host string
    if port = os.Getenv("PORT"); len(port) == 0 { // default port 8000, all interfaces
        port = "8000"
    }
    host = os.Getenv("HOST") // this can be empty
    addr := host + ":" + port
    log.Println("Listening on address", addr)
    var mainHandler proxyHandler
    http.ListenAndServe(addr, mainHandler)
}

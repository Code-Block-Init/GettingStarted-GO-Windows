Getting started for "Go Language" in Windows 8.1 x64 <br>
<b>Step 1: Installation:-</b><br>
Download the <a href="https://storage.googleapis.com/golang/go1.4.2.windows-amd64.msi">zipped package</a> and unzip it.<br>
<b>Step 2: Setting up environment variables:</b><br>
1. variable name: <code>path</code><br>
variable value: <code>D:\go</code><br>
2. variable name: <code>GOROOT</code><br>
variable value: <code>D:\go</code><br>
<b>Step 3: Testing whether it's working or not </b><br>
<code>hello.go</code> <br>
```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
Download <a href="http://git-scm.com/download/win">Git for Windows</a><br>
<code>$cd "/d/go"</code><br>
<code>$go run bin/hello.go</code><br>
<br>
This repo is dedicated to <code>GoLang</code> basics.

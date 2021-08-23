# jsleak
jsleak is a tool to identify sensitive data in JS files through regex patterns. Although it's built for this, you can use it to identify anything as long as you have a regex pattern for it.

### How to install

Directly:
```
{your package manager} install pkg-config libpcre++-dev
go get github.com/0xTeles/jsleak/v2/jsleak
```
Compiled:  [release page](https://github.com/0xTeles/jsleak/releases/tag/jsleak_v2.1)

### How to use
```
-json string
        [+] Json output file
  -pattern string
        [+] File contains patterns to test
  -timeout int
        [+] Timeout for request in seconds (default 5)
  -verbose
        [+] Verbose Mode
```
### Demo

```
cat urls.txt | jsleak -pattern regex.txt
[+] Url: http://localhost/index.js
[+] Pattern: p([a-z]+)ch
[+] Match: peach
```


### To Do
- [x] Fix output
- [ ] Add more patterns
- [x] Add stdin
- [ ] Implement JSON input
- [x] Fix patterns
- [x] Implement PCRE

### Regex list
- https://github.com/odomojuli/RegExAPI
- https://github.com/KaioGomesx/JSScanner/blob/main/regex.txt
### Inspired by 
- Necessity
- https://github.com/0x240x23elu/JSScanner
- https://github.com/KaioGomesx/JSScanner
### Thanks
[@fepame](https://twitter.com/Highustavo), [@gustavorobertux](https://twitter.com/gustavorobertux), [@Jhounx](https://github.com/Jhounx), [@arthurair_es](https://twitter.com/arthurair_es)

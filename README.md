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
Usage of jsleak:
  -pattern string
        File contains patterns to test
  -url string
        JS endpoint to test
```
### Demo

```
jsleak -url http://localhost/index.js -pattern regex.txt
[+] Url: http://localhost/index.js
[+] Pattern: p([a-z]+)ch
[+] Match: peach
```


### To Do
- [x] Fix output
- [ ] Add more patterns
- [ ] Add stdin
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
[@fepame](https://twitter.com/Highustavo), [@gustavorobertux](https://twitter.com/gustavorobertux)

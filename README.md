# add-anchor-links-to-html

<p align="center">
  <img src=./description.png />
</p>

## Motivation

When I convert .epub to .html with pandoc, I get a very long html document.
After reading some sections in it, I want to suspend reading and remember which chapter to start reading tomorrow.
However, the exported html doc does not have anchor links...
Thus, I add them to html with this script.

オライリー から買った電子書籍を pandoc で epub ファイルを変換すると、めちゃくちゃ長い 1 つの HTML ファイルになる。
この HTML ファイルにはアンカーリンクがないので、途中まで読んだ部分を記録するのがとても面倒だ。
だから、 epub ファイルを HTML ファイルに変換したあとは、このスクリプトでアンカーリンクを付けることにした。


## Getting started

Here is example html file to be input.

```sh
$ go get -u github.com/iwasa-kosui/add-anchor-links-to-html
$ cat example.html
<html><head>
    
</head>
<body>
<div class="post">
    <h1 id="aaa">aaa</h1>
    <h3 id="bbb">bbb</h3>
</div>
</body></html>
```

### outputs the result

```sh
$ ./add-id-links-to-html -i example.html
<html><head>
    
</head>
<body>
<div class="post">
    <a href="#aaa"><h1 id="aaa">aaa</h1></a>
    <a href="#bbb"><h3 id="bbb">bbb</h3></a>
</div>
</body></html>
```

### saves to the file

```
$ ./add-id-links-to-html -i example.html -o output.html
```



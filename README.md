# tools

A collection of small command line utilities:

* comicsdigest - send email with favorite comics
* hxd - a small hexdumper
* imgsize - show sizes of jpeg or png images
* kwed-dl - download latest tracks from remix.kwed.dk
* qotd - Quote of the Day
* remtilde - remove vim backup files (\*~)
* sauk - a small static web server
* tolower - rename all files in current directory to lowercase filenames (so ABC.jpg -> abc.jpg)
* waves - create cos and sin tables
* wi - "w"here "i"s (locate file)
* zen - a small "project management" (ahem) tool


## sauk

If invoked without argument sauk binds to port 8080 and serves files from
the current directory.

The following options can be given to sauk:

```
 -d path         Document root  
 -h              Help  
 -p num          Listen on port number  
```

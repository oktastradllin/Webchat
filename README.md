# Webchat
Ini adalah repositori pembuatan web chat sederhana dengan golang Rest &amp; Websocket

## Tabel Konten
  * Tech
  * API
      * PULL API
      * PUSH API 
  * Troubleshoot
    * [Author](#author)
  * [License](#license)



## Tech

Menggunakan Golang dan json file sebagai penyimpanan pesan - pesan lama 

## API

### PULL API
Menggunakan REST third party pada golang bila belum terinstall silahkan gunakan 

```
$ go get github.com/gorilla/mux
```

#### 1. API to send message

![post-api](http://www.juara.net.s3.amazonaws.com/post-api.png)

Rest handler post 
```
r := mux.NewRouter()
r.HandleFunc("/message", allMessage).Methods("POST")
```
Param akan di validasi dulu dan bila berhasil akan di masukkan ke dalam json file 
dan bila sukses akan seperti ini

![post-api](http://www.juara.net.s3.amazonaws.com/post-message.png?v=2)



#### 2. API to retrieve messages

![get-api](http://www.juara.net.s3.amazonaws.com/get-api.png)

Rest handler get 
```
r := mux.NewRouter()
r.HandleFunc("/message", allMessage).Methods("GET")
```
Semua chat yang lama akan di ambil dari allMessage.json  
![post-api](http://www.juara.net.s3.amazonaws.com/get-message.png)

### PUSH API
Menggunakan Websocket third party pada golang bila belum terinstall silahkan gunakan 

```
$ $ go get github.com/gorilla/websocket
```



#### 3. API to retrieve message at realtime

![get-api](http://www.juara.net.s3.amazonaws.com/websocket.png)

akses home akan masuk ke dalam template dengan tampilan chat sederhana

```
fs := http.FileServer(http.Dir("./public"))
	r.Handle("/", fs)
```
lalu pesan akan dikirim ke websocket /ws oleh javascript dan di proses oleh handleConnection function

```
r.HandleFunc("/ws", handleConnections)
```
dan terakhir buat goroutine untuk memproses pesan chat yang masuk 

```
go handleMessages()
```


![get-api](http://www.juara.net.s3.amazonaws.com/web-chat.png)


### TroubleShoot

#### Menjalankan Golang Listener

```
$ go run main.go
```


#### Setting Path Untuk JSON FILE DB

```
var path = "./Allmessage.json" -- path untuk penyimpanan db
```




### Author
Florentinus Oktavian : okta.stradllin@gmail.com.

License
----




[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>

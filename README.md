# netpaste
`netpaste` is a simple paste bin that allows clients to post pastes over `netcat`.
- built in `Golang 1.6` (i.e. the server is cross-platform)
- supports both UNIX and Windows line endings
- example usage: `echo hello | nc <domain> <port>`


![Demo](https://i.imgur.com/i0zzAsX.gif)



### Installation
Download the files.

    git clone https://github.com/petercunha/netpaste master
  
  
Change dirs into master

    cd master/
    

Open up `main.go` and enter your desired connection settings.
```go
func main() {
	// Specify the host and port you'd like to listen on
	listen("localhost", "3333")
}
```


Now, let's build it.

    go build


Ok. Good job. Now, lets make a paste directory in `/var/www/html/`

    mkdir /var/www/html/paste
    chown -R www-data:www-data /var/www/html/paste/

Nice! Almost done. Open up `/etc/apache2/apache.conf` and add this
```
<Directory /var/www/html/paste>
	DirectoryIndex index.txt
</Directory>
```



### Run the server
Everythings set up and configured correctly. Now you can run it by doing

    sudo ./master



### Test the server
On the same machine, open up a terminal and do

    echo Hello world! | nc localhost <port>

In response you should get

    http://localhost/paste/XXXXXXXX/

Go there in your browser and you should see `Hello world!`

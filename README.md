# StreamingServerPoC

1. `git clone https://github.com/npowanusorn/StreamingServerPoC.git`
2. `go run content.go`
3. `go run lb.go`
4. try downloading from `http://[local-ip]:8080/video/manifest.mpd`
5. compare with downloading from `http://[local-ip]:9090/manifest.mpd`
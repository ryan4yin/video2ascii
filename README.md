# Video2ASCII

Yet another version of video2ascii, writtern in Go.

other versions:

- [video2chars(python)](https://github.com/ryan4yin/video2chars)
- [Video2ASCII.jl(julia)](https://github.com/ryan4yin/Video2ASCII.jl)
- [Video2ASCII.jl(rust)](https://github.com/ryan4yin/video2ascii-rs)

>P.S. Video2ASCII is my personal "hello world" project, 
every time I learn a new programming language, I will implement it in that languange.


## Dependencies

then install dependencies:

for `opensuse`:

```shell
sudo zypper in libvpx-devel libopus-devel
sudo zypper in opencv opencv-devel
sudo zypper in clang clang-devel

# use mpv to play audio
sudo zypper in mpv
```

for macOS:

```shell
brew install opencv

# use mpv to play audio
brew install mpv
```

see [gocv](https://github.com/hybridgroup/gocv)'s docs for more details.

## How to Run

```
go mod download
go run cmd/main.go
```


## Reference

- [gocv](https://github.com/hybridgroup/gocv)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [asciiplayer](https://github.com/qeesung/asciiplayer)

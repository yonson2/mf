# mehflix - anime from your terminal to your video player.

_mf_ (pronounced _meh-flix_) is a terminal utility to stream anime directly to your video player. It finds its content from [nyaa.si](https://nyaa.si), turns the torrent into a filestream and pipes its contents to a compatible player.

## Installation

Precompiled binaries are available for windows, mac and linux, just head to the [releases](https://github.com/yonson2/mf/releases) page and download the one that suits your platform.

If you want to compile from source, first clone this repository:

```
git clone git@github.com:yonson2/mf.git
```

then run `go build`

```
go build .
```

You should now see a `mf` executable.


### Dependencies

[mf](https://github.com/yonson2/mf) relies on either [mpv](https://mpv.io), [mplayer](http://www.mplayerhq.hu) or [vlc](https://www.videolan.org) to be installed on your system.

## Usage

### Basic usage

If you wanted to stream the best result available (based on the amount of seeders) there are no extra parameters needed, just do `mf <name>` substituting `<name>` by your query, for example, if you want to search for the movie Big Buck Bunny:

```
mf big buck bunny
```

### Advanced usage

To choose which result to stream simply add `-s` as a parameter:

```
mf -s big buck bunny
```

In a similar fashion you can choose how many results you want to see with the `-n` parameter:

```
mf -s -n 10 big buck bunny
```

### Planned Features
 - Keep track of shows searched/episode watched


#leetflix - watch content online directly from your terminal

_leetflix_ is a small terminal utility that queries the internet for multimedia content hosted using the bittorrent protocol and tries to stream it directly to your video player.

## Installation

Precompiled binaries are available for windows, mac and linux, just head to the [releases](https://github.com/yonson2/leetflix/releases) page and download the one that suits your platform.

If you want to compile from source, first clone this repository:

```
git clone git@github.com:yonson2/leetflix.git
```

then run `go build`

```
go build .
```

You should now see a `leetflix` executable.


### Dependencies

`leetflix` relies on either `mpv`, `mplayer` or `vlc` to be installed on your system.

## Usage

###Basic usage

If you wanted to stream the best result available (based on the amount of seeders) there are no extra parameters needed, just do `leetflix <name>` substituting `<name>` by your query, for example, if you want to search for the movie Big Buck Bunny:

```
leetflix big buck bunny
```

###Advanced usage

To choose which result to stream simply add `-s` as a parameter:

```
leetflix -s big buck bunny
```

In a similar fashion you can choose how many results you want to see with the `-n` parameter:

```
leetflix -s -n 10 big buck bunny
```


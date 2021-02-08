# mehflix - video streaming to your video player without leaving your terminal.

_mf_ (pronounced _meɪflɪkz_) is a terminal utility to stream multimedia content directly to your video player. It finds its content from torrent indexers, turns the torrent into a filestream and pipes its contents to a compatible player.

## Usage

### Basic usage

mehflix can be run both as a standalone GUI application and from the terminal, if no arguments are provided when launching the app a GUI will spawn.

If you wanted to stream the best result available (based on the amount of seeders) there are no extra parameters needed, just do `mf <name>` substituting `<name>` by your query, for example, if you want to search for the movie Big Buck Bunny:

```
mf big buck bunny
```
By default `mf` only searches for anime content, to also search for other type of content try using the `-g` flag, like so:
```bash
mf -g big buck bunny
```
### Screenshots
![mehflix](/assets/screenshot.png?raw=true "mehflix GUI")

### Advanced usage

To choose which result to stream simply add `-s` as a parameter:

```
mf -s big buck bunny
```

In a similar fashion you can choose how many results you want to see with the `-n` parameter:

```
mf -s -n 10 big buck bunny
```
You can also always see a list of all available options and a brief description of their purpose with the `--help` command.

## Installation

Precompiled binaries are available for windows, mac and linux, just head to the [releases](https://github.com/yonson2/mf/releases) page and download the one that suits your platform.

If you want to compile from source, first clone this repository:

```
git clone git@github.com:yonson2/mf.git
```

then run `go build`

```
go build
```

You should now see a `mf` executable.

### Dependencies

[mf](https://github.com/yonson2/mf) relies on either [mpv](https://mpv.io), [mplayer](http://www.mplayerhq.hu) or [vlc](https://www.videolan.org) to be installed on your system.

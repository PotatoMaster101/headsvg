# Head SVG
Quick and dirty Minecraft player head to SVG converter. Can be used to generate player head icons.

## Usage
```
headsvg [-h|--help] -n|--name "<value>" [-s|--scale <integer>] [-p|--percent]

Arguments:
  -h  --help   Print help information
  -n  --name   Player username
  -s  --scale  SVG pixel scale. Default: 100
  -p  --percent  Whether to output x/y/width/height as a percentage. Default: false
```

## Building
```
$ go build -ldflags '-s -w' cmd/headsvg/headsvg.go
```

## Running Unit Tests
```
$ go test ./...
```

## Examples
```
$ headsvg -n PotatoMaster101 -s 50 > PotatoMaster101.svg
```

# Not working because of Iceborn

# jewel-extractor

Monster Hunter World Jewel extractor

This software extracts jewels you have from your save data to use at [the skill simulator](https://mhw.wiki-db.com/sim/?hl=en)


## Usage

```
jewel-extractor path/to/SAVEDATA1000
```

Copy the output and paste it to the simulator's import window

## Options

```
Usage of ./jewel-extractor:
  -l string
        language of output (default "en")
  -s int
        save slot (default 1)
```

### Supported languages

English, Japanese, Korean, Chinese


## How to build

We use [packr](https://github.com/gobuffalo/packr) to embed jewel information

```
go get github.com/gobuffalo/packr
go get github.com/gobuffalo/packr/packr
packr build
```

For other platform

```
packr
GOOS="darwin" GOARCH="amd64" go build
```

### References

https://github.com/TanukiSharp/MHWSaveUtils
https://github.com/gatheringhallstudios/MHWorldData
https://www.reddit.com/r/MonsterHunter/comments/812b8n/jewel_trashing_cheat_sheet_v4_updated_for_jewel/

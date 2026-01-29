# Go Screenshot CLI (v0.1.0)

Screenshot URL → ImgBB URL (upload) or PNG.

## Install
```bash
git clone https://github.com/Ntrakiyski/go-cli
cd go-cli
go build -o screenshot ./cmd/screenshot
sudo cp screenshot /usr/local/bin/  # Optional
```

## Usage
```
./screenshot --url=https://example.com --width=1920 --height=1080 [--upload] [--key=your_imgbb_key]

Flags:
  -u,--url     Target URL (required)
  -w,--width   Viewport width (default 1920)
  -h,--height  Viewport height (default 1080)
  -k,--key     ImgBB API key (or IMGBB_API_KEY env)
  -U,--upload  Upload to ImgBB (default true) → prints URL
  -v,--verbose Verbose logs

Examples:
  export IMGBB_API_KEY=your_free_key_from_imgbb.com
  ./screenshot --url=https://golang.org --width=375 --height=667 --upload
  # → ✅ ImgBB URL: https://i.ibb.co/ABC123 (expires ~10min)
```

## Dev
```
go mod tidy
go build -o screenshot ./cmd/screenshot
./screenshot --url=https://example.com --width=1200 --height=800 --upload
```

Requires Chrome/Chromium + free imgbb.com API key.

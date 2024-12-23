# Go Animate Coin

This script will take a circular image, say a coin, and rotate it in the perspective towards you as if the coin is being spun around clockwise facing you.

## Sample

|                       From                       |                        To                         |
|:------------------------------------------------:|:-------------------------------------------------:|
| ![-input path](sample/apario-xrp-coin@144px.png) | ![~output path](sample/apario-xrp-coin@144px.gif) |
|    `-input sample/apario-xrp-coin@144px.png`     |     `-output sample/apario-xrp-coin@144.gif`      |
|            **Transparent PNG Image**             |               _Animated GIF Image_                | 

[View $APARIO on XRP Ledger via XPMarket.com](https://xpmarket.com/token/APARIO-rU16Gt85z6ZM84vTgb7D82QueJ26HvhTz2)

## Usage

```bash

[~/go-animate-coin]$ ./go-animate-coin -help
Usage of ./go-animate-coin:
  -input string
    	Path to input PNG image
  -output string
    	Path for output GIF (default "animated.gif")
  -v	Show version
  -y	Override Y/N prompts with Y response
```

> **NOTE**: You cannot use JPG/JPEG or WebP images with this program.

### Override in action

When you supply an `-output` path that already exists, you'll receive an error...

```bash
[~/go-animate-coin]$ ./go-animate-coin -input sample/apario-xrp-coin@144px.png -output sample/apario-xrp-coin@144px.gif
File sample/apario-xrp-coin@144px.gif already exists. Overwrite? [y/N]: n
2024/12/23 10:39:08 operation cancelled by user
```

When you pass the `-y` flag in this case, you won't get this prompt.

```bash
[~/go-animate-coin]$ ./go-animate-coin -input sample/apario-xrp-coin@144px.png -output sample/apario-xrp-coin@144px.gif -y
Animation saved to sample/apario-xrp-coin@144px.gif
```

## License

This is released under the MIT License.
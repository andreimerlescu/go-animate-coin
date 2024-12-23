# Go Animate Coin

This script will take a circular image, say a coin, and rotate it in the perspective 
towards you as if the coin is being spun around clockwise facing you.

I wanted to create a GIF animation of a new XPMarket.com meme coin called $APARIO and 
I couldn't find anything online so I built one. 

I used Claude AI to help me debug the transformFrame func, but overall, it was a fun 
exercise that took me a few hours to do and I am happy to share it with you under the 
MIT License so you may use it as well.

## Sample

|                       From                       |                        To                         |
|:------------------------------------------------:|:-------------------------------------------------:|
| ![-input path](sample/apario-xrp-coin@144px.png) | ![~output path](sample/apario-xrp-coin@144px.gif) |
|    `-input sample/apario-xrp-coin@144px.png`     |     `-output sample/apario-xrp-coin@144.gif`      |
|            **Transparent PNG Image**             |               _Animated GIF Image_                | 

[View $APARIO on XRP Ledger via XPMarket.com](https://xpmarket.com/token/APARIO-rU16Gt85z6ZM84vTgb7D82QueJ26HvhTz2)

## Install

```bash
go install github.com/andreimerlescu/go-animate-coin@latest
```

Or use **GitLab** if GitHub is unavailable (_it's now owned by Microsoft, so it happens often now_): 

```bash
go install gitlab.com/andreimerlescu/go-animate-coin@latest
```

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

## Contribute

```bash
git clone git@github.com:andreimerlescu/go-animate-coin.git
git checkout develop
git pull origin develop
git checkout -b feature/my-new-feature
```

Submit the change as a pull request to the repository. The GitHub way to do it is to fork the repo,
edit the files you need, then submit the commit and the interface will create the PR in this repo
and integrate the changes from your forked repo into this repo automatically.

You can also use **GitLab** if you wish, 

```bash
git clone git@gitlab.com:andreimerlescu/go-animate-coin.git
```

If you use this project for your XRP, SOL, ETH, or other cryptocurrency, please send a tip: 

Confidently tip any amount with almost no fees (_even if you're just giving just your $0.02_):

- **XRP** _(preferred)_: rK1cJc5Jrhauhae4R84RwPVoRTWcDC3iwS
- **XLM** _(receiving)_: GBYMH7R35BZGFSXYSCUYR55RUH6Q3HAUV5QHV6FLZUPFYUVTAESJPEWK
- **HBAR** _(collecting)_: 0.0.7718744
- **SOL** _(liquidating)_: 3fFWxmTYHZR1ALKoh3aj3Fyoh8c6AnwKqff4M5USFQMH

Due to swap fees, please tip any amount over $10 if you're using :

- **BTC** _(don't hold any)_: 1EKZJRaFrybikUbvJaT8yHP9qzLfoWdiMc
- **ETH** _(fees way too high)_: 0x780A20EfD2884Fb2081D7501a6575Da82d0F027c

Thank you for your support.

## License

This is released under the MIT License.
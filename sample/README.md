# Sample

This image is Public Domain "CC BY-SA". Attributed to Project Apario LLC.

![APARIO Token on XPR Ledger](apario-xrp-coin@144px.gif)

[View $APARIO on XRP Ledger via XPMarket.com](https://xpmarket.com/token/APARIO-rU16Gt85z6ZM84vTgb7D82QueJ26HvhTz2)

## How this image was generated

First we'll need the **PNG** image itself. This will be the `apario-xrp-coin@369px.png` that we'll
convert into an animated gif using this program. We can use the ` -help` argument to see how to
use the `go-animate-coin` program.

```bash
[~/go-animate-coin/sample]$ stat apario-xrp-coin@144px.png
  File: apario-xrp-coin@144px.png
  Size: 17700     	Blocks: 40         IO Block: 4096   regular file
Device: fd02h/64770d	Inode: 1112857921  Links: 1
Access: (0644/-rw-r--r--)  Uid: ( 1000/  andrei)   Gid: ( 1002/  andrei)
Context: unconfined_u:object_r:user_home_t:s0
Access: 2024-12-22 21:27:40.096007941 -0500
Modify: 2024-12-22 20:11:17.716464095 -0500
Change: 2024-12-22 20:11:17.716464095 -0500
 Birth: 2024-12-22 20:11:17.714462173 -0500

[~/go-animate-coin]$ ./go-animate-coin -help
Usage of ./go-animate-coin:
  -input string
    	Path to input PNG image
  -output string
    	Path for output GIF (default "animated.gif")
  -v	Show version
  -y	Override Y/N prompts with Y response

```

## Generating new animation

```bash
[~/go-animate-coin/sample]$ ./go-animate-coin -input apario-xrp-coin@369px.png -output apario.gif
Animation saved to apario.gif
```

## Result

![APARIO Token Animation](apario.gif)






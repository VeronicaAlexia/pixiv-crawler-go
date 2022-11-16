<h1 align="center">
  <img src="./docs/logo.svg" alt="Pixiv_logo" width ="400">
  <br><a href="https://www.pixiv.net/">pixiv</a> crawler go<br>  
</h1> 

### pixiv login.

- authentication method is no longer supported to pixiv .
- The Pixiv app now logs in through `https://accounts.pixiv.net/login`
- but this page is protected by Google reCAPTCHA, which seems impossible to circumvent.
- so, you can't use this crawler to with login account,but you can use this crawler to web get the account token to
  login.

## start crawler with command line arguments

```
NAME:
   image downloader - download image from pixiv 

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   2.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -d value, --download value    Input IllustID to download illusts.
   -u value, --url value         Input pixiv url to download illusts
   -a value, --author value      Input AuthorID to download Author illusts. (default: 0)
   --user value, --userid value  Input user id to change default user id (default: 0)
   -m value, --max value         Input max thread number (default: 16)
   -f, --following               Download illusts from following users
   -r, --recommend               Download recommend illusts
   -s, --stars                   download stars illusts.
   --rk, --ranking               Download ranking illusts.
   -l, --login                   login pixiv account
   -n, --novel                   download novel
   --help, -h                    show help
   --version, -v                 print the version

```

## install pixiv crawler go
 

```Clone
 gh repo clone VeronicaAlexia/pixiv-crawler-go
```

## about command line arguments and usage

## NAME pixivlib

- **login account**
    - ``` -l / --login```
- **download image**
    - ```-d / --download <image_id> ```
- **download image**
    - ```-u / --url <url> ```
- **download novel**
    - ```-d / --download <novel_id> -n / --novel```
- **download author illustrations**
    - ``` -a / --author <author_id> ```
- **change the thread number**
    - ``` -m / --max ```
- **download collect illustrations**
    - ``` -s / --start ```
- **download recommend illustrations**
    - ``` -r / --recommend```
- **search illustrations**
    - ``` -s / --search <search_word> ```
- **ranking illustrations**
    - ``` -r / --rkaning ```
- **login account**
    - ``` -l / --login```  

| functions                                    | complete |
|----------------------------------------------|----------|
| download picture by image id                 | ✅        |
| command line                                 | ✅        |
| download picture by image name               | ✅        |
| download collect illustrations               | ✅        |
| download recommend illustrations             | ✅        |
| multi-threading                              | ✅        |
| asynchronous                                 | ❌        |
| browser automatically login pixiv on startup | ✅        |
| download illustrations by tag name           | ✅        |

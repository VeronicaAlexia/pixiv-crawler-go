
<h1 align="center">
  <img src="./docs/logo.svg" alt="Pixiv_logo" width ="400">
  <br>pixiv crawler for <a href="https://www.pixiv.net/">pixiv.net</a><br>  
</h1> 

### pixiv login.

- authentication method is no longer supported to pixiv .
- The Pixiv app now logs in through `https://accounts.pixiv.net/login`
- but this page is protected by Google reCAPTCHA, which seems impossible to circumvent.
- so, you can't use this crawler to with login account,but you can use this crawler to web get the account token to
  login. 

## start crawler with command line arguments

```
GLOBAL OPTIONS:
   -d value, --download value    input IllustID to download illusts
   -u value, --url value         input pixiv url to download illusts
   -a value, --author value      input author id to download illusts list
   --user value, --userid value  input user id to search user info
   -f, --following               download following author illusts
   -r, --recommend               download recommend illusts
   -s, --stars                   download stars illusts
   --rk, --ranking               download illusts from ranking
   --help, -h                    show args help message 
   --version, -v                 print the program version 

```

## install

``` pip install pixivlib ```

## about command line arguments and usage

## NAME pixivlib

- **login account**
    - ``` -l / --login```
- **download image**
    - ```-d / --download <image_id> ```
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
    - ``` -k / --rkaning ```
- **clear cache**
    - ``` -c / --clear_cache```

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

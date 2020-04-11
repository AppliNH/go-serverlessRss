# What is it

A golang REST API that you can query passing a RSS Feed URI, in order to read your news.

For now, it doesn't return JSON, it just returns text. I mainly use this app with curl, through a shell.

# How do I set it up

- First, pull it
- cd inside it
- Run `docker build -t [name] .`
- Run `docker run -d --env PORT=8080 -p 8080:8080 [name]`
- There you go, it should run on your localhost:8080

# How do I use it

## In order to simply read the news from a curl

If you query `localhost:8080/api/v1/rss/{YOUR_RSS_URI}` you should get the list of the articles, with articles' number.

Then, query `localhost:8080/api/v1/rss/{YOUR_RSS_URI}/item/{ARTICLE_NUMBER}` to "open" an article, which means getting access to its description as well as the URI link to read it in your Web Browser.

## In order to get the news in a JSON format

If you query `localhost:8080/api/json/rss` with the POST method, you can pass the following data :
```JSON
{
    "name": "FranceInfo",
    "uri": "https://www.francetvinfo.fr/monde.rss"
}
```
so you get in return the articles in a JSON array.
Each article contains the following properties : 
```JSON
{
        "id": Int,
        "title": String,
        "desc": String,
        "link": String
},
```


## Now deployed thanks to Heroku !

You can use it from here :
https://dry-tor-91544.herokuapp.com

# List of things that will be done 

- Deploy to Heroku `[done !]`
- Add a method so you can get JSON out of it `[done]`
- Reorganize the code because it's a mess
- Add a post method `[done !]`

# Why

Just wanted to practice go.

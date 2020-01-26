# What is it

A golang REST API that you can query passing a RSS Feed URI, in order to read your news.

For now, it doesn't return JSON, it just returns text. I mainly use this app with curl, through a shell.

# How do I set it up

- First, pull it
- cd inside it
- Run `docker build -t [name] .`
- Run `docker run -d -p 8080:8080 [name]`
- There you go, it should run on your localhost:8080

# How do I use it

## In order to simply read the news from a curl

If you query `localhost:8080/api/v1/rss/{YOUR_RSS_URI}` you should get the list of the articles, with articles' number.

Then, query `localhost:8080/api/v1/rss/{YOUR_RSS_URI}/item/{ARTICLE_NUMBER}` to "open" an article, which means getting access to its description as well as the URI link to read it in your Web Browser.

## In order to get the news in a JSON format

If you query `localhost:8080/api/json/rss/{YOUR_RSS_URI}` you should get a JSON array of the articles.
Each article contains the following properties : Id, Title, Description and Link.


# List of things that will be done 

- Deploy to Netlify `[in progress]`
- Add a method so you can get JSON out of it `[done]`
- Possibility to constitute a News Library

# Why

Just wanted to practice go.

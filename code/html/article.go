package html

const ArticleTemplate = `<!doctype html>
<html lang="en">
    <head>
        <title>` + BlogName + ` - {{Title}}</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="author" content="` + BlogName + `">
        <link
            rel="stylesheet"
            href="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/styles/atom-one-light.min.css"
            media="screen" />
          <link
            rel="stylesheet"
            href="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/styles/atom-one-dark.min.css"
            media="screen and (prefers-color-scheme: dark)" />
        <script src="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/highlight.min.js"></script>
        <script src="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/languages/clojure.min.js"></script>
        <script src="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.2.0/build/languages/clojure-repl.min.js"></script>
        <script>hljs.highlightAll();</script>
		` + CSS + `
    </head>

    <body>
        <nav>
            <a href="/">Home</a>
        </nav>

        <h1>{{Title}}</h1>

        <div>
            <h4>{{Date}}</h4>

            <div>

			{{Body}}

            <p><i>-` + BlogName + `</i></p>
            </div>
        </div>

        <br>
        <br>
    </body>
</html>
`

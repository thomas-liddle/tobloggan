package html

const ListingTemplate = `<!doctype html>
<html lang="en">
    <head>
        <title>Your Name Here</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="author" content="Your Name Here">
        <meta name="description" content="description">
        <link rel="canonical" href="https://your-domain-here.com">
        ` + CSS + `
    </head>

    <body>
        <nav>
            <a href="/about/">About</a>
        </nav>
        <h1>Your Name Here</h1>
        <p>Something about yourself and this website here.</p>
		<ul>
		{{Listing}}
		</ul>
        <br>
        <br>
    </body>
</html>
`

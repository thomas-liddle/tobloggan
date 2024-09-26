package html

const ListingTemplate = `<!doctype html>
<html lang="en">
    <head>
        <title>` + BlogName + `</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="author" content="` + BlogName + `">
        <meta name="description" content="` + BlogDescription + `">
        ` + CSS + `
    </head>

    <body>
        <h1>` + BlogName + `</h1>
        <p>` + BlogDescription + `</p>
		<ul>
		{{Listing}}
		</ul>
        <br>
        <br>
    </body>
</html>
`

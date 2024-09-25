# tobloggan

_Blogging made so easy you might as well be coasting smoothly down a snowy mountain track without a care in the world._

- Run `make test` to execute tests.
- Run `make dev` to generate the blog from the `/content` directory and open a web browser to view the pages in the `/generated` directory.
- Run `make publish` to generate the blog from the `/content` directory to the `/docs` directory. Will also commit and push the blog (to kickoff a build of github pages, if configured).

## Github Pages Configuration

![github-pages-config](./assets/github-pages-setup.png)


## Content Files

Here's the format of a content file:

```text
{
    "slug": "/the-path/of-the/article",
    "title": "The Title of the Article",
    "date": "2024-09-25T00:00:00Z"
}

+++

## Markdown content here
```

1. All content files must end in `.md`.
2. Content files should each have unique `"slug"` values.
3. Dates must be formatted as shown above.
4. The separator between the JSON object and the markdown content (`+++`) is REQUIRED.
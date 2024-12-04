# Welcome to the Markdown Server with HTMX

This is a simple server that renders Markdown files as HTML and uses HTMX for dynamic content loading.

## Features

- Automatically converts Markdown to HTML
- Uses Gin framework for routing
- Uses HTMX for dynamic content loading
- Simple and easy to extend

<button hx-post="/clicked" hx-swap="innerHTML">Click me</button>

# pitcher

**pitcher** - a presentation generator by `jonot-cyber`

[Try it Live!](https://jonot.me/pitcher)
[Github Link](https://github.com/jonot-cyber/pitcher)

# About

Pitcher is a (bad) program to create a presentation from a markdown file. Featuring:

- A limited subset of markdown!
  - Images
  - Links
  - Bold
  - Italics
  - Code
- The horror!

# The Horror

While the code as it is fits my needs just fine, it isn't the cleanest. There are a lot of hacks that rely on how permissive browsers are with HTML errors, and some other hacks just to get around doing work.

# Keybinds

`Space` - go forward a slide
`Left Arrow Key` - go forward a slide
`Right arrow key` - go back a slide

# Usage

Run `go build` to compile the project

`pitcher test.md > output.html`

`output.html` will now contain your presentation

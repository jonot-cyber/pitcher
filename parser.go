// Fun Fact! You don't actually _need_ to scroll down!
// Seriously, you won't get anything from it.
// It will just hurt










































package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

//go:embed template.html
var template string

var headerRegex = regexp.MustCompile("^# (.*)$")
var imageRegex = regexp.MustCompile(`!\[(.*)\]\((.*)\)`)
var linkRegex = regexp.MustCompile(`\[(.*)\]\((.*)\)`)
var boldRegex = regexp.MustCompile(`\*\*(.*)\*\*`)
var italicRegex = regexp.MustCompile(`_(.*)_|\*(.*)\*`)
var codeRegex = regexp.MustCompile("`(.*)`")
var listItemRegex = regexp.MustCompile(`^\s*- (.*)$`)

func Parse(markdown io.Reader) {
    var sections [][]string

    scanner := bufio.NewScanner(markdown)

    var section []string
    for scanner.Scan() {
        if strings.HasPrefix(scanner.Text(), "# ") && len(section) != 0 {
            sections = append(sections, section)
            section = make([]string, 0)
        }
        section = append(section, scanner.Text())
    }
    sections = append(sections, section)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var documentTitle string
    var b strings.Builder
    for i, section := range sections {
        output, title := parseSection(section, i)
        if i == 0 {
            documentTitle = title
        }
        b.WriteString(output)
    }
    fmt.Printf(template, documentTitle, b.String())
}

func parseSection(section []string, i int) (string, string) {
    title := parseTitle(section[0])
    return fmt.Sprintf(`<div id="%d" class="invisible"><h1>%s</h1><section>%s</section></div>`, i, title, parseRest(section[1:])), title
}

func parseRest(section []string) string {
    var b strings.Builder
    for _, line := range section {
        if len(line) == 0 {
            continue
        }
        isListItem := false
        line = imageRegex.ReplaceAllString(line, `<img src="$2" alt="$1"/>`)
        line = linkRegex.ReplaceAllString(line, `<a href="$2">$1</a>`)
        line = boldRegex.ReplaceAllString(line, `<strong>$1</strong>`)
        line = italicRegex.ReplaceAllString(line, `<em>$1</em>`)
        line = codeRegex.ReplaceAllString(line, `<pre>$1</pre>`)
        if listItemRegex.MatchString(line) {
            line = listItemRegex.ReplaceAllString(line, "<ul><li>$1</li></ul>") // I don't like this any more than you do.
            isListItem = true
        }
        b.WriteString(fmt.Sprintf("<span>%s</span>", line))
        if !isListItem {
            b.WriteString("<br>")
        }
    }
    return b.String()
}

func parseTitle(line string) string {
    if title := headerRegex.FindStringSubmatch(line); title != nil {
        return title[1]
    }
    return "ERR"
}

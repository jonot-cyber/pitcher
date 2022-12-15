package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var headerRegex = regexp.MustCompile("^# (.*)$")
var imageRegex = regexp.MustCompile(`!\[(.*)\]\((.*)\)`)
var linkRegex = regexp.MustCompile(`\[(.*)\]\((.*)\)`)
var boldRegex = regexp.MustCompile(`\*\*(.*)\*\*`)
var italicRegex = regexp.MustCompile(`_(.*)_|\*(.*)\*`)
var codeRegex = regexp.MustCompile("`(.*)`")
var listItemRegex = regexp.MustCompile(`^- (.*)$`)

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

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for i, section := range sections {
        output := parseSection(section, i)
        fmt.Println(output)
    }
}

func parseSection(section []string, i int) string {
    title := parseTitle(section[0])
    return fmt.Sprintf(`<div id="%d"><h1>%s</h1><section>%s</section></div>`, i, title, parseRest(section[1:]))
}

func parseRest(section []string) string {
    var b strings.Builder
    for _, line := range section {
        line = imageRegex.ReplaceAllString(line, `<img src="$2" alt="$1"/>`)
        line = linkRegex.ReplaceAllString(line, `<a href="$2">$1</a>`)
        line = listItemRegex.ReplaceAllString(line, `<li>$1</li>`)
        line = boldRegex.ReplaceAllString(line, `<strong>$1</strong>`)
        line = italicRegex.ReplaceAllString(line, `<em>$1</em>`)
        line = codeRegex.ReplaceAllString(line, `<pre>$1</pre>`)
        b.WriteString(fmt.Sprintf("<p>%s</p>", line))
    }
    return b.String()
}

func parseTitle(line string) string {
    if title := headerRegex.FindStringSubmatch(line); title != nil {
        return title[1]
    }
    return "ERR"
}

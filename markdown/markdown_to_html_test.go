package markdown_test

import (
	"testing"

	"github.com/gouniverse/base/markdown"
)

func TestMarkdownToHtml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "whitespace only",
			input:    "   \n\t   ",
			expected: "",
		},
		{
			name:     "simple paragraph",
			input:    "Hello world",
			expected: "<p>Hello world</p>\n",
		},
		{
			name:     "heading",
			input:    "# Heading 1",
			expected: "<h1 id=\"heading-1\">Heading 1</h1>\n",
		},
		{
			name:     "multiple paragraphs",
			input:    "First paragraph\n\nSecond paragraph",
			expected: "<p>First paragraph</p>\n<p>Second paragraph</p>\n",
		},
		{
			name:     "bold text",
			input:    "**bold text**",
			expected: "<p><strong>bold text</strong></p>\n",
		},
		{
			name:     "italic text",
			input:    "*italic text*",
			expected: "<p><em>italic text</em></p>\n",
		},
		{
			name:     "code block",
			input:    "```go\nfunc main() {}\n```",
			expected: "<pre><code class=\"language-go\">func main() {}\n</code></pre>\n",
		},
		{
			name:     "unordered list",
			input:    "- Item 1\n- Item 2",
			expected: "<ul>\n<li>Item 1</li>\n<li>Item 2</li>\n</ul>\n",
		},
		{
			name:     "ordered list",
			input:    "1. First\n2. Second",
			expected: "<ol>\n<li>First</li>\n<li>Second</li>\n</ol>\n",
		},
		{
			name:     "link",
			input:    "[Google](https://www.google.com)",
			expected: "<p><a href=\"https://www.google.com\">Google</a></p>\n",
		},
		{
			name:     "image",
			input:    "![alt text](image.jpg)",
			expected: "<p><img src=\"image.jpg\" alt=\"alt text\" /></p>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := markdown.MarkdownToHtml(tt.input)
			if result != tt.expected {
				t.Errorf("MarkdownToHtml() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestMarkdownToHtmlComplex(t *testing.T) {
	input := `# Main Title

## Section 1
This is a paragraph with **bold text** and *italic text*. You can also use ` + "`inline code`" + ` within text.

### Code Example
` + "```go" + `
func main() {
    fmt.Println("Hello, World!")
}
` + "```" + `

### Lists
1. First ordered item
2. Second ordered item
   - Nested unordered item
   - Another nested item
3. Third ordered item

### Links and Images
Here's a [link to Google](https://www.google.com) and an image:

![Example Image](https://example.com/image.jpg)

### Blockquote
> This is a blockquote
> With multiple lines
> And some **bold** text

### Table
| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |

### Horizontal Rule
---

Final paragraph with a [reference link][1].

[1]: https://example.com/reference`

	expected := `<h1 id="main-title">Main Title</h1>

<h2 id="section-1">Section 1</h2>
<p>This is a paragraph with <strong>bold text</strong> and <em>italic text</em>. You can also use <code>inline code</code> within text.</p>

<h3 id="code-example">Code Example</h3>
<pre><code class="language-go">func main() {
    fmt.Println("Hello, World!")
}
</code></pre>

<h3 id="lists">Lists</h3>
<ol>
<li>First ordered item</li>
<li>Second ordered item
<ul>
<li>Nested unordered item</li>
<li>Another nested item</li>
</ul>
</li>
<li>Third ordered item</li>
</ol>

<h3 id="links-and-images">Links and Images</h3>
<p>Here's a <a href="https://www.google.com">link to Google</a> and an image:</p>

<p><img src="https://example.com/image.jpg" alt="Example Image" /></p>

<h3 id="blockquote">Blockquote</h3>
<blockquote>
<p>This is a blockquote
With multiple lines
And some <strong>bold</strong> text</p>
</blockquote>

<h3 id="table">Table</h3>
<table>
<thead>
<tr>
<th>Header 1</th>
<th>Header 2</th>
</tr>
</thead>
<tbody>
<tr>
<td>Cell 1</td>
<td>Cell 2</td>
</tr>
<tr>
<td>Cell 3</td>
<td>Cell 4</td>
</tr>
</tbody>
</table>

<h3 id="horizontal-rule">Horizontal Rule</h3>
<hr />

<p>Final paragraph with a <a href="https://example.com/reference">reference link</a>.</p>
`

	result := markdown.MarkdownToHtml(input)
	if result != expected {
		t.Errorf("MarkdownToHtml() = %q, want %q", result, expected)
	}
}

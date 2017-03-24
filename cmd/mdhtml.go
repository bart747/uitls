// Copyright © 2017 Bartosz Wieczorek
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/russross/blackfriday"
	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "mdhtml",
	Short: "From a Markdown file make an HTML page with internal CSS styling for amazing readability.",
	Long: `Use Markdown file path as argument.
	It will create one '.html' file. CSS will be embedded inside it (internal CSS).`,
	Run: func(cmd *cobra.Command, args []string) {

		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		output := blackfriday.MarkdownBasic(file)
		l := len(output)

		var head = `<head><meta charset="utf-8"><title>` +
			args[0] + `</title>` + style + `</head>`

		outStr := "<!DOCTYPE html><html>" + head + "<body>" +
			string(output[:l]) + "</body></html>"

		newFile, err := os.Create(args[0] + ".html")
		if err != nil {
			panic(err)
		}
		defer newFile.Close()

		newFile.WriteString(outStr)
		if err != nil {
			panic(err)
		}

		fmt.Println("make called")

		if _, err := os.Stat(args[0] + ".html"); os.IsNotExist(err) {
			fmt.Println(err)
		}

		if _, err := os.Stat(args[0] + ".html"); err == nil {
			fmt.Println("HTML file created")
		}
	},
}

func init() {
	RootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

var style = `
<style>
html {
    box-sizing: border-box;
    font-size: 20px;
    line-height: 1.5;
}

*, *:before, *:after {
    box-sizing: inherit;
}

body {
    background-color: #fcfafa;
    color: #333;
    font-family: "Georgia", serif;
    font-size: 20px;
    max-width: 35em;
    margin: 2.5em auto;
    padding: 0 0.5em;
}

@media screen and (max-width: 568px) {
    body {
        font-size: 18px;
    }
}

audio,
canvas,
iframe,
img,
svg,
video {
    vertical-align: middle;
}

a {
    color: #2233dd;
    text-decoration: none;
    transition: color 0.1s ease-in;
}

a:active, a:hover {
    color: navy;
}

p {
    text-align: left;
    margin: 0 auto 1em;
}

h1,
h2,
h3,
h4,
h5 {
    font-family: "Roboto", sans-serif;
    color: #222;
    margin: 0 0 1em;
}

h1 {
    font-size: 2em;
    font-weight: bold;
    line-height: 1;
    margin: 1em 0;
    letter-spacing: -0.01em;
}

h2 {
    font-weight: bold;
    font-size: 1.5em;
    font-weight: 600;
    line-height: 1.25;
    color: #222;
    margin: 1.33em 0 0.66em;
}

h3 {
    font-size: 1.25em;
    font-weight: 600;
    line-height: 1.2;
    margin: 1.6em 0 0.8em
}

h4 {
    font-size: 1em;
    font-weight: 600;
    line-height: 1.5;
    margin: 2em 0 1em;
}

h5 {
    font-size: 1em;
    font-weight: normal;
    line-height: 1.5;
    margin: 2em 0 1em;
    text-transform: uppercase;
    letter-spacing: 0.02em;
}

strong {
    font-weight: bold;
}

blockquote {
    border-left: 1px dotted #777;
    margin: 0;
    padding-left: 1em;
}

pre {
    overflow-x: auto;
    margin: 1em 0;
}

pre,
code {
    font-size: 16px;
    font-family: monospace;
}
@media screen and (max-width: 768px) {
    pre,
    code {
        font-size: 14px;
    }
}

p code,
li code {
    color: #004010;
    font-size: 0.8em;
    font-family: monospace;
    padding: 0 0.25em;
}
</style>`

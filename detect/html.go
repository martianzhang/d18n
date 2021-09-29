/*
 * Copyright(c)  2021 Lianjia, Inc.  All Rights Reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package detect

import (
	"bufio"
	"os"

	"d18n/common"

	"golang.org/x/net/html"
)

func (d *DetectStruct) detectHTML() error {
	var err error

	fd, err := os.Open(d.CommonConfig.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := bufio.NewReaderSize(fd, d.CommonConfig.MaxBufferSize)
	token := html.NewTokenizer(r)

	var row []string
	for {

		t := token.Next()
		if t == html.ErrorToken {
			break
		}

		tag, _ := token.TagName()
		switch t {
		case html.StartTagToken:
			switch string(tag) {
			case "th", "td":
				token.Next()
				row = append(row, html.UnescapeString(string(token.Raw())))
			case "tr":
				d.Status.Lines++
			}
		case html.EndTagToken:
			switch string(tag) {
			case "tr":
				// check column names
				if d.Status.Lines == 1 {
					if !d.CommonConfig.NoHeader && d.CommonConfig.Schema == "" {
						for _, r := range row {
							d.Status.Header = append(d.Status.Header, common.HeaderColumn{Name: r})
						}
					}
					checkFileHeader(d.Status, d.Status.Header)

					// truncate row after new line
					row = []string{}

					if !d.CommonConfig.NoHeader {
						continue
					}
				}

				// check value
				for j, value := range row {
					d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], checkValue(value)...)
				}

				// truncate row after new line
				row = []string{}
			}
		}

		// SkipLines
		if d.Status.Lines <= d.CommonConfig.SkipLines {
			continue
		}
		if d.CommonConfig.Limit > 0 &&
			(d.Status.Lines-d.CommonConfig.SkipLines) > d.CommonConfig.Limit {
			break
		}

	}

	return err
}

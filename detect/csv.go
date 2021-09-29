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
	"encoding/csv"
	"io"
	"os"

	"d18n/common"
)

func (d *DetectStruct) detectCSV() error {

	var err error

	fd, err := os.Open(d.CommonConfig.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := csv.NewReader(fd)
	r.Comma = d.CommonConfig.Comma
	for {
		d.Status.Lines++

		// read row
		row, err := r.Read()
		if err == io.EOF { // end of file
			break
		} else if err != nil {
			return err
		}

		// check column names
		if d.Status.Lines == 1 {
			if !d.CommonConfig.NoHeader && d.CommonConfig.Schema == "" {
				for _, r := range row {
					d.Status.Header = append(d.Status.Header, common.HeaderColumn{Name: r})
				}
			}
			checkFileHeader(d.Status, d.Status.Header)
			if !d.CommonConfig.NoHeader {
				continue
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

		// check value
		for j, value := range row {
			d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], checkValue(value)...)
		}
	}

	return err
}

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
package mask

import (
	"fmt"
	"testing"

	"d18n/common"

	"github.com/kr/pretty"
)

func TestParseMaskConfig(t *testing.T) {
	orgMask := defaultMaskConfig

	file := common.TestPath + "/test/mask.csv"
	err := ParseMaskConfig(file)
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(defaultMaskConfig)

	defaultMaskConfig = orgMask
}

func TestGenRSAKey(t *testing.T) {
	privatekey, publicKey, err := genRSAKey()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(privatekey), string(publicKey))
}

func TestGenECCKey(t *testing.T) {
	privatekey, publicKey, err := genECCKey()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(privatekey), string(publicKey))
}
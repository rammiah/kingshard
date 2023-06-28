// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package hack

import (
	"strconv"
	"strings"
	"unsafe"
)

// no copy to change slice to string
// use your own risk
//
//go:inline
func String(b []byte) (s string) {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// no copy to change string to slice
// use your own risk
//
//go:inline
func Slice(s string) (b []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func IsSqlSep(r rune) bool {
	return r == ' ' || r == ',' ||
		r == '\t' || r == '/' ||
		r == '\n' || r == '\r'
}

func ArrayToString(array []int) string {
	if len(array) == 0 {
		return ""
	}
	var strArray []string
	for _, v := range array {
		strArray = append(strArray, strconv.FormatInt(int64(v), 10))
	}

	return strings.Join(strArray, ", ")
}

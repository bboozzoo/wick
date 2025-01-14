/*
*
* Copyright 2021-2022 Simple Things Inc.
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*
 */

package core

import (
	"strings"
	"testing"
)

const (
	privateKeyHex = "b99067e6e271ae300f3f5d9809fa09288e96f2bcef8dd54b7aabeb4e579d37ef"
)

func TestPrivateHexToKeyPair(t *testing.T) {
	publicKey, privateKey := getKeyPair(privateKeyHex)

	if publicKey == nil {
		t.Errorf("public key is nil")
	}

	if privateKey == nil {
		t.Errorf("private key is nil")
	}
}

func TestUrlSanitization(t *testing.T) {
	url := sanitizeURL("rs://localhost:8080/")
	if !strings.HasPrefix(url, "tcp") {
		t.Error("url sanitization failed")
	}

	url = sanitizeURL("rss://localhost:8080/")
	if !strings.HasPrefix(url, "tcp") {
		t.Error("url sanitization failed")
	}
}

func TestListToWampList(t *testing.T) {
	inputList := []string{"string", "1", "1.1", "true", "[\"group_1\",\"group_2\", 1, true]",
		"{\"firstKey\":\"value\", \"secondKey\":2}",
		"[{\"firstKey\":\"value\", \"secondKey\":2}, {\"firstKey\":\"value\", \"secondKey\":2}]"}
	wampList := listToWampList(inputList)

	if len(wampList) != len(inputList) {
		t.Error("error in list conversion")
	}

	if _, canConvert := wampList[1].(int); canConvert == false {
		t.Error("error in list conversion")
	}

	if _, canConvert := wampList[2].(float64); canConvert == false {
		t.Error("error in list conversion")
	}

	if _, canConvert := wampList[3].(bool); canConvert == false {
		t.Error("error in list conversion")
	}

	if _, canConvert := wampList[4].([]interface{}); canConvert == false {
		t.Error("error in list conversion")
	}

	if _, canConvert := wampList[5].(map[string]interface{}); canConvert == false {
		t.Error("error in list conversion")
	}

	if _, canConvert := wampList[6].([]map[string]interface{}); canConvert == false {
		t.Error("error in list conversion")
	}
}

func TestDictToWampDict(t *testing.T) {
	inputDict := map[string]string{"string": "string", "int": "1", "float": "1.1", "bool": "true",
		"list": "[\"group_1\",\"group_2\", 1, true]", "json": "{\"firstKey\":\"value\", \"secondKey\":2}",
		"jsonList": "[{\"firstKey\":\"value\", \"secondKey\":2}, {\"firstKey\":\"value\", \"secondKey\":2}]"}
	wampDict := dictToWampDict(inputDict)
	if len(inputDict) != len(wampDict) {
		t.Error("error in map conversion")
	}

	if _, canConvert := wampDict["int"].(int); canConvert == false {
		t.Error("error in map conversion")
	}

	if _, canConvert := wampDict["float"].(float64); canConvert == false {
		t.Error("error in map conversion")
	}

	if _, canConvert := wampDict["bool"].(bool); canConvert == false {
		t.Error("error in map conversion")
	}

	if _, canConvert := wampDict["list"].([]interface{}); canConvert == false {
		t.Error("error in map conversion")
	}

	if _, canConvert := wampDict["json"].(map[string]interface{}); canConvert == false {
		t.Error("error in map conversion")
	}

	if _, canConvert := wampDict["jsonList"].([]map[string]interface{}); canConvert == false {
		t.Error("error in dict conversion")
	}
}

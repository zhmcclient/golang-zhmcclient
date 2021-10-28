/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

import (
	"errors"
	"net/url"
)

func BuildUrlFromUri(uri string, query map[string]string) (*url.URL, error) {
	url, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := url.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		url.RawQuery = q.Encode()
	}
	return url, nil
}

func GenerateErrorFromResponse(status int, responseBodyStream []byte) error {
	// TODO, generate error messages
	return errors.New("Unknown error")
}

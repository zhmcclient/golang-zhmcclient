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
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type ErrorBody struct {
	Reason  int    `json:"reason"`
	Message string `json:"message"`
}

func BuildUrlFromQuery(url *url.URL, query map[string]string) (*url.URL, error) {
	if query != nil {
		q := url.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		url.RawQuery = q.Encode()
	}
	return url, nil
}

func RetrieveBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)
	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	return bytes, err
}

func GenerateErrorFromResponse(status int, responseBodyStream []byte) error {
	errBody := &ErrorBody{}
	_ = json.Unmarshal(responseBodyStream, errBody)

	if errBody.Message != "" {
		return errors.New(errBody.Message)
	}

	switch status {
	case http.StatusBadRequest: // 400
		switch errBody.Reason {
		case 1:
			return errors.New("The request included an unrecognized or unsupported query parameter.")
		case 2:
			return errors.New("A required request header is missing or invalid.")
		case 3:
			return errors.New("A required request body is missing.")
		case 4:
			return errors.New("A request body was specified when not expected.")
		case 5:
			return errors.New("A required request body field is missing.")
		case 6:
			return errors.New("The request body contains an unrecognized field (i.e. one that is not listed as either required or optional in the specification for the request body format for the operation).")
		case 7:
			return errors.New("The data type of a field in the request body is not as expected, or its value is not in the range permitted.")
		case 8:
			return errors.New("The value of a field does not provide a unique value for the corresponding data model property as required.")
		case 9:
			return errors.New("The request body is not a well-formed JSON document.")
		case 10:
			return errors.New("An unrecognized X-* header field was specified.")
		case 11:
			return errors.New("The length of the supplied request body does not match the value specified in the Content-Length header.")
		case 13:
			return errors.New("The maximum number of logged in user sessions for this user ID has been reached; no more are allowed.")
		case 14:
			return errors.New("Query parameters on the request are malformed or specify a value that is invalid for this operation. Common causes include the inability to successfully decode a parameter element, the presented parameters are not in the expected key=value format, the value is not a valid regular expression, a required parameter is missing, multiple instances of a parameter are present on an operation that does not permit multiple instances of that parameter, or the value is not a valid enum for the operation.")
		case 15:
			return errors.New("The request body contains a field whose presence or value is inconsistent with the presence or value of another field in the request body. A prerequisite condition or dependency among request body fields is not met.")
		case 18:
			return errors.New("The request body contains a field whose presence or value is inconsistent with the type of the object. Such a requirement is often described in an object's data model as the field having a prerequisite condition on a 'type', 'family', or similar property that identifies an object as being of a particular type. Such a property is typically, but not necessarily, immutable.")
		case 19:
			return errors.New("The request body contains a field whose corresponding data model property is no longer writable. In certain earlier HMC and/or SE versions the property is writable, but later versions do not allow changing the property through the Web Services APIs. This could be due to a change in the underlying system-management model, or the property may have become obsolete.")
		case 20:
			return errors.New("The request body contains a field or value that is directly or indirectly dependent on the version of the SE that is targeted by or associated with the request operation, and that SE is not at a version that supports or provides the field or value.")
		default:
			return errors.New("Bad Request.")

		}
	case http.StatusForbidden: // 403
		switch errBody.Reason {
		case 1:
			return errors.New("The user under which the API request was authenticated does not have the required authority to perform the requested action.")
		case 3:
			return errors.New("The ensemble is not operating at the management enablement level required to perform this operation.")
		case 4:
			return errors.New("The request requires authentication but no X-API-Session header was specified in the request.")
		case 5:
			return errors.New("An X-API-Session header was provided but the session id specified in that header is not valid.")
		case 301:
			return errors.New("The operation cannot be performed because it targets a CPC that does not support Web Services API operations.")
		default:
			return errors.New("The API user does not have the required permission for this operation.")

		}
	case http.StatusNotFound: // 404
		switch errBody.Reason {
		case 1:
			return errors.New("The request URI does not designate an existing resource of the expected type, or designates a resource for which the API user does not have object-access permission. For URIs that contain object ID and/or element ID components, this reason code may be used for issues accessing the resource identified by the first (leftmost) such ID in the URI.")
		case 2:
			return errors.New("A URI in the request body does not designate an existing resource of the expected type, or designates a resource for which the API user does not have object-access permission. For URIs that contain object ID and/or element ID components, this reason code may be used for issues accessing the resource identified by the first (leftmost) such ID in the URI.")
		case 3:
			return errors.New("The request URI designates a resource or operation that is not available on the Alternate HMC.")
		case 4:
			return errors.New("The object designated by the request URI does not support the requested operation.")
		case 5:
			return errors.New("The request URI does not designate an existing resource of the expected type, or designates a resource for which the API user does not have object-access permission. More specifically, this reason code indicates issues accessing the resource identified by the element ID component in the URI. Such an element ID is typically the second (counting left to right) ID component in the URI.")
		case 6:
			return errors.New("A URI in the request body does not designate an existing resource of the expected type, or designates a resource for which the API user does not have object-access permission. More specifically this reason code indicates issues accessing the resource identified by the element ID component in the URI. Such an element ID is typically the second (counting left to right) ID component in the URI.")
		default:
			return errors.New("Not Found")
		}
	case http.StatusConflict: // 409
		switch errBody.Reason {
		case 1:
			return errors.New("The operation cannot be performed because the object designated by the request URI is not in the correct state.")
		case 2:
			return errors.New("The operation cannot be performed because the object designated by the request URI is currently busy performing some other operation.")
		case 3:
			return errors.New("The operation cannot be performed because the object designated by the request URI is currently locked to prevent disruptive changes from being made.")
		case 4:
			return errors.New("The operation cannot be performed because the CPC designated by the request URI is currently enabled for DPM.")
		case 5:
			return errors.New("The operation cannot be performed because the CPC designated by the request URI is currently not enabled for DPM.")
		case 6:
			return errors.New("The operation cannot be performed because the object hosting the object designated by the request URI is not in the correct state.")
		case 8:
			return errors.New("The operation cannot be performed because the request would result in the object being placed into a state that is inconsistent with its data model or other requirements. The request body contains a field whose presence or value is inconsistent with the current state of the object or some aspect of the system, and thus a prerequisite condition or dependency would no longer be met.")
		case 9:
			return errors.New("The operation cannot be completed because it is attempting to update an effective property when the object is not in a state in which effective properties are applicable. More specifically, the request body contains one or more fields which correspond to a property marked with the (e) qualifier in the data model, and the object's effective-properties-apply property is false.")
		case 10:
			return errors.New("The operation cannot be performed because the affected SE is in the process of being shut down.")
		case 11:
			return errors.New("The operation cannot be performed because it requires a fully authenticated session and the API session that issued it is only partially authenticated.")
		case 12:
			return errors.New("The operation cannot be performed because a feature that prohibits the operation is currently enabled. The error-details field of the response body contains an error-feature-info object identifying the feature whose current enablement status is invalid for the operation. The error-feature-info object is described in the next table.")
		case 13:
			return errors.New("The operation cannot be performed because a feature required by the operation is currently disabled. The error-details field of the response body contains an error-feature-info object identifying the feature whose current enablement status is invalid for the operation. The error-feature-info object is described in the next table.")
		default:
			return errors.New("Conflict")
		}
	/* 50x */
	case http.StatusInternalServerError: // 500
		return errors.New("An internal processing error has occurred and no additional details are documented.")
	case http.StatusServiceUnavailable: // 503
		switch errBody.Reason {
		case 1:
			return errors.New("The request could not be processed because the HMC is not currently communicating with an SE needed to perform the requested operation.")
		case 2:
			return errors.New("The request could not be processed because the SE is not currently communicating with an element of a zBX needed to perform the requested operation.")
		case 3:
			return errors.New("This request would exceed the limit on the number of concurrent API requests allowed.")
		default:
			return errors.New("The request could not be processed because the HMC is not currently communicating with an SE needed to perform the requested operation.")
		}
	}
	return errors.New("HTTP Error: " + fmt.Sprint(status))
}

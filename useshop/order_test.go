package useshop

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	str := "QxpQB1QUA0MRRA1D1L31hYifgI+z29+0gqSq3PuKhuqCGh1BUw8SSUBcFYS9pobbnhFJFVJbFlYQSk4XWRsgKBoUEwdfCwdZDEQNQ1lFFhFBCTkYbRsXXRdMBwVWFw4fTV1UEFgJFh4BCVpDHRMHDFNaCRULFhddF0wHBVZ5EhcWW14OEkpEVgsURBV/UA8EEAlHZFlRD1wLVhUZQV4MCVxLdQZEBw9cQFxsGhNBEA5WRgZDf1UOXUYCFVYGSgsPChodQUETB14WD0MYEwtTTRBADkITDkEaSBpHRwpaBkQCCQFTTTtKEgsWFVsTAFtTHANLBx8EQRRGUUR4DFsKCl0aC1McRApRERJ5AFxUQFsQfwxWX1NBFEZXRVEGSyILV01fFxJcRAFSVRlRARNOQ11BAVJDdxZKFl1ZVhobWURta3VBHEQJQgYDRS9eE1hDCgtdDxMYQVcWXFJHN1YXB1QaCxgSDxJVDzJYFVBdQFsDA1UbE0ACQEYCBxlBSgsPSEhYDVdEXABORF8AX1UOCFxURw0CGEFcDUtUWhZXF0QCCExPEhYOXwwDFVsTAFBSAQBXBQUGUgpGFBVGF1gXAxoCE4a88YKKzkQbQ0VYDwRBRwRaQRZZCVIKDgRaC1NQDxQTFkIKJVEOClUAUlpAWxBbEUNBR1lkS2QYQQZKF1YIDR8OSRMDVREOWBEfUg0MbhwVVkhZBlYQZBhbDE0KAEEVRBFcOkkSTkRCE11yDQxCXwBDVBZZGgxMQ0UQAz9JZBdFBkMSVgBXSFoYRFQHElpcFRlSWw5kS0hWTA5cDRJkF0MGRBMUXk8TRQ1tHkBNEEkMRxMOQQBcAA8NQUQ="

	key := "uP84RLiamc"
	md5key := GetSha512(key)
	err, res := StrDecrypt(str, md5key)
	fmt.Println(err, "\n", res)
}

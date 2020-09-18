package filehash

import (
	"crypto/md5"
	"io"
	"reflect"
)

func ReaderHasMD5(r io.Reader, md5Hash [md5.Size]byte) (bool, error) {
	h := md5.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(h.Sum(nil), md5Hash[:]), nil
}

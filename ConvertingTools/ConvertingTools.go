package ConvertingTools

import "bytes"
import "encoding/binary"

func ConvertUint32ToBytesLittleEndian(val uint32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, val)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
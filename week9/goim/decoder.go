package goim

import (
	"bufio"
	"errors"
	"io"
)

/*
	该函数只负责解码，不负责处理开始，结束。如果读到EOF则作为error返回由外部处理。
*/
func Decode(r io.Reader, handle func(p Package)) error {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	for {
		p := Package{}
		err := decodeLength(&p, br)
		if err != nil {
			return err
		}
		err = decodeHeader(&p, br)
		if err != nil {
			return err
		}
		p.Body = make([]byte, p.Length-uint(p.HeaderLength))
		n, err := br.Read(p.Body)
		if err != nil {
			return err
		}
		if n < len(p.Body) {
			return errors.New("can't read enough data")
		}
		go handle(p)
	}
}

type Operation uint

const (
	Auth      Operation = 1
	HeartBeat Operation = 2
	Message   Operation = 3
)

type Package struct {
	Length       uint
	HeaderLength uint16
	Version      uint16
	Operation    Operation
	SequenceID   uint
	Body         []byte
}

func decodeLength(p *Package, r io.Reader) error {
	lengths := make([]byte, 6)
	n, err := r.Read(lengths)
	if err != nil {
		return err
	}
	if n < len(lengths) {
		return errors.New("can't read enough data")
	}
	p.Length = decodeUint(lengths)
	p.HeaderLength = decodeUint16(lengths[4:])
	return nil
}

func decodeHeader(p *Package, r io.Reader) error {
	header := make([]byte, p.HeaderLength-6) // 去掉6位表示长度的
	n, err := r.Read(header)
	if err != nil {
		return err
	}
	if n < len(header) {
		return errors.New("can't read enough data")
	}
	p.Version = decodeUint16(header)
	p.Operation = Operation(decodeUint(header[2:]))
	p.SequenceID = decodeUint(header[6:])
	return nil
}

func decodeUint(b []byte) uint {
	var v uint = 0
	for i := 0; i < 4; i++ {
		v = v<<8 + uint(b[i])
	}
	return v
}

func decodeUint16(b []byte) uint16 {
	var v uint16 = 0
	for i := 0; i < 2; i++ {
		v = v<<8 + uint16(b[i])
	}
	return v
}

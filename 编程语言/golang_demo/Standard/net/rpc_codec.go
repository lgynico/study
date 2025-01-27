package net

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/rpc"
)

type JsonServerCodec struct {
	rwc    io.ReadWriteCloser
	dec    *json.Decoder
	enc    *json.Encoder
	encBuf *bufio.Writer
	closed bool
}

func NewJsonRpcCodec(conn io.ReadWriteCloser) *JsonServerCodec {
	buf := bufio.NewWriter(conn)
	return &JsonServerCodec{
		rwc:    conn,
		dec:    json.NewDecoder(conn),
		enc:    json.NewEncoder(buf),
		encBuf: buf,
	}
}

func (c *JsonServerCodec) ReadRequestHeader(r *rpc.Request) error {
	return c.dec.Decode(r)
}

func (c *JsonServerCodec) ReadRequestBody(body any) error {
	return c.dec.Decode(body)
}

func (c *JsonServerCodec) WriteResponse(r *rpc.Response, body any) (err error) {
	if err = c.enc.Encode(r); err != nil {
		if c.encBuf.Flush() == nil {
			// Gob couldn't encode the header. Should not happen, so if it does,
			// shut down the connection to signal that the connection is broken.
			log.Println("rpc: gob error encoding response:", err)
			c.Close()
		}
		return
	}
	if err = c.enc.Encode(body); err != nil {
		if c.encBuf.Flush() == nil {
			// Was a gob problem encoding the body but the header has been written.
			// Shut down the connection to signal that the connection is broken.
			log.Println("rpc: gob error encoding body:", err)
			c.Close()
		}
		return
	}
	return c.encBuf.Flush()
}

func (c *JsonServerCodec) Close() error {
	if c.closed {
		// Only call c.rwc.Close once; otherwise the semantics are undefined.
		return nil
	}
	c.closed = true
	return c.rwc.Close()
}

type JsonClientCodec struct {
	rwc    io.ReadWriteCloser
	dec    *json.Decoder
	enc    *json.Encoder
	encBuf *bufio.Writer
}

func NewJsonClientCodec(conn io.ReadWriteCloser) *JsonClientCodec {
	encBuf := bufio.NewWriter(conn)
	return &JsonClientCodec{
		rwc:    conn,
		dec:    json.NewDecoder(conn),
		enc:    json.NewEncoder(encBuf),
		encBuf: encBuf,
	}
}

func (c *JsonClientCodec) WriteRequest(r *rpc.Request, body any) (err error) {
	if err = c.enc.Encode(r); err != nil {
		return
	}
	if err = c.enc.Encode(body); err != nil {
		return
	}
	return c.encBuf.Flush()
}

func (c *JsonClientCodec) ReadResponseHeader(r *rpc.Response) error {
	return c.dec.Decode(r)
}

func (c *JsonClientCodec) ReadResponseBody(body any) error {
	return c.dec.Decode(body)
}

func (c *JsonClientCodec) Close() error {
	return c.rwc.Close()
}

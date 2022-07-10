###  问题一

总结几种 socket 粘包的解包方式: fix length, delimiter based, length field based frame decoder。尝试举例其应用

发送数据过程：应用程序发送消息包，消息包以数据流的形式放入缓冲区，等缓冲区的数据流到达一定阈值后，再发送到网络上

接受数据过程：接受到网络过来的数据流，放入缓冲区，缓冲区的数据流到达一定阈值后，通知应用程序进行读取数据

在数据发送和接受的过程中，都是对数据流进行操作，而数据流本身没有任何开始和结束的边界。因此正确地解析出数据包，就要知道数据在流中的开始和结束位置。

#### fix length frame decoder

数据发送方每次发送固定长度的数据，且不超出缓冲区，接收方获取同样长度的数据来解码拼成一个数据包。

#### delimiter based frame decoder

数据发送方在数据中添加特殊的分隔符来标记边界，接收方读到分隔符时解码拼成一个数据包。

#### length field based frame decoder

数据发送方在消息包头添加长度信息，接收方获取指定长度的数据来解码拼成一个数据包。

###  

### 问题二

实现一个从 socket connection 中解码出 goim 协议的解码器。

```
func Decode(buf []byte) (Idecoder, error) {
	decoder := &Decoder{}

	decoder.packetLen = binary.BigEndian.Uint32(buf[_packOffset : _packOffset+_packSize])
	binary.BigEndian.Uint16(buf[_headerOffset : _headerOffset+_headerSize])
	decoder.version = binary.BigEndian.Uint16(buf[_verOffset : _verOffset+_verSize])
	decoder.operation = binary.BigEndian.Uint32(buf[_opOffset : _opOffset+_opSize])
	decoder.sequence = binary.BigEndian.Uint32(buf[_seqOffset : _seqOffset+_seqSize])

	if decoder.packetLen > _maxPackSize {
		return nil, errors.New("Error Packet Length ")
	}

	if _bodyLen := int(decoder.packetLen - uint32(decoder.headerLen)); _bodyLen > 0 {
		decoder.body = buf[_bodyOffset : _bodyOffset+_bodyLen]
	}

	return decoder, nil
}
```



```
type Idecoder interface {
	PacketLen() uint32
	HeaderLen() uint16
	Version() uint16
	Operation() uint32
	Sequence() uint32
	Body() []byte
}
```



```
const (
	// MaxBodySize max body size
	MaxBodySize = uint32(1 << 12) // 4096
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + uint32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_bodyOffset   = _seqOffset + _seqSize
)
```



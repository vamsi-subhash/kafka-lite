package service

const (
	API_KEY_PRODUCE             = 0
	API_KEY_FETCH               = 1
	API_KEY_METADATA            = 3
	API_KEY_UPDATE_METADATA_KEY = 6
)

const (
	ERROR                  = 1
	METADATA_REQUEST_ERROR = 2
)

type Request struct {
	apiKey         int16
	apiVersion     int16
	correlationId  int32
	clientId       string
	requestMessage []byte
}

type Response struct {
	correlationId   int32
	responseMessage []byte
}

type Error struct {
	Code int
	Msg  string
}

type Message struct {
	Crc        int32
	MagicByte  int8
	Attributes int8
	Key        []byte
	Value      []byte
}

type MessageSet struct {
	Offset      int64
	MessageSize int32
	Message     []Message
}

type Node struct {
	Id   int
	Host string
	Port string
}

type Topic struct {
	Name           string
	NoOfPartitions int
}

type Partition struct {
	Id   int
	Desc string
}

type TopicPartition struct {
	Topic        Topic
	Partition    Partition
	LeaderNode   Node
	ReplicaNodes []Node
}

type TopicMetadata struct {
	TopicName       string
	TopicPartitions []TopicPartition
}

type MetadataRequest struct {
	TopicNames []string
}

type MetadataResponse struct {
	Metadata map[string]TopicMetadata
}
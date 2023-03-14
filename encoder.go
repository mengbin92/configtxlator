package configtxlator

import (
	"os"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-config/protolator"
	"github.com/pkg/errors"
)

func EncodeProto(msgName string, input, output *os.File) error {
	msgType := proto.MessageType(msgName)
	if msgType == nil {
		return errors.Errorf("message of type %s unknown", msgType)
	}
	msg := reflect.New(msgType.Elem()).Interface().(proto.Message)

	err := protolator.DeepUnmarshalJSON(input, msg)
	if err != nil {
		return errors.Wrapf(err, "error decoding input")
	}

	out, err := proto.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "error marshaling")
	}

	_, err = output.Write(out)
	if err != nil {
		return errors.Wrapf(err, "error writing output")
	}

	return nil
}

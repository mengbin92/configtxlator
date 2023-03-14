package configtxlator

import (
	"io/ioutil"
	"os"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-config/protolator"
	"github.com/pkg/errors"
)

func DecodeProto(msgName string, input, output *os.File) error {
	msgType := proto.MessageType(msgName)
	if msgType == nil {
		return errors.Errorf("message of type %s unknown", msgType)
	}
	msg := reflect.New(msgType.Elem()).Interface().(proto.Message)

	in, err := ioutil.ReadAll(input)
	if err != nil {
		return errors.Wrapf(err, "error reading input")
	}

	err = proto.Unmarshal(in, msg)
	if err != nil {
		return errors.Wrapf(err, "error unmarshalling")
	}

	err = protolator.DeepMarshalJSON(output, msg)
	if err != nil {
		return errors.Wrapf(err, "error encoding output")
	}
	return nil
}

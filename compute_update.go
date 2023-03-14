package configtxlator

import (
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/pkg/errors"
)

func ComputeUpdt(original, updated, output *os.File, channelID string) error {
	origIn, err := ioutil.ReadAll(original)
	if err != nil {
		return errors.Wrapf(err, "error reading original config")
	}

	origConf := &cb.Config{}
	err = proto.Unmarshal(origIn, origConf)
	if err != nil {
		return errors.Wrapf(err, "error unmarshalling original config")
	}

	updtIn, err := ioutil.ReadAll(updated)
	if err != nil {
		return errors.Wrapf(err, "error reading updated config")
	}

	updtConf := &cb.Config{}
	err = proto.Unmarshal(updtIn, updtConf)
	if err != nil {
		return errors.Wrapf(err, "error unmarshalling updated config")
	}

	cu, err := Compute(origConf, updtConf)
	if err != nil {
		return errors.Wrapf(err, "error computing config update")
	}

	cu.ChannelId = channelID

	outBytes, err := proto.Marshal(cu)
	if err != nil {
		return errors.Wrapf(err, "error marshaling computed config update")
	}

	_, err = output.Write(outBytes)
	if err != nil {
		return errors.Wrapf(err, "error writing config update to output")
	}

	return nil
}

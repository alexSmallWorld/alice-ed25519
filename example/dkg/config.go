// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dkg

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/edwards"
	"github.com/getamis/alice/crypto/tss/dkg"
	"github.com/getamis/alice/example/config"
	"github.com/getamis/sirius/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type DKGConfig struct {
	Port      int64   `yaml:"port"`
	Rank      uint32  `yaml:"rank"`
	Threshold uint32  `yaml:"threshold"`
	Peers     []int64 `yaml:"peers"`
}

type DKGResult struct {
	Share  string               `yaml:"share"`
	Pubkey config.Pubkey        `yaml:"pubkey"`
	BKs    map[string]config.BK `yaml:"bks"`
}

func readDKGConfigFile(filaPath string) (*DKGConfig, error) {
	c := &DKGConfig{}
	yamlFile, err := ioutil.ReadFile(filaPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func writeDKGResult(id string, result *dkg.Result) error {
	dkgResult := &DKGResult{
		Share: result.Share.String(),
		Pubkey: config.Pubkey{
			X: result.PublicKey.GetX().String(),
			Y: result.PublicKey.GetY().String(),
		},
		BKs: make(map[string]config.BK),
	}
	for peerID, bk := range result.Bks {
		dkgResult.BKs[peerID] = config.BK{
			X:    bk.GetX().String(),
			Rank: bk.GetRank(),
		}
	}
	err := config.WriteYamlFile(dkgResult, getFilePath(id))
	if err != nil {
		log.Error("Cannot write YAML file", "err", err)
		return err
	}
	// Build public key.
	pubkey := edwards.NewPublicKey(edwards.Edwards(), result.PublicKey.GetX(), result.PublicKey.GetY())
	log.Info("dkg result", "publickey", hex.EncodeToString(pubkey.Serialize()), "Share", result.Share.String())
	writePublicKey2File(hex.EncodeToString(pubkey.Serialize()), id)
	return nil
}

func getFilePath(id string) string {
	return fmt.Sprintf("dkg/%s-ed25519-output.yaml", id)
}

func writePublicKey2File(str, id string) {
	userFile := fmt.Sprintf("dkg/publicKey-%s.txt", id)
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		log.Info("writePublicKey2File failed", userFile, err)
		return
	}
	fout.WriteString(str + "\n")
	log.Info("writePublicKey2File success", userFile, str)
}

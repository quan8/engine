/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package mem

import (
	"sync"

	"errors"

	"github.com/it-chain/engine/consensus/pbft"
)

var ConsensusAlreadyExistError = errors.New("Consensus Already Exist")
var EmptyConsensusIdError = errors.New("empty Consensus Id")
var LoadConsensusError = errors.New("There is no consensus for loading")

type ConsensusRepository struct {
	consensus *pbft.Consensus
	sync.RWMutex
}

func NewConsensusRepository() ConsensusRepository {
	return ConsensusRepository{
		consensus: nil,
		RWMutex:   sync.RWMutex{},
	}
}
func (repo *ConsensusRepository) Save(consensus pbft.Consensus) error {

	repo.Lock()
	defer repo.Unlock()

	if repo.consensus != nil {
		return ConsensusAlreadyExistError
	}
	repo.consensus = &consensus

	return nil
}
func (repo *ConsensusRepository) Load() (*pbft.Consensus, error) {

	if repo.consensus == nil {
		return nil, LoadConsensusError
	}

	return repo.consensus, nil
}

func (repo *ConsensusRepository) Remove() {

	if repo.consensus != nil {
		repo.consensus = nil
	}
}
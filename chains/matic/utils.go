/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package matic

import (
	"github.com/polynetwork/bridge-common/chains/matic/cosmos"
	"github.com/polynetwork/poly/common"
)

func ParseEpochSwitch(data []byte) (info *cosmos.CosmosEpochSwitchInfo, err error) {
	if len(data) == 0 {
		return
	}
	info = new(cosmos.CosmosEpochSwitchInfo)
	err = info.Deserialization(common.NewZeroCopySource(data))
	if err != nil {
		return nil, err
	}
	return
}
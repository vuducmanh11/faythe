// Copyright (c) 2019 Kien Nguyen-Tuan <kiennt2609@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"crypto"
	"time"

	"github.com/pkg/errors"

	"github.com/vCloud-DFTBA/faythe/pkg/common"
)

// Scaler represents a Scaler object
type Scaler struct {
	Query       string                 `json:"query"`
	Duration    string                 `json:"duration"`
	Description string                 `json:"description,omitempty"`
	Interval    string                 `json:"interval"`
	Actions     map[string]*ActionHTTP `json:"actions"`
	Tags        []string               `json:"tags"`
	Active      bool                   `json:"active"`
	ID          string                 `json:"id,omitempty"`
	Alert       *Alert                 `json:"alert,omitempty"`
	Cooldown    string                 `json:"cooldown"`
}

// Validate returns nil if all fields of the Scaler have valid values.
func (s *Scaler) Validate() error {
	for _, a := range s.Actions {
		if err := a.Validate(); err != nil {
			return err
		}
	}

	if s.Query == "" {
		return errors.Errorf("required field %+v is missing or invalid", s.Query)
	}

	if _, err := time.ParseDuration(s.Duration); err != nil {
		return err
	}

	if _, err := time.ParseDuration(s.Interval); err != nil {
		return err
	}

	if s.Cooldown == "" {
		s.Cooldown = "600s"
	}
	if _, err := time.ParseDuration(s.Cooldown); err != nil {
		return err
	}

	s.ID = common.Hash(s.Query, crypto.MD5)

	return nil
}

// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actionLogService

import (
	"context"
	_ "embed"

	"github.com/ServiceWeaver/weaver"
	dapr "github.com/dapr/go-sdk/client"
)

var (
	// set the environment as instructions.
	pubsubName = "actionLog"
	topicName  = "cart"
)

type ActionLogService interface {
	ConvertAndSend(ctx context.Context) error
}

type impl struct {
	weaver.Implements[ActionLogService]
	dclient dapr.Client
}

func (s *impl) Init(context.Context) error {
	client, err := dapr.NewClient()
	s.dclient = client
	return err
}

func (s *impl) ConvertAndSend(ctx context.Context) error {
	publishEventsData := []interface{}{"multi-ping", "multi-pong"}
	publish := s.dclient.PublishEvents(ctx, pubsubName, topicName, publishEventsData)
	return publish.Error
}

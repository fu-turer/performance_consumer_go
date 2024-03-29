// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pulsar

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/sirupsen/logrus"
	"performance_consumer_go/conf"
)

func Start() error {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: fmt.Sprintf("pulsar://%s:%d", conf.PulsarHost, conf.PulsarPort),
	})
	if err != nil {
		return err
	}
	go startConsumer(client)
	return nil
}

func startConsumer(client pulsar.Client) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            conf.PulsarTopic,
		Type:             pulsar.Failover,
		SubscriptionName: conf.PulsarSubscriptionName,
	})
	if err != nil {
		logrus.Errorf("create consumer %s error: %v", conf.PulsarTopic, err)
	}
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			logrus.Errorf("receive message %s error: %v", conf.PulsarTopic, err)
		} else {
			logrus.Debugf("receive message %s success, messageID: %s", conf.PulsarTopic, msg.ID())
		}
	}
}

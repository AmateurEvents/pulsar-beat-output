/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package pulsar

import (
	"context"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/codec"
	"github.com/elastic/beats/libbeat/publisher"
)

type client struct {
	clientOptions   pulsar.ClientOptions
	producerOptions pulsar.ProducerOptions
	pulsarClient    pulsar.Client
	producer        pulsar.Producer
	observer        outputs.Observer
	beat            beat.Info
	config          *pulsarConfig
	codec           codec.Codec
}

func newPulsarClient(
	beat beat.Info,
	observer outputs.Observer,
	clientOptions pulsar.ClientOptions,
	producerOptions pulsar.ProducerOptions,
	config *pulsarConfig,
) (*client, error) {
	c := &client{
		clientOptions:   clientOptions,
		producerOptions: producerOptions,
		observer:        observer,
		beat:            beat,
		config:          config,
	}
	return c, nil
}

func (c *client) Connect() error {
	var err error
	c.pulsarClient, err = pulsar.NewClient(c.clientOptions)
	logp.Info("start create pulsar client")
	if err != nil {
		logp.Debug("pulsar", "Create pulsar client failed: %v", err)
		return err
	}
	logp.Info("start create pulsar producer")
	c.producer, err = c.pulsarClient.CreateProducer(c.producerOptions)
	if err != nil {
		logp.Debug("pulsar", "Create pulsar producer failed: %v", err)
		return err
	}
	logp.Info("start create encoder")
	c.codec, err = codec.CreateEncoder(c.beat, c.config.Codec)
	if err != nil {
		logp.Debug("pulsar", "Create encoder failed: %v", err)
		return err
	}

	return nil
}

func (c *client) Close() error {
	c.pulsarClient.Close()
	return nil
}

func (c *client) Publish(batch publisher.Batch) error {
	defer batch.ACK()
	events := batch.Events()
	c.observer.NewBatch(len(events))
	dropped := 0
	logp.Debug("pulsar", "Pulsar received events: %d", len(events))
	wg := sync.WaitGroup{}
	for i := range events {
		wg.Add(1)
		event := &events[i]
		serializedEvent, err := c.codec.Encode(c.beat.Beat, &event.Content)
		if err != nil {
			dropped++
			logp.Err("Failed event: %v, error: %v", event, err)
		}

		buf := make([]byte, len(serializedEvent))
		copy(buf, serializedEvent)

		logp.Debug("pulsar", "Pulsar success encode events: %v", string(serializedEvent))
		c.producer.SendAsync(context.Background(), &pulsar.ProducerMessage{
			Payload: buf,
		}, func(msgID pulsar.MessageID, message *pulsar.ProducerMessage, e error) {
			if e != nil {
				dropped++
				logp.Err("produce send failed: %v", err)
			}
			wg.Done()
		})
		logp.Debug("pulsar", "Pulsar success send event: %d", i)
	}
	c.producer.Flush()
	wg.Wait()
	c.observer.Dropped(dropped)
	c.observer.Acked(len(events) - dropped)
	logp.Debug("pulsar", "Pulsar success send events: %d", len(events))
	return nil
}

func (c *client) String() string {
	return "file(" + c.clientOptions.URL + ")"
}

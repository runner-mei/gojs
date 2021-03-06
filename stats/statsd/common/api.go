/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2019 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package common

import (
	"time"

	"github.com/runner-mei/gojs/stats"
)

// Sample defines a sample type
type Sample struct {
	Type   stats.MetricType  `json:"type"`
	Metric string            `json:"metric"`
	Time   time.Time         `json:"time"`
	Value  float64           `json:"value"`
	Tags   map[string]string `json:"tags,omitempty"`
}

func generateDataPoint(sample stats.Sample) *Sample {
	return &Sample{
		Type:   sample.Metric.Type,
		Metric: sample.Metric.Name,
		Time:   sample.Time,
		Value:  sample.Value,
		Tags:   sample.Tags.CloneTags(),
	}
}

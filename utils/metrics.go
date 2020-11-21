package utils

import (
	"fmt"
)

var (
	MetricTimeToFetch = func(id string) string {
		return fmt.Sprintf("%s/name:time_to_fetch", id)
	}
	MetricMsgsRcvd = func(id string) string {
		return fmt.Sprintf("%s/name:msgs_rcvd", id)
	}
	MetricDataSent = func(id string) string {
		return fmt.Sprintf("%s/name:data_sent", id)
	}
	MetricDataRcvd = func(id string) string {
		return fmt.Sprintf("%s/name:data_rcvd", id)
	}
	MetricDupDataRcvd = func(id string) string {
		return fmt.Sprintf("%s/name:dup_data_rcvd", id)
	}
	MetricBlksSent = func(id string) string {
		return fmt.Sprintf("%s/name:blks_sent", id)
	}
	MetricBlksRcvd = func(id string) string {
		return fmt.Sprintf("%s/name:blks_rcvd", id)
	}
	MetricDupBlksRcvd = func(id string) string {
		return fmt.Sprintf("%s/name:dup_blks_rcvd", id)
	}
)

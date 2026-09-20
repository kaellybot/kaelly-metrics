package insights

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

// A refused message never reaches the consumer, so the counter is the only signal
// besides the logs that a producer is publishing without a game.
func TestCountRejectedMessage(t *testing.T) {
	prom, ok := NewPrometheusMetrics().(*prom)
	if !ok {
		t.Fatal("NewPrometheusMetrics() did not return a *prom")
	}

	prom.CountRejectedMessage("JOB_GET_BOOK_REQUEST")
	prom.CountRejectedMessage("JOB_GET_BOOK_REQUEST")
	prom.CountRejectedMessage("ALIGN_SET_REQUEST")

	expected := `
		# HELP kaelly_rejected_messages_total Messages refused by the AMQP guards before reaching the consumer, by type.
		# TYPE kaelly_rejected_messages_total counter
		kaelly_rejected_messages_total{type="ALIGN_SET_REQUEST"} 1
		kaelly_rejected_messages_total{type="JOB_GET_BOOK_REQUEST"} 2
	`

	if err := testutil.CollectAndCompare(prom.rejected, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected metrics:\n%v", err)
	}
}

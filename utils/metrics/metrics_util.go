package metrics

import (
	"errors"
	"math/big"
	"strconv"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/she-protocol/she-chain/x/evm/types"
)

// Measures the time taken to execute a sudo msg
// Metric Names:
//
//	she_sudo_duration_miliseconds
//	she_sudo_duration_miliseconds_count
//	she_sudo_duration_miliseconds_sum
func MeasureSudoExecutionDuration(start time.Time, msgType string) {
	metrics.MeasureSinceWithLabels(
		[]string{"she", "sudo", "duration", "milliseconds"},
		start.UTC(),
		[]metrics.Label{telemetry.NewLabel("type", msgType)},
	)
}

// Measures failed sudo execution count
// Metric Name:
//
//	she_sudo_error_count
func IncrementSudoFailCount(msgType string) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "sudo", "error", "count"},
		1,
		[]metrics.Label{telemetry.NewLabel("type", msgType)},
	)
}

// Gauge metric with shed version and git commit as labels
// Metric Name:
//
//	shed_version_and_commit
func GaugeShedVersionAndCommit(version string, commit string) {
	telemetry.SetGaugeWithLabels(
		[]string{"shed_version_and_commit"},
		1,
		[]metrics.Label{telemetry.NewLabel("shed_version", version), telemetry.NewLabel("commit", commit)},
	)
}

// she_tx_process_type_count
func IncrTxProcessTypeCounter(processType string) {
	metrics.IncrCounterWithLabels(
		[]string{"she", "tx", "process", "type"},
		1,
		[]metrics.Label{telemetry.NewLabel("type", processType)},
	)
}

// Measures the time taken to process a block by the process type
// Metric Names:
//
//	she_process_block_miliseconds
//	she_process_block_miliseconds_count
//	she_process_block_miliseconds_sum
func BlockProcessLatency(start time.Time, processType string) {
	metrics.MeasureSinceWithLabels(
		[]string{"she", "process", "block", "milliseconds"},
		start.UTC(),
		[]metrics.Label{telemetry.NewLabel("type", processType)},
	)
}

// Measures the time taken to execute a sudo msg
// Metric Names:
//
//	she_tx_process_type_count
func IncrDagBuildErrorCounter(reason string) {
	metrics.IncrCounterWithLabels(
		[]string{"she", "dag", "build", "error"},
		1,
		[]metrics.Label{telemetry.NewLabel("reason", reason)},
	)
}

// Counts the number of concurrent transactions that failed
// Metric Names:
//
//	she_tx_concurrent_delivertx_error
func IncrFailedConcurrentDeliverTxCounter() {
	metrics.IncrCounterWithLabels(
		[]string{"she", "tx", "concurrent", "delievertx", "error"},
		1,
		[]metrics.Label{},
	)
}

// Counts the number of operations that failed due to operation timeout
// Metric Names:
//
//	she_log_not_done_after_counter
func IncrLogIfNotDoneAfter(label string) {
	metrics.IncrCounterWithLabels(
		[]string{"she", "log", "not", "done", "after"},
		1,
		[]metrics.Label{
			telemetry.NewLabel("label", label),
		},
	)
}

// Measures the time taken to execute a sudo msg
// Metric Names:
//
//	she_deliver_tx_duration_miliseconds
//	she_deliver_tx_duration_miliseconds_count
//	she_deliver_tx_duration_miliseconds_sum
func MeasureDeliverTxDuration(start time.Time) {
	metrics.MeasureSince(
		[]string{"she", "deliver", "tx", "milliseconds"},
		start.UTC(),
	)
}

// Measures the time taken to execute a batch tx
// Metric Names:
//
//	she_deliver_batch_tx_duration_miliseconds
//	she_deliver_batch_tx_duration_miliseconds_count
//	she_deliver_batch_tx_duration_miliseconds_sum
func MeasureDeliverBatchTxDuration(start time.Time) {
	metrics.MeasureSince(
		[]string{"she", "deliver", "batch", "tx", "milliseconds"},
		start.UTC(),
	)
}

// she_oracle_vote_penalty_count
func SetOracleVotePenaltyCount(count uint64, valAddr string, penaltyType string) {
	metrics.SetGaugeWithLabels(
		[]string{"she", "oracle", "vote", "penalty", "count"},
		float32(count),
		[]metrics.Label{
			telemetry.NewLabel("type", penaltyType),
			telemetry.NewLabel("validator", valAddr),
		},
	)
}

// she_epoch_new
func SetEpochNew(epochNum uint64) {
	metrics.SetGauge(
		[]string{"she", "epoch", "new"},
		float32(epochNum),
	)
}

// Measures throughput
// Metric Name:
//
//	she_throughput_<metric_name>
func SetThroughputMetric(metricName string, value float32) {
	telemetry.SetGauge(
		value,
		"she", "throughput", metricName,
	)
}

// Measures number of new websocket connects
// Metric Name:
//
//	she_websocket_connect
func IncWebsocketConnects() {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "websocket", "connect"},
		1,
		nil,
	)
}

// Measures number of times a denom's price is updated
// Metric Name:
//
//	she_oracle_price_update_count
func IncrPriceUpdateDenom(denom string) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "oracle", "price", "update"},
		1,
		[]metrics.Label{telemetry.NewLabel("denom", denom)},
	)
}

// Measures throughput per message type
// Metric Name:
//
//	she_throughput_<metric_name>
func SetThroughputMetricByType(metricName string, value float32, msgType string) {
	telemetry.SetGaugeWithLabels(
		[]string{"she", "loadtest", "tps", metricName},
		value,
		[]metrics.Label{telemetry.NewLabel("msg_type", msgType)},
	)
}

// Measures the number of times the total block gas wanted in the proposal exceeds the max
// Metric Name:
//
//	she_failed_total_gas_wanted_check
func IncrFailedTotalGasWantedCheck(proposer string) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "failed", "total", "gas", "wanted", "check"},
		1,
		[]metrics.Label{telemetry.NewLabel("proposer", proposer)},
	)
}

// Measures the number of times the total block gas wanted in the proposal exceeds the max
// Metric Name:
//
//	she_failed_total_gas_wanted_check
func IncrValidatorSlashed(proposer string) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "failed", "total", "gas", "wanted", "check"},
		1,
		[]metrics.Label{telemetry.NewLabel("proposer", proposer)},
	)
}

// Measures number of times a denom's price is updated
// Metric Name:
//
//	she_oracle_price_update_count
func SetCoinsMinted(amount uint64, denom string) {
	telemetry.SetGaugeWithLabels(
		[]string{"she", "mint", "coins"},
		float32(amount),
		[]metrics.Label{telemetry.NewLabel("denom", denom)},
	)
}

// Measures the number of times the total block gas wanted in the proposal exceeds the max
// Metric Name:
//
//	she_tx_gas_counter
func IncrGasCounter(gasType string, value int64) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "tx", "gas", "counter"},
		float32(value),
		[]metrics.Label{telemetry.NewLabel("type", gasType)},
	)
}

// Measures the number of times optimistic processing runs
// Metric Name:
//
//	she_optimistic_processing_counter
func IncrementOptimisticProcessingCounter(enabled bool) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "optimistic", "processing", "counter"},
		float32(1),
		[]metrics.Label{telemetry.NewLabel("enabled", strconv.FormatBool(enabled))},
	)
}

// Measures RPC endpoint request throughput
// Metric Name:
//
//	she_rpc_request_counter
func IncrementRpcRequestCounter(endpoint string, connectionType string, success bool) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "rpc", "request", "counter"},
		float32(1),
		[]metrics.Label{
			telemetry.NewLabel("endpoint", endpoint),
			telemetry.NewLabel("connection", connectionType),
			telemetry.NewLabel("success", strconv.FormatBool(success)),
		},
	)
}

func IncrementErrorMetrics(scenario string, err error) {
	if err == nil {
		return
	}
	var assocErr types.AssociationMissingErr
	if errors.As(err, &assocErr) {
		IncrementAssociationError(scenario, assocErr)
		return
	}
	// add other error types to handle as metrics
}

func IncrementAssociationError(scenario string, err types.AssociationMissingErr) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "association", "error"},
		1,
		[]metrics.Label{
			telemetry.NewLabel("scenario", scenario),
			telemetry.NewLabel("type", err.AddressType()),
		},
	)
}

// Measures the RPC request latency in milliseconds
// Metric Name:
//
//	she_rpc_request_latency_ms
func MeasureRpcRequestLatency(endpoint string, connectionType string, startTime time.Time) {
	metrics.MeasureSinceWithLabels(
		[]string{"she", "rpc", "request", "latency_ms"},
		startTime.UTC(),
		[]metrics.Label{
			telemetry.NewLabel("endpoint", endpoint),
			telemetry.NewLabel("connection", connectionType),
		},
	)
}

// IncrProducerEventCount increments the counter for events produced.
// This metric counts the number of events produced by the system.
// Metric Name:
//
//	she_loadtest_produce_count
func IncrProducerEventCount(msgType string) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "loadtest", "produce", "count"},
		1,
		[]metrics.Label{telemetry.NewLabel("msg_type", msgType)},
	)
}

// IncrConsumerEventCount increments the counter for events consumed.
// This metric counts the number of events consumed by the system.
// Metric Name:
//
//	she_loadtest_consume_count
func IncrConsumerEventCount(msgType string) {
	telemetry.IncrCounterWithLabels(
		[]string{"she", "loadtest", "consume", "count"},
		1,
		[]metrics.Label{telemetry.NewLabel("msg_type", msgType)},
	)
}

func AddHistogramMetric(key []string, value float32) {
	metrics.AddSample(key, value)
}

// Gauge for gas price paid for transactions
// Metric Name:
//
// she_evm_effective_gas_price
func HistogramEvmEffectiveGasPrice(gasPrice *big.Int) {
	AddHistogramMetric(
		[]string{"she", "evm", "effective", "gas", "price"},
		float32(gasPrice.Uint64()),
	)
}

// Gauge for block base fee
// Metric Name:
//
// she_evm_block_base_fee
func GaugeEvmBlockBaseFee(baseFee *big.Int, blockHeight int64) {
	metrics.SetGauge(
		[]string{"she", "evm", "block", "base", "fee"},
		float32(baseFee.Uint64()),
	)
}

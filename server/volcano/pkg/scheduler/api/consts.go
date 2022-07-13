package api

const (
	// NodeSelectorFitFailed means node could not fit the selector of pod of the queue
	NodeSelectorFitFailed = "node(s) selector fit queue failed"

	PoolBindingNodeLabelKeyFormat = "openi.octopus.resourcepool.%s"
	DefaultPoolName               = "common-pool"
	PoolBindingNodeLabelValue     = "bound"
)

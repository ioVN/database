package internal

import "errors"

var (
	// ErrLoadBalancedWithMultipleHosts is returned when loadBalanced=true is specified in a URI with multiple hosts.
	ErrLoadBalancedWithMultipleHosts = errors.New("loadBalanced cannot be set to true if multiple hosts are specified")
	// ErrLoadBalancedWithReplicaSet is returned when loadBalanced=true is specified in a URI with the replicaSet option.
	ErrLoadBalancedWithReplicaSet = errors.New("loadBalanced cannot be set to true if a replica set name is specified")
	// ErrLoadBalancedWithDirectConnection is returned when loadBalanced=true is specified in a URI with the directConnection option.
	ErrLoadBalancedWithDirectConnection = errors.New("loadBalanced cannot be set to true if the direct connection option is specified")
	// ErrSRVMaxHostsWithReplicaSet is returned when srvMaxHosts > 0 is specified in a URI with the replicaSet option.
	ErrSRVMaxHostsWithReplicaSet = errors.New("srvMaxHosts cannot be a positive value if a replica set name is specified")
	// ErrSRVMaxHostsWithLoadBalanced is returned when srvMaxHosts > 0 is specified in a URI with loadBalanced=true.
	ErrSRVMaxHostsWithLoadBalanced = errors.New("srvMaxHosts cannot be a positive value if loadBalanced is set to true")
)

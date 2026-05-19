package loadbalancer

import "sync"

type LoadBalancer struct {
	mu      sync.Mutex
	servers []string
	current int
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
	}
}

func (lb *LoadBalancer) NextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	server := lb.servers[lb.current]

	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

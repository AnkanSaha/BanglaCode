package system

import (
	"BanglaCode/src/object"
	"net"
	"sync"
	"time"
)

var (
	interfaceCache      []net.Interface
	interfaceCacheTime  time.Time
	interfaceCacheTTL   = 5 * time.Second
	interfaceCacheMutex sync.RWMutex
)

func getInterfaces() ([]net.Interface, error) {
	interfaceCacheMutex.RLock()
	if time.Since(interfaceCacheTime) < interfaceCacheTTL && interfaceCache != nil {
		defer interfaceCacheMutex.RUnlock()
		return interfaceCache, nil
	}
	interfaceCacheMutex.RUnlock()

	// Need to refresh cache
	interfaceCacheMutex.Lock()
	defer interfaceCacheMutex.Unlock()

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	interfaceCache = interfaces
	interfaceCacheTime = time.Now()
	return interfaces, nil
}

func init() {
	// ==================== Network Information ====================

	// network_interface (নেটওয়ার্ক ইন্টারফেস) - Get all network interfaces
	// Returns array of maps with interface info
	registerBuiltin("network_interface", func(args ...object.Object) object.Object {
		interfaces, err := getInterfaces()
		if err != nil {
			return newError("failed to get network interfaces: %s", err.Error())
		}

		elements := make([]object.Object, 0, len(interfaces))
		for _, iface := range interfaces {
			info := make(map[string]object.Object)
			info["naam"] = &object.String{Value: iface.Name}
			info["mtu"] = &object.Number{Value: float64(iface.MTU)}
			info["mac"] = &object.String{Value: iface.HardwareAddr.String()}

			elements = append(elements, &object.Map{Pairs: info})
		}

		return &object.Array{Elements: elements}
	})

	// ip_address (আইপি ঠিকানা) - Get IP addresses for an interface
	registerBuiltin("ip_address", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("ip_address requires 1 argument (interface name)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("interface name must be STRING, got %s", args[0].Type())
		}

		ifaceName := args[0].(*object.String).Value

		iface, err := net.InterfaceByName(ifaceName)
		if err != nil {
			return newError("failed to get interface: %s", err.Error())
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return newError("failed to get addresses: %s", err.Error())
		}

		elements := make([]object.Object, 0, len(addrs))
		for _, addr := range addrs {
			elements = append(elements, &object.String{Value: addr.String()})
		}

		return &object.Array{Elements: elements}
	})

	// ip_shokal (আইপি সকল) - Get all IP addresses from all interfaces
	registerBuiltin("ip_shokal", func(args ...object.Object) object.Object {
		interfaces, err := getInterfaces()
		if err != nil {
			return newError("failed to get network interfaces: %s", err.Error())
		}

		var allAddrs []object.Object
		for _, iface := range interfaces {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				allAddrs = append(allAddrs, &object.String{Value: addr.String()})
			}
		}

		return &object.Array{Elements: allAddrs}
	})

	// mac_address (ম্যাক ঠিকানা) - Get MAC address for an interface
	registerBuiltin("mac_address", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("mac_address requires 1 argument (interface name)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("interface name must be STRING, got %s", args[0].Type())
		}

		ifaceName := args[0].(*object.String).Value

		iface, err := net.InterfaceByName(ifaceName)
		if err != nil {
			return newError("failed to get interface: %s", err.Error())
		}

		return &object.String{Value: iface.HardwareAddr.String()}
	})

	// network_gateway (নেটওয়ার্ক গেটওয়ে) - Get default gateway IP
	// Note: This is platform-specific and complex to implement
	// Returning placeholder for now
	registerBuiltin("network_gateway", func(args ...object.Object) object.Object {
		return newError("network_gateway not implemented yet")
	})

	// dns_server (ডিএনএস সার্ভার) - Get DNS server addresses
	// Note: This is platform-specific and complex to implement
	// Returning placeholder for now
	registerBuiltin("dns_server", func(args ...object.Object) object.Object {
		return newError("dns_server not implemented yet")
	})
}

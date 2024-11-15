// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	commonv1alpha1 "github.com/onmetal/onmetal-api/api/common/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
	"github.com/onmetal/onmetal-api/broker/common/cleaner"
	machinebrokerv1alpha1 "github.com/onmetal/onmetal-api/broker/machinebroker/api/v1alpha1"
	"github.com/onmetal/onmetal-api/broker/machinebroker/apiutils"
	ori "github.com/onmetal/onmetal-api/ori/apis/machine/v1alpha1"
	onmetalapiannotations "github.com/onmetal/onmetal-api/utils/annotations"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (s *Server) prepareOnmetalVirtualIP(virtualIPSpec *ori.VirtualIPSpec) (*networkingv1alpha1.VirtualIP, error) {
	ip, err := commonv1alpha1.ParseIP(virtualIPSpec.Ip)
	if err != nil {
		return nil, err
	}

	onmetalVirtualIP := &networkingv1alpha1.VirtualIP{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: s.cluster.Namespace(),
			Name:      s.cluster.IDGen().Generate(),
		},
		Spec: networkingv1alpha1.VirtualIPSpec{
			Type:     networkingv1alpha1.VirtualIPTypePublic,
			IPFamily: ip.Family(),
		},
		Status: networkingv1alpha1.VirtualIPStatus{
			IP: &ip,
		},
	}
	onmetalapiannotations.SetExternallyMangedBy(onmetalVirtualIP, machinebrokerv1alpha1.MachineBrokerManager)
	apiutils.SetPurpose(onmetalVirtualIP, machinebrokerv1alpha1.NetworkInterfacePurpose)
	return onmetalVirtualIP, nil
}

func (s *Server) prepareOnmetalLoadBalancerTarget(lbTgt *ori.LoadBalancerTargetSpec) (*machinebrokerv1alpha1.LoadBalancerTarget, error) {
	typ, err := s.convertORILoadBalancerType(lbTgt.LoadBalancerType)
	if err != nil {
		return nil, err
	}

	ip, err := commonv1alpha1.ParseIP(lbTgt.Ip)
	if err != nil {
		return nil, err
	}

	var ports []machinebrokerv1alpha1.LoadBalancerPort
	for _, port := range lbTgt.Ports {
		protocol, err := s.convertORIProtocol(port.Protocol)
		if err != nil {
			return nil, err
		}

		ports = append(ports, machinebrokerv1alpha1.LoadBalancerPort{
			Protocol: protocol,
			Port:     port.Port,
			EndPort:  port.EndPort,
		})
	}

	return &machinebrokerv1alpha1.LoadBalancerTarget{
		LoadBalancerType: typ,
		IP:               ip,
		Ports:            ports,
	}, nil
}

func (s *Server) prepareOnmetalLoadBalancerTargets(lbTargets []*ori.LoadBalancerTargetSpec) ([]machinebrokerv1alpha1.LoadBalancerTarget, error) {
	var res []machinebrokerv1alpha1.LoadBalancerTarget
	for _, lbTgt := range lbTargets {
		tgt, err := s.prepareOnmetalLoadBalancerTarget(lbTgt)
		if err != nil {
			return nil, err
		}

		res = append(res, *tgt)
	}
	return res, nil
}

func (s *Server) prepareOnmetalNATGatewayTarget(nat *ori.NATSpec) (*machinebrokerv1alpha1.NATGatewayTarget, error) {
	ip, err := commonv1alpha1.ParseIP(nat.Ip)
	if err != nil {
		return nil, fmt.Errorf("error parsing nat ip: %w", err)
	}

	return &machinebrokerv1alpha1.NATGatewayTarget{
		IP:      ip,
		Port:    nat.Port,
		EndPort: nat.EndPort,
	}, nil
}

func (s *Server) prepareOnmetalNATGatewayTargets(nats []*ori.NATSpec) ([]machinebrokerv1alpha1.NATGatewayTarget, error) {
	var res []machinebrokerv1alpha1.NATGatewayTarget
	for _, nat := range nats {
		tgt, err := s.prepareOnmetalNATGatewayTarget(nat)
		if err != nil {
			return nil, fmt.Errorf("error preparing nat target: %w", err)
		}

		res = append(res, *tgt)
	}
	return res, nil
}

func (s *Server) prepareAggregateOnmetalNetworkInterface(log logr.Logger, networkInterface *ori.NetworkInterface) (*AggregateOnmetalNetworkInterface, error) {

	var onmetalVirtualIP *networkingv1alpha1.VirtualIP
	if virtualIPSpec := networkInterface.Spec.VirtualIp; virtualIPSpec != nil {
		v, err := s.prepareOnmetalVirtualIP(virtualIPSpec)
		if err != nil {
			return nil, err
		}

		onmetalVirtualIP = v
	}

	ips, err := s.parseIPs(networkInterface.Spec.Ips)
	if err != nil {
		return nil, err
	}

	prefixes, err := s.parseIPPrefixes(networkInterface.Spec.Prefixes)
	if err != nil {
		return nil, err
	}

	lbTgts, err := s.prepareOnmetalLoadBalancerTargets(networkInterface.Spec.LoadBalancerTargets)
	if err != nil {
		return nil, err
	}

	natGatewayTgts, err := s.prepareOnmetalNATGatewayTargets(networkInterface.Spec.Nats)
	if err != nil {
		return nil, err
	}

	var virtualIPSource *networkingv1alpha1.VirtualIPSource
	if onmetalVirtualIP != nil {
		virtualIPSource = &networkingv1alpha1.VirtualIPSource{
			VirtualIPRef: &corev1.LocalObjectReference{Name: onmetalVirtualIP.Name},
		}
	}
	onmetalNetworkInterface := &networkingv1alpha1.NetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: s.cluster.Namespace(),
			Name:      s.cluster.IDGen().Generate(),
		},
		Spec: networkingv1alpha1.NetworkInterfaceSpec{
			IPFamilies: s.getOnmetalIPsIPFamilies(ips),
			IPs:        s.onmetalIPsToOnmetalIPSources(ips),
			Prefixes:   s.onmetalPrefixesToOnmetalPrefixSources(prefixes),
			VirtualIP:  virtualIPSource,
		},
	}
	if err := apiutils.SetObjectMetadata(onmetalNetworkInterface, networkInterface.Metadata); err != nil {
		return nil, err
	}
	apiutils.SetManagerLabel(onmetalNetworkInterface, machinebrokerv1alpha1.MachineBrokerManager)

	onmetalNetworkInterfaceConfig := &AggregateOnmetalNetworkInterface{
		NetworkInterface: onmetalNetworkInterface,
		Network: &networkingv1alpha1.Network{
			Spec: networkingv1alpha1.NetworkSpec{Handle: networkInterface.Spec.Network.Handle},
		},
		VirtualIP:           onmetalVirtualIP,
		LoadBalancerTargets: lbTgts,
		NATGatewayTargets:   natGatewayTgts,
	}

	peerings := []networkingv1alpha1.NetworkPeeringStatus{}
	//TODO fix duplicate code
	for _, peeringSpec := range networkInterface.Spec.Network.Peerings {
		peering := networkingv1alpha1.NetworkPeeringStatus{
			Name:          peeringSpec.Name,
			NetworkHandle: peeringSpec.Handle,
			Phase:         networkingv1alpha1.NetworkPeeringPhase(peeringSpec.Phase),
		}

		prefixes := []commonv1alpha1.IPPrefix{}
		if len(peeringSpec.Prefixes) > 0 {
			for _, prefix := range peeringSpec.Prefixes {
				prefixes = append(prefixes, commonv1alpha1.MustParseIPPrefix(prefix))
			}
		}

		if len(prefixes) > 0 {
			peering.Prefixes = &prefixes
		}

		peerings = append(peerings, peering)
	}

	onmetalNetworkInterfaceConfig.Network.Status.Peerings = peerings

	return onmetalNetworkInterfaceConfig, nil
}

func (s *Server) setOnmetalVirtualIPStatusIP(ctx context.Context, onmetalVirtualIP *networkingv1alpha1.VirtualIP, ip *commonv1alpha1.IP) error {
	baseOnmetalVirtualIP := onmetalVirtualIP.DeepCopy()
	onmetalVirtualIP.Status.IP = ip
	if err := s.cluster.Client().Status().Patch(ctx, onmetalVirtualIP, client.MergeFrom(baseOnmetalVirtualIP)); err != nil {
		return fmt.Errorf("error patching onmetal virtual ip status: %w", err)
	}
	return nil
}

func (s *Server) createOnmetalVirtualIP(ctx context.Context, log logr.Logger, c *cleaner.Cleaner, onmetalVirtualIP *networkingv1alpha1.VirtualIP) error {
	ip := *onmetalVirtualIP.Status.IP

	log.V(1).Info("Creating onmetal virtual ip")
	if err := s.cluster.Client().Create(ctx, onmetalVirtualIP); err != nil {
		return fmt.Errorf("error creating onmetal virtual ip: %w", err)
	}
	c.Add(func(ctx context.Context) error {
		if err := s.cluster.Client().Delete(ctx, onmetalVirtualIP); client.IgnoreNotFound(err) != nil {
			return fmt.Errorf("error deleting virtual ip: %w", err)
		}
		return nil
	})

	log.V(1).Info("Patching onmetal virtual ip status")
	if err := s.setOnmetalVirtualIPStatusIP(ctx, onmetalVirtualIP, &ip); err != nil {
		return fmt.Errorf("error setting onmetal virtual ip status ip: %w", err)
	}

	return nil
}

func (s *Server) createOnmetalNetworkInterface(ctx context.Context, log logr.Logger, onmetalNetworkInterface *AggregateOnmetalNetworkInterface) (retErr error) {
	c := cleaner.New()
	defer cleaner.CleanupOnError(ctx, c, &retErr)

	network, networkTransaction, err := s.networks.BeginCreate(ctx, onmetalNetworkInterface.Network)
	if err != nil {
		return fmt.Errorf("error getting network: %w", err)
	}
	c.Add(RollbackTransactionIgnoreClosedFunc(networkTransaction))
	onmetalNetworkInterface.Network = network

	onmetalNetworkInterface.NetworkInterface.Spec.NetworkRef = corev1.LocalObjectReference{Name: network.Name}

	if onmetalVirtualIP := onmetalNetworkInterface.VirtualIP; onmetalVirtualIP != nil {
		if err := s.createOnmetalVirtualIP(ctx, log, c, onmetalVirtualIP); err != nil {
			return err
		}
	}

	log.V(1).Info("Creating onmetal network interface")
	if err := s.cluster.Client().Create(ctx, onmetalNetworkInterface.NetworkInterface); err != nil {
		return fmt.Errorf("error creating onmetal network interface: %w", err)
	}
	c.Add(cleaner.DeleteObjectIfExistsFunc(s.cluster.Client(), onmetalNetworkInterface.NetworkInterface))

	if err := networkTransaction.Commit(onmetalNetworkInterface.NetworkInterface); err != nil {
		return fmt.Errorf("error committing network: %w", err)
	}
	c.Add(func(ctx context.Context) error {
		return s.networks.Delete(ctx, onmetalNetworkInterface.Network.Spec.Handle, onmetalNetworkInterface.NetworkInterface.Name)
	})

	if onmetalNetworkInterface.VirtualIP != nil {
		log.V(1).Info("Patching onmetal virtual ip to be controlled by onmetal network interface")
		if err := apiutils.PatchControlledBy(ctx, s.cluster.Client(), onmetalNetworkInterface.NetworkInterface, onmetalNetworkInterface.VirtualIP); err != nil {
			return fmt.Errorf("error patching onmetal virtual ip to be controlled by onmetal network interface: %w", err)
		}
	}

	log.V(1).Info("Creating load balancers")
	for _, lbTgt := range onmetalNetworkInterface.LoadBalancerTargets {
		if err := s.loadBalancers.Create(ctx, network, lbTgt, onmetalNetworkInterface.NetworkInterface); err != nil {
			return fmt.Errorf("error creating load balancer: %w", err)
		}
		c.Add(func(ctx context.Context) error {
			return s.loadBalancers.Delete(ctx, network.Spec.Handle, lbTgt, onmetalNetworkInterface.NetworkInterface)
		})
	}

	log.V(1).Info("Creating nat gateways")
	for _, natGatewayTgt := range onmetalNetworkInterface.NATGatewayTargets {
		if err := s.natGateways.Create(ctx, network, natGatewayTgt, onmetalNetworkInterface.NetworkInterface); err != nil {
			return fmt.Errorf("error creating nat gateway: %w", err)
		}
		c.Add(func(ctx context.Context) error {
			return s.natGateways.Delete(ctx, network.Spec.Handle, natGatewayTgt.IP, onmetalNetworkInterface.NetworkInterface)
		})
	}

	log.V(1).Info("Patching onmetal network interface as created")
	if err := apiutils.PatchCreated(ctx, s.cluster.Client(), onmetalNetworkInterface.NetworkInterface); err != nil {
		return fmt.Errorf("error patching onmetal network interface as created: %w", err)
	}

	return nil
}

func (s *Server) CreateNetworkInterface(ctx context.Context, req *ori.CreateNetworkInterfaceRequest) (res *ori.CreateNetworkInterfaceResponse, retErr error) {
	log := s.loggerFrom(ctx)

	onmetalNetworkInterface, err := s.prepareAggregateOnmetalNetworkInterface(log, req.NetworkInterface)
	if err != nil {
		return nil, fmt.Errorf("error getting onmetal network interface config: %w", err)
	}

	if err := s.createOnmetalNetworkInterface(ctx, log, onmetalNetworkInterface); err != nil {
		return nil, fmt.Errorf("error creating onmetal network interface: %w", err)
	}

	networkInterface, err := s.convertAggregateOnmetalNetworkInterface(onmetalNetworkInterface)
	if err != nil {
		return nil, fmt.Errorf("error converting onmetal network interface: %w", err)
	}

	return &ori.CreateNetworkInterfaceResponse{
		NetworkInterface: networkInterface,
	}, nil
}
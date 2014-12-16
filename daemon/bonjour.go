package daemon

import (
	"log"
	"net"

	"github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/socketplane/bonjour"
	"github.com/socketplane/socketplane/ipam"
	"github.com/socketplane/socketplane/ovs"
)

const DOCKER_CLUSTER_SERVICE = "_docker._cluster"
const DOCKER_CLUSTER_SERVICE_PORT = 9999 //TODO : fix this
const DOCKER_CLUSTER_DOMAIN = "local"

func Bonjour(intfName string) {
	b := bonjour.Bonjour{
		ServiceName:   DOCKER_CLUSTER_SERVICE,
		ServiceDomain: DOCKER_CLUSTER_DOMAIN,
		ServicePort:   DOCKER_CLUSTER_SERVICE_PORT,
		InterfaceName: intfName,
		BindToIntf:    true,
		Notify:        notify{},
	}
	b.Start()
}

type notify struct{}

func (n notify) NewMember(addr net.IP) {
	log.Println("New Member Added : ", addr)
	ipam.Join(addr.String())
	ovs.AddPeer(addr.String())
}
func (n notify) RemoveMember(addr net.IP) {
	log.Println("Member Left : ", addr)
	ovs.DeletePeer(addr.String())
}
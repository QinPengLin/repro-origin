package cluster

import "github.com/QinPengLin/repro-origin/rpc"

type ConfigDiscovery struct {
	funDelNode  FunDelNode
	funSetNode  FunSetNode
	localNodeId string
}

func (discovery *ConfigDiscovery) InitDiscovery(localNodeId string, funDelNode FunDelNode, funSetNode FunSetNode) error {
	discovery.localNodeId = localNodeId
	discovery.funDelNode = funDelNode
	discovery.funSetNode = funSetNode

	//解析本地其他服务配置
	_, nodeInfoList, _, err := GetCluster().readLocalClusterConfig(rpc.NodeIdNull)
	if err != nil {
		return err
	}

	for _, nodeInfo := range nodeInfoList {
		if nodeInfo.NodeId == localNodeId {
			continue
		}

		discovery.funSetNode(&nodeInfo)
	}

	return nil
}

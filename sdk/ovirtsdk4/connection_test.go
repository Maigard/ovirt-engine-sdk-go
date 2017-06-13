package ovirtsdk4

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"
	conn, err := NewConnection(
		inputRawURL, "admin@internal", "qwer1234",
		"", true, "", false,
		uint64(10*time.Second), true)
	if err != nil {
		t.Errorf("connection failed, reason %s", err.Error())
	}
	// ovRequest := NewOvRequest("GET", "/clusters", nil, nil, "")
	result, err := conn.SystemService().ClustersService().List(false, false, 100, "", nil, nil, false)
	var clusters Clusters
	if ovResp, ok := result.(*OvResponse); ok {
		t.Logf("response code is %d", ovResp.Code)
		// t.Logf("response body is %s", ovResp.Body)
		xml.Unmarshal([]byte(ovResp.Body), &clusters)
		t.Logf("clusters length is %d", len(clusters.Clusters))
	}
	if len(clusters.Clusters) > 0 {
		for _, cluster := range clusters.Clusters {
			t.Logf("cluster(%v): %+v", cluster.Id, cluster.Cpu)
		}
	}

	// ovResponse, err := conn.Send(ovRequest)
	// t.Logf("response %s", string(ovResponse.Body))
}

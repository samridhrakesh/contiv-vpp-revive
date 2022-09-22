// Code generated by 'create-test-data.sh' on Mon Jan 14 05:16:59 PST 2019. DO NOT EDIT.

package testdata

func getRawK8sNodeTestData() []string {
	return []string{

		`{
			"addresses": [
				{
					"address": "10.20.0.2",
					"type": "NodeInternalIP"
				},
				{
					"address": "k8s-master",
					"type": "NodeHostName"
				}
			],
			"name": "k8s-master",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.11.5",
				"OperatingSystem": "linux",
				"boot_ID": "30896f50-fb56-4b24-bfa1-90e42799b993",
				"container_runtime_version": "docker://18.3.0",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.11.5",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "8B506DD7-F713-44AB-88FD-9D83CD812A24"
			},
			"pod_CIDR": "10.0.0.0/24"
		}`,
		`{
			"addresses": [
				{
					"address": "10.20.0.10",
					"type": "NodeInternalIP"
				},
				{
					"address": "k8s-worker1",
					"type": "NodeHostName"
				}
			],
			"name": "k8s-worker1",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.11.5",
				"OperatingSystem": "linux",
				"boot_ID": "ffe0a278-bb86-4852-b1ef-db713a47bcdd",
				"container_runtime_version": "docker://18.3.0",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.11.5",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "D4912B1C-C960-4E5F-A267-27C1DDD1CDDD"
			},
			"pod_CIDR": "10.0.1.0/24"
		}`,
		`{
			"addresses": [
				{
					"address": "10.20.0.11",
					"type": "NodeInternalIP"
				},
				{
					"address": "k8s-worker2",
					"type": "NodeHostName"
				}
			],
			"name": "k8s-worker2",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.11.5",
				"OperatingSystem": "linux",
				"boot_ID": "84c0fcaf-9060-4d54-9a63-1024a8a7f3cb",
				"container_runtime_version": "docker://18.3.0",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.11.5",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "B56E7987-7776-48D5-A468-ACDE70B49BC0"
			},
			"pod_CIDR": "10.0.2.0/24"
		}`,
	}
}
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/logger"
)

type nodeLogsOptions struct {
	privateKeyPath string
	username       string
	nodeIps        []string
	osFamily       string
}

type nodeLogCommand struct {
	resourceName string
	command      string
}

func NewNodeLogCommands(osFamily string) []nodeLogCommand {
	if osFamily == "bottlerocket" {
		return []nodeLogCommand{
			{
				resourceName: "kubelet",
				command:      "sudo sheltie journalctl -u kubelet.service --no-pager",
			},
			{
				resourceName: "containerd",
				command:      "sudo sheltie journalctl -u containerd.service --no-pager",
			},
			{
				resourceName: "etcd",
				command:      "sudo sheltie cat /var/log/containers/$(sudo sheltie ls /var/log/containers/ | grep etcd)",
			},
			{
				resourceName: "cloud-init",
				command:      "sudo sheltie journalctl _COMM=host-ctr --no-pager",
			},
			{
				resourceName: "manifests",
				command:      "sudo sheltie cat /etc/kubernetes/manifests/*",
			},
		}
	}

	return []nodeLogCommand{
		{
			resourceName: "kubelet",
			command:      "sudo journalctl -u kubelet.service --no-pager",
		},
		{
			resourceName: "containerd",
			command:      "sudo journalctl -u containerd.service --no-pager",
		},
		{
			resourceName: "etcd-unstacked",
			command:      "sudo journalctl -u etcd.service --no-pager",
		},
		{
			resourceName: "etcd-stacked",
			command:      "sudo crictl logs $(sudo crictl ps --name etcd -q)",
		},
		{
			resourceName: "cloud-init",
			command:      "sudo cat /var/log/cloud-init-output.log",
		},
		{
			resourceName: "manifests",
			command:      "sudo cat /etc/kubernetes/manifests/*",
		},
	}
}

const (
	logOutputNameFormat = "%s-%s.log"
)

var nodeLogsCmd = &cobra.Command{
	Use:   "nodelogs",
	Short: "Get logs from cluster nodes",
	Long:  "Get container logs from cluster nodes",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := nodeLogs(cmd.Context())
		if err != nil {
			log.Fatalf("Error getting container logs from cluster nodes: %v", err)
		}
		return nil
	},
}

var nodeLogsOpt = &nodeLogsOptions{}

func init() {
	rootCmd.AddCommand(nodeLogsCmd)
	nodeLogsCmd.Flags().StringVarP(&nodeLogsOpt.privateKeyPath, "privatekey", "k", "", "Private key file path")
	nodeLogsCmd.Flags().StringVarP(&nodeLogsOpt.username, "username", "u", "", "Username to access the nodes")
	nodeLogsCmd.Flags().StringSliceVarP(&nodeLogsOpt.nodeIps, "nodeip", "i", make([]string, 0), "IP addresses of the nodes")
	nodeLogsCmd.Flags().StringVarP(&nodeLogsOpt.osFamily, "osfamily", "o", "", "OS family of the nodes (bottlerocket/ubuntu/rhel)")

	_ = nodeLogsCmd.MarkFlagRequired("privatekey")
}

func nodeLogs(ctx context.Context) error {
	writer, _ := filewriter.NewWriter("nodelogs")
	ssh := executables.NewLocalExecutablesBuilder().BuildSSHExecutable()
	nodeLogCommands := NewNodeLogCommands(nodeLogsOpt.osFamily)

	for _, nodeIP := range nodeLogsOpt.nodeIps {
		for _, nlc := range nodeLogCommands {
			stdout, err := ssh.RunCommand(ctx, nodeLogsOpt.privateKeyPath, nodeLogsOpt.username, nodeIP, nlc.command)
			if err != nil {
				stdout = *bytes.NewBufferString(err.Error())
			}
			if err := writeNodeLogsToFile(writer, stdout, nodeIP, nlc.resourceName); err != nil {
				return fmt.Errorf("writing logs to file: %v", err)
			}
		}
	}

	return nil
}

func writeNodeLogsToFile(writer filewriter.FileWriter, out bytes.Buffer, nodeIP, logName string) error {
	filename := fmt.Sprintf(logOutputNameFormat, logName, nodeIP)
	path, err := writer.Write(filename, out.Bytes())
	if err != nil {
		return err
	}
	logger.Info("Successfully fetched node logs", "path", path)
	return nil
}

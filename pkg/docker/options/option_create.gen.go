// Code generated by 'option-gen'. DO NOT EDIT.

package options

import (
	prompt "github.com/gohumble/go-prompt"
)

var Create = []prompt.Suggest{
	prompt.Suggest{Text: "--add-host list", Description: "Add a custom host-to-IP mapping (host:ip)"},
	prompt.Suggest{Text: "-a", Description: "Attach to STDIN, STDOUT or STDERR"},
	prompt.Suggest{Text: "--attach list", Description: "Attach to STDIN, STDOUT or STDERR"},
	prompt.Suggest{Text: "--blkio-weight uint16", Description: "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)"},
	prompt.Suggest{Text: "--blkio-weight-device list", Description: "Block IO weight (relative device weight) (default [])"},
	prompt.Suggest{Text: "--cap-add list", Description: "Add Linux capabilities"},
	prompt.Suggest{Text: "--cap-drop list", Description: "Drop Linux capabilities"},
	prompt.Suggest{Text: "--cgroup-parent string", Description: "Optional parent cgroup for the container"},
	prompt.Suggest{Text: "--cidfile string", Description: "Write the container ID to the file"},
	prompt.Suggest{Text: "--cpu-period int", Description: "Limit CPU CFS (Completely Fair Scheduler) period"},
	prompt.Suggest{Text: "--cpu-quota int", Description: "Limit CPU CFS (Completely Fair Scheduler) quota"},
	prompt.Suggest{Text: "--cpu-rt-period int", Description: "Limit CPU real-time period in microseconds"},
	prompt.Suggest{Text: "--cpu-rt-runtime int", Description: "Limit CPU real-time runtime in microseconds"},
	prompt.Suggest{Text: "-c", Description: "CPU shares (relative weight)"},
	prompt.Suggest{Text: "--cpu-shares int", Description: "CPU shares (relative weight)"},
	prompt.Suggest{Text: "--cpus decimal", Description: "Number of CPUs"},
	prompt.Suggest{Text: "--cpuset-cpus string", Description: "CPUs in which to allow execution (0-3, 0,1)"},
	prompt.Suggest{Text: "--cpuset-mems string", Description: "MEMs in which to allow execution (0-3, 0,1)"},
	prompt.Suggest{Text: "--device list", Description: "Add a host device to the container"},
	prompt.Suggest{Text: "--device-cgroup-rule list", Description: "Add a rule to the cgroup allowed devices list"},
	prompt.Suggest{Text: "--device-read-bps list", Description: "Limit read rate (bytes per second) from a device (default [])"},
	prompt.Suggest{Text: "--device-read-iops list", Description: "Limit read rate (IO per second) from a device (default [])"},
	prompt.Suggest{Text: "--device-write-bps list", Description: "Limit write rate (bytes per second) to a device (default [])"},
	prompt.Suggest{Text: "--device-write-iops list", Description: "Limit write rate (IO per second) to a device (default [])"},
	prompt.Suggest{Text: "--disable-content-trust", Description: "Skip image verification (default true)"},
	prompt.Suggest{Text: "--dns list", Description: "Set custom DNS servers"},
	prompt.Suggest{Text: "--dns-option list", Description: "Set DNS options"},
	prompt.Suggest{Text: "--dns-search list", Description: "Set custom DNS search domains"},
	prompt.Suggest{Text: "--domainname string", Description: "Container NIS domain name"},
	prompt.Suggest{Text: "--entrypoint string", Description: "Overwrite the default ENTRYPOINT of the image"},
	prompt.Suggest{Text: "-e", Description: "Set environment variables"},
	prompt.Suggest{Text: "--env list", Description: "Set environment variables"},
	prompt.Suggest{Text: "--env-file list", Description: "Read in a file of environment variables"},
	prompt.Suggest{Text: "--expose list", Description: "Expose a port or a range of ports"},
	prompt.Suggest{Text: "--gpus gpu-request", Description: "GPU devices to add to the container ('all' to pass all GPUs)"},
	prompt.Suggest{Text: "--group-add list", Description: "Add additional groups to join"},
	prompt.Suggest{Text: "--health-cmd string", Description: "Command to run to check health"},
	prompt.Suggest{Text: "--health-interval duration", Description: "Time between running the check (ms|s|m|h) (default 0s)"},
	prompt.Suggest{Text: "--health-retries int", Description: "Consecutive failures needed to report unhealthy"},
	prompt.Suggest{Text: "--health-timeout duration", Description: "Maximum time to allow one check to run (ms|s|m|h) (default 0s)"},
	prompt.Suggest{Text: "--help", Description: "Print usage"},
	prompt.Suggest{Text: "-h", Description: "Container host name"},
	prompt.Suggest{Text: "--hostname string", Description: "Container host name"},
	prompt.Suggest{Text: "--init", Description: "Run an init inside the container that forwards signals and reaps processes"},
	prompt.Suggest{Text: "-i", Description: "Keep STDIN open even if not attached"},
	prompt.Suggest{Text: "--interactive", Description: "Keep STDIN open even if not attached"},
	prompt.Suggest{Text: "--ip string", Description: "IPv4 address (e.g., 172.30.100.104)"},
	prompt.Suggest{Text: "--ip6 string", Description: "IPv6 address (e.g., 2001:db8::33)"},
	prompt.Suggest{Text: "--ipc string", Description: "IPC mode to use"},
	prompt.Suggest{Text: "--isolation string", Description: "Container isolation technology"},
	prompt.Suggest{Text: "--kernel-memory bytes", Description: "Kernel memory limit"},
	prompt.Suggest{Text: "-l", Description: "Set meta data on a container"},
	prompt.Suggest{Text: "--label list", Description: "Set meta data on a container"},
	prompt.Suggest{Text: "--label-file list", Description: "Read in a line delimited file of labels"},
	prompt.Suggest{Text: "--link list", Description: "Add link to another container"},
	prompt.Suggest{Text: "--link-local-ip list", Description: "Container IPv4/IPv6 link-local addresses"},
	prompt.Suggest{Text: "--log-driver string", Description: "Logging driver for the container"},
	prompt.Suggest{Text: "--log-opt list", Description: "Log driver options"},
	prompt.Suggest{Text: "--mac-address string", Description: "Container MAC address (e.g., 92:d0:c6:0a:29:33)"},
	prompt.Suggest{Text: "-m", Description: "Memory limit"},
	prompt.Suggest{Text: "--memory bytes", Description: "Memory limit"},
	prompt.Suggest{Text: "--memory-reservation bytes", Description: "Memory soft limit"},
	prompt.Suggest{Text: "--memory-swap bytes", Description: "Swap limit equal to memory plus swap: '-1' to enable unlimited swap"},
	prompt.Suggest{Text: "--memory-swappiness int", Description: "Tune container memory swappiness (0 to 100) (default -1)"},
	prompt.Suggest{Text: "--mount mount", Description: "Attach a filesystem mount to the container"},
	prompt.Suggest{Text: "--name string", Description: "Assign a name to the container"},
	prompt.Suggest{Text: "--network network", Description: "Connect a container to a network"},
	prompt.Suggest{Text: "--network-alias list", Description: "Add network-scoped alias for the container"},
	prompt.Suggest{Text: "--no-healthcheck", Description: "Disable any container-specified HEALTHCHECK"},
	prompt.Suggest{Text: "--oom-kill-disable", Description: "Disable OOM Killer"},
	prompt.Suggest{Text: "--oom-score-adj int", Description: "Tune host's OOM preferences (-1000 to 1000)"},
	prompt.Suggest{Text: "--pid string", Description: "PID namespace to use"},
	prompt.Suggest{Text: "--pids-limit int", Description: "Tune container pids limit (set -1 for unlimited)"},
	prompt.Suggest{Text: "--privileged", Description: "Give extended privileges to this container"},
	prompt.Suggest{Text: "-p", Description: "Publish a container's port(s) to the host"},
	prompt.Suggest{Text: "--publish list", Description: "Publish a container's port(s) to the host"},
	prompt.Suggest{Text: "-P", Description: "Publish all exposed ports to random ports"},
	prompt.Suggest{Text: "--publish-all", Description: "Publish all exposed ports to random ports"},
	prompt.Suggest{Text: "--read-only", Description: "Mount the container's root filesystem as read only"},
	prompt.Suggest{Text: "--restart string", Description: "Restart policy to apply when a container exits (default \"no\")"},
	prompt.Suggest{Text: "--rm", Description: "Automatically remove the container when it exits"},
	prompt.Suggest{Text: "--runtime string", Description: "Runtime to use for this container"},
	prompt.Suggest{Text: "--security-opt list", Description: "Security Options"},
	prompt.Suggest{Text: "--shm-size bytes", Description: "Size of /dev/shm"},
	prompt.Suggest{Text: "--stop-signal string", Description: "Signal to stop a container (default \"SIGTERM\")"},
	prompt.Suggest{Text: "--stop-timeout int", Description: "Timeout (in seconds) to stop a container"},
	prompt.Suggest{Text: "--storage-opt list", Description: "Storage driver options for the container"},
	prompt.Suggest{Text: "--sysctl map", Description: "Sysctl options (default map[])"},
	prompt.Suggest{Text: "--tmpfs list", Description: "Mount a tmpfs directory"},
	prompt.Suggest{Text: "-t", Description: "Allocate a pseudo-TTY"},
	prompt.Suggest{Text: "--tty", Description: "Allocate a pseudo-TTY"},
	prompt.Suggest{Text: "--ulimit ulimit", Description: "Ulimit options (default [])"},
	prompt.Suggest{Text: "-u", Description: "Username or UID (format: <name|uid>[:<group|gid>])"},
	prompt.Suggest{Text: "--user string", Description: "Username or UID (format: <name|uid>[:<group|gid>])"},
	prompt.Suggest{Text: "--userns string", Description: "User namespace to use"},
	prompt.Suggest{Text: "--uts string", Description: "UTS namespace to use"},
	prompt.Suggest{Text: "-v", Description: "Bind mount a volume"},
	prompt.Suggest{Text: "--volume list", Description: "Bind mount a volume"},
	prompt.Suggest{Text: "--volume-driver string", Description: "Optional volume driver for the container"},
	prompt.Suggest{Text: "--volumes-from list", Description: "Mount volumes from the specified container(s)"},
	prompt.Suggest{Text: "-w", Description: "Working directory inside the container"},
	prompt.Suggest{Text: "--workdir string", Description: "Working directory inside the container"},
}

// Code generated by 'option-gen'. DO NOT EDIT.

package options

import (
	prompt "github.com/gohumble/go-prompt"
)

var Build = []prompt.Suggest{
	prompt.Suggest{Text: "--add-host list", Description: "Add a custom host-to-IP mapping (host:ip)"},
	prompt.Suggest{Text: "--build-arg list", Description: "Set build-time variables"},
	prompt.Suggest{Text: "--cache-from strings", Description: "Images to consider as cache sources"},
	prompt.Suggest{Text: "--cgroup-parent string", Description: "Optional parent cgroup for the container"},
	prompt.Suggest{Text: "--compress", Description: "Compress the build context using gzip"},
	prompt.Suggest{Text: "--cpu-period int", Description: "Limit the CPU CFS (Completely Fair Scheduler) period"},
	prompt.Suggest{Text: "--cpu-quota int", Description: "Limit the CPU CFS (Completely Fair Scheduler) quota"},
	prompt.Suggest{Text: "-c", Description: "CPU shares (relative weight)"},
	prompt.Suggest{Text: "--cpu-shares int", Description: "CPU shares (relative weight)"},
	prompt.Suggest{Text: "--cpuset-cpus string", Description: "CPUs in which to allow execution (0-3, 0,1)"},
	prompt.Suggest{Text: "--cpuset-mems string", Description: "MEMs in which to allow execution (0-3, 0,1)"},
	prompt.Suggest{Text: "-f", Description: "Name of the Dockerfile (Default is 'PATH/Dockerfile')"},
	prompt.Suggest{Text: "--file string", Description: "Name of the Dockerfile (Default is 'PATH/Dockerfile')"},
	prompt.Suggest{Text: "--force-rm", Description: "Always remove intermediate containers"},
	prompt.Suggest{Text: "--iidfile string", Description: "Write the image ID to the file"},
	prompt.Suggest{Text: "--isolation string", Description: "Container isolation technology"},
	prompt.Suggest{Text: "--label list", Description: "Set metadata for an image"},
	prompt.Suggest{Text: "-m", Description: "Memory limit"},
	prompt.Suggest{Text: "--memory bytes", Description: "Memory limit"},
	prompt.Suggest{Text: "--memory-swap bytes", Description: "Swap limit equal to memory plus swap: '-1' to enable unlimited swap"},
	prompt.Suggest{Text: "--network string", Description: "Set the networking mode for the RUN instructions during build (default \"default\")"},
	prompt.Suggest{Text: "--no-cache", Description: "Do not use cache when building the image"},
	prompt.Suggest{Text: "--pull", Description: "Always attempt to pull a newer version of the image"},
	prompt.Suggest{Text: "-q", Description: "Suppress the build output and print image ID on success"},
	prompt.Suggest{Text: "--quiet", Description: "Suppress the build output and print image ID on success"},
	prompt.Suggest{Text: "--rm", Description: "Remove intermediate containers after a successful build (default true)"},
	prompt.Suggest{Text: "--security-opt strings", Description: "Security options"},
	prompt.Suggest{Text: "--shm-size bytes", Description: "Size of /dev/shm"},
	prompt.Suggest{Text: "-t", Description: "Name and optionally a tag in the 'name:tag' format"},
	prompt.Suggest{Text: "--tag list", Description: "Name and optionally a tag in the 'name:tag' format"},
	prompt.Suggest{Text: "--target string", Description: "Set the target build stage to build."},
	prompt.Suggest{Text: "--ulimit ulimit", Description: "Ulimit options (default [])"},
}

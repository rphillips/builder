package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	k8sversion "k8s.io/apimachinery/pkg/version"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"

	"github.com/openshift/builder/pkg/build/builder/cmd"
	"github.com/openshift/builder/pkg/version"
)

var (
	s2iBuilderLong = templates.LongDesc(`
		Perform a Source-to-Image build

		This command executes a Source-to-Image build using arguments passed via the environment.
		It expects to be run inside of a container.`)

	dockerBuilderLong = templates.LongDesc(`
		Perform a Docker build

		This command executes a Docker build using arguments passed via the environment.
		It expects to be run inside of a container.`)

	gitCloneLong = templates.LongDesc(`
		Perform a Git clone

		This command executes a Git clone using arguments passed via the environment.
		It expects to be run inside of a container.`)

	manageDockerfileLong = templates.LongDesc(`
		Manipulates a dockerfile for a docker build.

		This command updates a dockerfile based on build inputs.
		It expects to be run inside of a container.`)

	extractImageContentLong = templates.LongDesc(`
		Extracts files from existing images.

		This command extracts files from existing images to use as input to a build.
		It expects to be run inside of a container.`)
)

// NewCmdVersion provides a shim around version for
// non-client packages that require version information
func NewCmdVersion(fullName string, versionInfo k8sversion.Info, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display version",
		Long:  "Display version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(out, "%s %v\n", fullName, versionInfo)
		},
	}

	return cmd
}

// NewCommandS2IBuilder provides a CLI handler for S2I build type
func NewCommandS2IBuilder(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Run a Source-to-Image build",
		Long:  s2iBuilderLong,
		Run: func(c *cobra.Command, args []string) {
			err := cmd.RunS2IBuild(c.OutOrStderr())
			kcmdutil.CheckErr(err)
		},
	}

	cmd.AddCommand(NewCmdVersion(name, version.Get(), os.Stdout))
	return cmd
}

// NewCommandDockerBuilder provides a CLI handler for Docker build type
func NewCommandDockerBuilder(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Run a Docker build",
		Long:  dockerBuilderLong,
		Run: func(c *cobra.Command, args []string) {
			err := cmd.RunDockerBuild(c.OutOrStderr())
			kcmdutil.CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdVersion(name, version.Get(), os.Stdout))
	return cmd
}

// NewCommandGitClone manages cloning the git source for a build.
// It also manages binary build input content.
func NewCommandGitClone(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Git clone source code",
		Long:  gitCloneLong,
		Run: func(c *cobra.Command, args []string) {
			err := cmd.RunGitClone(c.OutOrStderr())
			kcmdutil.CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdVersion(name, version.Get(), os.Stdout))
	return cmd
}

func NewCommandManageDockerfile(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Manage a dockerfile for a docker build",
		Long:  manageDockerfileLong,
		Run: func(c *cobra.Command, args []string) {
			err := cmd.RunManageDockerfile(c.OutOrStderr())
			kcmdutil.CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdVersion(name, version.Get(), os.Stdout))
	return cmd
}

func NewCommandExtractImageContent(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Extract build input content from existing images",
		Long:  extractImageContentLong,
		Run: func(c *cobra.Command, args []string) {
			err := cmd.RunExtractImageContent(c.OutOrStderr())
			kcmdutil.CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdVersion(name, version.Get(), os.Stdout))
	return cmd
}

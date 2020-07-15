package thing

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	Test    bool
	CmdLine string
	output  string
}

func NewRunner() *Runner {
	return &Runner{Test: false}
}

func (r *Runner) Command(command string, args ...string) (string, error) {
	if r.Test {
		r.CmdLine = fmt.Sprintf("%s %s", command, strings.Join(args, " "))
		return r.CmdLine, nil
	}

	output, err := exec.Command(command, args...).CombinedOutput()
	r.output = fmt.Sprintf("%s %s", command, strings.Join(args, " "))

	if err != nil {
		return "", fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}

	return string(output), nil
}

func (r *Runner) YumUpdate() string {
	if r.Test == true {
		r.Command("yum", "update -y")
		return r.CmdLine
	}
	r.Command("yum", "update -y")
	return r.output
}

func (r *Runner) YumInstall() string {
	if r.Test == true {
		r.Command("yum", "install epel-release -y")
		return r.CmdLine
	}
	r.Command("yum", "install epel-release -y")
	return r.output
}

func (r *Runner) RubyInstall() string {
	if r.Test == true {
		r.Command("yum", "install ruby -y")
		return r.CmdLine
	}
	r.Command("yum", "install ruby -y")
	return r.output
}

func (r *Runner) JekyllInstall() string {
	if r.Test == true {
		r.Command("gem", "install jekyll")
		return r.CmdLine
	}
	r.Command("gem", "install jekyll")
	return r.output
}

func (r *Runner) BundlerInstall() string {
	if r.Test == true {
		r.Command("gem", "install bundler")
		return r.CmdLine
	}
	r.Command("gem", "install bundler")
	return r.output
}

func (r *Runner) RubyVersion() string {
	if r.Test == true {
		r.Command("ruby", "2.6")
		return r.CmdLine
	}
	r.Command("ruby", "2.6")
	return r.output
}

func (r *Runner) JekyllVersion() string {
	if r.Test == true {
		r.Command("jekyll", "4.*")
		return r.CmdLine
	}
	r.Command("jekyll", "4.*")
	return r.output
}

func (r *Runner) BundlerVersion() string {
	if r.Test == true {
		r.Command("Bundler", "2.*")
		return r.CmdLine
	}
	r.Command("Bundler", "2.*")
	return r.output
}

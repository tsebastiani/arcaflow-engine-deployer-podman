package podman

import (
	"io"

	log "go.arcalot.io/log/v2"
	"go.flow.arcalot.io/podmandeployer/internal/cliwrapper"
)

type CliPlugin struct {
	wrapper        cliwrapper.CliWrapper
	containerImage string
	containerName  string
	config         *Config
	logger         log.Logger
	stdin          io.WriteCloser
	stdout         io.ReadCloser
}

// TODO: unwrap the whole config

func (p *CliPlugin) Write(b []byte) (n int, err error) {
	return p.stdin.Write(b)
}

func (p *CliPlugin) Read(b []byte) (n int, err error) {
	return p.stdout.Read(b)
}

func (p *CliPlugin) Close() error {
	if err := p.wrapper.KillAndClean(p.containerName); err != nil {
		return err
	}

	if err := p.stdin.Close(); err != nil {
		p.logger.Errorf("failed to close stdin pipe")
	} else {
		p.logger.Infof("stdin pipe successfully closed")
	}
	if err := p.stdout.Close(); err != nil {
		p.logger.Infof("failed to close stdout pipe")
	} else {
		p.logger.Infof("stdout pipe successfully closed")
	}
	return nil
}

func (p *CliPlugin) ID() string {
	return p.containerName
}

package printers

import (
	"io"

	node "github.com/ernoaapa/eliot/pkg/api/services/node/v1"
	pods "github.com/ernoaapa/eliot/pkg/api/services/pods/v1"
	"github.com/ernoaapa/eliot/pkg/config"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// YamlPrinter is ResourcePrinter implementation which writes output in YAML format
type YamlPrinter struct {
}

// NewYamlPrinter creates new YamlPrinter instance
func NewYamlPrinter() *YamlPrinter {
	return &YamlPrinter{}
}

// PrintPods takes list of pods and prints to Writer in YAML format
func (p *YamlPrinter) PrintPods(pods []*pods.Pod, w io.Writer) error {
	if err := writeAsYml(pods, w); err != nil {
		return errors.Wrap(err, "Failed to write pods yaml")
	}
	return nil
}

// PrintNodes takes list of nodes and prints to Writer in YAML format
func (p *YamlPrinter) PrintNodes(nodes []*node.Info, w io.Writer) error {
	if err := writeAsYml(nodes, w); err != nil {
		return errors.Wrap(err, "Failed to write nodes yaml")
	}
	return nil
}

// PrintNode takes node info and prints to Writer in YAML format
func (p *YamlPrinter) PrintNode(node *node.Info, w io.Writer) error {
	if err := writeAsYml(node, w); err != nil {
		return errors.Wrap(err, "Failed to write node yaml")
	}
	return nil
}

// PrintPod takes Pod and prints to Writer in YAML format
func (p *YamlPrinter) PrintPod(pod *pods.Pod, w io.Writer) error {
	if err := writeAsYml(pod, w); err != nil {
		return errors.Wrap(err, "Failed to write pod yaml")
	}
	return nil
}

// PrintConfig takes Config and prints to Writer in YAML format
func (p *YamlPrinter) PrintConfig(config *config.Config, w io.Writer) error {
	if err := writeAsYml(config, w); err != nil {
		return errors.Wrap(err, "Failed to write config yaml")
	}
	return nil
}

func writeAsYml(in interface{}, w io.Writer) error {
	data, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

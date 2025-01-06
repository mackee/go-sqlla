package sqlla

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"strings"
)

type Plugin struct {
	Name    string
	Outpath string
	Args    map[string]string
	Table   *Table
}

func (p *Plugin) WriteCode() error {
	w, err := os.Create(p.Outpath)
	if err != nil {
		return fmt.Errorf("fail to create: plugin=%s: %w", p.Name, err)
	}
	defer w.Close()

	buf := &bytes.Buffer{}
	ptmpl := tmpl.Lookup(p.Name)
	if ptmpl == nil {
		return fmt.Errorf("template not found: plugin=%s", p.Name)
	}
	if err := ptmpl.Execute(buf, p); err != nil {
		return fmt.Errorf("fail to render: %w", err)
	}
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		if _, err := w.Write(buf.Bytes()); err != nil {
			return fmt.Errorf("fail to write: plugin=%s: %w", p.Name, err)
		}
		return fmt.Errorf("fail to format: plugin=%s: %w", p.Name, err)
	}
	if _, err := w.Write(bs); err != nil {
		return fmt.Errorf("fail to write: plugin=%s: %w", p.Name, err)
	}
	return nil
}

type Plugins []*Plugin

func (p Plugins) WriteCode() error {
	for _, plugin := range p {
		if err := plugin.WriteCode(); err != nil {
			return err
		}
	}
	return nil
}

var errThisCommentIsNotPlugin = errors.New("this comment is not plugin")

func parsePluginsByComments(comments []string) (Plugins, error) {
	plugins := make(Plugins, 0, len(comments))
	for _, comment := range comments {
		plugin, err := parsePluginByComment(comment)
		if errors.Is(err, errThisCommentIsNotPlugin) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("fail to parse plugin: args=%s %w", comment, err)
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}

func parsePluginByComment(comment string) (*Plugin, error) {
	pluginStr := strings.TrimPrefix(comment, "//sqlla:plugin ")
	if pluginStr == comment {
		return nil, errThisCommentIsNotPlugin
	}
	nameArgs := strings.Split(pluginStr, " ")
	if len(nameArgs) == 0 {
		return nil, errors.New("plugin name is not specified")
	}
	name := nameArgs[0]
	args := make(map[string]string, len(nameArgs)-1)
	for _, arg := range nameArgs[1:] {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			return nil, errors.New("invalid argument")
		}
		args[kv[0]] = kv[1]
	}
	if _, ok := args["outpath"]; !ok {
		return nil, errors.New("outpath is not specified")
	}
	outpath := args["outpath"]
	delete(args, "outpath")

	return &Plugin{
		Name:    name,
		Outpath: outpath,
		Args:    args,
	}, nil
}

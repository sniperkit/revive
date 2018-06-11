package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/cep21/xdgbasedir"
	"github.com/jinzhu/configor"

	"gopkg.in/yaml.v2"
)

var (
	configDirectoryPath string = ""
	configFormats              = []string{"yaml", "json", "toml"}
)

// Config contains configuration information
type Config struct {
	Formatters []lint.Formatter `json:"formatters" yaml:"formatters" toml:"formatters" xml:"formatters" ini:"formatters"`
	Rules      []lint.Rule      `json:"rules" yaml:"rules" toml:"rules" xml:"rules" ini:"rules"`

	Default struct {
		Formatters []lint.Formatter `json:"formatters" yaml:"formatters" toml:"formatters" xml:"formatters" ini:"formatters"`
		Rules      []lint.Rule      `json:"rules" yaml:"rules" toml:"rules" xml:"rules" ini:"rules"`
	} `json:"default" yaml:"default" toml:"default" xml:"default" ini:"default"`
}

func init() {
	baseDir, err := xdgbasedir.ConfigHomeDirectory()
	if err != nil {
		log.Fatal("Can't find XDG BaseDirectory")
	} else {
		configDirectoryPath = path.Join(baseDir, ProgramName)
	}
}

var defaultRules = []lint.Rule{
	&rule.VarDeclarationsRule{},
	&rule.PackageCommentsRule{},
	&rule.DotImportsRule{},
	&rule.BlankImportsRule{},
	&rule.ExportedRule{},
	&rule.VarNamingRule{},
	&rule.IndentErrorFlowRule{},
	&rule.IfReturnRule{},
	&rule.RangeRule{},
	&rule.ErrorfRule{},
	&rule.ErrorNamingRule{},
	&rule.ErrorStringsRule{},
	&rule.ReceiverNamingRule{},
	&rule.IncrementDecrementRule{},
	&rule.ErrorReturnRule{},
	&rule.UnexportedReturnRule{},
	&rule.TimeNamingRule{},
	&rule.ContextKeysType{},
	&rule.ContextAsArgumentRule{},
}

var allRules = append([]lint.Rule{
	&rule.ArgumentsLimitRule{},
	&rule.CyclomaticRule{},
	&rule.FileHeaderRule{},
	&rule.EmptyBlockRule{},
	&rule.SuperfluousElseRule{},
}, defaultRules...)

var allFormatters = []lint.Formatter{
	&formatter.Stylish{},
	&formatter.Friendly{},
	&formatter.JSON{},
	&formatter.Default{},
}

func getFormatters() map[string]lint.Formatter {
	result := map[string]lint.Formatter{}
	for _, f := range allFormatters {
		result[f.Name()] = f
	}
	return result
}

func getLintingRules(config *lint.Config) []lint.Rule {
	rulesMap := map[string]lint.Rule{}
	for _, r := range allRules {
		rulesMap[r.Name()] = r
	}

	lintingRules := []lint.Rule{}
	for name := range config.Rules {
		rule, ok := rulesMap[name]
		if !ok {
			fail("cannot find rule: " + name)
		}
		lintingRules = append(lintingRules, rule)
	}

	return lintingRules
}

func parseConfig(path string) *lint.Config {
	config := &lint.Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fail("cannot read the config file")
	}
	_, err = toml.Decode(string(file), config)
	if err != nil {
		fail("cannot parse the config file: " + err.Error())
	}
	return config
}

func normalizeConfig(config *lint.Config) {
	if config.Confidence == 0 {
		config.Confidence = 0.8
	}
	severity := config.Severity
	if severity != "" {
		for k, v := range config.Rules {
			if v.Severity == "" {
				v.Severity = severity
			}
			config.Rules[k] = v
		}
	}
}

func getConfig() *lint.Config {
	config := defaultConfig()
	if configPath != "" {
		config = parseConfig(configPath)
	}
	normalizeConfig(config)
	return config
}

func getFormatter() lint.Formatter {
	formatters := getFormatters()
	formatter := formatters["default"]
	if formatterName != "" {
		f, ok := formatters[formatterName]
		if !ok {
			fail("unknown formatter " + formatterName)
		}
		formatter = f
	}
	return formatter
}

func DefaultConfig() *lint.Config {
	defaultConfig := lint.Config{
		Confidence: 0.0,
		Severity:   lint.SeverityWarning,
		Rules:      map[string]lint.RuleConfig{},
	}
	for _, r := range defaultRules {
		defaultConfig.Rules[r.Name()] = lint.RuleConfig{}
	}
	return &defaultConfig
}

func normalizeSplit(strs []string) []string {
	res := []string{}
	for _, s := range strs {
		t := strings.Trim(s, " \t")
		if len(t) > 0 {
			res = append(res, t)
		}
	}
	return res
}

func getPackages() [][]string {
	globs := normalizeSplit(flag.Args())
	if len(globs) == 0 {
		globs = append(globs, ".")
	}

	packages, err := dots.ResolvePackages(globs, normalizeSplit(excludePaths))
	if err != nil {
		fail(err.Error())
	}

	return packages
}

// ReadConfig reads the configuration information
func ReadConfig() (*Config, error) {
	file := configFilePath()

	var config Config
	if _, err := os.Stat(file); err == nil {
		// Read and unmarshal file only if it exists
		f, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		// Earlier configurations have higher priority
		err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&config, "config.toml")

		// err = yaml.Unmarshal(f, &config)
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}

// WriteConfig writes the configuration information
func (config *Config) WriteConfig() error {
	err := os.MkdirAll(configDirectoryPath, 0700)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configFilePath(), data, 0600)
}

func configFilePath() string {
	return path.Join(configDirectoryPath, fmt.Sprintf("%s", ProgramName))
}

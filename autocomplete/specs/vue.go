// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["vue"] = model.Subcommand{
		Name:        []string{"vue"},
		Description: `Vue cli tools`,
		Options: []model.Option{{
			Name:        []string{"-V", "--version"},
			Description: `Output the version number`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Output usage information`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"create"},
			Description: `Create a new project powered by vue-cli-service`,
			Args: []model.Arg{{
				Name: "app-name",
			}},
			Options: []model.Option{{
				Name:        []string{"-p", "--preset"},
				Description: `Skip prompts and use saved or remote preset`,
				Args: []model.Arg{{
					Name: "presetName",
				}},
			}, {
				Name:        []string{"-d", "--default"},
				Description: `Skip prompts and use default preset`,
			}, {
				Name:        []string{"-i", "--inlinePreset"},
				Description: `Skip prompts and use inline JSON string as preset`,
				Args: []model.Arg{{
					Name: "json",
				}},
			}, {
				Name:        []string{"-m", "--packageManager"},
				Description: `Use specified npm client when installing dependencies`,
				Args: []model.Arg{{
					Name: "command",
				}},
			}, {
				Name:        []string{"-r", "--registry"},
				Description: `Use specified npm registry when installing dependencies (only for npm)`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"-g", "--git"},
				Description: `Force git initialization with initial commit message`,
				Args: []model.Arg{{
					Name: "message",
				}},
			}, {
				Name:        []string{"-n", "--no-git"},
				Description: `Skip git initialization`,
			}, {
				Name:        []string{"-f", "--force"},
				Description: `Overwrite target directory if it exists`,
			}, {
				Name:        []string{"--merge"},
				Description: `Merge target directory if it exists`,
			}, {
				Name:        []string{"-c", "--clone"},
				Description: `Use git clone when fetching remote preset`,
			}, {
				Name:        []string{"-X", "--proxy"},
				Description: `Use specified proxy when creating project`,
				Args: []model.Arg{{
					Name: "proxyUrl",
				}},
			}, {
				Name:        []string{"-b", "--bar"},
				Description: `Scaffold project without beginner instructions`,
			}, {
				Name:        []string{"--skipGetStarted"},
				Description: `Skip displaying "Get started" instructions`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"add"},
			Description: `Install a plugin and invoke its generator in an already created project`,
			Args: []model.Arg{{
				Name: "plugin",
			}},
			Options: []model.Option{{
				Name:        []string{"--registry"},
				Description: `Use specified npm registry when installing dependencies (only for npm)`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"invoke"},
			Description: `Invoke the generator of a plugin in an already created project`,
			Args: []model.Arg{{
				Name: "plugin",
			}},
			Options: []model.Option{{
				Name:        []string{"--registry"},
				Description: `Use specified npm registry when installing dependencies (only for npm)`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"inspect"},
			Description: `Inspect the webpack config in a project with vue-cli-service`,
			Args:        []model.Arg{{}},
			Options: []model.Option{{
				Name: []string{"--mode"},
				Args: []model.Arg{{
					Name: "mode",
				}},
			}, {
				Name:        []string{"--rule"},
				Description: `Inspect a specific module rule`,
				Args: []model.Arg{{
					Name: "ruleName",
				}},
			}, {
				Name:        []string{"--plugin"},
				Description: `Inspect a specific plugin`,
				Args: []model.Arg{{
					Name: "pluginName",
				}},
			}, {
				Name:        []string{"--rules"},
				Description: `List all module rule names`,
			}, {
				Name:        []string{"--plugins"},
				Description: `List all plugin names`,
			}, {
				Name:        []string{"-v", "--verbose"},
				Description: `Show full function definitions in output`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"serve"},
			Description: `Serve a .js or .vue file in development mode with zero config`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
			Options: []model.Option{{
				Name:        []string{"-o", "--open"},
				Description: `Open browser`,
			}, {
				Name:        []string{"-c", "--copy"},
				Description: `Copy local url to clipboard`,
			}, {
				Name:        []string{"-p", "--port"},
				Description: `Port used by the server (default: 8080 or next available port)`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"build"},
			Description: `Build a .js or .vue file in production mode with zero config`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
			}},
			Options: []model.Option{{
				Name:        []string{"-t", "--target"},
				Description: `Build target (app | lib | wc | wc-async, default: app)`,
				Args: []model.Arg{{
					Name: "target",
				}},
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Name for lib or web-component mode (default: entry filename)`,
				Args: []model.Arg{{
					Name: "name",
				}},
			}, {
				Name:        []string{"-d", "--destination"},
				Description: `Output directory (default: dist)`,
				Args: []model.Arg{{
					Name: "dir",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"ui"},
			Description: `Start and open the vue-cli ui`,
			Options: []model.Option{{
				Name:        []string{"-H", "--host"},
				Description: `Host used for the UI server (default: localhost)`,
				Args: []model.Arg{{
					Name: "host",
				}},
			}, {
				Name:        []string{"-p", "--port"},
				Description: `Port used for the UI server (by default search for available port)`,
				Args: []model.Arg{{
					Name: "port",
				}},
			}, {
				Name:        []string{"-D", "--dev"},
				Description: `Run in dev mode`,
			}, {
				Name:        []string{"--quiet"},
				Description: `Don't output starting messages`,
			}, {
				Name:        []string{"--headless"},
				Description: `Don't open browser on start and output port`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"init"},
			Description: `Generate a project from a remote template (legacy API, requires @vue/cli-init)`,
			Args: []model.Arg{{
				Name: "template",
			}, {
				Name: "app-name",
			}},
			Options: []model.Option{{
				Name:        []string{"-c", "--clone"},
				Description: `Use git clone when fetching remote template`,
			}, {
				Name:        []string{"--offline"},
				Description: `Use cached template`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"config"},
			Description: `Inspect and modify the config`,
			Args: []model.Arg{{
				Name: "value",
			}},
			Options: []model.Option{{
				Name:        []string{"-g", "--get"},
				Description: `Get value from option`,
				Args: []model.Arg{{
					Name: "path",
				}},
			}, {
				Name:        []string{"-s", "--set"},
				Description: `Set option value`,
				Args: []model.Arg{{
					Name: "value",
				}},
			}, {
				Name:        []string{"-d", "--delete"},
				Description: `Delete option from config`,
				Args: []model.Arg{{
					Name: "path",
				}},
			}, {
				Name:        []string{"-e", "--edit"},
				Description: `Open config with default editor`,
			}, {
				Name:        []string{"--json"},
				Description: `Outputs JSON result only`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"outdated"},
			Description: `(experimental) check for outdated vue cli service / plugins`,
			Options: []model.Option{{
				Name:        []string{"--next"},
				Description: `Also check for alpha / beta / rc versions when upgrading`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `(experimental) upgrade vue cli service / plugins`,
			Args: []model.Arg{{
				Name: "plugin-name",
			}},
			Options: []model.Option{{
				Name:        []string{"-t", "--to"},
				Description: `Upgrade <package-name> to a version that is not latest`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}, {
				Name:        []string{"-f", "--from"},
				Description: `Skip probing installed plugin, assuming it is upgraded from the designated version`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}, {
				Name:        []string{"-r", "--registry"},
				Description: `Use specified npm registry when installing dependencies`,
				Args: []model.Arg{{
					Name: "url",
				}},
			}, {
				Name:        []string{"--all"},
				Description: `Upgrade all plugins`,
			}, {
				Name:        []string{"--next"},
				Description: `Also check for alpha / beta /rc versions when upgrading`,
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"migrate"},
			Description: `(experimental) run migrator for an already-installed cli plugin`,
			Args: []model.Arg{{
				Name: "plugin-name",
			}},
			Options: []model.Option{{
				Name:        []string{"-f", "--from"},
				Description: `The base version for the migrator to migrate from`,
				Args: []model.Arg{{
					Name: "version",
				}},
			}, {
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}, {
			Name:        []string{"info"},
			Description: `Print debugging information about your environment`,
			Options: []model.Option{{
				Name:        []string{"-h", "--help"},
				Description: `Output usage information`,
			}},
		}},
	}
}
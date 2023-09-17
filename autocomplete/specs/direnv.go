// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["direnv"] = model.Subcommand{
		Name:        []string{"direnv"},
		Description: `Unclutter your .profile`,
		Options: []model.Option{{
			Name:        []string{"--version"},
			Description: `Prints the version or checks that direnv is older than VERSION_AT_LEAST`,
			Args: []model.Arg{{
				Name:       "VERSION_AT_LEAST",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--help"},
			Description: `Help for direnv`,
			Args: []model.Arg{{
				Name:        "SHOW_PRIVATE",
				Description: `Any string`,
				IsOptional:  true,
			}},
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"allow"},
			Description: `Grants direnv to load the given .envrc`,
			Args: []model.Arg{{
				Name:       "PATH_TO_RC",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
		}, {
			Name:        []string{"deny"},
			Description: `Revokes the authorization of a given .envrc`,
			Args: []model.Arg{{
				Name:       "PATH_TO_RC",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
		}, {
			Name:        []string{"edit"},
			Description: `Opens PATH_TO_RC or the current .envrc into an $EDITOR and allow the file to be loaded afterwards`,
			Args: []model.Arg{{
				Name:       "PATH_TO_RC",
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
		}, {
			Name:        []string{"exec"},
			Description: `Executes a command after loading the first .envrc found in DIR`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "DIR",
			}, {
				Name:      "COMMAND",
				IsCommand: true,
			}},
		}, {
			Name:        []string{"fetchurl"},
			Description: `Fetches a given URL into direnv's CAS`,
			Args: []model.Arg{{
				Name: "url",
			}, {
				Name:        "integrity-hash",
				Description: `Check if the "integrity hash" is equal to the hash value of the file obtained from the "url"`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Help for direnv`,
			Args: []model.Arg{{
				Name:        "SHOW_PRIVATE",
				Description: `Any string`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"hook"},
			Description: `Used to setup the shell hook`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
			}},
		}, {
			Name:        []string{"prune"},
			Description: `Removes old allowed files`,
		}, {
			Name:        []string{"reload"},
			Description: `Triggers an env reload`,
		}, {
			Name:        []string{"status"},
			Description: `Prints some debug status information`,
		}, {
			Name:        []string{"stdlib"},
			Description: `Displays the stdlib available in the .envrc execution context`,
		}, {
			Name:        []string{"version"},
			Description: `Prints the version or checks that direnv is older than VERSION_AT_LEAST`,
			Args: []model.Arg{{
				Name:       "VERSION_AT_LEAST",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"apply_dump"},
			Description: `Accepts a filename containing "direnv dump" output and generates a series of bash export statements to apply the given env`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}, {
			Name:        []string{"show_dump"},
			Description: `Show the data inside of a dump for debugging purposes`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "DUMP",
			}},
		}, {
			Name:        []string{"dotenv"},
			Description: `Transforms a .env file to evaluatable "export KEY=PAIR" statements`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
				IsOptional: true,
			}, {
				Name:      "PATH_TO_DOTENV",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"dump"},
			Description: `Used to export the inner bash state at the end of execution`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
				IsOptional: true,
			}, {
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "FILE",
				Description: `Overwrites by dump data`,
				IsOptional:  true,
			}},
		}, {
			Name:        []string{"export"},
			Description: `Loads an .envrc and prints the diff in terms of exports`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
			}},
		}, {
			Name:        []string{"watch"},
			Description: `Adds a path to the list that direnv watches for changes`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
			}, {
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "PATH",
				IsVariadic: true,
			}},
		}, {
			Name:        []string{"watch-dir"},
			Description: `Recursively adds a directory to the list that direnv watches for changes`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
			}, {
				Templates: []model.Template{model.TemplateFolders},
				Name:      "DIR",
			}},
		}, {
			Name:        []string{"watch-list"},
			Description: `Pipe pairs of "mtime path" to stdin to build a list of files to watch`,
			Args: []model.Arg{{
				Name: "SHELL",
				Suggestions: []model.Suggestion{{
					Name: []string{`bash`},
				}, {
					Name: []string{`zsh`},
				}, {
					Name: []string{`fish`},
				}, {
					Name: []string{`tcsh`},
				}, {
					Name: []string{`elvish`},
				}},
				IsOptional: true,
			}},
		}, {
			Name:        []string{"current"},
			Description: `Reports whether direnv's view of a file is current (or stale)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "PATH",
			}},
		}},
	}
}
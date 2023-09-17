// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["mdls"] = model.Subcommand{
		Name:        []string{"mdls"},
		Description: `Lists the metadata attributes for the specified file`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths},
			Name:       "file",
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"--name", "-name"},
			Description: `Print only the matching metadata attribute value.  Can be used multiple times`,
			Args: []model.Arg{{
				Name:        "attributeName",
				Description: `Metadata attribute name`,
				Suggestions: []model.Suggestion{{Name: []string{`_kMDItemDisplayNameWithExtensions`}}, {Name: []string{`kMDItemContentCreationDate`}}, {Name: []string{`kMDItemContentCreationDate_Ranking`}}, {Name: []string{`kMDItemContentModificationDate`}}, {Name: []string{`kMDItemContentModificationDate_Ranking`}}, {Name: []string{`kMDItemContentType`}}, {Name: []string{`kMDItemContentTypeTree`}}, {Name: []string{`kMDItemDateAdded`}}, {Name: []string{`kMDItemDateAdded_Ranking`}}, {Name: []string{`kMDItemDisplayName`}}, {Name: []string{`kMDItemDocumentIdentifier`}}, {Name: []string{`kMDItemFSContentChangeDate`}}, {Name: []string{`kMDItemFSCreationDate`}}, {Name: []string{`kMDItemFSCreatorCode`}}, {Name: []string{`kMDItemFSFinderFlags`}}, {Name: []string{`kMDItemFSHasCustomIcon`}}, {Name: []string{`kMDItemFSInvisible`}}, {Name: []string{`kMDItemFSIsExtensionHidden`}}, {Name: []string{`kMDItemFSIsStationery`}}, {Name: []string{`kMDItemFSLabel`}}, {Name: []string{`kMDItemFSName`}}, {Name: []string{`kMDItemFSNodeCount`}}, {Name: []string{`kMDItemFSOwnerGroupID`}}, {Name: []string{`kMDItemFSOwnerUserID`}}, {Name: []string{`kMDItemFSSize`}}, {Name: []string{`kMDItemFSTypeCode`}}, {Name: []string{`kMDItemInterestingDate_Ranking`}}, {Name: []string{`kMDItemKind`}}, {Name: []string{`kMDItemLogicalSize`}}, {Name: []string{`kMDItemPhysicalSize`}}},
			}},
			ExclusiveOn: []string{"--plist", "-plist"},
		}, {
			Name:        []string{"--raw", "-raw"},
			Description: `Print raw attribute data in the order that was requested. Fields will be separated with a ASCII NUL character, suitable for piping to xargs(1) -0`,
			ExclusiveOn: []string{"--plist", "-plist"},
		}, {
			Name:        []string{"--nullMarker", "-nullMarker"},
			Description: `Sets a marker string to be used when a requested attribute is null. Only used in -raw mode.  Default is '(null)'`,
			ExclusiveOn: []string{"--plist", "-plist"},
		}, {
			Name:        []string{"--plist", "-plist"},
			Description: `Output attributes in XML format to file. Use - to write to stdout option. Incompatible with options -raw, -nullMarker, and -name`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "stdout or file",
				Description: `XML output location`,
				Suggestions: []model.Suggestion{{
					Name:        []string{`-`},
					Description: `Writes to stdout`,
				}},
			}, {
				Templates:   []model.Template{model.TemplateFilepaths},
				Name:        "file",
				Description: `File to read from`,
			}},
			ExclusiveOn: []string{"--raw", "-raw", "--nullMarker", "-nullMarker", "--name", "-name"},
		}},
	}
}
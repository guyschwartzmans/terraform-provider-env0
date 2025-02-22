package env0

import (
	"errors"
	"regexp"
	"testing"

	"github.com/env0/terraform-provider-env0/client"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestUnitModuleResource(t *testing.T) {
	resourceType := "env0_module"
	resourceName := "test"
	resourceNameImport := resourceType + "." + resourceName
	accessor := resourceAccessor(resourceType, resourceName)

	author := client.User{
		Name: "user0",
	}

	module := client.Module{
		ModuleName:     "name1",
		ModuleProvider: "provider1",
		Repository:     "repository1",
		Description:    "description1",
		TokenId:        "tokenid",
		TokenName:      "tokenname",
		IsGitlab:       false,
		OrganizationId: "org1",
		Author:         author,
		AuthorId:       "author1",
		Id:             uuid.NewString(),
		Path:           "path1",
	}

	updatedModule := client.Module{
		ModuleName:         "name2",
		ModuleProvider:     "provider1",
		Repository:         "repository1",
		Description:        "description1",
		BitbucketClientKey: stringPtr("1234"),
		OrganizationId:     "org1",
		Author:             author,
		AuthorId:           "author1",
		Id:                 module.Id,
		Path:               "path2",
	}

	t.Run("Success", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     module.ModuleName,
						"module_provider": module.ModuleProvider,
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
						"path":            module.Path,
					}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(accessor, "id", module.Id),
						resource.TestCheckResourceAttr(accessor, "module_name", module.ModuleName),
						resource.TestCheckResourceAttr(accessor, "module_provider", module.ModuleProvider),
						resource.TestCheckResourceAttr(accessor, "repository", module.Repository),
						resource.TestCheckResourceAttr(accessor, "description", module.Description),
						resource.TestCheckResourceAttr(accessor, "token_id", module.TokenId),
						resource.TestCheckResourceAttr(accessor, "token_name", module.TokenName),
						resource.TestCheckResourceAttr(accessor, "path", module.Path),
					),
				},
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":          updatedModule.ModuleName,
						"module_provider":      updatedModule.ModuleProvider,
						"repository":           updatedModule.Repository,
						"description":          updatedModule.Description,
						"bitbucket_client_key": *updatedModule.BitbucketClientKey,
						"path":                 updatedModule.Path,
					}),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(accessor, "id", updatedModule.Id),
						resource.TestCheckResourceAttr(accessor, "module_name", updatedModule.ModuleName),
						resource.TestCheckResourceAttr(accessor, "module_provider", updatedModule.ModuleProvider),
						resource.TestCheckResourceAttr(accessor, "repository", updatedModule.Repository),
						resource.TestCheckResourceAttr(accessor, "description", updatedModule.Description),
						resource.TestCheckResourceAttr(accessor, "token_id", ""),
						resource.TestCheckResourceAttr(accessor, "token_name", ""),
						resource.TestCheckResourceAttr(accessor, "bitbucket_client_key", *updatedModule.BitbucketClientKey),
						resource.TestCheckResourceAttr(accessor, "path", updatedModule.Path),
					),
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {
			mock.EXPECT().ModuleCreate(client.ModuleCreatePayload{
				ModuleName:     module.ModuleName,
				ModuleProvider: module.ModuleProvider,
				Repository:     module.Repository,
				Description:    module.Description,
				TokenId:        module.TokenId,
				TokenName:      module.TokenName,
				Path:           module.Path,
			}).Times(1).Return(&module, nil)

			mock.EXPECT().ModuleUpdate(updatedModule.Id, client.ModuleUpdatePayload{
				ModuleName:           updatedModule.ModuleName,
				ModuleProvider:       updatedModule.ModuleProvider,
				Repository:           updatedModule.Repository,
				Description:          updatedModule.Description,
				TokenId:              "",
				TokenName:            "",
				IsGitlab:             false,
				GithubInstallationId: nil,
				BitbucketClientKey:   *updatedModule.BitbucketClientKey,
				Path:                 updatedModule.Path,
			}).Times(1).Return(&updatedModule, nil)

			gomock.InOrder(
				mock.EXPECT().Module(module.Id).Times(2).Return(&module, nil),
				mock.EXPECT().Module(module.Id).Times(1).Return(&updatedModule, nil),
			)

			mock.EXPECT().ModuleDelete(module.Id).Times(1)
		})
	})

	t.Run("Create Failure", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     module.ModuleName,
						"module_provider": module.ModuleProvider,
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
					}),
					ExpectError: regexp.MustCompile("could not create module: error"),
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {
			mock.EXPECT().ModuleCreate(client.ModuleCreatePayload{
				ModuleName:     module.ModuleName,
				ModuleProvider: module.ModuleProvider,
				Repository:     module.Repository,
				Description:    module.Description,
				TokenId:        module.TokenId,
				TokenName:      module.TokenName,
			}).Times(1).Return(nil, errors.New("error"))
		})
	})

	t.Run("Create Failure - Invalid ModuleName", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     "bad!!name^",
						"module_provider": module.ModuleProvider,
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
					}),
					ExpectError: regexp.MustCompile("must match pattern"),
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {})
	})

	t.Run("Create Failure - Invalid ModuleProvider", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     module.ModuleName,
						"module_provider": "ab_c",
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
					}),
					ExpectError: regexp.MustCompile("must match pattern"),
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {})
	})

	t.Run("Update Failure", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     module.ModuleName,
						"module_provider": module.ModuleProvider,
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
						"path":            module.Path,
					}),
				},
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":          updatedModule.ModuleName,
						"module_provider":      updatedModule.ModuleProvider,
						"repository":           updatedModule.Repository,
						"description":          updatedModule.Description,
						"bitbucket_client_key": *updatedModule.BitbucketClientKey,
						"path":                 updatedModule.Path,
					}),
					ExpectError: regexp.MustCompile("could not update module: error"),
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {
			mock.EXPECT().ModuleCreate(client.ModuleCreatePayload{
				ModuleName:     module.ModuleName,
				ModuleProvider: module.ModuleProvider,
				Repository:     module.Repository,
				Description:    module.Description,
				TokenId:        module.TokenId,
				TokenName:      module.TokenName,
				Path:           module.Path,
			}).Times(1).Return(&module, nil)

			mock.EXPECT().ModuleUpdate(updatedModule.Id, client.ModuleUpdatePayload{
				ModuleName:           updatedModule.ModuleName,
				ModuleProvider:       updatedModule.ModuleProvider,
				Repository:           updatedModule.Repository,
				Description:          updatedModule.Description,
				TokenId:              "",
				TokenName:            "",
				IsGitlab:             false,
				GithubInstallationId: nil,
				BitbucketClientKey:   *updatedModule.BitbucketClientKey,
				Path:                 updatedModule.Path,
			}).Times(1).Return(nil, errors.New("error"))

			mock.EXPECT().Module(module.Id).Times(2).Return(&module, nil)
			mock.EXPECT().ModuleDelete(module.Id).Times(1)
		})
	})

	t.Run("Import By Name", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     module.ModuleName,
						"module_provider": module.ModuleProvider,
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
						"path":            module.Path,
					}),
				},
				{
					ResourceName:      resourceNameImport,
					ImportState:       true,
					ImportStateId:     module.ModuleName,
					ImportStateVerify: true,
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {
			mock.EXPECT().ModuleCreate(client.ModuleCreatePayload{
				ModuleName:     module.ModuleName,
				ModuleProvider: module.ModuleProvider,
				Repository:     module.Repository,
				Description:    module.Description,
				TokenId:        module.TokenId,
				TokenName:      module.TokenName,
				Path:           module.Path,
			}).Times(1).Return(&module, nil)
			mock.EXPECT().Modules().Times(1).Return([]client.Module{module}, nil)
			mock.EXPECT().Module(module.Id).Times(2).Return(&module, nil)
			mock.EXPECT().ModuleDelete(module.Id).Times(1)
		})
	})

	t.Run("Import By Id", func(t *testing.T) {
		testCase := resource.TestCase{
			Steps: []resource.TestStep{
				{
					Config: resourceConfigCreate(resourceType, resourceName, map[string]interface{}{
						"module_name":     module.ModuleName,
						"module_provider": module.ModuleProvider,
						"repository":      module.Repository,
						"description":     module.Description,
						"token_id":        module.TokenId,
						"token_name":      module.TokenName,
						"path":            module.Path,
					}),
				},
				{
					ResourceName:      resourceNameImport,
					ImportState:       true,
					ImportStateId:     module.Id,
					ImportStateVerify: true,
				},
			},
		}

		runUnitTest(t, testCase, func(mock *client.MockApiClientInterface) {
			mock.EXPECT().ModuleCreate(client.ModuleCreatePayload{
				ModuleName:     module.ModuleName,
				ModuleProvider: module.ModuleProvider,
				Repository:     module.Repository,
				Description:    module.Description,
				TokenId:        module.TokenId,
				TokenName:      module.TokenName,
				Path:           module.Path,
			}).Times(1).Return(&module, nil)
			mock.EXPECT().Module(module.Id).Times(3).Return(&module, nil)
			mock.EXPECT().ModuleDelete(module.Id).Times(1)
		})
	})

}

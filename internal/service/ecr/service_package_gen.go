// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAuthorizationToken,
			TypeName: "aws_ecr_authorization_token",
		},
		{
			Factory:  DataSourceImage,
			TypeName: "aws_ecr_image",
		},
		{
			Factory:  DataSourceRepository,
			TypeName: "aws_ecr_repository",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLifecyclePolicy,
			TypeName: "aws_ecr_lifecycle_policy",
		},
		{
			Factory:  ResourcePullThroughCacheRule,
			TypeName: "aws_ecr_pull_through_cache_rule",
		},
		{
			Factory:  ResourceRegistryPolicy,
			TypeName: "aws_ecr_registry_policy",
		},
		{
			Factory:  ResourceRegistryScanningConfiguration,
			TypeName: "aws_ecr_registry_scanning_configuration",
		},
		{
			Factory:  ResourceReplicationConfiguration,
			TypeName: "aws_ecr_replication_configuration",
		},
		{
			Factory:  ResourceRepository,
			TypeName: "aws_ecr_repository",
		},
		{
			Factory:  ResourceRepositoryPolicy,
			TypeName: "aws_ecr_repository_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECR
}

var ServicePackage = &servicePackage{}

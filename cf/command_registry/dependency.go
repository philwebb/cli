package command_registry

import (
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/i18n/detection"
	"github.com/cloudfoundry/cli/cf/terminal"
)

type Dependency struct {
	Ui          terminal.UI
	Config      core_config.ReadWriter
	RepoLocator api.RepositoryLocator
	Detector    detection.Detector
}

// type RepositoryLocator struct {
// 	AuthRepo                        authentication.AuthenticationRepository
// 	CurlRepo                        api.CurlRepository
// 	EndpointRepo                    api.RemoteEndpointRepository
// 	OrganizationRepo                organizations.CloudControllerOrganizationRepository
// 	QuotaRepo                       quotas.CloudControllerQuotaRepository
// 	SpaceRepo                       spaces.CloudControllerSpaceRepository
// 	AppRepo                         applications.CloudControllerApplicationRepository
// 	AppBitsRepo                     application_bits.CloudControllerApplicationBitsRepository
// 	AppSummaryRepo                  api.CloudControllerAppSummaryRepository
// 	AppInstancesRepo                app_instances.CloudControllerAppInstancesRepository
// 	AppEventsRepo                   app_events.CloudControllerAppEventsRepository
// 	AppFilesRepo                    api_app_files.CloudControllerAppFilesRepository
// 	DomainRepo                      api.CloudControllerDomainRepository
// 	RouteRepo                       api.CloudControllerRouteRepository
// 	StackRepo                       stacks.CloudControllerStackRepository
// 	ServiceRepo                     api.CloudControllerServiceRepository
// 	ServiceBindingRepo              api.CloudControllerServiceBindingRepository
// 	ServiceSummaryRepo              api.CloudControllerServiceSummaryRepository
// 	UserRepo                        api.CloudControllerUserRepository
// 	PasswordRepo                    password.CloudControllerPasswordRepository
// 	LogsRepo                        api.LogsRepository
// 	LogsNoaaRepo                    api.LogsNoaaRepository
// 	AuthTokenRepo                   api.CloudControllerServiceAuthTokenRepository
// 	ServiceBrokerRepo               api.CloudControllerServiceBrokerRepository
// 	ServicePlanRepo                 api.CloudControllerServicePlanRepository
// 	ServicePlanVisibilityRepo       api.ServicePlanVisibilityRepository
// 	UserProvidedServiceInstanceRepo api.CCUserProvidedServiceInstanceRepository
// 	BuildpackRepo                   api.CloudControllerBuildpackRepository
// 	BuildpackBitsRepo               api.CloudControllerBuildpackBitsRepository
// 	SecurityGroupRepo               security_groups.SecurityGroupRepo
// 	StagingSecurityGroupRepo        staging.StagingSecurityGroupsRepo
// 	RunningSecurityGroupRepo        running.RunningSecurityGroupsRepo
// 	SecurityGroupSpaceBinder        securitygroupspaces.SecurityGroupSpaceBinder
// 	SpaceQuotaRepo                  space_quotas.SpaceQuotaRepository
// 	FeatureFlagRepo                 feature_flags.FeatureFlagRepository
// 	EnvironmentVariableGroupRepo    environment_variable_groups.EnvironmentVariableGroupsRepository
// 	CopyAppSourceRepo               copy_application_source.CopyApplicationSourceRepository
// }

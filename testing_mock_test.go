package marathon

import (
	"net/url"
	"time"

	"github.com/dmajere/go-marathon"
	"github.com/stretchr/testify/mock"
)

type MarathonMock struct {
	mock.Mock
}

func (m *MarathonMock) ListApplications(v url.Values) ([]string, error) {
	args := m.Called(v)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MarathonMock) ApplicationVersions(name string) (*marathon.ApplicationVersions, error) {
	args := m.Called(name)
	return args.Get(0).(*marathon.ApplicationVersions), args.Error(1)
}

func (m *MarathonMock) HasApplicationVersion(name, version string) (bool, error) {
	args := m.Called(name, version)
	return args.Bool(0), args.Error(1)
}

func (m *MarathonMock) SetApplicationVersion(name string, version *marathon.ApplicationVersion) (*marathon.DeploymentID, error) {
	args := m.Called(name, version)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}

func (m *MarathonMock) ApplicationOK(name string) (bool, error) {
	args := m.Called(name)
	return args.Bool(0), args.Error(1)
}

func (m *MarathonMock) CreateApplication(application *marathon.Application) (*marathon.Application, error) {
	args := m.Called(application)
	return args.Get(0).(*marathon.Application), args.Error(1)
}

func (m *MarathonMock) DeleteApplication(name string) (*marathon.DeploymentID, error) {
	args := m.Called(name)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}

func (m *MarathonMock) UpdateApplication(application *marathon.Application, force bool) (*marathon.DeploymentID, error) {
	args := m.Called(application, force)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}

func (m *MarathonMock) ApplicationDeployments(name string) ([]*marathon.DeploymentID, error) {
	args := m.Called(name)
	return args.Get(0).([]*marathon.DeploymentID), args.Error(1)
}

func (m *MarathonMock) ScaleApplicationInstances(name string, instances int, force bool) (*marathon.DeploymentID, error) {
	args := m.Called(name, instances, force)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}
func (m *MarathonMock) RestartApplication(name string, force bool) (*marathon.DeploymentID, error) {
	args := m.Called(name, force)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}

func (m *MarathonMock) Applications(v url.Values) (*marathon.Applications, error) {
	args := m.Called(v)
	return args.Get(0).(*marathon.Applications), args.Error(1)
}
func (m *MarathonMock) Application(name string) (*marathon.Application, error) {
	args := m.Called(name)
	return args.Get(0).(*marathon.Application), args.Error(1)
}
func (m *MarathonMock) ApplicationByVersion(name, version string) (*marathon.Application, error) {
	args := m.Called(name, version)
	return args.Get(0).(*marathon.Application), args.Error(1)
}
func (m *MarathonMock) WaitOnApplication(name string, timeout time.Duration) error {
	args := m.Called(name, timeout)
	return args.Error(0)
}
func (m *MarathonMock) Tasks(application string) (*marathon.Tasks, error) {
	args := m.Called(application)
	return args.Get(0).(*marathon.Tasks), args.Error(1)
}
func (m *MarathonMock) AllTasks(opts *marathon.AllTasksOpts) (*marathon.Tasks, error) {
	args := m.Called(opts)
	return args.Get(0).(*marathon.Tasks), args.Error(1)
}

func (m *MarathonMock) TaskEndpoints(name string, port int, healthCheck bool) ([]string, error) {
	args := m.Called(name, port, healthCheck)
	return args.Get(0).([]string), args.Error(1)
}
func (m *MarathonMock) KillApplicationTasks(applicationID string, opts *marathon.KillApplicationTasksOpts) (*marathon.Tasks, error) {
	args := m.Called(applicationID, opts)
	return args.Get(0).(*marathon.Tasks), args.Error(1)
}
func (m *MarathonMock) KillTask(taskID string, opts *marathon.KillTaskOpts) (*marathon.Task, error) {
	args := m.Called(taskID, opts)
	return args.Get(0).(*marathon.Task), args.Error(1)
}

func (m *MarathonMock) KillTasks(taskIDs []string, opts *marathon.KillTaskOpts) error {
	args := m.Called(taskIDs, opts)
	return args.Error(0)
}
func (m *MarathonMock) Groups() (*marathon.Groups, error) {
	args := m.Called()
	return args.Get(0).(*marathon.Groups), args.Error(1)
}
func (m *MarathonMock) Group(name string) (*marathon.Group, error) {
	args := m.Called(name)
	return args.Get(0).(*marathon.Group), args.Error(1)
}
func (m *MarathonMock) CreateGroup(group *marathon.Group) error {
	args := m.Called(group)
	return args.Error(0)
}
func (m *MarathonMock) DeleteGroup(name string) (*marathon.DeploymentID, error) {
	args := m.Called(name)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}
func (m *MarathonMock) UpdateGroup(id string, group *marathon.Group) (*marathon.DeploymentID, error) {
	args := m.Called(id, group)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}
func (m *MarathonMock) HasGroup(name string) (bool, error) {
	args := m.Called(name)
	return args.Bool(0), args.Error(1)
}
func (m *MarathonMock) WaitOnGroup(name string, timeout time.Duration) error {
	args := m.Called(name, timeout)
	return args.Error(0)
}
func (m *MarathonMock) Deployments() ([]*marathon.Deployment, error) {
	args := m.Called()
	return args.Get(0).([]*marathon.Deployment), args.Error(1)
}
func (m *MarathonMock) DeleteDeployment(id string, force bool) (*marathon.DeploymentID, error) {
	args := m.Called(id, force)
	return args.Get(0).(*marathon.DeploymentID), args.Error(1)
}
func (m *MarathonMock) HasDeployment(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
func (m *MarathonMock) WaitOnDeployment(id string, timeout time.Duration) error {
	args := m.Called(id, timeout)
	return args.Error(0)
}
func (m *MarathonMock) Subscriptions() (*marathon.Subscriptions, error) {
	args := m.Called()
	return args.Get(0).(*marathon.Subscriptions), args.Error(1)
}

func (m *MarathonMock) AddEventsListener(channel marathon.EventsChannel, filter int) error {
	args := m.Called(channel, filter)
	return args.Error(0)
}
func (m *MarathonMock) RemoveEventsListener(channel marathon.EventsChannel) {
	m.Called(channel)
}

func (m *MarathonMock) Unsubscribe(callbackURL string) error {
	args := m.Called(callbackURL)
	return args.Error(0)
}
func (m *MarathonMock) GetMarathonURL() string {
	args := m.Called()
	return args.String(0)
}
func (m *MarathonMock) Ping() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}
func (m *MarathonMock) Info() (*marathon.Info, error) {
	args := m.Called()
	return args.Get(0).(*marathon.Info), args.Error(1)
}
func (m *MarathonMock) Leader() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
func (m *MarathonMock) AbdicateLeader() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

package config

type MockConfig struct {
	JenkinsURL   string
	JenkinsUser  string
	JenkinsToken string
}

func (m *MockConfig) GetJenkinsURL() string {
	return m.JenkinsURL
}

func (m *MockConfig) GetJenkinsUser() string {
	return m.JenkinsUser
}

func (m *MockConfig) GetJenkinsToken() string {
	return m.JenkinsToken
}

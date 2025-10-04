package types

type Instance struct {
	UUID string `json:"uuid" mapstructure:"uuid"`
	Name string `json:"name" mapstructure:"name"`
	Version string `json:"version" mapstructure:"version"`
	Changelog string `json:"changelog" mapstructure:"changelog"`
	InstanceType string  `json:"instance_type" mapstructure:"instance_type"`
	Containers []Container `json:"containers" mapstructure:"containers"`
}

type Container struct {
	ContentType string `json:"content_type" mapstructure:"content_type"`
	Content []Content `json:"content" mapstructure:"content"`
}

type Content struct {
	File string `json:"file" mapstructure:"file"`
	Url string `json:"url" mapstructure:"url"`
}

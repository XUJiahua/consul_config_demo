module consul_config_demo

go 1.14

require (
	github.com/fsnotify/fsnotify v1.4.7
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/viper v1.6.3
)

replace github.com/spf13/viper v1.6.3 => github.com/XUJiahua/viper v1.6.3-0.20200422030214-fdcfc3ee2bdc

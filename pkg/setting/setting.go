package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	//the file name of config file
	vp.SetConfigName("config")
	//the path that config file in
	vp.AddConfigPath("configs/")
	//the type of config file
	vp.SetConfigType("yaml")

	//read the config.yaml in configs/
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, nil
}

//@params k : key in config
//@Params v : read value to any interface

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}

	return nil
}

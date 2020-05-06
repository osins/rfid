package env

import (
	"os"

	"github.com/joho/godotenv"
)

// NewENV 创建全局变量配置对象
func NewENV() ENV {
	return &env{}
}

// ENV env 全局变量配置接口
type ENV interface {
	Load() error
}

type env struct {
}

// Load 载入全局变量配置文档
func (e *env) Load() error {
	env, err := e.getEnvPath()
	if err != nil {
		return err
	}

	return godotenv.Load(env)
}

// env exists 判断配置文件是否存在
func (e *env) getEnvPath() (env string, err error) {
	if !e.exists(".env") {
		if e.exists("/usr/local/rfid/xlslr5603/.env") {
			return "/usr/local/rfid/xlslr5603/.env", nil
		} else {
			return "E:\\go\\github.com\\wangsying\\rfid\\xlslr5603\\.env", nil
		}
	} else {
		return ".env", nil
	}
}

// exists 判断文件是否存在
func (e *env) exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

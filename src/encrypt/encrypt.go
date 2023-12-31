package encrypt

import (
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const phaseName = "encrypt"

type InitEncryptPhase struct {
	Multilevel map[string]interface{}
}

func (c *InitEncryptPhase) Init() error {
	enable := c.Multilevel["enable"]
	if enable != nil && enable.(bool) {
		p, err := exec.LookPath(fmt.Sprintf("encrypt_%s_%s", runtime.GOOS, runtime.GOARCH))
		if err != nil {
			logger.Log.Errorf("Failed to lookpath passwordencrypt_%s_%s: %s", runtime.GOOS, runtime.GOARCH, err.Error())
			return err
		}
		secret := c.Multilevel["secret"].(string)
		parts := c.Multilevel["parts"].([]interface{})
		args := []string{"decrypt", "-t", secret}
		args = append(args, "-p")
		for i := range parts {
			args = append(args, parts[i].(string))
		}
		cmd := exec.Command(p, args...)
		cmd.Env = os.Environ()

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			logger.Log.Errorf("Failed to get stdout: %s", err.Error())
			return err
		}
		defer func(stdout io.ReadCloser) {
			err := stdout.Close()
			if err != nil {
				logger.Log.Errorf("Failed to close stdout: %s", err.Error())
				return
			}
		}(stdout)
		if err := cmd.Start(); err != nil {
			logger.Log.Errorf("Failed to start cmd: %s", err.Error())
			return err
		}
		bs, err := ioutil.ReadAll(stdout)
		if err != nil {
			logger.Log.Errorf("Failed to read stdout info: %s", err.Error())
			return err
		}
		if err := cmd.Wait(); err != nil {
			logger.Log.Errorf("Failed to wait stdout info: %s", err.Error())
			return err
		}
		str := string(bs)
		str = strings.TrimPrefix(str, "\n")

		viper.Set("encrypt.key", string(bs))
	}
	return nil
}

func (c *InitEncryptPhase) PhaseName() string {
	return phaseName
}

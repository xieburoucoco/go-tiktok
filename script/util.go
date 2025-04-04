package script

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"go-tiktok/consts"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed js/script.js
var xbScriptContent string

type INodeScriptUtil interface {
	GetXbScript(ctx context.Context) (NodeScript, error)
	GetXbValueByNodeCmd(ctx context.Context, url, userAgent string) (string, error)
}

type SNodeScriptUtil struct{}

type NodeScript struct {
	Code string
	Path string
}

func New() INodeScriptUtil {
	return &SNodeScriptUtil{}
}

func (s *SNodeScriptUtil) GetXbValueByNodeCmd(ctx context.Context, url, userAgent string) (string, error) {
	value := ""
	if err := CheckNodeEnv(ctx); err != nil {
		return value, err
	}
	xbScript, err := s.GetXbScript(ctx)
	if err != nil {
		return value, err
	}
	cmd := exec.Command("node", xbScript.Path, url, userAgent)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	if err := cmd.Run(); err != nil {
		return value, fmt.Errorf("node execution failed: %w\nStderr: %s", err, stderrBuf.String())
	}
	rawOutput := stdoutBuf.String()
	value = strings.TrimSpace(rawOutput)
	if value == "" {
		return "", fmt.Errorf("empty output from script")
	}

	return value, nil
}

func (s *SNodeScriptUtil) GetXbScript(ctx context.Context) (NodeScript, error) {
	res := NodeScript{}
	workingDir, err := os.Getwd()
	if err != nil {
		return res, err
	}
	tempPath := filepath.Join(workingDir, consts.XB_SCRIPT_PATH)
	scriptDir := filepath.Dir(consts.XB_SCRIPT_PATH)
	if err := os.MkdirAll(scriptDir, 0755); err != nil {
		return res, fmt.Errorf("failed to create script directory: %w", err)
	}
	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		err := os.WriteFile(tempPath, []byte(xbScriptContent), 0777)
		if err != nil {
			return res, err
		}
	}

	res.Path = tempPath
	res.Code = xbScriptContent
	return res, nil
}

func CheckNodeEnv(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "node", "-v")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("node.js environment not found or not working: %v", err)
	}
	return nil
}

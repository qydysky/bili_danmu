package roomsignal

import (
	"fmt"
	"os/exec"
	"strings"
	"sync/atomic"
	"time"

	comp "github.com/qydysky/part/component2"
	pe "github.com/qydysky/part/errors/v2"
	fc "github.com/qydysky/part/funcCtrl"
)

type Inter interface {
	FiliterRoomId(configs any, roomId int) interface {
		LoginChange(isLogin bool) error
		Begin() error
		Fin() error
		TitleChange(oldTitle, newTitle string) error
		AfterRec(videoType string, liveDir string, duration time.Duration) error
	}
}

func init() {
	comp.RegisterOrPanic[Inter](`roomSignal`, impl{})
}

var ActRoomSignal = pe.Action[struct {
	NoRoomFound pe.Error
}](`ActRoomSignal`)

type impl struct {
	vm map[string]any
}

func (t impl) FiliterRoomId(configs any, roomId int) interface {
	LoginChange(isLogin bool) error
	Begin() error
	Fin() error
	TitleChange(oldTitle, newTitle string) error
	AfterRec(videoType string, liveDir string, duration time.Duration) error
} {
	if v, ok := configs.([]any); ok {
		for i := 0; i < len(v); i++ {
			if vm, ok := v[i].(map[string]any); ok {
				if roomid, ok := vm["roomid"].(float64); ok && int(roomid) == roomId {
					t.vm = vm
					break
				}
			}
		}
	}
	return &t
}

var (
	loginStateInit atomic.Bool
	lastLoginState atomic.Bool
	loginRunning   fc.SkipFunc
)

func (t *impl) LoginChange(isLogin bool) error {
	if loginStateInit.CompareAndSwap(false, true) {
		lastLoginState.Store(isLogin)
		return nil
	} else if lastLoginState.Swap(isLogin) == isLogin {
		return nil
	}
	defer loginRunning.UnSet()
	if loginRunning.NeedSkip() {
		return nil
	}

	var (
		loginChange, _ = t.vm["loginChange"].([]any)
		runDir, _      = t.vm["runDir"].(string)
	)
	if len(loginChange) < 1 {
		return nil
	}
	var cmds []string
	for i := 0; i < len(loginChange); i++ {
		if cmd, ok := loginChange[i].(string); ok && cmd != "" {
			cmds = append(cmds, cmd)
		}
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Dir = runDir
	return cmd.Run()
}

var beginRunning fc.SkipFunc

func (t *impl) Begin() error {
	defer beginRunning.UnSet()
	if beginRunning.NeedSkip() {
		return nil
	}
	var (
		begin, _  = t.vm["begin"].([]any)
		runDir, _ = t.vm["runDir"].(string)
	)
	if len(begin) < 1 {
		return nil
	}
	var cmds []string
	for i := 0; i < len(begin); i++ {
		if cmd, ok := begin[i].(string); ok && cmd != "" {
			cmds = append(cmds, cmd)
		}
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Dir = runDir
	return cmd.Run()
}

var finRunning fc.SkipFunc

func (t *impl) Fin() error {
	defer finRunning.UnSet()
	if finRunning.NeedSkip() {
		return nil
	}
	var (
		fin, _    = t.vm["fin"].([]any)
		runDir, _ = t.vm["runDir"].(string)
	)
	if len(fin) < 1 {
		return nil
	}
	var cmds []string
	for i := 0; i < len(fin); i++ {
		if cmd, ok := fin[i].(string); ok && cmd != "" {
			cmds = append(cmds, cmd)
		}
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Dir = runDir
	return cmd.Run()
}

func (t *impl) TitleChange(oldTitle, newTitle string) error {
	var (
		titleChange, _ = t.vm["titleChange"].([]any)
		runDir, _      = t.vm["runDir"].(string)
	)
	if len(titleChange) < 1 {
		return nil
	}
	var cmds []string
	for i := 0; i < len(titleChange); i++ {
		if cmd, ok := titleChange[i].(string); ok && cmd != "" {
			cmd = strings.ReplaceAll(cmd, `{oldTitle}`, oldTitle)
			cmd = strings.ReplaceAll(cmd, `{newTitle}`, newTitle)
			cmds = append(cmds, cmd)
		}
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Dir = runDir
	return cmd.Run()
}

func (t *impl) AfterRec(videoType string, liveDir string, duration time.Duration) error {
	var (
		after, _  = t.vm["afterRec"].([]any)
		runDir, _ = t.vm["runDir"].(string)
	)
	if len(after) < 1 {
		return nil
	}
	var cmds []string
	for i := 0; i < len(after); i++ {
		if cmd, ok := after[i].(string); ok && cmd != "" {
			cmd = strings.ReplaceAll(cmd, `{type}`, videoType)
			cmd = strings.ReplaceAll(cmd, `{liveDir}`, liveDir)
			cmd = strings.ReplaceAll(cmd, `{recSec}`, fmt.Sprintf("%0.0f", duration.Seconds()))
			cmds = append(cmds, cmd)
		}
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Dir = runDir
	return cmd.Run()
}

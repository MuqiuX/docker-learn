package cmd

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use: "run",
	Short: "",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		// 重新启动自己到一个新的进程，
		// 也就是dkr
		// 执行init命令, 进行容器的初始化，并执行传递的命令
		re := exec.Command("/proc/self/exe", append([]string{"init"}, args...)...)

		// 将新进程的输入输出，绑定到操作系统
		re.Stdin = os.Stdin
		re.Stdout = os.Stdout
		re.Stderr = os.Stderr

		// 设置新进程的部分namespace初始化
		// 独立的进程pids
		// 独立的uts
		// 为mount提供新命名空间
		re.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS,
			// 阻止与主机的命名空间共享
			Unshareflags: syscall.CLONE_NEWNS,
		}

		// 开始执行
		if err := re.Run(); err != nil {
			log.Fatal("ERROR: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
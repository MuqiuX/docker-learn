package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/MuqiuX/dkr/core/utils"
	"github.com/spf13/cobra"
)

// 用于初始化一个容器
var initCmd = &cobra.Command{
	Use: "init",
	Short: "",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		// 设置新容器主机名
		utils.Must(syscall.Sethostname([]byte("container")))
		// 根目录
		utils.Must(syscall.Chroot("/"))
		// 默认目录
		utils.Must(syscall.Chdir("/"))

		// 挂载proc
		syscall.Mount("proc", "proc", "proc", 0, "")

		re := exec.Command(args[0], args[1:]...)

		re.Stdin = os.Stdin
		re.Stdout = os.Stdout
		re.Stderr = os.Stderr

		if err := re.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hgen",
	Short: "a tool for wood to generate go code with mysql data table",
	Long:  "a tool for wood to generate go code with mysql data table",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if password == "" {
			fmt.Println("no such flag --password")
			return
		}
		if dbname == "" {
			fmt.Println("no such flag --dbname")
			return
		}
		if TableName == "" {
			fmt.Println("no such flag --TableName")
			return
		}
		var err error
		mssql, err = NewMysql(username, dbname, ip, password, port)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 生成变量对应的值
		generateVar()
		// 写入model文件
		modelstring, err := generateModel()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = WriteFile("backend/model", TableName+".go", modelstring)
		if err != nil {
			fmt.Println(err)
			return
		}

		if backend {
			apistring := generateApi()
			err = WriteFile("backend/api/"+TagName, TableName+".go", apistring)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		// admin
		if admin {
			// api
			vueapistring := generateAdminApi()
			err = WriteFile("admin/src/api/"+TagName, TableName+".ts", vueapistring)
			if err != nil {
				fmt.Println(err)
				return
			}
			adminmodel, err := generateAdminStruct()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = WriteFile("admin/src/api/"+TagName+"/model", TableName+".ts", adminmodel)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	cfgFile  string
	port     int
	username string
	password string
	ip       string
	dbname   string
	admin    bool
	backend  bool
)

func init() {
	initConfig()
	rootCmd.PersistentFlags().IntVar(&port, "port", 3306, "mysql port")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "mysql username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "123456ab", "mysql password")
	rootCmd.PersistentFlags().StringVarP(&TableName, "TableName", "t", "", "table name")
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip")
	rootCmd.PersistentFlags().StringVarP(&dbname, "dbname", "d", "", "dbname")
	rootCmd.PersistentFlags().StringVarP(&TagName, "TagName", "g", "", "TagName")
	rootCmd.PersistentFlags().BoolVarP(&admin, "admin", "a", false, "is admin")
	rootCmd.PersistentFlags().BoolVarP(&backend, "backend", "b", false, "is backend")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("hgen.yaml")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hgen",
	Short: "A tool to generate go code with mysql data table",
	Long:  `A tool to generate go code with mysql data table`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if DbName == "" {
			fmt.Println("no such flag --dbname")
			return
		}
		if TableName == "" {
			fmt.Println("no such flag --tablename")
			return
		}
		var err error
		SQL, err = NewMysql()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 写入model文件
		if IsApi {
			// api
			gostructs, err := generateGoStruct()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = WriteFile("internal/model", TableName+".go", gostructs)
			if err != nil {
				fmt.Println(err)
				return
			}

			goapis := generateGoApi()
			err = WriteFile("internal/app/api/"+TagName, TableName+".go", goapis)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if IsWeb {
			// web
			tsapis, err := generateTsApi()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = WriteFile("web/admin/src/api/"+TagName, TableName+".ts", tsapis)
			if err != nil {
				fmt.Println(err)
				return
			}
			// 模板
			vuetemps := generateVueTemp()
			err = WriteFile("web/admin/src/views/"+TagName+"/"+TableName, "list.vue", vuetemps)
			if err != nil {
				fmt.Println(err)
				return
			}
			// model
			vuemodels, err := generateVueModel()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = WriteFile("web/admin/src/views/"+TagName+"/"+TableName+"/model", "_list.ts", vuemodels)
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
	cfgFile   string
	Dsn       string
	DbName    string
	IsApi     bool
	IsWeb     bool
	TableName string
	TagName   string
)

func init() {
	initConfig()
	rootCmd.PersistentFlags().StringVarP(&Dsn, "DSN", "l", "root:123456ab@tcp(127.0.0.1:3306)%s?charset=utf8mb4&parseTime=True&loc=Local", "mysql dsn")
	rootCmd.PersistentFlags().StringVarP(&DbName, "dbname", "d", "", "dbname")
	rootCmd.PersistentFlags().StringVarP(&TableName, "tablename", "t", "", "table name")
	rootCmd.PersistentFlags().StringVarP(&TableName, "TagName", "g", "", "tag name")
	rootCmd.PersistentFlags().BoolVarP(&IsApi, "api", "a", false, "is api")
	rootCmd.PersistentFlags().BoolVarP(&IsWeb, "web", "w", false, "is web")
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

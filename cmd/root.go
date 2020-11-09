package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

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
		if tablename == "" {
			fmt.Println("no such flag --tablename")
			return
		}
		var err error
		mssql, err = NewMysql(username, dbname, ip, password, port)
		if err != nil {
			fmt.Println(err)
			return
		}
		tables := strings.Split(tablename, ",")
		err = os.MkdirAll("hgen", 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v := range tables {
			f, err := os.Create("hgen/" + v + "_dao.go")
			f.WriteString(`package hgen
			`)
			if err != nil {
				fmt.Println(err)
				return
			}
			daostring, err := generatedao(v)
			if err != nil {
				fmt.Println(err)
				f.Close()
				os.Remove("hgen/" + v + "_dao.go")
				return
			}
			_, err = f.WriteString(daostring)
			if err != nil {
				fmt.Println(err)
				f.Close()
				os.Remove("hgen/" + v + "_dao.go")
				return
			}
			err = f.Sync()
			if err != nil {
				fmt.Println(err)
				f.Close()
				os.Remove("hgen/" + v + "_dao.go")
				return
			}
			f2, err := os.Create("hgen/" + v + "_api.go")
			if err != nil {
				fmt.Println(err)
				return
			}
			f2.WriteString(`package hgen
			`)
			apistring := generateapi(v)
			_, err = f2.WriteString(apistring)
			if err != nil {
				fmt.Println(err)
				f2.Close()
				os.Remove("hgen/" + v + "_api.go")
				return
			}
			err = f2.Sync()
			if err != nil {
				fmt.Println(err)
				f2.Close()
				os.Remove("hgen/" + v + "_api.go")
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
	port      int
	tablename string
	username  string
	password  string
	ip        string
	dbname    string
	tags      string
	comment   bool
	cache     bool
)

func init() {
	rootCmd.PersistentFlags().IntVar(&port, "port", 3306, "mysql port")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "mysql username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "123456ab", "mysql password")
	rootCmd.PersistentFlags().StringVarP(&tablename, "tablename", "t", "", "table names,you yan use , join mult")
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip")
	rootCmd.PersistentFlags().StringVarP(&dbname, "dbname", "d", "", "dbname")
	rootCmd.PersistentFlags().StringVarP(&tags, "tags", "g", "", "tags")
	rootCmd.PersistentFlags().BoolVarP(&comment, "comment", "c", false, "swagger comment")
	rootCmd.PersistentFlags().BoolVarP(&cache, "cache", "e", false, "is cache")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config.yaml")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

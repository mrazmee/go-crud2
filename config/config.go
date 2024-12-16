package config

import (
	"log"

	"github.com/spf13/viper"

)
//field struct disini di cocokan dengan kebutuhan untuk koneksi ke database
type Config struct{
	DBHost string `mapstructure:"DBHost"` //untuk tag disini bertujuan untuk memetakan ke dalam variabel lingkungan
	DBPort string `mapstructure:"DBPort"`
	DBUser string `mapstructure:"DBUser"`
	DBPassword string `mapstructure:"DBPassword"`
	DBName string `mapstructure:"DBName"`
}

func LoadConfig() (config Config, err error){
	viper.AddConfigPath(".") //path tempat dimana project dijalankan, "." disini berarti root folder
	viper.SetConfigName("config") //nama file
	viper.SetConfigType("env") //type filenya

	viper.AutomaticEnv() //mengaktifkan pencarian

	if err = viper.ReadInConfig(); err != nil { //membaca file yang ditemukan
		log.Fatalf("Error reading congfig file %s", err) 
	}

	if err := viper.Unmarshal(&config); err != nil { //mengubah nilai yang ditemukan ke nilai dalam struct
		log.Fatalf("Unable to decode %v", err)
	}

	return
}

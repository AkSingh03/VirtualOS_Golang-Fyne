package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(w fyne.Window){
	
	//Api 

	res , err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=new york&units=metric&APPID=e0a98ac0c57b7e77678958a6376fabdc")
	if err!= nil{

		fmt.Println(err)
	} 

	defer res.Body.Close()

	Body , err := ioutil.ReadAll(res.Body)
	if err!= nil{

		fmt.Println(err)
	} 


	weather , err := UnmarshalWeather(Body)
	if err!= nil{

		fmt.Println(err)
	} 
	 //UI

	img:= canvas.NewImageFromFile("wea.jpg")
	img.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("Weather Details" , color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	label2 := canvas.NewText(fmt.Sprintf("Name %s" , weather.Name) , color.White)
	
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f" + "KM/Hr" , weather.Wind.Speed) , color.White)

	label4 := canvas.NewText(fmt.Sprintf("Temperature %.2f" + "°C" , weather.Main.Temp) , color.White)

	label5 := canvas.NewText(fmt.Sprintf("FeelsLike %.2f"+ "°C" , weather.Main.FeelsLike) , color.White)

	label6 := canvas.NewText(fmt.Sprintf("MaximumTemperature %.2f" + "°C" , weather.Main.TempMax) , color.White)

	label7 := canvas.NewText(fmt.Sprintf("MinimumTemperature %.2f" + "°C" , weather.Main.TempMin) , color.White)

    weatherContainer:=container.NewVBox(
			label1,
			img,
			label2,
			label3,
			label4,
			label5,
			label6,
			label7,
			container.NewGridWithColumns(1,
			),
	)
	

	w.SetContent(container.NewBorder(DeskBtn,nil,nil,nil,weatherContainer),)
	w.Show()
}



// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()


func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}


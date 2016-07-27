package main

import (
	"time"
	"html/template"
)


// ##### Structs ##############################################################

// Represents an "instance" record
type Instance struct {
	Id       	int64 		`db:"id"`
	Domain     	string 		`db:"domain"`
	Host     	string 		`db:"host"`
	Timestamp	time.Time	`db:"timestamp"`
}

// Base fields that all database tables support
type Base struct {
	Id      	int64 		`db:"id"`
	Domain  	string 		`db:"domain"`
	Host    	string 		`db:"host"`
	UtcTime		time.Time	`db:"timestamp"`
	UtcTimeStr	string		`db:"-"`
}

// Represents an "autorun" record
type Autorun struct {
	Id       		int64 			`db:"id"`
	Instance     	int64 			`db:"instance"`
	FilePath    	string 			`db:"file_path"`
	FileName    	string 			`db:"file_name"`
	FileDirectory	string 			`db:"file_directory"`
	Location 		string			`db:"location"`
	LocationStr 	template.HTML	`db:"-"`
	ItemName 		string 			`db:"item_name"`
	Enabled 		bool 			`db:"enabled"`
	Profile 		string 			`db:"profile"`
	LaunchString 	string 			`db:"launch_string"`
	Description 	string 			`db:"description"`
	Company 		string 			`db:"company"`
	Signer 			string 			`db:"signer"`
	VersionNumber 	string 			`db:"version_number"`
	Time 			time.Time		`db:"time"`
	TimeStr 		string			`db:"-"`
	Sha256			string			`db:"sha256"`
	Md5				string			`db:"md5"`
	OtherData		template.HTML	`db:"-"`
}

// Represents an "alert" record
type Alert struct {
	Base
	AutorunId		int64			`db:"autorun_id"`
	Instance     	int64 			`db:"instance"`
	FilePath    	string 			`db:"file_path"`
	FileName    	string 			`db:"file_name"`
	FileDirectory	string 			`db:"file_directory"`
	Location 		string			`db:"location"`
	LocationStr 	template.HTML	`db:"-"`
	ItemName 		string 			`db:"item_name"`
	Enabled 		bool 			`db:"enabled"`
	Profile 		string 			`db:"profile"`
	LaunchString 	string 			`db:"launch_string"`
	Description 	string 			`db:"description"`
	Company 		string 			`db:"company"`
	Signer 			string 			`db:"signer"`
	VersionNumber 	string 			`db:"version_number"`
	Time 			time.Time		`db:"time"`
	TimeStr 		string			`db:"-"`
	Sha256			string			`db:"sha256"`
	Md5				string			`db:"md5"`
	OtherData		template.HTML	`db:"-"`
}

// Represents an "export" record
type Export struct {
	Id       		int64 			`db:"id"`
	DataType 		string 			`db:"data_type"`
	FileName 		string 			`db:"file_name"`
	Updated			time.Time		`db:"updated"`
	OtherData		template.HTML	`db:"-"`
}
## AngularJS
#### Remove the # in the url
``` javascript
$locationProvider.html5Mode(true)
```

## Recept Template
##### Type
* Julmat
* Julgodis
* Förrätt
* Huvudrätt
* Efterrätt
* Frukost
* Dryck
* Övrigt
```json
{
	"ID": "template", 
	"Type": "template",
	"Name": "Template",
	"Description": "template template template",

	"Ingrediences": [
		{
			"Name": "template",
			"Amount": "1 L"
		},
		{
			"Name": "",
			"Amount": "Lag"
		},
		{
			"Name": "template",
			"Amount": "2 L"
		}
	],

	"Making": [
		"Gör så här",
		"Och såhär", 
		"Sedan såhär"
	],


	"Notes": [
		"Detta är min åsikt!"
	],

	"Image": "/recources/images/pic01.jpg"
}
```

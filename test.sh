#!/bin/bash

echo tests | figlet | lolcat

./backend &

token=$(curl --request POST \
  --url http://localhost:8080/login \
  --header 'content-type: application/json' \
  --data '{
	"username": "m0on",
	"passHash": "7f965560c9f2ce126407eda7c7dbbdb75037ef4d"
}')

sleep 2

echo $token

sleep 2

# curl --request GET \
#   --url http://localhost:8080/foods

curl --request POST \
  --url http://localhost:8080/food \
  --header "authorization: $token" \
  --header 'content-type: application/json' \
  --data '{
	"Name": "Cibatta",
	"Category": "Dairy",
	"Visual": "White loaf or bun/roll similar to standard white bread, but is generally handmade so it is more rounded. Has a crust but it is less thick than standard white bread and generally dusted with flour. Inside is cream/tan coloured with larger air pockets than standard white loaves. Lower quality ones will be just the same as white loaves inside.",
	"Texture": "Soft but can be slightly more doughy inside. Different to standard white bread, kind of smoother/less dry inside (but isnt wet/damp or anything). Goes stale faster than standard bread and will get really hard when it does. Toasts well but slightly differently to standard bread",
	"Smell": "Yeasty",
	"Taste": "Very similar to standard white bread but with a slight sourdough like taste (not as strong as standard sourdough breads)",
	"Nutrients": [
		"carbs"
	]
}'

sleep 2

curl --request GET \
  --url http://localhost:8080/food/Cibatta

sleep 2

curl --request POST \
  --url http://localhost:8080/food/edit \
  --header "authorization: Bearer $token" \
  --header 'content-type: application/json' \
  --data '{
	"Name": "Cibatta",
	"Category": "Grains",
	"Visual": "White loaf or bun/roll similar to standard white bread, but is generally handmade so it is more rounded. Has a crust but it is less thick than standard white bread and generally dusted with flour. Inside is cream/tan coloured with larger air pockets than standard white loaves. Lower quality ones will be just the same as white loaves inside.",
	"Texture": "Soft but can be slightly more doughy inside. Different to standard white bread, kind of smoother/less dry inside (but isnt wet/damp or anything). Goes stale faster than standard bread and will get really hard when it does. Toasts well but slightly differently to standard bread",
	"Smell": "Yeasty",
	"Taste": "",
	"Nutrients": [
		"carbs"
	]
}'

sleep 2

curl --request GET \
  --url http://localhost:8080/food/Cibatta

echo SUCCESS | figlet | lolcat
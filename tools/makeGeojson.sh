#!/usr/bin/env bash

#curl -s http://localhost:8088/travels | jq '[.Directories[] | { name: .name, modified: .modified, path: .path, thumb: .thumb, pictures: .items }]'

ALBUMS=( $(curl -s http://localhost:8088/travels | jq -c '[.Directories[] | .name]' | sed 's/"//g' | sed 's/,/ /g' | sed 's/\[//' | sed 's/\]//') )

#echo $ALBUMS

#for ALBUM in "${ALBUMS[@]}";do
#  echo $ALBUM
#  curl -s http://localhost:8088/travels/$ALBUM | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' > $ALBUM-geojson.json
#done

curl -s http://localhost:8088/travels/Colima | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > colima-geojson.json
curl -s http://localhost:8088/travels/Colombia | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > colombia-geojson.json
curl -s http://localhost:8088/travels/MexicoCity | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > mexicocity-geojson.json
curl -s http://localhost:8088/travels/Monterrey-2014-08 | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > monterrey-09-2014-geojson.json
curl -s http://localhost:8088/travels/Monterrey-2015-03| jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > monterrey-03-2015-geojson.json
curl -s http://localhost:8088/travels/Monterrey-2015-07 | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > monterrey-07-2015-geojson.json
curl -s http://localhost:8088/travels/PuertoVallarta | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > puero-vallarta-geojson.json
curl -s http://localhost:8088/travels/Sayulita | jq '[.[][]  | {type: "Feature", geometry: {type: "Point", coordinates: [.exif.long, .exif.lat]}, properties: {name: .name, taken: .exif.taken, thumb: .thumb, path: .path} }]' | sed '2,17d' > sayulita-geojson.json

exit 0

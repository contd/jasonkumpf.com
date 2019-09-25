---
title: Functional Javascript Snips
lang: en-US
tags:
	- javascript
---

# {{ $page.title }}

While working my way through learning RxJS one of the recommended links on their site lead me to a pretty exhaustive set of samples.  Most of them I was already very familiar with which get you familiar with array map and reduce.  Below is where I will show the ones I liked so much I wanted to write them down and annotate a bit.

## Converting Arrays to Trees

We have two arrays of objects (`lists` and `videos`) which have a relationship by `id`.  Using a map to iterate through each list item.  For each list item we use `filter` to itterate each video and returning only the ones with a mathcing parent `listId`.  That returns an array which is then iterated again with map in order to build a new object for each video, removing the `listId` property since its no longer needed.

```javascript
var lists = [
		{
			"id": 5434364,
			"name": "New Releases"
		},
		{
			"id": 65456475,
			name: "Thrillers"
		}
	],
	videos = [
		{
			"listId": 5434364,
			"id": 65432445,
			"title": "The Chamber"
		},
		{
			"listId": 5434364,
			"id": 675465,
			"title": "Fracture"
		},
		{
			"listId": 65456475,
			"id": 70111470,
			"title": "Die Hard"
		},
		{
			"listId": 65456475,
			"id": 654356453,
			"title": "Bad Boys"
		}
	];

var newList = lists.map(function(list) {
	return {
		name: list.name,
		videos:
			videos.
				filter(function(video) {
					return video.listId === list.id;
				}).
				map(function(video) {
					return {id: video.id, title: video.title};
				})
	};
});
```
## Converting Arrays to Deeper Trees

```javascript
var lists = [
		{
			"id": 5434364,
			"name": "New Releases"
		},
		{
			"id": 65456475,
			name: "Thrillers"
		}
	],
	videos = [
		{
			"listId": 5434364,
			"id": 65432445,
			"title": "The Chamber"
		},
		{
			"listId": 5434364,
			"id": 675465,
			"title": "Fracture"
		},
		{
			"listId": 65456475,
			"id": 70111470,
			"title": "Die Hard"
		},
		{
			"listId": 65456475,
			"id": 654356453,
			"title": "Bad Boys"
		}
	],
	boxarts = [
		{ videoId: 65432445, width: 130, height:200, url:"http://cdn-0.nflximg.com/images/2891/TheChamber130.jpg" },
		{ videoId: 65432445, width: 200, height:200, url:"http://cdn-0.nflximg.com/images/2891/TheChamber200.jpg" },
		{ videoId: 675465, width: 200, height:200, url:"http://cdn-0.nflximg.com/images/2891/Fracture200.jpg" },
		{ videoId: 675465, width: 120, height:200, url:"http://cdn-0.nflximg.com/images/2891/Fracture120.jpg" },
		{ videoId: 675465, width: 300, height:200, url:"http://cdn-0.nflximg.com/images/2891/Fracture300.jpg" },
		{ videoId: 70111470, width: 150, height:200, url:"http://cdn-0.nflximg.com/images/2891/DieHard150.jpg" },
		{ videoId: 70111470, width: 200, height:200, url:"http://cdn-0.nflximg.com/images/2891/DieHard200.jpg" },
		{ videoId: 654356453, width: 200, height:200, url:"http://cdn-0.nflximg.com/images/2891/BadBoys200.jpg" },
		{ videoId: 654356453, width: 140, height:200, url:"http://cdn-0.nflximg.com/images/2891/BadBoys140.jpg" }
	],
	bookmarks = [
		{ videoId: 65432445, time: 32432 },
		{ videoId: 675465, time: 3534543 },
		{ videoId: 70111470, time: 645243 },
		{ videoId: 654356453, time: 984934 }
	];

return lists.map(function(list) {
	return {
		name: list.name,
		videos:
			videos.
				filter(function(video) {
					return video.listId === list.id;
				}).
				concatMap(function(video) {
					return Array.zip(
						bookmarks.filter(function(bookmark) {
							return bookmark.videoId === video.id;
						}),
						boxarts.filter(function(boxart) {
							return boxart.videoId === video.id;
						}).
						reduce(function(acc,curr) {
							return acc.width * acc.height < curr.width * curr.height ? acc : curr;
						}),
						function(bookmark, boxart) {
							return { id: video.id, title: video.title, time: bookmark.time, boxart: boxart.url };
						});
			})
	};
});
```

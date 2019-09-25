---
title: Today I Learned Javascript
lang: en-US
tags:
  - javascript
---

# {{ $page.title }}

## Functional Expressions in D&D

I've been having fun with functions while creating a JavaScript version of the 5th edition D&D character sheet.

In recreating the character bonus table I got thru 47 lines of inline conditionals before the formula slapped me in the face:
![](./OG_s.jpg)


Even in adding abilities to an array, I was still left with 6 variables and a declared function:

```javascript
var abilities = [9, 10, 15, 12, 9, 19];  //The [str, dex, con, int, wis, chr] of character
function abilityMod(score) { return Math.floor((score / 2) - 5); }
var strMod = abilityMod(abilities[0]);
var dexMod = abilityMod(abilities[1]);
var conMod = abilityMod(abilities[2]);
var intMod = abilityMod(abilities[3]);
var wisMod = abilityMod(abilities[4]);
var chrMod = abilityMod(abilities[5]);
```

Until I used an expression with arrays and `.map`:

```javascript
var abilities = [9, 10, 15, 12, 9, 19];
var mod = abilities.map(function(score) { return Math.floor((score / 2) - 5); });
```

## NPM won't run arbitrary scripts

NPM packages can declare named commands in the `scripts` section of their `package.json`. NPM scripts are a useful way to run repetitive tasks necessary for maintaining the package.

As an example, the ember.js `package.json` declares

- a `start` script to run a local web server,
- a `test` script to run the package's test suite, and
- a `build` script to transpile, concatenate, and minify the package.

To execute the `start` script, run:

```bash
npm start
```

To execute the `test` script, run:

```bash
npm test
```

This will execute the `pretest` script before the `test` script.

To execute the `build` script, you would probably expect to run:

```bash
npm build # womp womp
```

*This will not behave like `start` or `test`.*

NPM only supports a specific set of scripts and hooks, and `build` is not directly supported.

To execute the `build` script, run:

```bash
npm run-script build
```

It is also important to note that a `prebuild` hook *would not* fire before the `build` script's execution.

[package.json]: https://github.com/emberjs/ember.js/blob/48e115928bcb6b366a621370339354c44aad86b1/package.json
[start]: https://github.com/emberjs/ember.js/blob/48e115928bcb6b366a621370339354c44aad86b1/package.json#L10
[test]: https://github.com/emberjs/ember.js/blob/48e115928bcb6b366a621370339354c44aad86b1/package.json#L8
[pretest]: https://github.com/emberjs/ember.js/blob/48e115928bcb6b366a621370339354c44aad86b1/package.json#L9
[build]: https://github.com/emberjs/ember.js/blob/48e115928bcb6b366a621370339354c44aad86b1/package.json#L6
[scripts]: https://docs.npmjs.com/misc/scripts

## Prototype Your (Character) Class

Continuing with `functional Expressions in D&D`, I decided to use prototypes.

```javascript
function Character (name, hpBonus, profession, hd, castingAbility) {
	this.name = name;
	this.profession = profession;
	this.hd = hd;
	this.hpBonus = hpBonus;
	this.lvl = 1;
	this.profBonus = 2;
	this.castingAbility = castingAbility;
	this.exp = 0;
	this.abilities = new Array(6);
	this.mod = this.abilities.map(function (score){ return Math.floor((score / 2) - 5); } );
	this.ac = 10 + this.mod[1];
	this.hp = this.hd + this.mod[2] + this.dr;
	this.spellSave = 8 + this.profBonus + this.mod[castingAbility];
	this.spellAttack = this.profBonus + this.mod[castingAbility];
}
```

In using a prototype over an object we gain the ability to efficiently store identical functions common to all characters without bloating objects.

```javascript
Character.prototype = {
	levelUp: function (){
		this.lvl++;
		var conMod = this.mod[2];
		var lvlHP = (this.lvl-1)*(((this.hd/2)+1) + conMod);
		this.hp = this.hd + conMod + lvlHP + (this.hpBonus * this.lvl);
		if ((this.lvl - 1) % 4 == 0 ) {
			this.profBonus++;
			this.spellSave = 8 + this.profBonus + this.mod[this.castingAbility];
			this.spellAttack = this.profBonus + this.mod[this.castingAbility];
		}
},
	add: function (paramater, input){
		this.paramater.push(input);
	}
};
```

So I've added the main functionality of leveling up our character, as well the ability to add on items to an inventory, weapons list, or character attributes; should the user decide to add those optional values later.

Running a quick check, passing in & leveling up my sorcerer, Mordai:

```javascript
var abilities = [9, 10, 15, 12, 9, 19];
var Mordai = new Character("Mordai", 1, "Sorcerer", 6, 5);
Mordai.levelUp();
console.log(Mordai);
```

I get back the same results as my hand-calculated values for ability bonuses, level, proficiency bonus, and hit points.
